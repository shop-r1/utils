package models

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/shop-r1/utils/tools"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Category struct {
	gorm.Model
	ParentId    string         `json:"parent_id"`
	No          string         `sql:"-" json:"id"`
	Name        string         `sql:"type:varchar(100)" description:"名称" json:"name" validate:"required"`
	Alias       string         `sql:"type:varchar(100)" description:"别名" json:"alias"`
	Description string         `sql:"text" description:"描述" json:"description"`
	Sort        int            `sql:"default(0)" description:"排序" json:"sort"`
	Img         string         `sql:"type:varchar(255)" description:"图片" json:"img"`
	Tag         string         `sql:"type:varchar(255)" description:"商品标签" json:"tag"`
	PackRule    []byte         `sql:"type:json" description:"关联的物流规则ID" json:"-"`
	PackRules   []PackRule     `sql:"-" json:"pack_rules"`
	Children    []ShowCategory `gorm:"ForeignKey:ParentId" json:"children"`
}

type PackRule struct {
	CourierId   string   `json:"courier_id"`
	LeftRuleId  string   `json:"left_rule_id"`
	RightRuleId string   `json:"right_rule_id"`
	ObjectIds   []string `json:"object_ids"`
}

type SearchCategory struct {
	List      []Category `json:"list"`
	Total     int        `json:"total"`
	Page      int        `json:"page"`
	TotalPage int        `json:"total_page"`
	Limit     int        `json:"limit"`
}

func (c *Category) AfterFind() error {
	c.transform()
	return nil
}

func (c *Category) AfterSave() (err error) {
	c.transform()
	ruleIds := make([][]int, 0)
	ruleIds, err = getRuleIds(c.PackRules)
	err = add调用物流模块规则关联新增(int(c.ID), ruleIds)
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) BeforeSave() error {
	if c.No != "" && c.ID == 0 {
		var id int
		id, _ = strconv.Atoi(c.No)
		c.ID = uint(id)
	}
	if len(c.PackRules) > 0 && len(c.PackRule) == 0 {
		c.PackRule, _ = json.Marshal(c.PackRules)
	} else if len(c.PackRule) == 0 {
		c.PackRule = []byte(`[]`)
	}
	return nil
}

func (c *Category) transform() {
	c.No = strconv.Itoa(int(c.ID))
	_ = json.Unmarshal(c.PackRule, &c.PackRules)
}

func add调用物流模块规则关联新增(id int, ruleIds [][]int) (err error) {
	req := map[string]interface{}{
		"id":       id,
		"rule_ids": ruleIds,
	}
	err = tools.Call(context.TODO(), tools.LogisticsService, tools.CourierLinkCreateBatch, req, nil)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func getRuleIds(packRules []PackRule) (ruleIds [][]int, err error) {
	ruleIds = make([][]int, 0)
	var ids []int
	var id int
	for _, rule := range packRules {
		ids = make([]int, 0)
		if rule.LeftRuleId != "" {
			id = 0
			id, err = strconv.Atoi(rule.LeftRuleId)
			if err != nil {
				log.Error(err)
				return ruleIds, errors.New("规则ID必须为数字")
			}
			ids = append(ids, id)
		}

		if len(rule.ObjectIds) > 0 {
			for _, objectId := range rule.ObjectIds {
				id = 0
				id, err = strconv.Atoi(objectId)
				if err != nil {
					log.Error(err)
					return ruleIds, errors.New("可混装对象ID必须为数字")
				}
				ids = append(ids, id)
			}
		}

		if len(ids) > 0 {
			ruleIds = append(ruleIds, ids)
		}
	}
	return ruleIds, nil
}
