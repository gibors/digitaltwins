package service

import (
	"caidc_auto_devicetwins/config"
	"log"
	"time"

	repo "caidc_auto_devicetwins/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tenant struct {
	TenantID     primitive.ObjectID
	TenantName   string
	Organization string
	Location     string
	Issuer       string
	Audience     string
	Properties   map[interface{}]interface{}
	CreatedDate  time.Time
}

type ServiceConfig struct {
	ConfigParams config.Configuration
	repository   repo.Repository
}

func (sc *ServiceConfig) GetTenantInformation(device string) {
	tenantInfo := sc.repository.GetTenantInformation(device)

	log.Println(tenantInfo)

}
