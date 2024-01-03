// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameBotsPrompt = "bots_prompt"

// BotsPrompt mapped from table <bots_prompt>
type BotsPrompt struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:机器人初始设置ID" json:"id"`        // 机器人初始设置ID
	BotID     int64     `gorm:"column:bot_id;not null;comment:机器人ID 关联 bots.id" json:"bot_id"`              // 机器人ID 关联 bots.id
	Prompt    string    `gorm:"column:prompt;not null;comment:机器人初始设置" json:"prompt"`                       // 机器人初始设置
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
}

// TableName BotsPrompt's table name
func (*BotsPrompt) TableName() string {
	return TableNameBotsPrompt
}
