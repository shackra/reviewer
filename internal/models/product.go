package models

type Product struct {
	ID          string
	Name        string
	Description string
	ImgURL      string
	Reviews     []Review
}

type Review struct {
	ID     string
	Name   string
	Text   string
	Rating float32
}
