package model

//RequestBuilder structure
type RequestBuilder struct {
	params map[string]interface{}
}

func (r RequestBuilder) setDescription(description string) RequestBuilder {
	r.params["description"] = description
	return r
}

func (r RequestBuilder) setName(name string) RequestBuilder {
	r.params["name"] = name
	return r
}

func (r RequestBuilder) setCompany(company string) RequestBuilder {
	r.params["company"] = company
	return r
}

func (r RequestBuilder) setPhone(phone string) RequestBuilder {
	r.params["phone"] = phone
	return r
}

func (r RequestBuilder) setEmail(email string) RequestBuilder {
	r.params["email"] = email
	return r
}

func (r RequestBuilder) setLine1(line1 string) RequestBuilder {
	r.params["line1"] = line1
	return r
}

func (r RequestBuilder) setLine2(line2 string) RequestBuilder {
	r.params["line2"] = line2
	return r
}

func (r RequestBuilder) setCity(zip string) RequestBuilder {
	r.params["zip"] = zip
	return r
}

func (r RequestBuilder) setCountry(country string) RequestBuilder {
	r.params["country"] = country
	return r
}

func (r RequestBuilder) setMetadata(metadata map[string]string) RequestBuilder {
	r.params["metadata"] = metadata
	return r
}

func (r RequestBuilder) build() map[string]interface{} {
	return r.params
}
