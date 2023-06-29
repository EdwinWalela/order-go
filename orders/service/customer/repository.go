package customer

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Repository struct {
	DBConn  *firestore.Client
	context context.Context
}

func (r *Repository) GetById(uuid string) (Model, error) {
	model := Model{}
	iter := r.DBConn.Collection("customers").Where("Uuid", "==", uuid).Documents(r.context)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return model, err
		}
		if err := doc.DataTo(&model); err != nil {
			return model, fmt.Errorf("failed to map document to model: %v", err.Error())
		}
	}
	if model.Uuid == "" && model.Name == "" && model.Address == "" && model.Contact == "" {
		return model, fmt.Errorf("not found")
	}
	return model, nil
}

func (r *Repository) GetAll() ([]Model, error) {
	iter := r.DBConn.Collection("customers").Documents(r.context)
	customers := []Model{}
	for {
		customer := Model{}
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return customers, err
		}
		if err := doc.DataTo(&customer); err != nil {
			return customers, fmt.Errorf("failed to map document to model: %v", err.Error())
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (r *Repository) Update() {

}

func (r *Repository) Delete() {

}

func (r *Repository) Create(model Model) (string, error) {
	_, _, err := r.DBConn.Collection("customers").Add(r.context, model)

	if err != nil {
		return "", fmt.Errorf("failed to save customer: %v", err.Error())
	}
	return model.Uuid, nil
}
