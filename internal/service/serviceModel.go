package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Service struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	ServiceName string             `bson:"serviceName" json:"serviceName"`
	Path        string             `bson:"path" json:"path"`
	Hosts       []Host             `bson:"hosts" json:"hosts"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"UpdatedAt"`
}

type Host struct {
	Url      string `bson:"url" json:"url"`
	Port     int    `bson:"port" json:"port"`
	Protocol string `bson:"protocol" json:"protocol"`
}
