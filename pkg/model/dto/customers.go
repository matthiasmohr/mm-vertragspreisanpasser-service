package dto

type ListCustomersRequest struct {
	Pagination
}

type ListCustomersResponse struct {
	Pagination Pagination `json:"pagination"`
	Items      []Customer `json:"items"`
}

// CustomerFindRequest contains the data needed to fullfil a customer case search request.
type CustomerFindRequest struct {
	Pagination Pagination `json:"pagination"`
	FirstName  *string    `query:"firstName" json:"firstName"`
	LastName   *string    `query:"lastName" json:"lastName"`
	Email      *string    `query:"email" json:"email"`
}

func (cfr *CustomerFindRequest) Map() map[string]interface{} {
	m := map[string]interface{}{}

	if cfr.FirstName != nil {
		m["first_name"] = *cfr.FirstName
	}

	if cfr.LastName != nil {
		m["last_name"] = *cfr.LastName
	}

	if cfr.Email != nil {
		m["email"] = *cfr.Email
	}

	return m
}
