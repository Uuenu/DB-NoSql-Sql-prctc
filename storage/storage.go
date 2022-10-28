package storage

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Storager interface {
	SaveStudent(firstname, lastname, group, email string, use, yearBirth int, isLocal bool)
	UpdateStudent(StudentID string, UpdateParams map[string]interface{}) error
	TableStudents() ([]bson.M, error)
	TbStudentsSort(sortParam string) ([]bson.M, error)
	Student(studentID string) bson.M
}

type Student struct {
}
