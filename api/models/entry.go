package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// Entry is used by pop to map your entries database table to your go code.
type Entry struct {
	ID        uuid.UUID `json:"id"         db:"id"`
	CreatedAt int       `json:"created_at" db:"created_at"`
	UpdatedAt int       `json:"updated_at" db:"updated_at"`
	Systolic  int       `json:"systolic"   db:"systolic"  binding:"required"`
	Diastolic int       `json:"diastolic"  db:"diastolic" binding:"required"`
	Heartrate int       `json:"heartrate"  db:"heartrate" binding:"required"`
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

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Entry) Validate(tx *pop.Connection) (*validate.Errors, error) {
	// return validate.NewErrors(), nil
	return validate.Validate(
		&validators.IntIsPresent{Field: e.Systolic, Name: "Systolic"},
		&validators.IntIsPresent{Field: e.Diastolic, Name: "Diastolic"},
		&validators.IntIsPresent{Field: e.Heartrate, Name: "Heartrate"},
	), nil
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
