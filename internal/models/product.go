package models

type Product struct {
	ID            string   `bson:"_id,omitempty"`
	Name          string   `bson:"name"`
	Description   string   `bson:"description"`
	ImgURL        string   `bson:"img_url"`
	Reviews       []Review `bson:"reviews"`
	AverageRating float32  `bson:"average_rating"`
}

type Review struct {
	ID     string  `bson:"_id,omitempty"`
	Name   string  `bson:"name"`
	Text   string  `bson:"text"`
	Rating float32 `bson:"rating"`
}
