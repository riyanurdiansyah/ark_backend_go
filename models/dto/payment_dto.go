package dto

type PaymentMethodDTO struct {
	ID          int    `json:"id"`
	Value       int    `json:"value"`
	Chanel      string `json:"chanel"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Limit       int    `json:"limit"`
	Status      int    `json:"status"`
	Tipe        int    `json:"tipe"`
	Title       string `json:"title"`
	TitleType   string `json:"title_type"`
}
