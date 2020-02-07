package types

//
// Pagination is used when responding with
// a paginated list
//
type Pagination struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

//
// PaginatedResponse is used when responding with
// a paginated list
//
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}
