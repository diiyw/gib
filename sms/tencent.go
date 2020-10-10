package sms

import (
	"errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

type TencentSMS struct {
	appID      string
	credential *common.Credential
	profile    *profile.ClientProfile
	client     *sms.Client
}

func NewTencentSMS(secretId, secretKey string, appID string) (*TencentSMS, error) {
	credential := common.NewCredential(
		secretId,
		secretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 10
	cpf.SignMethod = "HmacSHA1"

	client, err := sms.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		return nil, err
	}
	return &TencentSMS{
		appID:      appID,
		credential: credential,
		profile:    cpf,
		client:     client,
	}, nil
}

func (t *TencentSMS) Send(templateId string, params, phones []string) error {
	req := sms.NewSendSmsRequest()
	req.SmsSdkAppid = common.StringPtr(t.appID)
	for i, phone := range phones {
		phones[i] = "86" + phone
	}
	req.PhoneNumberSet = common.StringPtrs(phones)
	req.TemplateParamSet = common.StringPtrs(params)
	req.TemplateID = common.StringPtr(templateId)

	resp, err := t.client.SendSms(req)
	if err != nil {
		return err
	}
	if *resp.Response.SendStatusSet[0].Code != "Ok" {
		return errors.New(*resp.Response.SendStatusSet[0].Message)
	}
	_ = resp.Response
	return nil
}
