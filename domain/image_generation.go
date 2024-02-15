package domain

type ImageGeneration struct {
	ID         string
	Url        string
	Censored   bool
	CreateTime int64
}

func (d *ImageGeneration) TableName() string {
	return "image_generation"
}
