package models

import (
	"backend/autogen"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	PointsTransaction struct {
		PrivateID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
		autogen.PointsTransaction `bson:",inline"`

		CreatedAt int64 `bson:"created_at" json:"created_at"`
	}
)

// ToJSON converts the model to JSON
func (s *PointsTransaction) ToJSON() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return data
}