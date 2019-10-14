package dao

import (
	"github.com/asdine/genji"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
)

type DomainsDAO struct {
	db     *genji.DB
	fields *model.DomainFields
}

func NewDomainsDAO(db *genji.DB) DomainsDAO {
	return DomainsDAO{
		db:     db,
		fields: model.NewDomainFields(),
	}
}

func (self DomainsDAO) Create(domain *model.Domain) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.CreateTx(tx, domain)
	})
}

func (self DomainsDAO) FindById(id string) (*model.Domain, error) {
	var retval *model.Domain
	err := self.db.View(func(tx *genji.Tx) error {
		domain, err := self.FindByIdTx(tx, id)
		if err == nil {
			retval = domain
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return retval, nil
}

func (self DomainsDAO) Replace(domain *model.Domain) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.ReplaceTx(tx, domain)
	})
}
