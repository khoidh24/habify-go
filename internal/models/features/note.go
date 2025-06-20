package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string             `json:"noteTitle" bson:"noteTitle"`
	Content          string             `json:"content" bson:"content"`
	CoverImage       string             `json:"coverImage" bson:"coverImage"`
	UserID           primitive.ObjectID `json:"userId" bson:"userId"`
	CategoryID       primitive.ObjectID `json:"categoryId" bson:"categoryId"`
	IsPublic         bool               `json:"isPublic" bson:"isPublic"`
	IsActive         bool               `json:"isActive" bson:"isActive"`
}
