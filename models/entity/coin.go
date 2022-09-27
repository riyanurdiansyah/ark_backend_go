package entity

type Coin struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement"`
	UserID      int    `gorm:"column:user_id"`
	IsOldUser   int    `gorm:"column:is_old_user"`
	IsCompleted int    `gorm:"column:is_completed"`
	CreatedAt   string `gorm:"column:created_at"`
	UpdatedAt   string `gorm:"column:updated_at"`
	Coin        int    `gorm:"column:coin"`
}
