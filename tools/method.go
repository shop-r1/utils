package tools

import (
	"github.com/shop-r1/royalpay"
	"github.com/shop-r1/utils/models"
)

var PaymentMethod = map[models.PaymentMethodType]func(key, secret, image, orderId string, body *royalpay.Body) (*royalpay.Result, error){
	models.MethodRoyalPayWechat: RoyalPayWechat,
	models.MethodRoyalPayAlipay: RoyalPayAliapay,
	models.MethodWechatVoucher:  Voucher,
	models.MethodOverage:        Overage,
}

func RoyalPayWechat(key, secret, image, orderId string, body *royalpay.Body) (result *royalpay.Result, e error) {
	pay := royalpay.NewPay(key, secret)
	body.Channel = royalpay.Wechat
	return pay.QrcodeOrder(orderId, body)
}

func RoyalPayAliapay(key, secret, image, orderId string, body *royalpay.Body) (result *royalpay.Result, e error) {
	pay := royalpay.NewPay(key, secret)
	body.Channel = royalpay.Wechat
	return pay.QrcodeOrder(orderId, body)
}

func Voucher(key, secret, image, orderId string, body *royalpay.Body) (result *royalpay.Result, e error) {
	result = &royalpay.Result{
		ReturnCode: "SUCCESS",
		OrderId:    orderId,
		QrcodeImg:  image,
		CodeUrl:    image,
	}
	return result, nil
}

func Overage(key, secret, image, orderId string, body *royalpay.Body) (result *royalpay.Result, e error) {
	result = &royalpay.Result{
		ReturnCode: "SUCCESS",
		OrderId:    orderId,
	}
	return result, nil
}
