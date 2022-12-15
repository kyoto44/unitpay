package unitpay

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"unitpay/pkg/models"
)

func CreatePayment(paymentType Type, account string, sum int, projectID int, description string, resultURL *string, secretKey string) *models.Payment {
	return &models.Payment{
		paymentType: paymentType,
		account:     account,
		sum:         sum,
		projectId:   projectID,
		desc:        description,
		resultUrl:   resultURL,
		secretKey:   secretKey,
	}
}

func (payment *models.Payment) SetSignature(signature string) {
	payment.signature = &signature
}

func (payment *models.Payment) GetSignature() string {
	return *models.payment.signature
}

func (payment *models.Payment) SetLocale(locale string) {
	payment.locale = &locale
}

func (payment *models.Payment) SetCurrency(currency string) {
	payment.currency = &currency
}

func (payment *models.Payment) GetCurrency() string {
	return *models.payment.currency
}

func (payment *models.Payment) SetIPAddress(IPAddress string) {
	payment.ip = &IPAddress
}

func (payment *models.Payment) SetBackURL(backUrl string) {
	payment.backUrl = &backUrl
}

func (payment *models.Payment) SetSubscription(subscription bool) {
	payment.subscription = &subscription
}

func (payment *models.Payment) SetSubscriptionID(subscriptionId int) {
	payment.subscriptionId = &subscriptionId
}

func (payment *models.Payment) SetPreAuth(preauth bool) {
	payment.preauth = &preauth
}

func (payment *models.Payment) SetPreAuthExpireLogic(preauthExpireLogic int) {
	payment.preauthExpireLogic = &preauthExpireLogic
}

func (payment *models.Payment) SetTestMode(test bool) {
	payment.test = &test
}

func (payment *models.Payment) CreateSignature() string {
	hashString := fmt.Sprintf("%s{up}%s{up}%d{up}%d{up}%s{up}%s", payment.paymentType, payment.account, payment.sum, payment.projectId, payment.desc, payment.secretKey)
	hash := sha256.Sum256([]byte(hashString))
	signature := hex.EncodeToString(hash[:])
	return signature
}
