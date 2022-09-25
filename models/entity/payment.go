package entity

type PaymentMethod struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement"`
	Value       int    `gorm:"column:value"`
	Chanel      string `gorm:"column:chanel"`
	Code        string `gorm:"column:code"`
	Description string `gorm:"column:description"`
	Image       string `gorm:"column:image"`
	Limit       int    `gorm:"column:limit"`
	Status      int    `gorm:"column:status"`
	Tipe        int    `gorm:"column:tipe"`
	Title       string `gorm:"column:title"`
	TitleType   string `gorm:"column:title_type"`
}
