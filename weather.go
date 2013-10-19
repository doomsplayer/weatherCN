package weather

import (
	"fmt"
	//"regexp"
	"io/ioutil"
	"net/http"
	//"os"
	"encoding/json"
)

type WeatherGetter struct {
	weather_url string
}

func New() *WeatherGetter {
	return &WeatherGetter{`http://m.weather.com.cn/data/`}
}

func (this *WeatherGetter) SetACode(acode string) {
	this.weather_url += (acode + `.html`)
}

func (this *WeatherGetter) GetRaw() (string, error) {
	if this.weather_url[len(this.weather_url)-5:] != `.html` {
		return ``, fmt.Errorf(`you must invoke SetACode first`)
	}
	re, err := http.Get(this.weather_url)
	if err != nil {
		return ``, err
	}
	defer re.Body.Close()
	b, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return ``, err
	}
	return string(b), nil
}

func (this *WeatherGetter) GetInfo() (*WeatherRet, error) {
	sret, err := this.GetRaw()
	if err != nil {
		return nil, err
	}
	ret := &WeatherRet{}
	err = json.Unmarshal([]byte(sret), &struct {
		WeatherInfo *WeatherRet `json:"weatherinfo"`
	}{
		ret})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type WeatherRet struct {
	City         string `json:"city"`
	CityEn       string `json:"city_en"`
	DateY        string `json:"date_y"`
	Date         string `json:"date"`
	Week         string `json:"week"`
	Fchh         string `json:"fchh"`
	CityId       string `json:"cityid"`
	Temperature1 string `json:"temp1"`
	Temperature2 string `json:"temp2"`
	Temperature3 string `json:"temp3"`
	Temperature4 string `json:"temp4"`
	Temperature5 string `json:"temp5"`
	Temperature6 string `json:"temp6"`
	Weather1     string `json:"weather1"`
	Weather2     string `json:"weather2"`
	Weather3     string `json:"weather3"`
	Weather4     string `json:"weather4"`
	Weather5     string `json:"weather5"`
	Weather6     string `json:"weather6"`
	Wind1        string `json:"wind1"`
	Wind2        string `json:"wind2"`
	Wind3        string `json:"wind3"`
	Wind4        string `json:"wind4"`
	Wind5        string `json:"wind5"`
	Wind6        string `json:"wind6"`
	C24          string `json:"index"`
	Wear24       string `json:"index_d"`
	C48          string `json:"index48"`
	Wear48       string `json:"index48_d"`
	UV24         string `json:"index_uv"`
	UV48         string `json:"index48_uv"`
	WashCar      string `json:"index_xc"`
	Travel       string `json:"index_tr"`
	Comfortable  string `json:"index_co"`
	MrngExse     string `json:"index_cl"`
	SClothes     string `json:"index_ls"`
	Allergy      string `json:"index_ag"`
}

//{
//    "weatherinfo": {
//        "city": "成都",
//        "city_en": "chengdu",
//        "date_y": "2013年10月19日",
//        "date": "",
//        "week": "星期六",
//        "fchh": "18",
//        "cityid": "101270101",
//        "temp1": "13℃~18℃",
//        "temp2": "13℃~17℃",
//        "temp3": "13℃~20℃",
//        "temp4": "13℃~22℃",
//        "temp5": "12℃~22℃",
//        "temp6": "13℃~21℃",
//        "tempF1": "55.4℉~64.4℉",
//        "tempF2": "55.4℉~62.6℉",
//        "tempF3": "55.4℉~68℉",
//        "tempF4": "55.4℉~71.6℉",
//        "tempF5": "53.6℉~71.6℉",
//        "tempF6": "55.4℉~69.8℉",
//        "weather1": "阴",
//        "weather2": "阵雨转阴",
//        "weather3": "阵雨转阴",
//        "weather4": "多云",
//        "weather5": "多云",
//        "weather6": "多云转阴",
//        "img1": "2",
//        "img2": "99",
//        "img3": "3",
//        "img4": "2",
//        "img5": "3",
//        "img6": "2",
//        "img7": "1",
//        "img8": "99",
//        "img9": "1",
//        "img10": "99",
//        "img11": "1",
//        "img12": "2",
//        "img_single": "2",
//        "img_title1": "阴",
//        "img_title2": "阴",
//        "img_title3": "阵雨",
//        "img_title4": "阴",
//        "img_title5": "阵雨",
//        "img_title6": "阴",
//        "img_title7": "多云",
//        "img_title8": "多云",
//        "img_title9": "多云",
//        "img_title10": "多云",
//        "img_title11": "多云",
//        "img_title12": "阴",
//        "img_title_single": "阴",
//        "wind1": "微风",
//        "wind2": "微风",
//        "wind3": "北风小于3级",
//        "wind4": "南风小于3级",
//        "wind5": "南风小于3级",
//        "wind6": "南风转北风小于3级",
//        "fx1": "微风",
//        "fx2": "微风",
//        "fl1": "小于3级",
//        "fl2": "小于3级",
//        "fl3": "小于3级",
//        "fl4": "小于3级",
//        "fl5": "小于3级",
//        "fl6": "小于3级",
//        "index": "较舒适",
//        "index_d": "建议着薄外套、开衫牛仔衫裤等服装。年老体弱者应适当添加衣物，宜着夹克衫、薄毛衣等。",
//        "index48": "较舒适",
//        "index48_d": "建议着薄外套、开衫牛仔衫裤等服装。年老体弱者应适当添加衣物，宜着夹克衫、薄毛衣等。",
//        "index_uv": "最弱",
//        "index48_uv": "最弱",
//        "index_xc": "不宜",
//        "index_tr": "适宜",
//        "index_co": "舒适",
//        "st1": "19",
//        "st2": "12",
//        "st3": "18",
//        "st4": "11",
//        "st5": "21",
//        "st6": "11",
//        "index_cl": "较适宜",
//        "index_ls": "不太适宜",
//        "index_ag": "极易发"
//    }
//}
