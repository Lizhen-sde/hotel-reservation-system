package types

type User struct {
	ID        string `bson:"_id" json:"id"`
	FirstName string `bson:"firstname" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
