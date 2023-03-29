package models

type Products struct {
	apiKey string
}

func New(apiKey string) *Products {
	return &Products{apiKey: apiKey}
}
