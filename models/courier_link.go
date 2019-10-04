package models

import "time"

type CourierLink struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	No          string          `sql:"-" json:"id"`
	LinkId      int             `gorm:"primary_key" json:"link_id"`
	LeftRuleId  int             `gorm:"primary_key" json:"left_rule_id"`
	RightRuleId int             `gorm:"primary_key" json:"right_rule_id"`
	LeftRule    CourierPackRule `gorm:"save_associations:false"`
	RightRule   CourierPackRule `gorm:"save_associations:false"`
}

type ObjectLinkCourier struct {
	Id      int     `json:"id"`
	RuleIds [][]int `json:"rule_ids"`
}
