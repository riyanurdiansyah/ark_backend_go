// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type PaymentMethodGql struct {
	ID          int    `json:"id"`
	Value       int    `json:"value"`
	Chanel      string `json:"chanel"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Limit       int    `json:"limit"`
	Status      bool   `json:"status"`
	Tipe        int    `json:"tipe"`
	Title       string `json:"title"`
	TitleType   string `json:"title_type"`
}

type PaymentMethodResponse struct {
	Code   int                 `json:"code"`
	Status bool                `json:"status"`
	Data   []*PaymentMethodGql `json:"data"`
}
