package device

import (
	"time"
)

type Tenant struct {
	ID           string
	Name         string
	Organization string
	Issuer       string
	location     string
	CreatedDate  time.Time
	IsDefault    bool
}
