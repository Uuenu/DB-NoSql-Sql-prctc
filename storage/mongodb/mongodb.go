package mongodb

import (
	"context"
	"fmt"
	"log"
	"test-go/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Client *mongo.Client
	DB     *mongo.Database
	Ctx    context.Context
}

const (
	dbURI = "mongodb://localhost:27017"
)

func New() *Storage {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil
	}
	database := client.Database("Universe")
	if err != nil {
		return nil
	}
	s := Storage{
		Client: client,
		DB:     database,
		Ctx:    context.TODO(),
	}

	return &s
}

func (s Storage) SaveStudent(firstname, lastname, group, email string, use, yearBirth int, isLocal bool) { // fix in future
	student := storage.Student{
		Firstname: firstname,
		Lastname:  lastname,
		Group:     group,
		Email:     email,
		Use:       use,
		YearBirth: yearBirth,
		IsLocal:   isLocal,
	}
	s.DB.Collection("Students").InsertOne(s.Ctx, student)
}

func (s Storage) UpdateStudent(studentID string, student storage.Student) {
	s.DB.Collection("Students").UpdateOne(s.Ctx, bson.M{"firstname": studentID}, bson.M{"$set": student})
}

func (s Storage) Student(studentID string) storage.Student {
	var result storage.Student
	s.DB.Collection("Students").FindOne(s.Ctx, bson.M{"firstname": studentID}).Decode(&result)
	return result
}

func (s Storage) TableStudents() ([]storage.Student, error) {
	opts := options.Find()
	opts.SetSort(bson.M{"use": -1})
	sortCursor, err := s.DB.Collection("Students").Find(s.Ctx, bson.D{{}}, opts)
	if err != nil {
		return nil, err
	}
	var studentsTable []storage.Student
	if err = sortCursor.All(s.Ctx, &studentsTable); err != nil {
		return nil, err
	}
	return studentsTable, nil
}

func (s Storage) TbStudentsSort(sortParam string) ([]storage.Student, error) {
	opts := options.Find()
	opts.SetSort(bson.M{sortParam: -1})
	sortCursor, err := s.DB.Collection("Students").Find(s.Ctx, bson.D{{}}, opts)
	if err != nil {
		return nil, err
	}
	var studentsTable []storage.Student
	if err = sortCursor.All(s.Ctx, &studentsTable); err != nil {
		return nil, err
	}
	return studentsTable, nil
}

func (s Storage) StudentsAge(limitAge int) {
	opts := options.Find()
	opts.SetSort(bson.M{"Age": -1})
	//opts.SetBatchSize(1)
	sortCursor, err := s.DB.Collection("Students").Find(s.Ctx, bson.M{"age": bson.M{"$lte": limitAge}}, opts)
	if err != nil {
		log.Printf("%s", err)
	}

	var episodesSorted []bson.M
	if err = sortCursor.All(s.Ctx, &episodesSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesSorted)

	filterCursor, err := s.DB.Collection("Students").Find(s.Ctx, bson.M{"age": 25})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(s.Ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesFiltered)
}
