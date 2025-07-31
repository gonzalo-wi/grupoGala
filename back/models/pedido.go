package models

import (
	"time"
)

type Pedido struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Combo     string    `json:"combo"`
	Direccion string    `json:"direccion"`
	Fecha     time.Time `json:"fecha"`
}
