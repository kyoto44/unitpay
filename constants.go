package unitpay

const (
	UnitpayAPIHost           = "unitpay.ru"
	UnitpayMethodInitPayment = "initPayment"
	UnitpayMethodGetPayment  = "getPayment"
	Error                    = "error"
)

type PaymentMethod string

const (
	Check   PaymentMethod = "check"
	Pay     PaymentMethod = "pay"
	PreAuth PaymentMethod = "preauth"
)

type PaymentType string

const (
	BankCard    PaymentType = "card"
	QIWI        PaymentType = "qiwi"
	SBP         PaymentType = "sbp"
	YandexPay   PaymentType = "yandexpay"
	ApplePay    PaymentType = "applepay"
	Paypal      PaymentType = "paypal"
	Webmoney    PaymentType = "webmoney"
	Mobile      PaymentType = "mc"
	WebmoneyWMR PaymentType = "webmoneyWmr"
	Yandex      PaymentType = "yandex"
	SamsungPay  PaymentType = "samsungpay"
	GooglePay   PaymentType = "googlepay"
)

type PaymentStatus string

const (
	Success    PaymentStatus = "success"
	Wait       PaymentStatus = "wait"
	ErrorPay   PaymentStatus = "error_pay"
	ErrorCheck PaymentStatus = "error_check"
	Refund     PaymentStatus = "refund"
	Secure     PaymentStatus = "secure"
)
