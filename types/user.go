package types

type User struct {
	ID        string `bson:"_id" json:"id,omitempty"`
	Firstname string `bson:"firstname" json:"firstname"`
	Lastname  string `bson:"lastname" json:"lastname"`
}
