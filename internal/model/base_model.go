package model

import "time"

// Model is the base model, like gorm.Model
type Model struct {
	ID        uint64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deletedAt"`
}

type BaseChain struct {
	Address   string `gorm:"column:address;type:varchar(66);NOT NULL" json:"address"`
	ChainID   uint64 `gorm:"column:chain_id;type:int(11);NOT NULL" json:"chainId"`
	ChainName string `gorm:"column:chain_name;type:varchar(50);NOT NULL" json:"chainName"`
}