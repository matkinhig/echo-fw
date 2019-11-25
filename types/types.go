package types

type Student struct {
	ID        int    `bson : "id"`
	Firstname string `bson : "first_name"`
	Lastname  string `bson : "last_name"`
	Age       int    `bson : "age"`
	Classname string `bson : "class_name`
	Email     string `bson : "email"`
}
