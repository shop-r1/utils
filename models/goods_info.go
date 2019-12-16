package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type GoodsInfo struct {
	gorm.Model
	Category         Category   `gorm:"save_associations:false" json:"category" validate:"-"`
	Brand            Brand      `gorm:"save_associations:false" json:"brand" validate:"-"`
	No               string     `sql:"-" json:"id"`
	CategoryId       string     `sql:"type:char(20);index" json:"category_id"`
	ParentCategoryId string     `sql:"type:char(20);index" json:"parent_category_id"`
	BrandId          string     `sql:"type:char(20);index" json:"brand_id"`
	Name             string     `sql:"type:varchar(255)" description:"名称" json:"name" validate:"required"`
	Album            string     `sql:"type:text" description:"相册" json:"album"`
	Albums           []string   `sql:"-" description:"相册(数组)" json:"albums"`
	Description      string     `sql:"type:text" description:"描述" json:"description"`
	Image            string     `sql:"type:varchar(255)" description:"图片" json:"image"`
	Video            string     `sql:"type:varchar(255)" description:"视频" json:"video"`
	Keywords         string     `sql:"type:varchar(255)" description:"关键字" json:"keywords"`
	BarCode          string     `sql:"type:varchar(100)" description:"条形码" json:"bar_code"`
	Content          string     `sql:"type:text" description:"详情内容" json:"content"`
	Weight           int        `sql:"type:integer;default(0)" description:"重量" json:"weight" validate:"required"`
	HasPackRule      bool       `description:"有打包规则" json:"has_pack_rule"`
	PackRule         []byte     `sql:"type:json" description:"关联的物流规则ID" json:"-"`
	PackRules        []PackRule `sql:"-" json:"pack_rules"`
	Unit             string     `sql:"type:varchar(20)" description:"包装单位" json:"unit"`
	GoodsType        GoodsType  `description:"商品类型 0:常规商品 1:组合商品" json:"goods_type"`
}

type SearchGoodsInfo struct {
	List      []GoodsInfo `json:"list"`
	Total     int         `json:"total"`
	Page      int         `json:"page"`
	TotalPage int         `json:"total_page"`
	Limit     int         `json:"limit"`
}

func (g *GoodsInfo) AfterFind() error {
	g.transform()
	return nil
}

func (g *GoodsInfo) BeforeSave() error {
	if len(g.Albums) > 0 && g.Album == "" {
		g.Album = strings.Join(g.Albums, ",")
	}
	if len(g.PackRules) > 0 && len(g.PackRule) == 0 {
		g.PackRule, _ = json.Marshal(g.PackRules)
	} else if len(g.PackRule) == 0 {
		g.PackRule = []byte(`[]`)
	}
	var id int
	if g.No != "" {
		id, _ = strconv.Atoi(g.No)
		g.ID = uint(id)
		id = 0
	}
	return nil
}

func (g *GoodsInfo) AfterSave() (err error) {
	g.transform()
	if g.HasPackRule {
		ruleIds := make([][]int, 0)
		ruleIds, err = getRuleIds(g.PackRules)
		err = add调用物流模块规则关联新增(int(g.ID), ruleIds)
		if err != nil {
			return err
		}
	} else {
		err = add调用物流模块规则关联新增(int(g.ID), nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GoodsInfo) transform() {
	g.No = strconv.Itoa(int(g.ID))
	if g.Album == "" {
		g.Albums = make([]string, 0)
	} else {
		g.Albums = strings.Split(g.Album, ",")
	}
	_ = json.Unmarshal(g.PackRule, &g.PackRules)
	//g.PackRule, _ = json.Marshal(g.PackRules)
}
