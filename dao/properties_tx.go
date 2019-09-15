package dao

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
	"github.com/drauschenbach/megalithicd/model"
)

func (self PropertiesDAO) GetTx(tx *genji.Tx, key string) (string, error) {
	propertyTable, err := tx.GetTable(model.PropertyTable)
	if err != nil {
		return "", err
	}
	f := model.NewPropertyFields()
	var retval string
	err = query.
		Select().
		From(propertyTable).
		Where(f.Key.Eq(key)).
		Limit(1).
		Run(tx).
		Iterate(func(recordID []byte, r record.Record) error {
			field, err := r.GetField(f.Value.Name())
			retval = string(field.Data)
			return err
		})
	return retval, err
}

func (self PropertiesDAO) SetTx(tx *genji.Tx, key, value string) error {
	if table, err := tx.GetTable(model.PropertyTable); err != nil {
		return err
	} else {
		_, err := table.Insert(&model.Property{Key: key, Value: value})
		return err
	}
}

func (self PropertiesDAO) SetIfKeyNotExistsTx(tx *genji.Tx, key, value string) (string, error) {
	table, err := tx.GetTable(model.PropertyTable)
	if err != nil {
		return "", err
	}

	f := model.NewPropertyFields()

	found := false
	var retval string
	if err := query.
		Select().
		From(table).
		Where(f.Key.Eq(key)).
		Run(tx).
		Iterate(func(recordID []byte, r record.Record) error {
			field, err := r.GetField(f.Value.Name())
			retval = string(field.Data)
			found = true
			return err
		}); err != nil {
		return "", err
	}

	if !found {
		_, err := table.Insert(&model.Property{
			Key:   key,
			Value: value,
		})
		return value, err
	}

	return retval, err
}
