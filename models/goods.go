package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type GoodsType int

const (
	GoodsTypeNormal   GoodsType = 0
	GoodsTypeAssemble GoodsType = 1
)

type Goods struct {
	gorm.Model
	No                   string                   `sql:"-" json:"id"`
	TenantId             string                   `gorm:"primary_key" sql:"type:char(20);index" description:"租户ID" json:"-" `
	CategoryId           string                   `sql:"type:char(20);index" json:"category_id"`
	ParentCategoryId     string                   `sql:"type:char(20);index" json:"parent_category_id"`
	BrandId              string                   `sql:"type:char(20);index" json:"brand_id"`
	Used                 bool                     `description:"领用" json:"used"`
	GoodsInfoId          string                   `sql:"type:char(20);index" json:"goods_info_id" description:"商品基础信息ID"`
	GoodsInfo            GoodsInfo                `gorm:"save_associations:false" json:"goods_info" validate:"-"`
	ShowCategory         ShowCategory             `gorm:"save_associations:false" json:"show_category" validate:"-"`
	ShowCategoryId       string                   `sql:"type:char(20);index" description:"显示分类ID" json:"show_category_id"`
	ParentShowCategory   ShowCategory             `gorm:"save_associations:false" json:"parent_show_category" validate:"-"`
	ParentShowCategoryId string                   `sql:"type:char(20);index" description:"顶级显示分类ID" json:"parent_show_category_id"`
	Alias                string                   `sql:"type:varchar(255)" description:"别名" json:"alias"`
	CommissionRmb        float64                  `sql:"type:DECIMAL(10, 2)" description:"佣金(人民币)" json:"commission_rmb"`
	BarCode              string                   `sql:"type:varchar(100)" description:"条形码" json:"bar_code"`
	Image                string                   `sql:"type:varchar(255)" description:"图片" json:"image"`
	Album                string                   `sql:"type:text" description:"相册" json:"album"`
	Albums               []string                 `sql:"-" description:"相册(数组)" json:"albums"`
	Video                string                   `sql:"type:varchar(255)" description:"视频" json:"video"`
	Content              string                   `sql:"type:text" description:"详情内容" json:"content"`
	Description          string                   `sql:"type:text" description:"描述" json:"description"`
	QualityPeriod        string                   `sql:"type:varchar(50)" description:"保质期" json:"quality_period"`
	Stage                []byte                   `sql:"type:json" description:"阶段" json:"-"`
	Stages               []int                    `sql:"-" json:"stages"`
	Show                 Status                   `sql:"type:integer;default(1)" description:"状态 1 上架 2 下架" json:"show"`
	Status               Status                   `sql:"type:integer;default(1)" description:"状态 1 启用 2 禁用" json:"status"`
	Specifications       []GoodsSpecification     `gorm:"ForeignKey:GoodsId;save_associations:false" description:"规格关联" json:"specifications"`
	Inventory            int                      `description:"库存" json:"inventory"`
	NeedInventory        bool                     `description:"是否需要库存" json:"need_inventory"`
	ClickNum             int                      `sql:"type:integer;default(0)" description:"点击数" json:"click_num"`
	BuyNum               int                      `sql:"type:integer;default(0)" description:"购买数" json:"buy_num"`
	SpecificationInfo    []byte                   `sql:"type:json" description:"规格选择参数" json:"-"`
	SpecificationInfoS   []SpecificationInfo      `sql:"-" description:"规格选择参数" json:"specification_infos"`
	HasSpecification     bool                     `description:"是否有属性" json:"has_specification"`
	Warehouses           []GoodsShippingWarehouse `gorm:"ForeignKey:GoodsId;save_associations:false" description:"发货仓库关联" json:"warehouses" validate:"-"`
	Metadata             []byte                   `description:"附加信息" json:"-"`
	Meta                 interface{}              `sql:"-" description:"附加信息结构" json:"meta"`
	Sort                 int                      `description:"排序" json:"sort"`
	Unit                 string                   `sql:"type:varchar(20)" description:"包装单位" json:"unit"`
	CustomPay            bool                     `description:"是否自定义支付方式" json:"custom_pay"`
	PaymentIds           string                   `sql:"type:text" description:"可用的支付方式" json:"-"`
	PaymentIdsArray      []string                 `sql:"-" json:"payment_ids"`
	ToppedAt             time.Time                `description:"置顶时间"`
	GoodsType            GoodsType                `description:"商品类型 0:常规商品 1:组合商品" json:"goods_type"`
	Assembles            []GoodsAssemble          `gorm:"ForeignKey:GoodsId;save_associations:false" json:"assembles"`
}

type GoodsAssemble struct {
	ID                uint      `gorm:"primary_key" json:"-"`
	GoodsId           int       `sql:"type:integer;index" json:"-" description:"商品ID"`
	GoodsIdNumber     string    `sql:"-" json:"goods_id" description:"商品ID"`
	GoodsInfoId       int       `sql:"type:integer;index" json:"-" description:"商品基础信息ID"`
	GoodsInfoIdNumber string    `sql:"-" json:"goods_info_id" description:"商品基础信息ID"`
	GoodsInfo         GoodsInfo `gorm:"save_associations:false" json:"goods_info" validate:"-"`
	Name              string    `sql:"type:varchar(255)" description:"名称" json:"name"`
	Image             string    `sql:"type:varchar(255)" description:"图片" json:"image"`
	Quantity          int       `json:"quantity" description:"数量"`
	CreatedAt         time.Time
}

type SearchKeyword struct {
	Name string `json:"name"`
}

type ResultKeyword struct {
	Id    string `json:"id"`
	Alias string `json:"name"`
}

type BatchUseGoods struct {
	ShowCategoryId string   `json:"show_category_id"`
	GoodsInfoIds   []string `json:"goods_info_ids"`
}

type CheckInventory struct {
	GoodsId         string `form:"goods_id" json:"goods_id"`
	SpecificationId string `form:"specification_id" json:"specification_id"`
	WarehouseId     string `form:"warehouse_id" json:"warehouse_id"`
	Stage           int    `form:"stage" json:"stage"`
	Quantity        int    `form:"quantity" json:"quantity" validate:"required"`
}

type MemberLevelPrice struct {
	Id    string  `json:"id"`
	Price float64 `json:"price"`
}

type GoodsShippingWarehouse struct {
	Id                   string             `gorm:"type:char(40);primary_key" json:"id"`
	GoodsId              string             `sql:"type:char(20);index" json:"goods_id"`
	WarehouseId          string             `sql:"type:char(20);index" json:"warehouse_id"`
	Warehouse            ShippingWarehouse  `json:"warehouse"`
	MemberLevelPrice     []MemberLevelPrice `sql:"-" description:"会员级别价格" json:"member_level_price"`
	MemberLevelPriceData []byte             `sql:"type:json" json:"-"`
	Price                float64            `sql:"type:DECIMAL(10, 2)" description:"基本售价" json:"price"`
	Init                 bool               `sql:"type:bool;index" description:"默认发货仓" json:"default"`
}

func (g *GoodsShippingWarehouse) AfterSave() error {
	g.transform()
	return nil
}

func (g *GoodsShippingWarehouse) AfterFind() error {
	g.transform()
	return nil
}

func (g *GoodsShippingWarehouse) BeforeSave() error {
	g.unTransform()
	return nil
}

func (g *GoodsShippingWarehouse) unTransform() {
	if len(g.MemberLevelPrice) > 0 {
		g.MemberLevelPriceData, _ = json.Marshal(g.MemberLevelPrice)
	} else {
		g.MemberLevelPriceData = []byte(`[]`)
	}
}

func (g *GoodsShippingWarehouse) transform() {
	_ = json.Unmarshal(g.MemberLevelPriceData, &g.MemberLevelPrice)
}

type GoodsSpecification struct {
	gorm.Model
	No             string   `sql:"-" json:"id"`
	Name           string   `sql:"type:varchar(255)" description:"规格名称" json:"name"`
	TenantId       string   `sql:"type:char(20);index" description:"租户ID" json:"-" `
	GoodsId        string   `sql:"type:char(20);index" description:"租户商品ID" json:"goods_id_int"`
	BarCode        string   `sql:"type:varchar(100)" description:"条形码" json:"bar_code"`
	Specification  string   `sql:"type:varchar(255)" description:"规格拼接" json:"specification"`
	Specifications []string `sql:"-" description:"规格" json:"specifications"`
	Ratio          float64  `sql:"type:DECIMAL(10, 2)" description:"价格浮动比例" json:"ratio"`
	Album          string   `sql:"type:text" description:"相册" json:"album"`
	Inventory      int      `description:"库存" json:"inventory"`
	DefaultSelect  bool     `sql:"column:default_select" description:"默认规格" json:"default"`
}

func (g *GoodsSpecification) BeforeSave() error {
	if len(g.Specifications) > 0 {
		g.Specification = strings.Join(g.Specifications, ",")
	}
	return nil
}

func (g *GoodsSpecification) AfterSave() error {
	g.transform()
	return nil
}

func (g *GoodsSpecification) AfterFind() error {
	g.transform()
	return nil
}

func (g *GoodsSpecification) transform() {
	g.No = strconv.Itoa(int(g.ID))
	if len(g.Specification) > 0 {
		g.Specifications = strings.Split(g.Specification, ",")
	}
}

type SpecificationInfo struct {
	Level         int                 `json:"level"`
	Item          string              `json:"item"`
	Specification GoodsSpecification  `json:"specification,omitempty"`
	Children      []SpecificationInfo `json:"children,omitempty"`
}

type SearchGoods struct {
	List      []Goods `json:"list"`
	Total     int     `json:"total"`
	Page      int     `json:"page"`
	TotalPage int     `json:"total_page"`
	Limit     int     `json:"limit"`
}

func (s *SpecificationInfo) G生成规格记录(tx *gorm.DB, tenantId, goodsId string, specifications ...string) ([]GoodsSpecification, error) {
	var err error
	gs := make([]GoodsSpecification, 0)
	if len(s.Children) > 0 {
		for _, ss := range s.Children {
			s := append(specifications, ss.Item)
			g, err := ss.G生成规格记录(tx, tenantId, goodsId, s...)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			gs = append(gs, g...)
		}
	} else {
		s.Specification.TenantId = tenantId
		s.Specification.GoodsId = goodsId
		s.Specification.Specifications = specifications
		if err = tx.Create(&s.Specification).Error; err != nil {
			log.Error(err)
			return nil, err
		}
	}
	return gs, nil
}

func (g *Goods) BeforeSave() (err error) {
	if len(g.Stage) == 0 {
		g.Stage = []byte(`[]`)
	}
	if len(g.SpecificationInfo) == 0 {
		g.SpecificationInfo = []byte(`[]`)
	}
	if g.Meta != nil {
		g.Metadata, _ = json.Marshal(g.Meta)
	} else {
		g.Metadata = []byte(`{}`)
	}
	return nil
}

func (g *Goods) AfterSave(db *gorm.DB) (err error) {
	if g.GoodsType == GoodsTypeAssemble {
		//组合商品, 保存关联数据
		//1、删除所有关联
		err = db.Where("goods_id = ?", g.ID).Delete(&GoodsAssemble{}).Error
		if err != nil {
			log.Error(err)
			return err
		}
		//2、保存现有关联
		for i, a := range g.Assembles {
			a.GoodsId = int(g.ID)
			a.GoodsInfoId, _ = strconv.Atoi(a.GoodsInfoIdNumber)
			err = db.Create(&a).Error
			if err != nil {
				log.Error(err)
				return err
			}
			g.Assembles[i] = a
		}
	}
	return nil
}

func (g *Goods) AfterFind() error {
	g.transform()
	return nil
}

func (g *Goods) BeforeUpdate(tx *gorm.DB) (err error) {
	err = tx.Where("goods_id = ?", g.ID).Unscoped().Delete(&GoodsShippingWarehouse{}).Error
	if err != nil {
		return err
	}
	if err = tx.Where("goods_id = ?", g.ID).Unscoped().Delete(GoodsSpecification{}).Error; err != nil {
		log.Error(err)
		return err
	}
	g.unTransform()
	return nil
}

func (g *Goods) saveLink(tx *gorm.DB) (err error) {
	//保存发货仓关联
	for i, warehouse := range g.Warehouses {
		warehouse.GoodsId = strconv.Itoa(int(g.ID))
		err = tx.Create(&warehouse).Error
		if err != nil {
			return err
		}
		g.Warehouses[i] = warehouse
	}
	//保存规格关联
	//规格结构整合
	specificationInfos := make([]SpecificationInfo, 0)
	for index, goodsSpecification := range g.Specifications {
		arr := strings.Split(goodsSpecification.Specification, ",")
		specificationInfos = append(specificationInfos, transformSpecification(arr, goodsSpecification))
		goodsSpecification.TenantId = g.TenantId
		goodsSpecification.GoodsId = strconv.Itoa(int(g.ID))
		goodsSpecification.DefaultSelect = index == 0
		if err = tx.Create(&goodsSpecification).Error; err != nil {
			log.Error(err)
			return err
		}
	}
	g.SpecificationInfoS = specificationInfos
	return nil
}

func (g *Goods) transform() {
	if g.PaymentIds != "" {
		g.PaymentIdsArray = strings.Split(g.PaymentIds, ",")
	} else {
		g.PaymentIdsArray = make([]string, 0)
	}
	if g.HasSpecification {
		_ = json.Unmarshal(g.SpecificationInfo, &g.SpecificationInfoS)
	}
	g.No = strconv.Itoa(int(g.ID))
	if len(g.Metadata) > 0 {
		_ = json.Unmarshal(g.Metadata, &g.Meta)
	}
	if len(g.Stage) > 0 {
		_ = json.Unmarshal(g.Stage, &g.Stages)
	}
	if g.Album != "" {
		g.Albums = strings.Split(g.Album, ",")
	}
}

func (g *Goods) unTransform() {
	g.Stage, _ = json.Marshal(g.Stages)
	g.Album = strings.Join(g.Albums, ",")
	g.PaymentIds = strings.Join(g.PaymentIdsArray, ",")
}

func transformSpecification(arr []string, o GoodsSpecification) SpecificationInfo {
	s := SpecificationInfo{}
	for i, a := range arr {
		s.Item = a
		s.Level = i
		if len(arr) > i+1 {
			s.Children = make([]SpecificationInfo, 0)
			s.Children = append(s.Children, transformSpecification(arr[i:], o))
			continue
		}
		s.Specification = o
	}
	return s
}

func (e *GoodsAssemble) AfterFind() error {
	e.GoodsInfoIdNumber = strconv.Itoa(e.GoodsInfoId)
	e.GoodsInfoIdNumber = strconv.Itoa(e.GoodsId)
	return nil
}

func (e *GoodsAssemble) BeforeSave() error {
	e.GoodsInfoId, _ = strconv.Atoi(e.GoodsInfoIdNumber)
	e.GoodsId, _ = strconv.Atoi(e.GoodsIdNumber)
	return nil
}
