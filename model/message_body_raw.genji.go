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
func (m *MessageBodyRaw) GetField(name string) (field.Field, error) {
	switch name {
	case "ID":
		return field.NewUint32("ID", m.ID), nil
	case "Body":
		return field.NewBytes("Body", m.Body), nil
	case "Unused":
		return field.NewUint32("Unused", m.Unused), nil
	}

	return field.Field{}, errors.New("unknown field")
}

// Iterate through all the fields one by one and pass each of them to the given function.
// It the given function returns an error, the iteration is interrupted.
func (m *MessageBodyRaw) Iterate(fn func(field.Field) error) error {
	var err error

	err = fn(field.NewUint32("ID", m.ID))
	if err != nil {
		return err
	}

	err = fn(field.NewBytes("Body", m.Body))
	if err != nil {
		return err
	}

	err = fn(field.NewUint32("Unused", m.Unused))
	if err != nil {
		return err
	}

	return nil
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (m *MessageBodyRaw) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "ID":
			m.ID, err = field.DecodeUint32(f.Data)
		case "Body":
			m.Body, err = field.DecodeBytes(f.Data)
		case "Unused":
			m.Unused, err = field.DecodeUint32(f.Data)
		}
		return err
	})
}

// PrimaryKey returns the primary key. It implements the table.PrimaryKeyer interface.
func (m *MessageBodyRaw) PrimaryKey() ([]byte, error) {
	return field.EncodeUint32(m.ID), nil
}

// Indexes creates a map containing the configuration for each index of the table.
func (m *MessageBodyRaw) Indexes() map[string]index.Options {
	return map[string]index.Options{
		"Unused": index.Options{Unique: false},
	}
}

// MessageBodyRawFields describes the fields of the MessageBodyRaw record.
// It can be used to select fields during queries.
type MessageBodyRawFields struct {
	ID     query.Uint32FieldSelector
	Body   query.BytesFieldSelector
	Unused query.Uint32FieldSelector
}

// NewMessageBodyRawFields creates a MessageBodyRawFields.
func NewMessageBodyRawFields() *MessageBodyRawFields {
	return &MessageBodyRawFields{
		ID:     query.Uint32Field("ID"),
		Body:   query.BytesField("Body"),
		Unused: query.Uint32Field("Unused"),
	}
}
