package customer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

type Repository struct {
	DBConn  *firestore.Client
	context context.Context
}

func (r *Repository) GetById(uuid string) (Model, error) {
	model := Model{}
	doc, err := r.DBConn.Doc("go-order/customers").Get(context.Background())

	if err != nil {
		return model, fmt.Errorf("failed to retrieve customer (%s): %v", uuid, err.Error())
	}
	docJson, err := json.MarshalIndent(doc.Data(), "", "  ")
	if err != nil {
		return model, fmt.Errorf("failed to marshal document to JSON")
	}

	err = json.Unmarshal(docJson, &model)

	if err != nil {
		return model, fmt.Errorf("failed to unmarshal customer JSON to model")
	}
	return model, nil
}

func (r *Repository) GetAll() {
	customers, err := r.DBConn.Doc("Demo/Test").Get(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	json, err := json.MarshalIndent(customers.Data(), "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(json))
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
