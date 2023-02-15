package entity

import (
	"gorm"
)

type Customer struct {
	gorm.Model
	Name       string // ต้องไม่เป็นค่าว่าง
	Email      string
	CustomerID string // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
}