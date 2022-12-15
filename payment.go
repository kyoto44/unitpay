package unitpay

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Payment struct {
	paymentType PaymentType
	account     string
	sum         int
	projectId   int
	resultUrl   *string
	desc        string
	secretKey   string

	currency           *string
	locale             *string
	signature          *string
	ip                 *string
	backUrl            *string
	subscription       *bool
	subscriptionId     *int
	preauth            *bool
	preauthExpireLogic *int
	test               *bool
}

type PaymentCreationResultSuccess struct {
	Message      string `json:"message"`
	PaymentId    int64  `json:"paymentId"`
	RecieptUrl   string `json:"receiptUrl"`
	ResponseType string `json:"type"`
	RedirectUrl  string `json:"redirectUrl,omitempty"`
	Response     string `json:"responce,omitempty"`
	InvoiceId    string `json:"invoiceId,omitempty"`
}

type PaymentCreationResultError struct {
	Message string `json:"message"`
}

type PaymentCreationResult struct {
	Success *PaymentCreationResultSuccess `json:"result,omitempty"`
	Error   *PaymentCreationResultError   `json:"error,omitempty"`
}

type PaymentResult struct {
	Method         string
	UnitpayId      int
	ProjectId      int
	Account        string
	PayerSum       int
	PayerCurrency  string
	Profit         int
	PaymentType    string
	OrderSum       int
	OrderCurrency  string
	Operator       string
	Date           time.Time
	ErrorMessage   string
	Test           string
	Phone          *int64
	ThreeDS        *int
	SubscriptionId *int
	Signature      *string
}

type PaymentResultResponseSuccess struct {
	Message string `json:"message"`
}

type PaymentResultResponseError struct {
	Message string `json:"message"`
}

type PaymentResultResponse struct {
	Success *PaymentResultResponseSuccess `json:"result,omitempty"`
	Error   *PaymentResultResponseError   `json:"error,omitempty"`
}

type PaymentInfoResponse struct {
	Status        PaymentStatus
	PaymentId     int
	ProjectId     int
	Account       string
	Purse         string
	Profit        int
	PaymentType   string
	OrderSum      int
	OrderCurrency string
	Date          time.Time
	PayerSum      int
	PayerCurrency string
	RecieptUrl    string
	ErrorMessage  *string
	Message       *string
}

func CreatePayment(paymentType PaymentType, account string, sum int, projectID int, description string, resultURL *string, secretKey string) *Payment {
	return &Payment{
		paymentType: paymentType,
		account:     account,
		sum:         sum,
		projectId:   projectID,
		desc:        description,
		resultUrl:   resultURL,
		secretKey:   secretKey,
	}
}

func (payment *Payment) SetSignature(signature string) {
	payment.signature = &signature
}

func (payment *Payment) GetSignature() string {
	return *payment.signature
}

func (payment *Payment) SetLocale(locale string) {
	payment.locale = &locale
}

func (payment *Payment) SetCurrency(currency string) {
	payment.currency = &currency
}

func (payment *Payment) GetCurrency() string {
	return *payment.currency
}

func (payment *Payment) SetIPAddress(IPAddress string) {
	payment.ip = &IPAddress
}

func (payment *Payment) SetBackURL(backUrl string) {
	payment.backUrl = &backUrl
}

func (payment *Payment) SetSubscription(subscription bool) {
	payment.subscription = &subscription
}

func (payment *Payment) SetSubscriptionID(subscriptionId int) {
	payment.subscriptionId = &subscriptionId
}

func (payment *Payment) SetPreAuth(preauth bool) {
	payment.preauth = &preauth
}

func (payment *Payment) SetPreAuthExpireLogic(preauthExpireLogic int) {
	payment.preauthExpireLogic = &preauthExpireLogic
}

func (payment *Payment) SetTestMode(test bool) {
	payment.test = &test
}

func (payment *Payment) CreateSignature() string {
	hashString := fmt.Sprintf("%s{up}%s{up}%d{up}%d{up}%s{up}%s", payment.paymentType, payment.account, payment.sum, payment.projectId, payment.desc, payment.secretKey)
	hash := sha256.Sum256([]byte(hashString))
	signature := hex.EncodeToString(hash[:])
	return signature
}

func CreatePaymentURL(payment *Payment) string {
	paymentUrl := &url.URL{
		Scheme: "https",
		Host:   UnitpayAPIHost,
		Path:   "/api",
	}

	paymentUrlParams := paymentUrl.Query()
	paymentUrlParams.Set("method", UnitpayMethodInitPayment)
	paymentUrlParams.Set("params[paymentType]", string(payment.paymentType))
	paymentUrlParams.Set("params[account]", payment.account)
	paymentUrlParams.Set("params[sum]", strconv.Itoa(payment.sum))
	paymentUrlParams.Set("params[projectId]", strconv.Itoa(payment.projectId))
	if payment.resultUrl != nil {
		paymentUrlParams.Set("params[resultUrl]", *payment.resultUrl)
	}
	paymentUrlParams.Set("params[desc]", payment.desc)
	paymentUrlParams.Set("params[secretKey]", payment.secretKey)
	paymentUrlParams.Set("params[signature]", payment.CreateSignature())

	if payment.test != nil {
		paymentUrlParams.Set("params[test]", "1")
	}
	if payment.currency != nil {
		paymentUrlParams.Set("params[currency]", *payment.currency)
	}
	if payment.locale != nil {
		paymentUrlParams.Set("params[locale]", *payment.locale)
	}
	if payment.signature != nil {
		paymentUrlParams.Set("params[signature]", *payment.signature)
	}
	if payment.ip != nil {
		paymentUrlParams.Set("params[ip]", *payment.ip)
	}
	if payment.backUrl != nil {
		paymentUrlParams.Set("params[backUrl]", *payment.backUrl)
	}
	if payment.subscription != nil {
		paymentUrlParams.Set("params[subscription]", strconv.FormatBool(*payment.subscription))
	}
	if payment.subscriptionId != nil {
		paymentUrlParams.Set("params[subscriptionId]", strconv.Itoa(*payment.subscriptionId))
	}
	if payment.preauth != nil {
		paymentUrlParams.Set("params[preauth]", strconv.FormatBool(*payment.preauth))
	}
	if payment.preauthExpireLogic != nil {
		paymentUrlParams.Set("params[preauthExpireLogic]", strconv.Itoa(*payment.preauthExpireLogic))
	}

	paymentUrl.RawQuery = paymentUrlParams.Encode()

	return paymentUrl.String()
}

func CreatePaymentInfoURL(paymentId string, secretKey string) string {
	paymentUrl := &url.URL{
		Scheme: "https",
		Host:   UnitpayAPIHost,
		Path:   "/api",
	}

	paymentUrlParams := paymentUrl.Query()
	paymentUrlParams.Set("method", UnitpayMethodGetPayment)
	paymentUrlParams.Set("params[paymentId]", paymentId)
	paymentUrlParams.Set("params[secretKey]", secretKey)

	paymentUrl.RawQuery = paymentUrlParams.Encode()

	return paymentUrl.String()
}