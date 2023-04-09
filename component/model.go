package component

import (
	"time"
)

type BaseModelUUID struct {
	UUID string `gorm:"primaryKey"`
}

type BaseModelHost struct {
	CreatedBy string `gorm:"column:created_by"`
}

type BaseModelTimestampCreatedAt struct {
	CreatedAt time.Time `gorm:"column:created_at"`
}

type BaseModelTimestampUpdatedAt struct {
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type BaseModelTimestamp struct {
	BaseModelTimestampCreatedAt
	BaseModelTimestampUpdatedAt
}

type BaseModelStatus struct {
	Status int
}

type BaseModel struct {
	BaseModelUUID
	BaseModelHost
	BaseModelTimestamp
}
