package device

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tenant struct {
	TenantID     primitive.ObjectID `json:"tenantId"`
	TenantName   string             `json:"tenantName"`
	Organization string             `json:"organization"`
	Location     string             `json:"location"`
	Issuer       string             `json:"issuer"`
	Endpoint     string             `json:"endpoint"`
	Audience     string             `json:"audience"`
	Properties   []Property         `json:"properties"`
	CreatedDate  time.Time          `json:"createdDate"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (t *Tenant) ToString() string {
	return fmt.Sprintf("Tenant [ID: %s - Name: %s ] ", t.TenantID, t.TenantName)
}
