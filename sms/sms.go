package sms

type Sms interface {
	Send(templateId string, params, phones []string) error
}
