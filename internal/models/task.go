package models

import (
	"time"

	"cloud.google.com/go/firestore"
)

type Task struct {
	ID          string    `firestore:"-" json:"id"`
	Title       string    `firestore:"title" json:"title"`
	Description string    `firestore:"description" json:"description"`
	StartDate   time.Time `firestore:"startDate" json:"startDate"`
	EndDate     time.Time `firestore:"endDate" json:"endDate"`
	Priority    int       `firestore:"priority" json:"priority"`
	Status      int       `firestore:"status" json:"status"`
	Subtasks    []Subtask `firestore:"-" json:"subtasks,omitempty"`
	Base
}

func (t *Task) ToFirestoreUpdate() []firestore.Update {
	return []firestore.Update{
		{Path: "title", Value: t.Title},
		{Path: "description", Value: t.Description},
		{Path: "startDate", Value: t.StartDate},
		{Path: "endDate", Value: t.EndDate},
		{Path: "priority", Value: t.Priority},
		{Path: "status", Value: t.Status},
		{Path: "updatedBy", Value: t.UpdatedBy},
		{Path: "updatedDate", Value: time.Now()},
		{Path: "rowVersion", Value: int(time.Now().Unix())},
	}
}
