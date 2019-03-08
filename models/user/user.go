package user

import "gopkg.in/mgo.v2/bson"

// User - model representing user
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
}
