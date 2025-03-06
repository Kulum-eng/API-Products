package domain

type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Size        string  `json:"size"`
    Color       string  `json:"color"`
    Category    string  `json:"category"`
    Stock       int     `json:"stock"`
}
