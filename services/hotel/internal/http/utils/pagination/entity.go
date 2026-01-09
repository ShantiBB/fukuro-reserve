package pagination

type Query struct {
	Page  uint64
	Limit uint64
}

type Links struct {
	Prev  *string `json:"prev"`
	Next  *string `json:"next"`
	First string  `json:"first"`
	Last  string  `json:"last"`
}
