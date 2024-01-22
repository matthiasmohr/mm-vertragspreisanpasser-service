package dto

// Pagination represents a request/response structure which contains pagination parameters.
type Pagination struct {
	Limit     int   `query:"limit" json:"limit" validate:"gte=1,required"`
	Offset    int   `query:"offset" json:"offset"`
	TotalRows int64 `json:"totalRows"`
}
