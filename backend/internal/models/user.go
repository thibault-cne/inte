package models

import (
	"backend/autogen"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		PrivateID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
		autogen.User `bson:",inline"`

		CreatedAt int64 `bson:"created_at" json:"created_at"`
	}
)

func (u *User) IsInteAdmin() bool {
	return u.User.Role == autogen.UserInteAdmin
}

func (u *User) IsAdmin() bool {
	return u.User.Role == autogen.UserAdmin
}

// ToJSON converts the model to JSON
func (u *User) ToJSON() []byte {
	data, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return data
}