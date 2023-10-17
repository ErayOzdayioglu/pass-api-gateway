package service

import (
	"context"
	"github.com/ErayOzdayioglu/api-gateway/internal/config/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ServiceRepository interface {
	Save(serviceRequest *AddServiceRequest) (*Service, error)
	FindByName(name string) (*Service, error)
}

type serviceRepositoryI struct {
	db *mongo.Database
}

func (r *serviceRepositoryI) Save(serviceRequest *AddServiceRequest) (*Service, error) {
	coll := database.GetServiceCollection(r.db)
	serviceEntity := &Service{
		ID:          primitive.NewObjectID(),
		ServiceName: serviceRequest.ServiceName,
		Hosts:       serviceRequest.Hosts,
		Path:        serviceRequest.Path,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err := coll.InsertOne(context.Background(), serviceEntity)
	if err != nil {
		return nil, err
	}
	return serviceEntity, nil
}

func (r *serviceRepositoryI) FindByName(name string) (*Service, error) {
	coll := database.GetServiceCollection(r.db)
	var serviceEntity *Service
	filter := bson.D{{"serviceName", name}}
	err := coll.FindOne(context.Background(), filter).Decode(&serviceEntity)
	if err != nil {
		return nil, err
	}
	return serviceEntity, nil
}

func NewServiceRepository(db *mongo.Database) ServiceRepository {
	return &serviceRepositoryI{
		db: db,
	}
}
