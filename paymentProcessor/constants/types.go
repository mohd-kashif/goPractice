package constants

type PaymentMethodType string

const (
	PaymentMethodCreditCard PaymentMethodType = "Credit Card"
	PaymentMethodDebitCard  PaymentMethodType = "Debit Card"
	PaymentMethodUPI        PaymentMethodType = "UPI"
	PaymentMethodNetBanking PaymentMethodType = "Net Banking"
)

type PaymentMethodStatus string

const (
	PaymentMethodStatusActive      PaymentMethodStatus = "Active"
	PaymentMethodStatusDeactivated PaymentMethodStatus = "Deactivated"
)
