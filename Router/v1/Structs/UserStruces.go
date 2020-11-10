package Structs

type WechatCode struct {
	Code string
}

type ToLogin struct {
	UserName string
	PassWord string
}

type Register struct {
	UserName string
	PassWord string
	SMSCode string
}