package tools

import (
	"encoding/json"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	Conf    map[string]interface{}
	AppConf *AppConfigs
)

type AppConfigs struct {
	configs map[string]AppConfig
	mutex   sync.Mutex
}

type AppConfig struct {
	Wx                Wx   `yaml:"wx"`
	GetSelfSkipVerify bool `yaml:"getSelfSkipVerify"`
}

func (e *AppConfigs) GetAppConfig(tenantId string) AppConfig {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	return e.configs[tenantId]
}

func (e *AppConfigs) SetAppConfig(tenantId string, appConfig AppConfig) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.configs[tenantId] = appConfig
}

type Wx struct {
	AppId     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
	MchId     string `yaml:"mchId"`
	ApiKey    string `yaml:"apiKey"`
	LoginHost string `yaml:"loginHost"`
}

func initConf() error {
	return config.Load(
		env.NewSource(),
		file.NewSource(file.WithPath("conf/app.yml")),
	)
}

func InitConf() {
	err := initConf()
	if err != nil {
		log.Fatal(err)
	}
	Conf = config.Map()
	appMap := make(map[string]interface{})
	err = config.Get("app").Scan(&appMap)
	if err != nil {
		log.Fatal(err)
	}
	AppConf = &AppConfigs{
		configs: make(map[string]AppConfig),
	}
	var appConfig AppConfig
	for k := range appMap {
		err = config.Get("app", k).Scan(&appConfig)
		if err != nil {
			log.Fatal(err)
		}
		AppConf.SetAppConfig(k, appConfig)
		appConfig = AppConfig{}
	}
}

//func GetGlobalConfig() (conf interface{}, err error) {
//	err = config.Load(consul.NewSource(consul.WithPrefix("config/tenant")))
//	if err != nil {
//		log.Error(err)
//		return
//	}
//	c := config.Map()
//	c = c["config"].(map[string]interface{})
//	conf = c["tenant"]
//	return
//}
//
//func GetAppConfig(tenantId string) (conf interface{}, err error) {
//	err = config.Load(consul.NewSource(consul.WithPrefix("config/" + tenantId)))
//	if err != nil {
//		log.Error(err)
//		return
//	}
//	c := config.Map()
//	c = c["config"].(map[string]interface{})
//	conf = c[tenantId]
//	return
//}

type ExchangeCurrencyReturn struct {
	ErrorCode int        `json:"error_code"`
	Reason    string     `json:"reason"`
	Result    []Currency `json:"result"`
}

type Currency struct {
	CurrencyF     string      `json:"currencyF"`
	CurrencyFName string      `json:"currencyF_Name"`
	CurrencyT     string      `json:"currencyT"`
	CurrencyTName string      `json:"currencyT_Name"`
	CurrencyFD    json.Number `json:"currencyFD"`
	Exchange      json.Number `json:"exchange"`
	Result        float64     `json:"result,string"`
	UpdateTime    time.Time   `json:"updateTime"`
}

//实时汇率查询换算
func ExchangeCurrency(uri, from, to, key string) {
	//请求地址

	//初始化参数
	param := url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("from", from) //转换汇率前的货币代码
	param.Set("to", to)     //转换汇率成的货币代码
	param.Set("key", key)   //应用APPKEY(应用详细页查询)

	//发送请求
	data, err := Get(uri, param)
	if err != nil {
		log.Error(err)
		return
	}
	var netReturn ExchangeCurrencyReturn
	_ = json.Unmarshal(data, &netReturn)
	if netReturn.ErrorCode != 0 {
		log.Error(netReturn.Reason)
		return
	}
	for _, r := range netReturn.Result {
		Conf[r.CurrencyF+"to"+r.CurrencyT] = r.Result
	}

}

// get 网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
