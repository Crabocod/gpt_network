package handlers

type Pagination struct {
	PageIndex      int `json:"pageIndex"`
	RecordsPerPage int `json:"recordsPerPage"`
}

type PaginationResponse struct {
	PageIndex      int `json:"pageIndex"`
	RecordsPerPage int `json:"recordsPerPage"`
	TotalRecords   int `json:"totalRecords"`
}
