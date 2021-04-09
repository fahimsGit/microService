package repository

import (
	"context"
	"log"
	"time"

	//courseproto "github.com/fahimsGit/basic-microservice/proto/course"
	pb "github.com/fahimsGit/basic-microservice/proto/student"
	"github.com/fahimsGit/basic-microservice/student/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type studentRepository struct {
	client *mongo.Client
	config config.DBConfig
}

// Create ...
func (repoService studentRepository) CreateStudent(student *pb.Student) (*pb.Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Minute)
	collection := repoService.client.Database(repoService.config.Dbname).Collection(repoService.config.Students)
	_, err := collection.InsertOne(ctx, student)
	if err != nil {
		return &pb.Student{}, err
	}
	return student, nil
}

func (repoService studentRepository) GetAllStudent() ([]*pb.Student, error) {
	var students []*pb.Student
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Minute)
	collection := repoService.client.Database(repoService.config.Dbname).Collection(repoService.config.Students)
	resultCursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return students, err
	}
	for resultCursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem pb.Student
		err := resultCursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		students = append(students, &elem)
	}
	return students, nil
}
func (repoService studentRepository) CreateCourseEnrollment(req *pb.RequestCreateCourseEnrollment, courseName string) (*pb.ResponseCreateCourseEnrollment, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	collection := repoService.client.Database(repoService.config.Dbname).Collection(repoService.config.Enrolment)
	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return &pb.ResponseCreateCourseEnrollment{}, err
	}
	return &pb.ResponseCreateCourseEnrollment{
		Id:         req.StudentId,
		Name:       courseName,
		CourseId:   req.StudentId,
		CourseName: courseName,
		Status: &pb.Status{
			Success: true,
			Error:   "",
		},
	}, nil

}

// NewMongoRepository returns a new todo repo for mongo database
func NewMongoRepository(dbClient *mongo.Client) StudentRepository {
	return &studentRepository{
		client: dbClient,
		config: config.New(),
	}
}
