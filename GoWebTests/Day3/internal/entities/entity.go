package entities

type Product struct {
	Id            uint    `json:"id"`
	Name          string  `json:"name"`
	Color         string  `json:"color"`
	Price         float64 `json:"price"`
	Stock         uint    `json:"stock"`
	Code          string  `json:"code"`
	Published     bool    `json:"published"`
	Creation_date string  `json:"creation_date"`
}

type UpdateProduct struct {
	Name          string  `json:"name"`
	Color         string  `json:"color"`
	Price         float64 `json:"price"`
	Stock         uint    `json:"stock"`
	Code          string  `json:"code"`
	Published     bool    `json:"published"`
	Creation_date string  `json:"creation_date"`
}

type CreateProduct struct {
	Name          string  `json:"name"`
	Color         string  `json:"color"`
	Price         float64 `json:"price"`
	Stock         uint    `json:"stock"`
	Code          string  `json:"code"`
	Published     bool    `json:"published"`
	Creation_date string  `json:"creation_date"`
}
