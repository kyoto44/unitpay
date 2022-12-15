package unitpay

import (
	"net/url"
	"strconv"
)

func CreateURL(payment *Payment) string {
	paymentUrl := &url.URL{
		Scheme: "https",
		Host:   d,
		Path:   "/api",
	}

	paymentUrlParams := paymentUrl.Query()
	paymentUrlParams.Set("method", u.InitPayment)
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

func CreateInfoURL(paymentId string, secretKey string) string {
	paymentUrl := &url.URL{
		Scheme: "https",
		Host:   APIHost,
		Path:   "/api",
	}

	paymentUrlParams := paymentUrl.Query()
	paymentUrlParams.Set("method", GetPayment)
	paymentUrlParams.Set("params[paymentId]", paymentId)
	paymentUrlParams.Set("params[secretKey]", secretKey)

	paymentUrl.RawQuery = paymentUrlParams.Encode()

	return paymentUrl.String()
}
