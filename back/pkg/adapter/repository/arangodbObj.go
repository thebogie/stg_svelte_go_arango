package repository

type ArangoDBContest struct {
	Name        string `json:"name"`
	Start       string `json:"start"`
	Startoffset string `json:"startoffset"`
	Stop        string `json:"stop"`
	Stopoffset  string `json:"stopoffset"`
}

type ArangoDBPlayedAt struct {
	Label string `json:"_label"`
	From  string `json:"_from"`
	To    string `json:"_to"`
}

type ArangoDBVenue struct {
	Address string `json:"address"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
}

type ArangoDBPlayer struct {
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type ArangoDBResultedIn struct {
	Label  string `json:"_label"`
	Place  string `json:"place"`
	Result string `json:"result"`
	From   string `json:"_from"`
	To     string `json:"_to"`
}

type ArangoDBPlayedWith struct {
	Label string `json:"_label"`
	From  string `json:"_from"`
	To    string `json:"_to"`
}
