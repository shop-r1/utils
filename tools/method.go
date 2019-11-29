package tools

import (
	"errors"
	"github.com/shop-r1/royalpay"
	"github.com/shop-r1/sandpay"
	"github.com/shop-r1/sandpay/pay"
	"github.com/shop-r1/sandpay/pay/params"
	"strconv"
	"time"
)

var PaymentMethod = map[string]func(key, secret, image, orderId string, body *royalpay.Body) (*royalpay.Result, error){
	"RoyalPayWechat": RoyalPayWechat,
	"RoyalPayAlipay": RoyalPayAliapay,
	"WechatVoucher":  Voucher,
	"Overage":        Overage,
	"SandPayAliapay": SandPayAliapay,
}

func RoyalPayWechat(key, secret, image, orderId string, body *royalpay.Body) (result *royalpay.Result, e error) {
	if key == "" || secret == "" {
		return nil, errors.New("商家对该支付方式未配置")
	}
	pay := royalpay.NewPay(key, secret)
	body.Channel = royalpay.Wechat
	result, e = pay.QrcodeOrder(orderId, body)
	if e == nil && len(body.Redirect) > 10 {
		result.CodeUrl = pay.Redirect(result.CodeUrl, body.Redirect)
	}
	return result, e
}

func RoyalPayAliapay(key, secret, image, orderId string, body *royalpay.Body) (result *royalpay.Result, e error) {
	if key == "" || secret == "" {
		return nil, errors.New("商家对该支付方式未配置")
	}
	pay := royalpay.NewPay(key, secret)
	body.Channel = royalpay.Alipay
	return pay.QrcodeNativeOrder(orderId, body)
}

func SandPayAliapay(key, secret, image, orderId string, body *royalpay.Body) (result *royalpay.Result, e error) {
	if key == "" {
		return nil, errors.New("商家对该支付方式未配置")
	}
	sandpay.Client.Config = pay.Config{
		Version:         "1.0",
		MerId:           key,
		PrivatePath:     "certs/server.key", //私钥文件
		CertPath:        "certs/server.crt", //公钥文件
		EncryptCertPath: "certs/sand.cer",   //导出的公钥文件
		ApiHost:         "https://cashier.sandpay.com.cn",
		NotifyUrl:       body.NotifyUrl,
		FrontUrl:        body.Redirect,
	}
	params := params.OrderPayParams{
		OrderNo:     strconv.Itoa(int(time.Now().UnixNano())),
		TotalAmount: body.Price,
		Subject:     body.Description,
		Body:        body.Description + string(body.Price),
		TxnTimeOut:  time.Now().Add(24 * time.Hour).Format("20060102150405"),
		ClientIp:    "127.0.0.1",
		PayMode:     params.PayModWeiXinMp,
		PayExtra: params.PayExtraWeiChat{
			SubAppId: "xxx",
			OpenId:   "xxx",
		},
	}
	e = pay.LoadCertInfo(&sandpay.Client.Config)
	if e != nil {
		return nil, e
	}
	sandPay := &sandpay.SandPay{
		Config: sandpay.Client.Config,
	}
	gotResp, err := sandPay.OrderPay(params)
	if err != nil {
		return nil, err
	}
	result = &royalpay.Result{
		OrderId: gotResp.Body.OrderCode,
		CodeUrl: gotResp.Body.QrCode,
	}
	return result, nil
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
