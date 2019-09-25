package models

type CourierLink struct {
	No          string `sql:"-" json:"id"`
	LinkId      int    `gorm:"primary_key" json:"link_id_int"`
	LeftRuleId  int    `gorm:"primary_key" json:"left_rule_id"`
	RightRuleId int    `gorm:"primary_key" json:"right_rule_id"`
	Rule        CourierPackRule
}

type ObjectLinkCourier struct {
	Id      int     `json:"id"`
	RuleIds [][]int `json:"rule_ids"`
}
