// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Record interface {
	IsRecord()
	GetWon() *int
	GetLost() *int
}

type Contest struct {
	Key         string     `json:"_key"`
	ID          string     `json:"_id"`
	Rev         string     `json:"_rev"`
	Start       string     `json:"start"`
	Startoffset string     `json:"startoffset"`
	Stop        string     `json:"stop"`
	Stopoffset  string     `json:"stopoffset"`
	Outcomes    []*Outcome `json:"outcomes"`
	Games       []*Game    `json:"games"`
	Venue       *Venue     `json:"venue,omitempty"`
}

type Game struct {
	Key  string `json:"_key"`
	ID   string `json:"_id"`
	Rev  string `json:"_rev"`
	Name string `json:"name"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginData struct {
	Token    string    `json:"token"`
	Userdata *UserData `json:"userdata"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Outcome struct {
	Key    string    `json:"_key"`
	ID     string    `json:"_id"`
	Rev    string    `json:"_rev"`
	Player *UserData `json:"player"`
	Place  string    `json:"place"`
	Result string    `json:"result"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Stats struct {
	Nemesis      *string    `json:"nemesis,omitempty"`
	WonLostRatio *int       `json:"won_lost_ratio,omitempty"`
	Record       Record     `json:"record,omitempty"`
	Contests     []*Contest `json:"contests,omitempty"`
}

type UserData struct {
	Key       string `json:"_key"`
	ID        string `json:"_id"`
	Rev       string `json:"_rev"`
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Venue struct {
	Key     string `json:"_key"`
	ID      string `json:"_id"`
	Rev     string `json:"_rev"`
	Address string `json:"address"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
}
