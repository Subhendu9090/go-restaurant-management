package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id         primitive.ObjectID `bson:"_id"`
	Order_date time.Time          `json:"order_date" validate:"required"`
	Order_id   string             `json:"order_id"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Table_id   *string            `json:"table_id" validate:"required"`
}
