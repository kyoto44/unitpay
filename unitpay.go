package unitpay

type Unitpay struct {
	ProjectId string
	SecretKey string
}

func New(projectId string, secretKey string) *Unitpay {
	return &Unitpay{
		SecretKey: secretKey,
		ProjectId: projectId,
	}
}
