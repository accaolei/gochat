package mp

import (
	"encoding/json"

	"github.com/accaolei/gochat/wx"
)

// LogisticsAddOrderReq 物流助手生成运单请求结构体
type LogisticsAddOrderReq struct {
	AddSource    int8               `json:"add_source"`    //订单来源，0为小程序订单，2为App或者H5订单，填2则不发物流服务通知
	OrderID      string             `json:"order_id"`      //订单号id，需要全局唯一,不超过512字节
	Openid       string             `json:"openid"`        //用户openid
	DeliveryID   string             `json:"delivery_id"`   //快递公司ID
	BizID        string             `json:"biz_id"`        //快递客户编码或者现付编码
	CustomRemark string             `json:"custom_remark"` //备注
	Sender       DeliverOrderSender `json:"sender"`        //发件人信息
	Receiver     DeliverOrderSender `json:"receiver"`      //收件人信息
	Shop         DeliverOrderShop   `json:"shop"`          //商品信息，会展示到物流服务通知和电子面单中
	Insured      struct {
		UseInsured   int `json:"use_insured"`   //是否报价，0不保，1保
		InsuredValue int `json:"insured_value"` //保价金额，单位是分
	} `json:"insured"` //是否报价
	Service struct {
	} `json:"service"` //服务类型
	ExpectTime int `json:"expect_time"` //Unix时间戳，单位秒，顺丰必传。0表示已事先约定好取件时间
}

// DeliverOrderSender 发件人数据结构
type DeliverOrderSender struct {
	Name     string `json:"name"`     //联系人姓名
	Tel      string `json:"tel"`      //电话
	Mobile   string `json:"mobile"`   //手机号，电话和手机不能同时为空
	Province string `json:"province"` //省份
	City     string `json:"city"`     //地区/市
	Area     string `json:"area"`     //区/县
	Address  string `json:"address"`  //详细信息
}

// DeliverOrderShop 商家数据结构
type DeliverOrderShop struct {
	WxaPath    string `json:"wxa_path"`    //商家小程序的路径
	ImgURL     string `json:"img_url"`     //产品图片缩略图rul
	GoodsName  string `json:"goods_name"`  // 商品名称，不超过128字节
	GoodsCount string `json:"goods_count"` // 商品数量
}

// DeliverOrderService 物流公司服务类型
type DeliverOrderService struct {
	ServiceType int    `json:"service_type"`
	ServiceName string `json:"service_name"`
}

// LogisticsAddOrderResult 物流助手生成运单返回结构体
type LogisticsAddOrderResult struct {
}

// LogisticsAddOrder 物流助手生成运单
func LogisticsAddOrder(dest *LogisticsAddOrderResult, req *LogisticsAddOrderReq) wx.Action {
	return wx.NewAPI(
		LogisticsAddOrderURL,
		wx.WithMethod(wx.MethodPost),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(req)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, dest)
		}),
	)
}
