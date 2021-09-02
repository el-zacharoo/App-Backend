package store

import (
	"context"
	"fmt"
	"log"

	"github.com/el-zacharoo/go-101/data"
	"go.mongodb.org/mongo-driver/bson"
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

// Person

func (s *Store) AddPerson(p data.Person) {
	insertResult, err := s.Customer.InsertOne(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) GetPerson(id string) (data.Person, error) {
	var p data.Person
	err := s.Customer.FindOne(
		context.Background(),
		bson.M{"id": id},
	).Decode(&p)
	if err != nil {
		return data.Person{}, err
	}

	return p, nil
}

func (s *Store) UpdatePerson(id string, p data.Person) {
	insertResult, err := s.Customer.ReplaceOne(context.Background(), bson.M{"id": id}, p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) DeleteUser(id string) error {
	removeResult, err := s.Customer.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {

		return err
	}
	fmt.Printf("\nRemoved a Single Document: %v\n", removeResult)
	return nil
}

// Organisation

func (s *Store) AddOrg(o data.Org) {
	insertResult, err := s.Org.InsertOne(context.Background(), o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) GetOrg(id string) (data.Org, error) {
	var o data.Org
	err := s.Org.FindOne(
		context.Background(),
		bson.M{"id": id},
	).Decode(&o)
	if err != nil {
		return data.Org{}, err
	}

	return o, nil
}

func (s *Store) UpdateOrg(id string, o data.Org) {
	insertResult, err := s.Org.ReplaceOne(context.Background(), bson.M{"id": id}, o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) DeleteOrg(id string) error {
	removeResult, err := s.Org.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nRemoved a Single Document: %v\n", removeResult)
	return nil
}
