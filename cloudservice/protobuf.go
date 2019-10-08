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
		DomainId:          pbAccount.DomainId,
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

func DomainFromProtobuf(pbDomain *emailproto.Domain) model.Domain {
	return model.Domain{
		Id:                pbDomain.Id,
		ServiceInstanceId: pbDomain.ServiceInstanceId,
		Name:              pbDomain.Name,
	}
}

func DomainsFromProtobuf(pbDomains []emailproto.Domain) []*model.Domain {
	domains := []*model.Domain{}
	for _, pbDomain := range pbDomains {
		domain := DomainFromProtobuf(&pbDomain)
		domains = append(domains, &domain)
	}
	return domains
}

func ServiceInstanceFromProtobuf(pbServiceInstance *emailproto.ServiceInstance) model.ServiceInstance {
	return model.ServiceInstance{
		Id:        pbServiceInstance.Id,
		ServiceId: pbServiceInstance.ServiceId,
		PlanId:    pbServiceInstance.PlanId,
	}
}

func SnapshotFromProtobuf(pbSnapshot *emailproto.Snapshot) model.Snapshot {
	return model.Snapshot{
		Id:   pbSnapshot.Id,
		Name: pbSnapshot.Name,
	}
}

func SnapshotsFromProtobuf(pbSnapshots []emailproto.Snapshot) []*model.Snapshot {
	snapshots := []*model.Snapshot{}
	for _, pbSnapshot := range pbSnapshots {
		snapshot := SnapshotFromProtobuf(&pbSnapshot)
		snapshots = append(snapshots, &snapshot)
	}
	return snapshots
}
