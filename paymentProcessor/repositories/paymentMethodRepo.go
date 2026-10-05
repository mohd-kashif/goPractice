package repositories

import (
	"errors"
	"goPractice/paymentProcessor/constants"
	"goPractice/paymentProcessor/models"
)

type PaymentMethodRepo struct {
	PaymentMethods map[constants.PaymentMethodType]*models.PaymentMethod
}

func NewPaymentMethodRepo() PaymentMethodRepo {
	return PaymentMethodRepo{
		PaymentMethods: make(map[constants.PaymentMethodType]*models.PaymentMethod),
	}
}

func (pmr *PaymentMethodRepo) AddNewPaymentMethod(name constants.PaymentMethodType) (*string, error) {
	_, exist := pmr.PaymentMethods[name]
	if exist {
		return nil, errors.New("payment method already exist")
	}

	paymentMethod := models.NewPaymentMethod(name)
	return &paymentMethod.ID, nil
}
