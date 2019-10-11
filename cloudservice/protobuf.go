package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
	"github.com/on-prem-net/emaild/model"
)

func AccountFromProtobuf(pbAccount *emailproto.Account) model.Account {
	return model.Account{
		Id:          pbAccount.Id,
		Name:        pbAccount.Name,
		DomainId:    pbAccount.DomainId,
		Email:       pbAccount.Email,
		First:       pbAccount.First,
		Last:        pbAccount.Last,
		DisplayName: pbAccount.DisplayName,
		Password:    pbAccount.Password,
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
		Id:   pbDomain.Id,
		Name: pbDomain.Name,
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

func EndpointFromProtobuf(pbEndpoint *emailproto.Endpoint) model.Endpoint {
	return model.Endpoint{
		Id:       pbEndpoint.Id,
		Protocol: pbEndpoint.Protocol,
		Type:     pbEndpoint.Type,
		Port:     uint16(pbEndpoint.Port),
		Path:     pbEndpoint.Path,
		Enabled:  pbEndpoint.Enabled,
	}
}

func EndpointsFromProtobuf(pbEndpoints []emailproto.Endpoint) []*model.Endpoint {
	endpoints := []*model.Endpoint{}
	for _, pbEndpoint := range pbEndpoints {
		endpoint := EndpointFromProtobuf(&pbEndpoint)
		endpoints = append(endpoints, &endpoint)
	}
	return endpoints
}

func SnapshotFromProtobuf(pbSnapshot *emailproto.Snapshot) model.Snapshot {
	return model.Snapshot{
		Id:       pbSnapshot.Id,
		Name:     pbSnapshot.Name,
		Engine:   pbSnapshot.Engine,
		Progress: pbSnapshot.Progress,
		Size:     pbSnapshot.Size,
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

func SnapshotToProtobuf(snapshot *model.Snapshot) *emailproto.Snapshot {
	if snapshot == nil {
		return nil
	}
	return &emailproto.Snapshot{
		Id:        snapshot.Id,
		ServiceId: snapshot.ServiceId,
		Name:      snapshot.Name,
		Engine:    snapshot.Engine,
		Progress:  snapshot.Progress,
		Size:      snapshot.Size,
	}
}
