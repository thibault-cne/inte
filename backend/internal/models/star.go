package models

import (
	"backend/autogen"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Star struct {
		PrivateID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
		autogen.Star `bson:",inline"`

		CreatedAt int64 `bson:"created_at" json:"created_at"`
	}
)

// ToJSON converts the model to JSON
func (s *Star) ToJSON() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return data
}