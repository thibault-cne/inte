package models

import (
	"backend/autogen"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	News struct {
		PrivateID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
		autogen.News `bson:",inline"`

		CreatedAt int64 `bson:"created_at" json:"created_at"`
	}
)

// ToJSON converts the model to JSON
func (s *News) ToJSON() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return data
}