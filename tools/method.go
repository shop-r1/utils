package tools

import "github.com/shop-r1/royalpay"

var PaymentMethod = map[string]func(key, secret, image, orderId string, body *royalpay.Body) (*royalpay.Result, error){
	"RoyalPayWechat": RoyalPayWechat,
	"RoyalPayAlipay": RoyalPayAliapay,
	"WechatVoucher":  Voucher,
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
