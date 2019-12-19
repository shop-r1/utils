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
	OrderService     = "go.micro.srv.order"

	//method
	TenantCreate       = "Tenant.Create"
	TenantUpdate       = "Tenant.Update"
	TenantDelete       = "Tenant.Delete"
	TenantRead         = "Tenant.Read"
	TenantSearch       = "Tenant.Search"
	TenantReadById     = "Tenant.ReadById"
	TenantReadByDomain = "Tenant.ReadByDomain"
	Oauth2Token        = "Oauth2.Token"
	Oauth2Introspect   = "Oauth2.Introspect"
	RoleCreate         = "Role.Create"
	RoleUpdate         = "Role.Update"
	RoleDelete         = "Role.Delete"
	RoleRead           = "Role.Read"
	RoleSearch         = "Role.Search"
	UserCreate         = "User.Create"
	UserUpdate         = "User.Update"
	UserDelete         = "User.Delete"
	UserRead           = "User.Read"
	UserSearch         = "User.Search"
	ConfigCreate       = "Config.Create"
	ConfigUpdate       = "Config.Update"
	ConfigDelete       = "Config.Delete"
	ConfigRead         = "Config.Read"
	ConfigSearch       = "Config.Search"
	//courier
	CourierCreate = "Courier.Create"
	CourierUpdate = "Courier.Update"
	CourierDelete = "Courier.Delete"
	CourierSearch = "Courier.Search"
	CourierRead   = "Courier.Read"
	//courier install
	CourierInstallInstall   = "CourierInstall.Install"
	CourierInstallUpdate    = "CourierInstall.Update"
	CourierInstallUninstall = "CourierInstall.Uninstall"
	CourierInstallSearch    = "CourierInstall.Search"
	CourierInstallRead      = "CourierInstall.Read"
	//courier template
	CourierTempalteCreate = "CourierTemplate.Create"
	CourierTemplateUpdate = "CourierTemplate.Update"
	CourierTemplateDelete = "CourierTemplate.Delete"
	CourierTemplateSearch = "CourierTemplate.Search"
	CourierTemplateRead   = "CourierTemplate.Read"
	//payment
	PaymentCreate = "Payment.Create"
	PaymentUpdate = "Payment.Update"
	PaymentDelete = "Payment.Delete"
	PaymentSearch = "Payment.Search"
	PaymentRead   = "Payment.Read"
	//payment install
	PaymentInstallInstall   = "PaymentInstall.Install"
	PaymentInstallUpdate    = "PaymentInstall.Update"
	PaymentInstallUninstall = "PaymentInstall.Uninstall"
	PaymentInstallSearch    = "PaymentInstall.Search"
	PaymentInstallRead      = "PaymentInstall.Read"
	//payment order
	PaymentOrderCreate        = "PaymentOrder.Create"
	PaymentOrderUpdate        = "PaymentOrder.Update"
	PaymentOrderCallback      = "PaymentOrder.Callback"
	PaymentOrderRead          = "PaymentOrder.Read"
	PaymentOrderReadByOrderId = "PaymentOrder.ReadByOrderId"
	//finance
	FinanceCreate                 = "Finance.Create"
	FinanceChangeOverage          = "Finance.ChangeOverage"
	FinanceFreezeOverage          = "Finance.FreezeOverage"
	FinanceDeductionFreezeOverage = "Finance.DeductionFreezeOverage"
	FinanceChangeGold             = "Finance.ChangeGold"
	FinanceFreezeGold             = "Finance.FreezeGold"
	FinanceDeductionFreezeGold    = "Finance.DeductionFreezeGold"
	FinanceRead                   = "Finance.Read"
	FinanceSearch                 = "Finance.Search"
	FinanceSearchLog              = "Finance.SearchLog"

	CourierLinkCreateBatch  = "CourierLink.CreateBatch"
	CourierLinkGeneratePack = "CourierLink.GeneratePack"
	CourierPackRuleCreate   = "CourierPackRule.Create"
	CourierPackRuleUpdate   = "CourierPackRule.Update"
	CourierPackRuleDelete   = "CourierPackRule.Delete"
	CourierPackRuleSearch   = "CourierPackRule.Search"
	CourierPackRuleRead     = "CourierPackRule.Read"
	BrandCreate             = "Brand.Create"
	BrandUpdate             = "Brand.Update"
	BrandDelete             = "Brand.Delete"
	BrandSearch             = "Brand.Search"
	BrandRead               = "Brand.Read"
	CategoryCreate          = "Category.Create"
	CategoryUpdate          = "Category.Update"
	CategoryDelete          = "Category.Delete"
	CategorySearch          = "Category.Search"
	CategoryRead            = "Category.Read"
	ClassCreate             = "Class.Create"
	ClassUpdate             = "Class.Update"
	ClassDelete             = "Class.Delete"
	ClassSearch             = "Class.Search"
	ClassRead               = "Class.Read"
	GoodsInfoCreate         = "GoodsInfo.Create"
	GoodsInfoUpdate         = "GoodsInfo.Update"
	GoodsInfoDelete         = "GoodsInfo.Delete"
	GoodsInfoSearch         = "GoodsInfo.Search"
	GoodsInfoRead           = "GoodsInfo.Read"
	//goods
	GoodsCreate         = "Goods.Create"
	GoodsUpdate         = "Goods.Update"
	GoodsUpdateOther    = "Goods.UpdateOther"
	GoodsCheckInventory = "Goods.CheckInventory"
	GoodsDelete         = "Goods.Delete"
	GoodsSearch         = "Goods.Search"
	GoodsList           = "Goods.List"
	GoodsRead           = "Goods.Read"
	GoodsBatchUse       = "Goods.BatchUse"
	GoodsSearchKeyword  = "Goods.SearchKeyword"
	GoodsSearchApp      = "Goods.SearchApp"
	GoodsReadApp        = "Goods.ReadApp"
	//collection
	CollectionCreate = "Collection.Create"
	CollectionCheck  = "Collection.Check"
	CollectionDelete = "Collection.Delete"
	CollectionSearch = "Collection.Search"
	//shopping cart
	ShoppingCartCreate         = "ShoppingCart.Create"
	ShoppingCartUpdate         = "ShoppingCart.Update"
	ShoppingCartCheck          = "ShoppingCart.Check"
	ShoppingCartCheckWarehouse = "ShoppingCart.CheckWarehouse"
	ShoppingCartDelete         = "ShoppingCart.Delete"
	ShoppingCartDeleteSeleted  = "ShoppingCart.DeleteSelected"
	ShoppingCartList           = "ShoppingCart.List"
	//show category
	ShowCategoryCreate          = "ShowCategory.Create"
	ShowCategoryUpdate          = "ShowCategory.Update"
	ShowCategoryDelete          = "ShowCategory.Delete"
	ShowCategorySearch          = "ShowCategory.Search"
	ShowCategoryRead            = "ShowCategory.Read"
	ShowCategorySearchParentApp = "ShowCategory.SearchParentApp"
	ShowCategoryChildrenIds     = "ShowCategory.ChildrenIds"
	ShippingWarehouseCreate     = "ShippingWarehouse.Create"
	ShippingWarehouseUpdate     = "ShippingWarehouse.Update"
	ShippingWarehouseDelete     = "ShippingWarehouse.Delete"
	ShippingWarehouseSearch     = "ShippingWarehouse.Search"
	ShippingWarehouseRead       = "ShippingWarehouse.Read"
	//member level
	MemberLevelCreate = "MemberLevel.Create"
	MemberLevelUpdate = "MemberLevel.Update"
	MemberLevelPatch  = "MemberLevel.Patch"
	MemberLevelDelete = "MemberLevel.Delete"
	MemberLevelSearch = "MemberLevel.Search"
	MemberLevelRead   = "MemberLevel.Read"
	//member
	MemberCreate       = "Member.Create"
	MemberUpdate       = "Member.Update"
	MemberDelete       = "Member.Delete"
	MemberSearch       = "Member.Search"
	MemberRead         = "Member.Read"
	MemberReadByOpenId = "Member.ReadByOpenId"
	MemberRegistry     = "Member.Registry"
	//sender
	SenderCreate = "Sender.Create"
	SenderUpdate = "Sender.Update"
	SenderDelete = "Sender.Delete"
	SenderSearch = "Sender.Search"
	SenderRead   = "Sender.Read"
	//consignee
	ConsigneeCreate      = "Consignee.Create"
	ConsigneeUpdate      = "Consignee.Update"
	ConsigneeDelete      = "Consignee.Delete"
	ConsigneeSearch      = "Consignee.Search"
	ConsigneeRead        = "Consignee.Read"
	ConsigneeReadDefault = "Consignee.ReadDefault"
	//功能圈
	//ui control
	FunctionCircleCreate = "FunctionCircle.Create"
	FunctionCircleUpdate = "FunctionCircle.Update"
	FunctionCircleDelete = "FunctionCircle.Delete"
	FunctionCircleSearch = "FunctionCircle.Search"
	FunctionCircleRead   = "FunctionCircle.Read"
	//订单
	OrderAppCreate        = "OrderApp.Create"
	OrderAppUpdateStatus  = "OrderApp.UpdateStatus"
	OrderAppDelete        = "OrderApp.Delete"
	OrderAppSearch        = "OrderApp.Search"
	OrderAppRead          = "OrderApp.Read"
	OrderAppReadStatus    = "OrderApp.ReadStatus"
	OrderAppUpdateSuccess = "OrderApp.UpdateSuccess"
	//订单包裹
	OrderUnitPackRead   = "OrderUnitPack.Read"
	OrderUnitPackUpdate = "OrderUnitPack.Update"
	//后台订单
	OrderCreate     = "Order.Create"
	OrderUpdate     = "Order.Update"
	OrderDelete     = "Order.Delete"
	OrderSearch     = "Order.Search"
	OrderRead       = "Order.Read"
	OrderUpdatePack = "Order.UpdatePack"
	OrderBatchShip  = "Order.BatchShip"
	OrderCountGroup = "Order.CountGroup"
	//活动
	ActivityCreate     = "Activity.Create"
	ActivityUpdate     = "Activity.Update"
	ActivityDelete     = "Activity.Delete"
	ActivityRead       = "Activity.Read"
	ActivitySearch     = "Activity.Search"
	ActivityEnableLink = "Activity.EnableLink"
	ActivityGoodsLink  = "Activity.GoodsLink"
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
