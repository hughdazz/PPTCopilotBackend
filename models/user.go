package models

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}
