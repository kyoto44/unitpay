package unitpay

import (
	"net/http"
	"strconv"
)

const (
	APIHost     = "unitpay.ru"
	InitPayment = "initPayment"
	GetPayment  = "getPayment"
	Error       = "error"
)

var UnitpayAllowedIPs = map[string]bool{"31.186.100.49": true, "52.29.152.23": true, "52.19.56.234": true}

type Unitpay struct {
	SecretKey string
	ProjectId int
	client    *http.Client
}

func NewClient(projectId string, secretKey string) *Unitpay {
	id, err := strconv.Atoi(projectId)
	if err != nil {
		return nil
	}

	return &Unitpay{
		SecretKey: secretKey,
		ProjectId: id,
		client:    &http.Client{},
	}
}

func (u *Unitpay) CheckIPAllowed(IPAddress string) bool {
	if _, allowed := UnitpayAllowedIPs[IPAddress]; allowed {
		return allowed
	}
	return false
}

// func (u *Unitpay) GetPaymentURL(paymentInfo *uni) (string, error) {
// 	paymentRequestURL := u.CreatePaymentURL(paymentInfo)
// 	resp, err := u.client.Get(paymentRequestURL)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)

// 	creationResult := &PaymentCreationResult{}
// 	err = json.Unmarshal(body, creationResult)
// 	if err != nil {
// 		return "", err
// 	}

// 	if creationResult.Error != nil {
// 		return "", errors.New(creationResult.Error.Message)
// 	}

// 	return creationResult.Success.RedirectUrl, err
// }

// func (u *Unitpay) GetPaymentStatus(paymentInfo *Payment) (string, error) {

// 	paymentInfoUrl := u.CreatePaymentInfoURL(info.PaymentId, *info.SecretKey)

// 	client := req.C()
// 	resp, err := client.R().Get(paymentInfoUrl)
// 	if err != nil {
// 		return false, err
// 	}

// 	paymentInfo := &unitpay.PaymentInfoResponse{}
// 	err = resp.UnmarshalJson(paymentInfo)
// 	if err != nil {
// 		return false, err
// 	}
// 	if paymentInfo.ErrorMessage != nil {
// 		return false, errors.New(*paymentInfo.ErrorMessage)
// 	}

// 	if paymentInfo.Status != unitpay.Success {
// 		return false, nil
// 	}

// 	return true, nil

// 	paymentRequestURL := u.CreatePaymentURL(paymentInfo)
// 	resp, err := u.client.Get(paymentRequestURL)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)

// 	creationResult := &PaymentCreationResult{}
// 	err = json.Unmarshal(body, creationResult)
// 	if err != nil {
// 		return "", err
// 	}

// 	if creationResult.Error != nil {
// 		return "", errors.New(creationResult.Error.Message)
// 	}

// 	return creationResult.Success.RedirectUrl, err
// }
