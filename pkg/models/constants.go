package models

type Method string

const (
	Check   Method = "check"
	Pay     Method = "pay"
	PreAuth Method = "preauth"
)

type Type string

const (
	BankCard    Type = "card"
	QIWI        Type = "qiwi"
	SBP         Type = "sbp"
	YandexPay   Type = "yandexpay"
	ApplePay    Type = "applepay"
	Paypal      Type = "paypal"
	Webmoney    Type = "webmoney"
	Mobile      Type = "mc"
	WebmoneyWMR Type = "webmoneyWmr"
	Yandex      Type = "yandex"
	SamsungPay  Type = "samsungpay"
	GooglePay   Type = "googlepay"
)

type Status string

const (
	Success    Status = "success"
	Wait       Status = "wait"
	ErrorPay   Status = "error_pay"
	ErrorCheck Status = "error_check"
	Refund     Status = "refund"
	Secure     Status = "secure"
)
