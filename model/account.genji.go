// Code generated by genji.
// DO NOT EDIT!

package model

import (
	"errors"

	"github.com/asdine/genji/field"
	"github.com/asdine/genji/index"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
)

// GetField implements the field method of the record.Record interface.
func (a *Account) GetField(name string) (field.Field, error) {
	switch name {
	case "Id":
		return field.NewString("Id", a.Id), nil
	case "Name":
		return field.NewString("Name", a.Name), nil
	case "Email":
		return field.NewString("Email", a.Email), nil
	case "First":
		return field.NewString("First", a.First), nil
	case "Last":
		return field.NewString("Last", a.Last), nil
	case "DisplayName":
		return field.NewString("DisplayName", a.DisplayName), nil
	case "Password":
		return field.NewBytes("Password", a.Password), nil
	}

	return field.Field{}, errors.New("unknown field")
}

// Iterate through all the fields one by one and pass each of them to the given function.
// It the given function returns an error, the iteration is interrupted.
func (a *Account) Iterate(fn func(field.Field) error) error {
	var err error

	err = fn(field.NewString("Id", a.Id))
	if err != nil {
		return err
	}

	err = fn(field.NewString("Name", a.Name))
	if err != nil {
		return err
	}

	err = fn(field.NewString("Email", a.Email))
	if err != nil {
		return err
	}

	err = fn(field.NewString("First", a.First))
	if err != nil {
		return err
	}

	err = fn(field.NewString("Last", a.Last))
	if err != nil {
		return err
	}

	err = fn(field.NewString("DisplayName", a.DisplayName))
	if err != nil {
		return err
	}

	err = fn(field.NewBytes("Password", a.Password))
	if err != nil {
		return err
	}

	return nil
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (a *Account) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "Id":
			a.Id, err = field.DecodeString(f.Data)
		case "Name":
			a.Name, err = field.DecodeString(f.Data)
		case "Email":
			a.Email, err = field.DecodeString(f.Data)
		case "First":
			a.First, err = field.DecodeString(f.Data)
		case "Last":
			a.Last, err = field.DecodeString(f.Data)
		case "DisplayName":
			a.DisplayName, err = field.DecodeString(f.Data)
		case "Password":
			a.Password, err = field.DecodeBytes(f.Data)
		}
		return err
	})
}

// PrimaryKey returns the primary key. It implements the table.PrimaryKeyer interface.
func (a *Account) PrimaryKey() ([]byte, error) {
	return field.EncodeString(a.Id), nil
}

// Indexes creates a map containing the configuration for each index of the table.
func (a *Account) Indexes() map[string]index.Options {
	return map[string]index.Options{
		"Email": index.Options{Unique: true},
	}
}

// AccountFields describes the fields of the Account record.
// It can be used to select fields during queries.
type AccountFields struct {
	Id          query.StringFieldSelector
	Name        query.StringFieldSelector
	Email       query.StringFieldSelector
	First       query.StringFieldSelector
	Last        query.StringFieldSelector
	DisplayName query.StringFieldSelector
	Password    query.BytesFieldSelector
}

// NewAccountFields creates a AccountFields.
func NewAccountFields() *AccountFields {
	return &AccountFields{
		Id:          query.StringField("Id"),
		Name:        query.StringField("Name"),
		Email:       query.StringField("Email"),
		First:       query.StringField("First"),
		Last:        query.StringField("Last"),
		DisplayName: query.StringField("DisplayName"),
		Password:    query.BytesField("Password"),
	}
}
