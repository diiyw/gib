package etcd

import (
	"context"
	"crypto/tls"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/etcd-io/etcd/clientv3"
	"time"
)

var (
	DefaultDialTimeout                  = 2 * time.Second
	DefaultKeepAliveTime                = 2 * time.Second
	DefaultKeepAliveTimeOut             = 6 * time.Second
	DefaultTLS              *tls.Config = nil
)

type Etcd struct {
	Client       *clientv3.Client
	Kv           clientv3.KV
	Lease        clientv3.Lease
	ServiceNodes map[string]string
}

func NewEtcd(endpoints []string) *Etcd {

	e := new(Etcd)

	e.ServiceNodes = make(map[string]string, 0)
	// 客户端配置
	config := clientv3.Config{
		Endpoints:            endpoints,
		DialTimeout:          DefaultDialTimeout,
		DialKeepAliveTime:    DefaultKeepAliveTime,
		DialKeepAliveTimeout: DefaultKeepAliveTimeOut,
	}

	if DefaultTLS != nil {
		config.TLS = DefaultTLS
	}

	// 建立连接
	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}

	e.Client = client
	e.Kv = clientv3.NewKV(client)
	e.Lease = clientv3.NewLease(client)
	return e
}

func (etcd *Etcd) RegisterService(key, val string) error {

	etcd.ServiceNodes = make(map[string]string, 0)

	// 申请一个5秒的租约
	leaseGrantResp, err := etcd.Lease.Grant(context.TODO(), 5)
	if err != nil {
		return err
	}

	// 5秒后会取消自动续租
	keepRespChan, err := etcd.Lease.KeepAlive(context.TODO(), leaseGrantResp.ID)
	if err != nil {
		return err
	}

	go func() {

		for {
			select {
			case _ = <-keepRespChan:
				if keepRespChan == nil {
					// 自动续约结束
					return
				}
			}
		}
	}()
	_, err = etcd.Kv.Put(context.TODO(), key, val, clientv3.WithLease(leaseGrantResp.ID))
	if err != nil {
		return err
	}
	return nil
}

func (etcd *Etcd) WatchService(key string) {

	// 首次获取一下值，后续监听节点变化
	getResp, err := etcd.Client.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}

	// 注册节点
	for _, key := range getResp.Kvs {
		etcd.ServiceNodes[string(key.Key)] = string(key.Value)
	}

	watchChan := etcd.Client.Watch(context.Background(), key, clientv3.WithPrefix())

	for wResp := range watchChan {
		for _, ev := range wResp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				etcd.ServiceNodes[string(ev.Kv.Key)] = string(ev.Kv.Value)
			case mvccpb.DELETE:
				delete(etcd.ServiceNodes, string(ev.Kv.Key))
			}
		}
	}
}
