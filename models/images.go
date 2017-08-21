package models

type Image struct {
	Filename string `json:"filename"`
}

type Images []*Image
