package repository

type ArangoDBContest struct {
	Start       string `json:"start"`
	Startoffset string `json:"startoffset"`
	Stop        string `json:"stop"`
	Stopoffset  string `json:"stopoffset"`
}

type ArangoDBPlayedAt struct {
	Label string `json:"label"`
	From  string `json:"_from"`
	To    string `json:"_to"`
}

type ArangoDBVenue struct {
	Address string `json:"address"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
}
