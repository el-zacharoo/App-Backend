package store

import (
	"context"
	"fmt"
	"log"

	"github.com/el-zacharoo/go-101/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Org      *mongo.Collection
	Customer *mongo.Collection
}

func (s *Store) Connect() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://zweston:yHoPhQbQBtzjaXGF@customerinfo.pipug.mongodb.net/")
	//ctx := context.Background()
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	s.Customer = client.Database("customer-details").Collection("data")
	fmt.Print("Connected to Mongo Database!\n")

	s.Org = client.Database("org-details").Collection("data")
	fmt.Print("Connected to Mongo Database!\n")
}

func (s *Store) AddPerson(p data.Person) {
	///var insertResult int
	insertResult, err := s.Customer.InsertOne(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

// organisation

func (s *Store) AddOrg(o data.Org) {
	///var insertResult int
	insertResult, err := s.Org.InsertOne(context.Background(), o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}
