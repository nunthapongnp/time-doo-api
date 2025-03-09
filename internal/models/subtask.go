package models

import (
	"time"

	"cloud.google.com/go/firestore"
)

type Subtask struct {
	ID          string `firestore:"-" json:"id"`
	TaskID      string `firestore:"-" json:"taskId"`
	Title       string `firestore:"title" json:"title"`
	Description string `firestore:"description" json:"description"`
	Status      int    `firestore:"status" json:"status"`
	Base
}

func (t *Subtask) ToFirestoreUpdate() []firestore.Update {
	return []firestore.Update{
		{Path: "title", Value: t.Title},
		{Path: "description", Value: t.Description},
		{Path: "status", Value: t.Status},
		{Path: "updatedBy", Value: t.UpdatedBy},
		{Path: "updatedDate", Value: time.Now()},
		{Path: "rowVersion", Value: int(time.Now().Unix())},
	}
}
