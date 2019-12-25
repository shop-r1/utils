package main

import (
	_ "github.com/lib/pq"
	"github.com/shop-r1/utils/models"
	"log"
)

func main() {
	models.InitDb(false, false)
	goodsInfos := make([]models.GoodsInfo, 0)
	models.Db.Model(&models.GoodsInfo{}).Where("goods_type = ?", models.GoodsTypeNormal).Find(&goodsInfos)
	tenants := make([]models.Tenant, 0)
	models.Db.Model(&models.Tenant{}).Where("system = ?", false).Find(&tenants)
	amount := len(tenants)
	var num int
	var count int
	var err error
	//开始检查是否同步
	for _, goods := range goodsInfos {
		num = 0
		models.Db.Model(&models.Goods{}).
			Where("goods_info_id = ?", goods.ID).Count(&num)
		if num >= amount {
			log.Printf("%d-%s同步正常", goods.ID, goods.Name)
			continue
		}
		//存在同步遗漏情况
		for _, tenant := range tenants {
			count = 0
			models.Db.Model(&models.Goods{}).
				Where("goods_info_id = ?", goods.ID).
				Where("tenant_id = ?", tenant.ID).Count(&count)
			if count <= 0 {
				log.Printf("租户%d-%s 商品 %d-%s同步遗漏, 马上进行同步...", tenant.ID, tenant.Name, goods.ID, goods.Name)
				//补漏
				err = models.Db.Create(&models.Goods{
					TenantId:          tenant.No,
					GoodsInfoId:       goods.No,
					Stage:             []byte(`[]`),
					SpecificationInfo: []byte(`[]`),
					Used:              false,
					Metadata:          []byte(`{}`),
					Status:            models.Enable,
					Show:              models.Disable,
					ToppedAt:          goods.CreatedAt,
					Alias:             goods.Name,
					BarCode:           goods.BarCode,
					Unit:              goods.Unit,
					Image:             goods.Image,
					Albums:            goods.Albums,
					CategoryId:        goods.CategoryId,
					ParentCategoryId:  goods.ParentCategoryId,
				}).Error
				if err != nil {
					log.Printf("租户%d-%s 商品 %d-%s同步失败", tenant.ID, tenant.Name, goods.ID, goods.Name)
					log.Println(err)
				} else {
					log.Printf("租户%d-%s 商品 %d-%s同步成功", tenant.ID, tenant.Name, goods.ID, goods.Name)
				}
			}
		}
	}
}
