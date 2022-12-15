package models

import "time"

type Payment struct {
	paymentType Type
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

type CreationResult struct {
	Success *struct {
		Message      string `json:"message"`
		PaymentId    int64  `json:"paymentId"`
		RecieptUrl   string `json:"receiptUrl"`
		ResponseType string `json:"type"`
		RedirectUrl  string `json:"redirectUrl,omitempty"`
		Response     string `json:"responce,omitempty"`
		InvoiceId    string `json:"invoiceId,omitempty"`
	} `json:"result,omitempty"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
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

type ResultResponse struct {
	Success *struct {
		Message string `json:"message"`
	} `json:"result,omitempty"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

type InfoResponse struct {
	Status        Status
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
