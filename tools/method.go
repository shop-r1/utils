package tools

import (
	"errors"
	"github.com/shop-r1/royalpay"
)

var PaymentMethod = map[string]func(key, secret, image, orderId string, body *royalpay.Body) (*royalpay.Result, error){
	"RoyalPayWechat": RoyalPayWechat,
	"RoyalPayAlipay": RoyalPayAliapay,
	"WechatVoucher":  Voucher,
	"Overage":        Overage,
}

func RoyalPayWechat(key, secret, image, orderId string, body *royalpay.Body) (result *royalpay.Result, e error) {
	if key == "" || secret == "" {
		return nil, errors.New("商家对该支付方式未配置")
	}
	pay := royalpay.NewPay(key, secret)
	body.Channel = royalpay.Wechat
	return pay.QrcodeNativeOrder(orderId, body)
}

func RoyalPayAliapay(key, secret, image, orderId string, body *royalpay.Body) (result *royalpay.Result, e error) {
	if key == "" || secret == "" {
		return nil, errors.New("商家对该支付方式未配置")
	}
	pay := royalpay.NewPay(key, secret)
	body.Channel = royalpay.Alipay
	return pay.QrcodeNativeOrder(orderId, body)
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
