package tools

import (
	"errors"
	"github.com/chanxuehong/wechat/mch/core"
	"github.com/chanxuehong/wechat/mch/pay"
	"github.com/shop-r1/royalpay"
	"github.com/shop-r1/sandpay"
	spay "github.com/shop-r1/sandpay/pay"
	"github.com/shop-r1/sandpay/pay/params"
	"strconv"
	"time"
)

var PaymentMethod = map[string]func(key, secret, appId, image, orderId string, body *royalpay.Body, params1 ...string) (*royalpay.Result, error){
	"RoyalPayWechat": RoyalPayWechat,
	"RoyalPayAlipay": RoyalPayAliapay,
	"WechatVoucher":  Voucher,
	"Overage":        Overage,
	"SandPayAliapay": SandPayAliapay,
	"WechatOfficial": WechatOfficial,
}

func WechatOfficial(key, secret, appId, image, orderId string, body *royalpay.Body, params1 ...string) (result *royalpay.Result, e error) {
	client := core.NewClient(appId, key, secret, nil)
	req := &pay.UnifiedOrderRequest{
		Body:           body.Description,
		OutTradeNo:     orderId,
		TotalFee:       int64(body.Price),
		SpbillCreateIP: "127.0.0.1",
		NotifyURL:      body.NotifyUrl,
		TradeType:      "JSAPI",
		OpenId:         body.Operator,
	}
	if len(params1) > 0 {
		req.SpbillCreateIP = params1[0]
	}
	resp, err := pay.UnifiedOrder2(client, req)
	if err != nil {
		return nil, err
	}
	timeStamp := strconv.Itoa(int(time.Now().Unix()))
	nonceStr := strconv.Itoa(int(time.Now().UnixNano()))
	signType := core.SignType_MD5
	pack := "prepay_id=" + resp.PrepayId
	paySign := core.JsapiSign(appId, timeStamp, nonceStr, pack, signType, secret)
	result = &royalpay.Result{}

	result.Params = map[string]interface{}{
		"appId":     appId,
		"timeStamp": timeStamp,
		"nonceStr":  nonceStr,
		"package":   pack,
		"signType":  signType,
		"paySign":   paySign,
	}
	return
}

func RoyalPayWechat(key, secret, appId, image, orderId string, body *royalpay.Body, params1 ...string) (result *royalpay.Result, e error) {
	if key == "" || secret == "" {
		return nil, errors.New("商家对该支付方式未配置")
	}
	pay2 := royalpay.NewPay(key, secret)
	body.Channel = royalpay.Wechat
	result, e = pay2.QrcodeOrder(orderId, body)
	if e == nil && len(body.Redirect) > 10 {
		result.CodeUrl = pay2.Redirect(result.CodeUrl, body.Redirect)
	}
	return result, e
}

func RoyalPayAliapay(key, secret, appId, image, orderId string, body *royalpay.Body, params1 ...string) (result *royalpay.Result, e error) {
	if key == "" || secret == "" {
		return nil, errors.New("商家对该支付方式未配置")
	}
	pay2 := royalpay.NewPay(key, secret)
	body.Channel = royalpay.Alipay
	return pay2.QrcodeNativeOrder(orderId, body)
}

func SandPayAliapay(key, secret, appId, image, orderId string, body *royalpay.Body, params1 ...string) (result *royalpay.Result, e error) {
	if key == "" {
		return nil, errors.New("商家对该支付方式未配置")
	}
	sandpay.Client.Config = spay.Config{
		Version:         "1.0",
		MerId:           key,
		PrivatePath:     "certs/server.key", //私钥文件
		CertPath:        "certs/server.crt", //公钥文件
		EncryptCertPath: "certs/sand.cer",   //导出的公钥文件
		ApiHost:         "https://cashier.sandpay.com.cn",
		NotifyUrl:       body.NotifyUrl,
		FrontUrl:        body.Redirect,
	}
	payParams := params.OrderPayParams{
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
	e = spay.LoadCertInfo(&sandpay.Client.Config)
	if e != nil {
		return nil, e
	}
	sandPay := &sandpay.SandPay{
		Config: sandpay.Client.Config,
	}
	gotResp, err := sandPay.OrderPayQrAlipay(payParams)
	if err != nil {
		return nil, err
	}
	result = &royalpay.Result{
		OrderId: gotResp.Body.OrderCode,
		CodeUrl: gotResp.Body.QrCode,
	}
	return result, nil
}

func Voucher(key, secret, appId, image, orderId string, body *royalpay.Body, params1 ...string) (result *royalpay.Result, e error) {
	result = &royalpay.Result{
		ReturnCode: "SUCCESS",
		OrderId:    orderId,
		QrcodeImg:  image,
		CodeUrl:    image,
	}
	return result, nil
}

func Overage(key, secret, appId, image, orderId string, body *royalpay.Body, params1 ...string) (result *royalpay.Result, e error) {
	result = &royalpay.Result{
		ReturnCode: "SUCCESS",
		OrderId:    orderId,
	}
	return result, nil
}
