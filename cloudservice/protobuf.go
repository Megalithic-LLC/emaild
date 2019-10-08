package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
	"github.com/on-prem-net/emaild/model"
)

func AccountFromProtobuf(pbAccount *emailproto.Account) model.Account {
	return model.Account{
		Id:                pbAccount.Id,
		ServiceInstanceId: pbAccount.ServiceInstanceId,
		Name:              pbAccount.Name,
		Email:             pbAccount.Email,
		First:             pbAccount.First,
		Last:              pbAccount.Last,
		DisplayName:       pbAccount.DisplayName,
		Password:          pbAccount.Password,
	}
}

func AccountsFromProtobuf(pbAccounts []emailproto.Account) []*model.Account {
	accounts := []*model.Account{}
	for _, pbAccount := range pbAccounts {
		account := AccountFromProtobuf(&pbAccount)
		accounts = append(accounts, &account)
	}
	return accounts
}
