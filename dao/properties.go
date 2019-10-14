package dao

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
	"github.com/asdine/genji/table"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
)

type PropertiesDAO struct {
	db *genji.DB
}

func NewPropertiesDAO(db *genji.DB) PropertiesDAO {
	return PropertiesDAO{db: db}
}

func (self PropertiesDAO) Get(key string) (string, error) {
	var retval string
	err := self.db.View(func(tx *genji.Tx) error {
		propertyTable, err := tx.GetTable(model.PropertyTable)
		if err != nil {
			return err
		}
		f := model.NewPropertyFields()
		return query.
			Select().
			From(propertyTable).
			Where(f.Key.Eq(key)).
			Limit(1).
			Run(tx).
			Iterate(func(recordId []byte, r record.Record) error {
				field, err := r.GetField(f.Value.Name())
				retval = string(field.Data)
				return err
			})
	})
	return retval, err
}

func (self PropertiesDAO) Set(key, value string) error {
	return self.db.Update(func(tx *genji.Tx) error {
		propertyTable, err := tx.GetTable(model.PropertyTable)
		if err != nil {
			return err
		}
		searchFor := &model.Property{Key: key}
		pk, err := searchFor.PrimaryKey()
		if err != nil {
			return err
		}
		if _, err := propertyTable.GetRecord(pk); err == table.ErrRecordNotFound {
			_, err := propertyTable.Insert(&model.Property{Key: key, Value: value})
			return err
		}
		property := model.Property{Key: key, Value: value}
		return propertyTable.Replace(pk, &property)
	})
}

func (self PropertiesDAO) SetIfKeyNotExists(key, value string) (string, error) {
	var retval string
	err := self.db.Update(func(tx *genji.Tx) error {
		table, err := tx.GetTable(model.PropertyTable)
		if err != nil {
			return err
		}

		f := model.NewPropertyFields()

		found := false
		if err := query.
			Select().
			From(table).
			Where(f.Key.Eq(key)).
			Run(tx).
			Iterate(func(recordId []byte, r record.Record) error {
				field, err := r.GetField(f.Value.Name())
				retval = string(field.Data)
				found = true
				return err
			}); err != nil {
			return err
		}

		if !found {
			_, err := table.Insert(&model.Property{
				Key:   key,
				Value: value,
			})
			retval = value
			return err
		}

		return nil
	})
	return retval, err
}
