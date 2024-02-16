package domain

type ImageGeneration struct {
	ID         string `json:"id"`
	Url        string `json:"url"`
	Censored   bool   `json:"censored"`
	CreateTime int64  `json:"create_time"`
}

func (d *ImageGeneration) TableName() string {
	return "image_generation"
}
