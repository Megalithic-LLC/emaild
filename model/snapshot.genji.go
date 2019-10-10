// Code generated by genji.
// DO NOT EDIT!

package model

import (
	"errors"

	"github.com/asdine/genji/field"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
)

// GetField implements the field method of the record.Record interface.
func (s *Snapshot) GetField(name string) (field.Field, error) {
	switch name {
	case "Id":
		return field.NewString("Id", s.Id), nil
	case "Name":
		return field.NewString("Name", s.Name), nil
	case "Engine":
		return field.NewString("Engine", s.Engine), nil
	case "Progress":
		return field.NewFloat32("Progress", s.Progress), nil
	case "Size":
		return field.NewUint64("Size", s.Size), nil
	}

	return field.Field{}, errors.New("unknown field")
}

// Iterate through all the fields one by one and pass each of them to the given function.
// It the given function returns an error, the iteration is interrupted.
func (s *Snapshot) Iterate(fn func(field.Field) error) error {
	var err error

	err = fn(field.NewString("Id", s.Id))
	if err != nil {
		return err
	}

	err = fn(field.NewString("Name", s.Name))
	if err != nil {
		return err
	}

	err = fn(field.NewString("Engine", s.Engine))
	if err != nil {
		return err
	}

	err = fn(field.NewFloat32("Progress", s.Progress))
	if err != nil {
		return err
	}

	err = fn(field.NewUint64("Size", s.Size))
	if err != nil {
		return err
	}

	return nil
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (s *Snapshot) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "Id":
			s.Id, err = field.DecodeString(f.Data)
		case "Name":
			s.Name, err = field.DecodeString(f.Data)
		case "Engine":
			s.Engine, err = field.DecodeString(f.Data)
		case "Progress":
			s.Progress, err = field.DecodeFloat32(f.Data)
		case "Size":
			s.Size, err = field.DecodeUint64(f.Data)
		}
		return err
	})
}

// PrimaryKey returns the primary key. It implements the table.PrimaryKeyer interface.
func (s *Snapshot) PrimaryKey() ([]byte, error) {
	return field.EncodeString(s.Id), nil
}

// SnapshotFields describes the fields of the Snapshot record.
// It can be used to select fields during queries.
type SnapshotFields struct {
	Id       query.StringFieldSelector
	Name     query.StringFieldSelector
	Engine   query.StringFieldSelector
	Progress query.Float32FieldSelector
	Size     query.Uint64FieldSelector
}

// NewSnapshotFields creates a SnapshotFields.
func NewSnapshotFields() *SnapshotFields {
	return &SnapshotFields{
		Id:       query.StringField("Id"),
		Name:     query.StringField("Name"),
		Engine:   query.StringField("Engine"),
		Progress: query.Float32Field("Progress"),
		Size:     query.Uint64Field("Size"),
	}
}