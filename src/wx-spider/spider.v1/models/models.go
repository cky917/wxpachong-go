package models

type Article struct {
	Url string
	Name string
}

type File struct {
	UpdateTime  string `json:"updateTime"`
	Result string `json:"result"`
}