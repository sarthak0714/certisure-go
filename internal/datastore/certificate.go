package datastore

import (
	"context"
	"time"

	"github.com/sarthak0714/certisure-go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CertificateDatastore interface {
	Create(cert *models.Certificate) error
	Delete(id string) error
	GetById(id string) (*models.Certificate, error)
	GetByGroup(groupId string) ([]*models.Certificate, error)
}

type MongoCertificateDatastore struct {
	collection *mongo.Collection
}

func NewMongoCertificateDatastore(db *mongo.Database) *MongoCertificateDatastore {
	return &MongoCertificateDatastore{collection: db.Collection("certificates")}

}

func (datastore *MongoCertificateDatastore) GetById(id string) (*models.Certificate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var cert models.Certificate
	err := datastore.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&cert)
	if err != nil {
		return nil, err
	}

	return &cert, nil
}

func (datastore *MongoCertificateDatastore) GetByGroup(group_id string) ([]*models.Certificate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := datastore.collection.Find(ctx, bson.M{"group_id": group_id})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	var certs []*models.Certificate

	for cursor.Next(ctx) {
		var cert models.Certificate
		if err = cursor.Decode(&cert); err != nil {
			return nil, err
		}
		certs = append(certs, &cert)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return certs, nil
}

func (datastore *MongoCertificateDatastore) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := datastore.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
