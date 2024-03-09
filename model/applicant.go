package model

import (
	"time"
)

type Applicant struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Dob       time.Time `json:"dob"`
	Sex       string    `json:"sex"`
}

type Attributes struct {
	Attributes Applicant `json:"attributes"`
}
type StrapiResp struct {
	Data Attributes
}
