package request

type WlUserSearch struct {
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	UserName   string `json:"userName"`
	Mobile     string `json:"mobile"`
	Status     string `json:"status"`
	Department int    `json:"department"`
}
