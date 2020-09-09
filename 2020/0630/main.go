package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

var jString = `{
	{
		"data": {
			"fields": [
				{
					"aliases": [
						"JrvswW8P"
					],
					"allowEmpty": false,
					"builtIn": false,
					"defaultValue": null,
					"fieldType": "option",
					"name": "\b客户信息",
					"options": [
						{
							"uuid": "6qtazQuA",
							"value": "虎牙直播"
						},
						{
							"uuid": "7B6VqBtc",
							"value": "亿嘉和"
						},
						{
							"uuid": "HqCeGtKW",
							"value": "中国电信"
						},
						{
							"uuid": "NWis7ujE",
							"value": "云象安全"
						},
						{
							"uuid": "CXkmPmmA",
							"value": "棱镜卫士"
						},
						{
							"uuid": "FD9NFacD",
							"value": "一知云"
						},
						{
							"uuid": "e2jQmK1K",
							"value": "四川航空（华云）"
						},
						{
							"uuid": "Ftzd2kee",
							"value": "一面数据"
						},
						{
							"uuid": "TK6WVcwW",
							"value": "讯飞陶云"
						},
						{
							"uuid": "BEi9S2nn",
							"value": "集奥聚合"
						},
						{
							"uuid": "F5TCoorW",
							"value": "猫玩"
						},
						{
							"uuid": "GkxaRG6H",
							"value": "机甲大师"
						},
						{
							"uuid": "Eed8vYr1",
							"value": "水母"
						},
						{
							"uuid": "95sz1gCH",
							"value": "北大英华"
						},
						{
							"uuid": "HpVegmr6",
							"value": "销售演示"
						},
						{
							"uuid": "RNuCQmXS",
							"value": "掌游"
						},
						{
							"uuid": "L21ekRt3",
							"value": "小猪短租"
						},
						{
							"uuid": "Pn23Cacd",
							"value": "虾皮"
						},
						{
							"uuid": "UGmYyPbw",
							"value": "荔枝 FM"
						},
						{
							"uuid": "CGwvDNFv",
							"value": "成都游戏"
						},
						{
							"uuid": "8q3wZXTC",
							"value": "追一"
						},
						{
							"uuid": "PN9FsqMA",
							"value": "B站"
						},
						{
							"uuid": "GtzDhNiU",
							"value": "上海伊邦医药"
						},
						{
							"uuid": "2TMrhHo8",
							"value": "文轩在线"
						},
						{
							"uuid": "VZ7CUbWE",
							"value": "ONES"
						},
						{
							"uuid": "FJWTmPHb",
							"value": "珠海金控"
						},
						{
							"uuid": "7g5nwY28",
							"value": "贵州茅台"
						},
						{
							"uuid": "9WupUgBc",
							"value": "玩蟹科技"
						},
						{
							"uuid": "PDSyEiuc",
							"value": "康佳集团"
						},
						{
							"uuid": "YL3Xo6pv",
							"value": "神州细胞"
						},
						{
							"uuid": "RWGQvVnb",
							"value": "数码星空"
						},
						{
							"uuid": "QekwqgrP",
							"value": "招商基金"
						},
						{
							"uuid": "Mqt3mxoZ",
							"value": "南瑞信通"
						},
						{
							"uuid": "NWohGmV6",
							"value": "杭州场景鹿科技"
						},
						{
							"uuid": "FacYQyJP",
							"value": "广汽财务"
						},
						{
							"uuid": "QSzBX8pt",
							"value": "华发集团"
						},
						{
							"uuid": "Qb1bDrhx",
							"value": "上海易连"
						},
						{
							"uuid": "MgkKj3mK",
							"value": "福利彩票"
						},
						{
							"uuid": "UE2rAniG",
							"value": "华为"
						},
						{
							"uuid": "NGwWUMQ2",
							"value": "中国核电"
						},
						{
							"uuid": "B5DtYQ1a",
							"value": "传智播客"
						},
						{
							"uuid": "H8rscNfn",
							"value": "立讯精密"
						},
						{
							"uuid": "MfoXSBZN",
							"value": "上证信息"
						},
						{
							"uuid": "FKFyMboZ",
							"value": "Desk"
						},
						{
							"uuid": "WMQKmFat",
							"value": "上海荣数"
						},
						{
							"uuid": "TFma9zcB",
							"value": "蔚来汽车"
						},
						{
							"uuid": "41epahjn",
							"value": "地平线"
						},
						{
							"uuid": "YMYVMLDV",
							"value": "优艾智合"
						},
						{
							"uuid": "VRg1YXkj",
							"value": "快手"
						},
						{
							"uuid": "R1ZYesSk",
							"value": "立思辰"
						},
						{
							"uuid": "3PNDFTZS",
							"value": "魔样科技"
						},
						{
							"uuid": "FtGAb4LZ",
							"value": "传音控股"
						},
						{
							"uuid": "9AgcUseq",
							"value": "平安寿险"
						},
						{
							"uuid": "9MDv9dMw",
							"value": "浪潮云"
						},
						{
							"uuid": "9Zd2U6vh",
							"value": "浪潮合作"
						},
						{
							"uuid": "CpDMo27i",
							"value": "ONES运维"
						},
						{
							"uuid": "YSyftgEn",
							"value": "感易智能"
						},
						{
							"uuid": "PdBt4qpf",
							"value": "海康威视"
						},
						{
							"uuid": "EPjZbUjz",
							"value": "联仁集团"
						},
						{
							"uuid": "7p9kBRGh",
							"value": "吉林祥云"
						},
						{
							"uuid": "VtgaSPn5",
							"value": "作业帮"
						},
						{
							"uuid": "TUBEdUmN",
							"value": "上海灿瑞"
						},
						{
							"uuid": "9XFougJn",
							"value": "本原网络"
						},
						{
							"uuid": "8yQKMjGH",
							"value": "滴普科技"
						},
						{
							"uuid": "NbSnCBSQ",
							"value": "盐城一之来"
						},
						{
							"uuid": "5DCz8RnQ",
							"value": "杭州巨玩科技"
						},
						{
							"uuid": "63kfgZoN",
							"value": "无锡正向"
						},
						{
							"uuid": "HrjmEbhm",
							"value": "深圳农村商业银行"
						},
						{
							"uuid": "RjzGZ95D",
							"value": "西安医保中心"
						},
						{
							"uuid": "TVdhqzBf",
							"value": "鸿合科技"
						},
						{
							"uuid": "Q7bXwwip",
							"value": "新华三软件"
						},
						{
							"uuid": "7NS6yECz",
							"value": "唐山启奥"
						},
						{
							"uuid": "XpBasEvj",
							"value": "富士康工业互联网"
						},
						{
							"uuid": "MyfDkR5t",
							"value": "成都久远银海"
						},
						{
							"uuid": "8CGUa72T",
							"value": "医惠科技有限公司"
						},
						{
							"uuid": "VCCYyrhj",
							"value": "北京热云科技有限公司"
						},
						{
							"uuid": "Cw6xkkqa",
							"value": "长安国际信托"
						},
						{
							"uuid": "MKvgTBow",
							"value": "国家超级计算无锡中心"
						},
						{
							"uuid": "WDFLgEeg",
							"value": "浙江华为"
						},
						{
							"uuid": "227dLrdq",
							"value": "燃灯计划"
						},
						{
							"uuid": "PPWNCYSV",
							"value": "中国民航中南管理局空中交通管理局"
						},
						{
							"uuid": "8dGMec2K",
							"value": "浙江省农村信用社联合社"
						},
						{
							"uuid": "Hmz1V4BF",
							"value": "中移（苏州）软件技术有限公司"
						},
						{
							"uuid": "GWtiKpAx",
							"value": "日电（中国）有限公司"
						},
						{
							"uuid": "8Potb6zt",
							"value": "小米"
						},
						{
							"uuid": "YVfqAXvS",
							"value": "北汽集团财务股份有限公司"
						},
						{
							"uuid": "S9Xh5yFu",
							"value": "浪潮软件"
						},
						{
							"uuid": "CCQTsvrm",
							"value": "上海万物工场智能科技有限公司"
						},
						{
							"uuid": "8JYaadSd",
							"value": "深圳市环球易购电子商务有限公司"
						},
						{
							"uuid": "KJs1KXEa",
							"value": "深圳市创梦天地科技有限公司"
						},
						{
							"uuid": "VhgcAEPb",
							"value": "华宇金信（北京）软件有限公司"
						},
						{
							"uuid": "LqPP8K9r",
							"value": "湖南皇爷食品有限公司"
						},
						{
							"uuid": "QpBZLWXm",
							"value": "深圳XX科技有限公司"
						},
						{
							"uuid": "8xr4wMTm",
							"value": "厦门市美亚柏科信息股份有限公司"
						},
						{
							"uuid": "HLKiZ5Ur",
							"value": "北京海纳川汽车零部件股份有限公司"
						},
						{
							"uuid": "TmoybsYb",
							"value": "慧聪网"
						},
						{
							"uuid": "C7jancG4",
							"value": "北京多来点信息技术有限公司"
						},
						{
							"uuid": "AtbyDd1U",
							"value": "晓多科技"
						},
						{
							"uuid": "TgtpAPuf",
							"value": "青牛软件"
						},
						{
							"uuid": "Sf9hkCY5",
							"value": "陕西移动平台研发部"
						},
						{
							"uuid": "YGnG2m2v",
							"value": "人民日报社"
						},
						{
							"uuid": "Vy77cHBR",
							"value": "广州伊的家网络科技有限公司"
						},
						{
							"uuid": "TmR1kHJK",
							"value": "深圳传音控股股份有限公司（上海）"
						},
						{
							"uuid": "T3hFWVE6",
							"value": "南京科远智慧科技集团股份有限公司"
						},
						{
							"uuid": "MDFbVXph",
							"value": "大象慧云信息技术有限公司"
						},
						{
							"uuid": "ANQrSNAY",
							"value": "中国工商银行软件开发中心珠海研发部"
						},
						{
							"uuid": "GTjYVKEJ",
							"value": "北京慧聪互联信息技术有限公司"
						},
						{
							"uuid": "I2zQcziW",
							"value": "青牛软件"
						},
						{
							"uuid": "dy9iDOJ8",
							"value": "中国工商银行软件开发中心珠海研发部"
						},
						{
							"uuid": "zlJtdX1Y",
							"value": "中国工商银行软件开发中心珠海研发部"
						},
						{
							"uuid": "Kyd9TA1s",
							"value": "中国工商银行软件开发中心珠海研发部"
						}
					],
					"required": false,
					"uuid": "JrvswW8P"
				}
			]
		}
	}`

var js = `{
    "data": {
        "tasks": [
            {
                "_U1Zf7epq": null,
                "name": "销售演示专用实例初装申请",
                "uuid": "DgdrtjQHNa9iBvG6"
            }
        ]
    }
}`

type GJson struct {
	Value string
}

func (g *GJson) Set(v string) {
	g.Value = v
}

func main() {
	value := gjson.Get(js, "data.tasks.0._U1Zf7epq").Uint()
	fmt.Println(value)

}

func TypeAccert(v interface{}) {
	g, ok := v.(*string)
	fmt.Println(ok)
	fmt.Println("a", *g)
}
