package dao

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/rs/xid"
)

func (self DomainsDAO) CreateTx(tx *genji.Tx, domain *model.Domain) error {
	if domainTable, err := tx.GetTable(model.DomainTable); err != nil {
		return err
	} else {
		if domain.Id == "" {
			domain.Id = xid.New().String()
		}
		_, err := domainTable.Insert(domain)
		return err
	}
}

func (self DomainsDAO) DeleteAllTx(tx *genji.Tx) error {
	domainTable, err := tx.GetTable(model.DomainTable)
	if err != nil {
		return err
	}
	return query.
		Delete().
		From(domainTable).
		Run(tx)
}

func (self DomainsDAO) FindByIdTx(tx *genji.Tx, id string) (*model.Domain, error) {
	domainTable, err := tx.GetTable(model.DomainTable)
	if err != nil {
		return nil, err
	}
	searchFor := &model.Domain{Id: id}
	pk, err := searchFor.PrimaryKey()
	if err != nil {
		return nil, err
	}
	r, err := domainTable.GetRecord(pk)
	if err != nil {
		return nil, err
	}
	var domain model.Domain
	err = domain.ScanRecord(r)
	return &domain, err
}

func (self DomainsDAO) ReplaceTx(tx *genji.Tx, domain *model.Domain) error {
	domainTable, err := tx.GetTable(model.DomainTable)
	if err != nil {
		return err
	}
	pk, err := domain.PrimaryKey()
	if err != nil {
		return err
	}
	return domainTable.Replace(pk, domain)
}
