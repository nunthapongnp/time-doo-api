package models

import (
	"time"
)

type Base struct {
	ID          string    `firestore:"-" json:"id"`
	CreatedBy   string    `firestore:"createdBy" json:"-"`
	CreatedDate time.Time `firestore:"createdDate" json:"-"`
	UpdatedBy   string    `firestore:"updatedBy" json:"-"`
	UpdatedDate time.Time `firestore:"updatedDate" json:"-"`
	RowVersion  int       `firestore:"rowVersion" json:"rowVersion"`
}
