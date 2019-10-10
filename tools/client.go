package tools

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

var (
	Client   client.Client
	Json     = client.WithContentType("application/json")
	Protobuf = client.WithContentType("application/protobuf")
)

const (
	//service
	TenantsService   = "go.micro.srv.tenant"
	LogisticsService = "go.micro.srv.logistics"
	GoodsService     = "go.micro.srv.goods"
	MemberService    = "go.micro.srv.member"
	UiControlService = "go.micro.srv.ui-control"
	PaymentService   = "go.micro.srv.payment"

	//method
	TenantCreate           = "Tenant.Create"
	TenantUpdate           = "Tenant.Update"
	TenantDelete           = "Tenant.Delete"
	TenantRead             = "Tenant.Read"
	TenantSearch           = "Tenant.Search"
	TenantReadByDomain     = "Tenant.ReadByDomain"
	Oauth2Token            = "Oauth2.Token"
	Oauth2Introspect       = "Oauth2.Introspect"
	RoleCreate             = "Role.Create"
	RoleUpdate             = "Role.Update"
	RoleDelete             = "Role.Delete"
	RoleRead               = "Role.Read"
	RoleSearch             = "Role.Search"
	UserCreate             = "User.Create"
	UserUpdate             = "User.Update"
	UserDelete             = "User.Delete"
	UserRead               = "User.Read"
	UserSearch             = "User.Search"
	CourierCreate          = "Courier.Create"
	CourierUpdate          = "Courier.Update"
	CourierDelete          = "Courier.Delete"
	CourierSearch          = "Courier.Search"
	CourierRead            = "Courier.Read"
	CourierLinkCreateBatch = "CourierLink.CreateBatch"
	CourierPackRuleCreate  = "CourierPackRule.Create"
	CourierPackRuleUpdate  = "CourierPackRule.Update"
	CourierPackRuleDelete  = "CourierPackRule.Delete"
	CourierPackRuleSearch  = "CourierPackRule.Search"
	CourierPackRuleRead    = "CourierPackRule.Read"
	BrandCreate            = "Brand.Create"
	BrandUpdate            = "Brand.Update"
	BrandDelete            = "Brand.Delete"
	BrandSearch            = "Brand.Search"
	BrandRead              = "Brand.Read"
	CategoryCreate         = "Category.Create"
	CategoryUpdate         = "Category.Update"
	CategoryDelete         = "Category.Delete"
	CategorySearch         = "Category.Search"
	CategoryRead           = "Category.Read"
	ClassCreate            = "Class.Create"
	ClassUpdate            = "Class.Update"
	ClassDelete            = "Class.Delete"
	ClassSearch            = "Class.Search"
	ClassRead              = "Class.Read"
	GoodsInfoCreate        = "GoodsInfo.Create"
	GoodsInfoUpdate        = "GoodsInfo.Update"
	GoodsInfoDelete        = "GoodsInfo.Delete"
	GoodsInfoSearch        = "GoodsInfo.Search"
	GoodsInfoRead          = "GoodsInfo.Read"
	//goods
	GoodsCreate   = "Goods.Create"
	GoodsUpdate   = "Goods.Update"
	GoodsDelete   = "Goods.Delete"
	GoodsSearch   = "Goods.Search"
	GoodsRead     = "Goods.Read"
	GoodsBatchUse = "Goods."
	//show category
	ShowCategoryCreate      = "ShowCategory.Create"
	ShowCategoryUpdate      = "ShowCategory.Update"
	ShowCategoryDelete      = "ShowCategory.Delete"
	ShowCategorySearch      = "ShowCategory.Search"
	ShowCategoryRead        = "ShowCategory.Read"
	ShippingWarehouseCreate = "ShippingWarehouse.Create"
	ShippingWarehouseUpdate = "ShippingWarehouse.Update"
	ShippingWarehouseDelete = "ShippingWarehouse.Delete"
	ShippingWarehouseSearch = "ShippingWarehouse.Search"
	ShippingWarehouseRead   = "ShippingWarehouse.Read"
	//member level
	MemberLevelCreate = "MemberLevel.Create"
	MemberLevelUpdate = "MemberLevel.Update"
	MemberLevelDelete = "MemberLevel.Delete"
	MemberLevelSearch = "MemberLevel.Search"
	MemberLevelRead   = "MemberLevel.Read"
	//member
	MemberCreate = "Member.Create"
	MemberUpdate = "Member.Update"
	MemberDelete = "Member.Delete"
	MemberSearch = "Member.Search"
	MemberRead   = "Member.Read"
	//功能圈
	//ui control
	FunctionCircleCreate = "FunctionCircle.Create"
	FunctionCircleUpdate = "FunctionCircle.Update"
	FunctionCircleDelete = "FunctionCircle.Delete"
	FunctionCircleSearch = "FunctionCircle.Search"
	FunctionCircleRead   = "FunctionCircle.Read"
)

func init() {
	service := micro.NewService()
	service.Init()
	Client = service.Client()
}

//调取服务
func Call(c context.Context, service, method string, req interface{}, rsp interface{}, reqOpts ...client.RequestOption) error {
	if len(reqOpts) == 0 {
		reqOpts = append(reqOpts, Json)
	}
	request := Client.NewRequest(service, method, req, reqOpts...)
	return Client.Call(c, request, rsp)
}
