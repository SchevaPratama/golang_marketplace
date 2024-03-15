package converter

import (
	"golang-marketplace/internal/entity"
	"golang-marketplace/internal/model"
)

func BankAccountConverter(bankAccount *entity.BankAccount) *model.BankAccountResponse {
	return &model.BankAccountResponse{
		BankAccountId:     bankAccount.ID,
		BankName:          bankAccount.BankName,
		BankAccountName:   bankAccount.Name,
		BankAccountNumber: bankAccount.Number,
	}
}
