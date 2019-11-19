package etcd

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewEtcd(t *testing.T) {
	e := NewEtcd([]string{"localhost:2379"})
	getResp, err := e.Kv.Get(context.TODO(), "hello")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(getResp.Kvs)
}

func TestEtcd_RegisterService(t *testing.T) {
	e := NewEtcd([]string{"localhost:2379"})
	if err := e.RegisterService("nodes", "001"); err != nil {
		t.Error(err)
	}
	fmt.Println("start lease...")
	time.Sleep(20 * time.Second)
	fmt.Println("end lease...")
}

func TestEtcd_WatchService(t *testing.T) {
	e := NewEtcd([]string{"localhost:2379"})
	e.WatchService("sync/nodes")
}
