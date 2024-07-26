package datastore

import (
	"github.com/sarthak0714/certisure-go/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CertificateDatastore interface {
	Create(cert *models.Certificate) error
	GetById(id string) (*models.Certificate, error)
	GetByGroup(groupId string) ([]*models.Certificate, error)
}

type MongoCertificateDatastore struct {
	db *mongo.Database
}
