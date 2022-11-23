package unitpay

import "strconv"

type Unitpay struct {
	SecretKey string
	ProjectId int
}

func New(secretKey string, projectId string) *Unitpay {
	id, err := strconv.Atoi(projectId)
	if err != nil {
		return nil
	}

	return &Unitpay{
		SecretKey: secretKey,
		ProjectId: id,
	}
}
