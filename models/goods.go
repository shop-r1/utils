package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Goods struct {
	gorm.Model
	No                 string                   `sql:"-" json:"id"`
	TenantId           uint                     `description:"租户ID" json:"-" `
	GoodsInfoId        uint                     `json:"goods_info_id_int" description:"商品基础信息ID"`
	GoodsInfoNo        string                   `sql:"-" json:"goods_info_id" description:"商品基础信息ID"`
	GoodsInfo          GoodsInfo                `json:"goods_info" validate:"-"`
	ShowCategory       ShowCategory             `json:"show_category" validate:"-"`
	ShowCategoryId     uint                     `description:"显示分类ID" json:"show_category_id_int"`
	ShowCategoryNo     string                   `description:"显示分类ID" json:"show_category_id"`
	Alias              string                   `sql:"type:varchar(255)" description:"别名" json:"alias"`
	PriceAud           float32                  `sql:"type:DECIMAL(10, 2)" description:"售价(澳币)" json:"price_aud"`
	PriceRmb           float32                  `sql:"type:DECIMAL(10, 2)" description:"售价(人民币)" json:"price_rmb"`
	CommissionRmb      float32                  `sql:"type:DECIMAL(10, 2)" description:"佣金(人民币)" json:"commission_rmb"`
	BarCode            string                   `sql:"type:varchar(100)" description:"条形码" json:"bar_code"`
	Image              string                   `sql:"type:varchar(255)" description:"图片" json:"image"`
	Album              string                   `sql:"type:text" description:"相册" json:"album"`
	Albums             []string                 `sql:"-" description:"相册(数组)" json:"albums"`
	Video              string                   `sql:"type:varchar(255)" description:"视频" json:"video"`
	Content            string                   `sql:"type:text" description:"详情内容" json:"content"`
	Description        string                   `sql:"type:text" description:"描述" json:"description"`
	QualityPeriod      string                   `sql:"type:varchar(50)" description:"保质期" json:"quality_period"`
	Stage              []byte                   `sql:"type:json" description:"阶段" json:"-"`
	Stages             []int                    `sql:"-" json:"stages"`
	Show               Status                   `sql:"type:integer;default(1)" description:"是否展示" json:"show"`
	Status             Status                   `sql:"type:integer;default(1)" description:"状态 1 上架 2 下架" json:"status"`
	Specifications     []GoodsSpecification     `gorm:"ForeignKey:GoodsId;save_associations:false" description:"规格关联" json:"specifications"`
	Inventory          int                      `description:"库存" json:"inventory"`
	ClickNum           int                      `sql:"type:integer;default(0)" description:"点击数" json:"click_num"`
	BuyNum             int                      `sql:"type:integer;default(0)" description:"购买数" json:"buy_num"`
	SpecificationInfo  []byte                   `sql:"type:json" description:"规格选择参数" json:"-"`
	SpecificationInfoS []SpecificationInfo      `sql:"-" description:"规格选择参数" json:"specification_infos"`
	HasSpecification   bool                     `description:"是否有属性" json:"has_specification"`
	Warehouses         []GoodsShippingWarehouse `gorm:"ForeignKey:GoodsId;save_associations:false" description:"发货仓库关联" json:"warehouses"`
}

type GoodsShippingWarehouse struct {
	gorm.Model
	GoodsId     uint              `json:"goods_id_int"`
	WarehouseId uint              `json:"warehouse_id_int"`
	GoodsNo     string            `sql:"-" json:"goods_id"`
	Warehouse   ShippingWarehouse `json:"warehouse"`
	WarehouseNo string            `sql:"-" json:"warehouse_id"`
	Price       float32           `sql:"type:DECIMAL(10, 2)" description:"售价" json:"price"`
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
	var id int
	if g.GoodsNo != "" && g.ID == 0 {
		id, _ = strconv.Atoi(g.GoodsNo)
		g.GoodsId = uint(id)
		id = 0
	}
	if g.WarehouseNo != "" && g.WarehouseId == 0 {
		id, _ = strconv.Atoi(g.WarehouseNo)
		g.WarehouseId = uint(id)
		id = 0
	}
}

func (g *GoodsShippingWarehouse) transform() {
	g.GoodsNo = strconv.Itoa(int(g.GoodsId))
	g.WarehouseNo = strconv.Itoa(int(g.WarehouseId))
}

type GoodsSpecification struct {
	gorm.Model
	No             string   `sql:"-" json:"id"`
	TenantId       uint     `description:"租户ID" json:"-" `
	GoodsId        uint     `description:"租户商品ID" json:"goods_id_int"`
	GoodsNo        string   `sql:"-" json:"goods_id" validate:"required"`
	BarCode        string   `sql:"type:varchar(100)" description:"条形码" json:"bar_code"`
	Specification  string   `sql:"type:varchar(255)" description:"规格拼接" json:"-"`
	Specifications []string `sql:"-" description:"规格" json:"specifications"`
	Ratio          float32  `sql:"type:DECIMAL(10, 2)" description:"价格浮动比例" json:"ratio"`
	Album          string   `sql:"type:text" description:"相册" json:"album"`
	Inventory      int      `description:"库存" json:"inventory"`
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
	g.GoodsNo = strconv.Itoa(int(g.GoodsId))
	if len(g.Specification) > 0 {
		g.Specifications = strings.Split(g.Specification, ",")
	}
}

type SpecificationInfo struct {
	Level         int                  `json:"level"`
	CategoryId    string               `json:"category_id"`
	ClassId       string               `json:"class_id"`
	Item          string               `json:"item"`
	Specification GoodsSpecification   `json:"specification,omitempty"`
	Children      []*SpecificationInfo `json:"children,omitempty"`
}

type SearchGoods struct {
	List      []Goods `json:"list"`
	Total     int     `json:"total"`
	Page      int     `json:"page"`
	TotalPage int     `json:"total_page"`
	Limit     int     `json:"limit"`
}

func (s *SpecificationInfo) G生成规格记录(tx *gorm.DB, tenantId, goodsId uint, specifications ...string) ([]GoodsSpecification, error) {
	var err error
	gs := make([]GoodsSpecification, 0)
	if len(s.Children) > 0 {
		for _, ss := range s.Children {
			s := append(specifications, ss.Item)
			fmt.Println(s)
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
	if g.HasSpecification {
		g.SpecificationInfo, err = json.Marshal(g.SpecificationInfoS)
		if err != nil {
			return fmt.Errorf("规格参数json序列户错误: %v", err)
		}
	} else {
		g.SpecificationInfoS = nil
	}
	return nil
}

func (g *Goods) AfterFind() error {
	g.transform()
	return nil
}

func (g *Goods) AfterSave(tx *gorm.DB) (err error) {
	g.transform()
	return nil
}

func (g *Goods) AfterCreate(tx *gorm.DB) (err error) {
	err = g.saveLink(tx)
	rb, _ := json.Marshal(g.SpecificationInfoS)
	err = tx.Model(g).Update("specification_info", rb).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return err
}

func (g *Goods) BeforeUpdate(tx *gorm.DB) (err error) {
	err = g.saveLink(tx)
	if err != nil {
		log.Error(err)
		return err
	}
	g.SpecificationInfo, _ = json.Marshal(g.SpecificationInfoS)
	return nil
}

func (g *Goods) saveLink(tx *gorm.DB) (err error) {
	//保存发货仓关联
	err = tx.Where("goods_id = ?", g.ID).Unscoped().Delete(&GoodsShippingWarehouse{}).Error
	if err != nil {
		return err
	}
	for i, warehouse := range g.Warehouses {
		warehouse.GoodsId = g.ID
		err = tx.Create(&warehouse).Error
		if err != nil {
			return err
		}
		g.Warehouses[i] = warehouse
	}
	//保存规格关联
	if err = tx.Where("goods_id = ?", g.ID).Unscoped().Delete(GoodsSpecification{}).Error; err != nil {
		log.Error(err)
		return err
	}
	goodsSpecifications := make([]GoodsSpecification, 0)
	for i, specificationInfo := range g.SpecificationInfoS {
		//整合结构
		gs, err := (&specificationInfo).G生成规格记录(tx, g.TenantId, g.ID, specificationInfo.Item)
		if err != nil {
			log.Error(err)
			return err
		}
		goodsSpecifications = append(goodsSpecifications, gs...)
		g.SpecificationInfoS[i] = specificationInfo
	}

	return nil
}

func (g *Goods) transform() {
	if g.HasSpecification {
		_ = json.Unmarshal(g.SpecificationInfo, &g.SpecificationInfoS)
	}
	g.No = strconv.Itoa(int(g.ID))
}
