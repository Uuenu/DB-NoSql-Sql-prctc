package storage

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Storager interface {
	SaveStudent(firstname, lastname, group, email string, use, yearBirth int, isLocal bool) error
	UpdateStudent(StudentID string, UpdateParams map[string]interface{}) error
	TableStudents() ([]bson.M, error)
	TbStudentsSort(sortParam string) ([]bson.M, error)
	Student(studentID string) bson.M
}

type Student struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
	Group     string `json:"group" bson:"group"`
	Email     string `json:"email" bson:"email"`
	Use       int    `json:"use" bson:"use"`
	YearBirth int    `json:"year_birth" bson:"year_birth"`
	IsLocal   bool   `json:"is_local" bson:"is_local"`
	//PasswordHash string `json:"-" bson:"password"`
}
