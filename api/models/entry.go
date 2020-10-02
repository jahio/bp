package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Entry is used by pop to map your entries database table to your go code.
type Entry struct {
	ID        uuid.UUID         `json:"id"         db:"id"`
	CreatedAt int               `json:"created_at" db:"created_at"`
	UpdatedAt int               `json:"updated_at" db:"updated_at"`
	Systolic  json.Number       `json:"systolic"   db:"systolic"  binding:"required"`
	Diastolic json.Number       `json:"diastolic"  db:"diastolic" binding:"required"`
	Heartrate json.Number       `json:"heartrate"  db:"heartrate" binding:"required"`
}

// String is not required by pop and may be deleted
func (e Entry) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Entries is not required by pop and may be deleted
type Entries []Entry

// String is not required by pop and may be deleted
func (e Entries) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

func (e *Entry) IsValid(errors *validate.Errors) {
	if e.Systolic == "" {
		errors.Add("Systolic", "is required")
	}
	if e.Diastolic == "" {
		errors.Add("Diastolic", "is required")
	}
	if e.Heartrate == "" {
		errors.Add("Heartrate", "is required")
	}
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Entry) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Entry) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
