package sms

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

type TencentSMS struct {
	credential *common.Credential
	profile    *profile.ClientProfile
	client     *sms.Client
}

func NewTencentSMS(secretId, secretKey string) (*TencentSMS, error) {
	credential := common.NewCredential(
		secretId,
		secretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	cpf.HttpProfile.ReqTimeout = 5
	cpf.SignMethod = "HmacSHA1"

	client, err := sms.NewClient(credential, "", cpf)
	if err != nil {
		return nil, err
	}
	return &TencentSMS{
		credential: credential,
		profile:    cpf,
		client:     client,
	}, nil
}

func (t *TencentSMS) Send(templateId string, params, phones []string) error {
	smsRequest := sms.NewSendSmsRequest()

	var phoneCache []*string
	for _, phone := range phones {
		phoneCache = append(phoneCache, &phone)
	}

	var paramsCache []*string
	for _, param := range params {
		paramsCache = append(paramsCache, &param)
	}

	smsRequest.PhoneNumberSet = paramsCache
	smsRequest.TemplateID = &templateId
	_, err := t.client.SendSms(smsRequest)
	if err != nil {
		return err
	}
	return nil
}
