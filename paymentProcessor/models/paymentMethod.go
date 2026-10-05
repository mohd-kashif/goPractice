package models

import (
	"goPractice/paymentProcessor/constants"
	"time"

	"github.com/google/uuid"
)

type PaymentMethod struct {
	ID         string
	MethodName constants.PaymentMethodType
	Status     constants.PaymentMethodStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewPaymentMethod(name constants.PaymentMethodType) PaymentMethod {
	return PaymentMethod{
		ID:         uuid.NewString(),
		MethodName: name,
		Status:     constants.PaymentMethodStatusDeactivated,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
