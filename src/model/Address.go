package model

import (
	"encoding/json"
	"strconv"
	"time"
)

//Address structure
type Address struct {
	ID           string            `json:"id"`
	description  string            `json:"description"`
	name         string            `json:"name"`
	company      string            `json:"company"`
	phone        string            `json:"phone"`
	email        string            `json:"email"`
	line1        string            `json:"line1"`
	line2        string            `json:"line2"`
	city         string            `json:"city"`
	state        string            `json:"state"`
	zip          string            `json:"zip"`
	country      string            `json:"country"`
	metadata     map[string]string `json:"metadata"`
	dateCreated  time.Time         `json:"dateCreated"`
	dateModified time.Time         `json:"dateModified"`
	deleted      bool              `json:"deleted"`
	object       string            `json:"object"`

	RequestBuilder
}

func (a Address) getID() string {
	return a.ID
}

func (a Address) getDescription() string {
	return a.description
}

func (a Address) getName() string {
	return a.name
}

func (a Address) getCompany() string {
	return a.company
}

func (a Address) getPhone() string {
	return a.phone
}

func (a Address) getEmail() string {
	return a.email
}

func (a Address) getLine1() string {
	return a.line1
}

func (a Address) getLine2() string {
	return a.line2
}

func (a Address) getCity() string {
	return a.city
}

func (a Address) getState() string {
	return a.state
}

func (a Address) getZip() string {
	return a.zip
}

func (a Address) getCountry() string {
	return a.country
}

func (a Address) getMetadata() map[string]string {
	return a.metadata
}

func (a Address) getDateCreated() time.Time {
	return a.dateCreated
}

func (a Address) getDateModified() time.Time {
	return a.dateModified
}

func (a Address) isDeleted() bool {
	return a.deleted
}

func (a Address) getObject() string {
	return a.object
}

func (a Address) String() string {
	m, _ := json.Marshal(a.metadata)
	metadata := string(m)
	return "Address{" +
		"id='" + a.ID + "'" +
		", description='" + a.description + "'" +
		", name='" + a.name + "'" +
		", company='" + a.company + "'" +
		", phone='" + a.phone + "'" +
		", email='" + a.email + "'" +
		", line1='" + a.line1 + "'" +
		", line2='" + a.line2 + "'" +
		", city='" + a.city + "'" +
		", state='" + a.state + "'" +
		", country='" + a.country + "'" +
		", metadata=" + metadata +
		", dateCreated=" + a.dateCreated.String() +
		", dateModified=" + a.dateModified.String() +
		", deleted=" + strconv.FormatBool(a.deleted) +
		", object='" + a.object + "'" +
		"}"
}
