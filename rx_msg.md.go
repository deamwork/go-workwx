// Code generated by sdkcodegen; DO NOT EDIT.

package workwx

// rxMessageCommon 接收消息的公共部分
type rxMessageCommon struct {
	// ToUserName 企业微信CorpID
	ToUserName string `xml:"ToUserName"`
	// FromUserName 成员UserID
	FromUserName string `xml:"FromUserName"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// MsgType 消息类型
	MsgType string `xml:"MsgType"`
	// MsgID 消息id，64位整型
	MsgID int64 `xml:"MsgId"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID int64 `xml:"AgentID"`
}

// msgTypeText 文本消息
const msgTypeText = "text"

// msgTypeImage 图片消息
const msgTypeImage = "image"

// msgTypeVoice 语音消息
const msgTypeVoice = "voice"

// msgTypeVideo 视频消息
const msgTypeVideo = "video"

// msgTypeLocation 位置消息
const msgTypeLocation = "location"

// msgTypeLink 链接消息
const msgTypeLink = "link"

// rxTextMessageSpecifics 接收的文本消息，特有字段
type rxTextMessageSpecifics struct {
	// Content 文本消息内容
	Content string `xml:"Content"`
}

// rxImageMessageSpecifics 接收的图片消息，特有字段
type rxImageMessageSpecifics struct {
	// PicURL 图片链接
	PicURL string `xml:"PicUrl"`
	// MediaID 图片媒体文件id，可以调用获取媒体文件接口拉取，仅三天内有效
	MediaID string `xml:"MediaId"`
}

// rxVoiceMessageSpecifics 接收的语音消息，特有字段
type rxVoiceMessageSpecifics struct {
	// MediaID 语音媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
	MediaID string `xml:"MediaId"`
	// Format 语音格式，如amr，speex等
	Format string `xml:"Format"`
}

// rxVideoMessageSpecifics 接收的视频消息，特有字段
type rxVideoMessageSpecifics struct {
	// MediaID 视频媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
	MediaID string `xml:"MediaId"`
	// ThumbMediaID 视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据，仅三天内有效
	ThumbMediaID string `xml:"ThumbMediaId"`
}

// rxLocationMessageSpecifics 接收的位置消息，特有字段
type rxLocationMessageSpecifics struct {
	// Lat 地理位置纬度
	Lat float64 `xml:"Location_X"`
	// Lon 地理位置经度
	Lon float64 `xml:"Location_Y"`
	// Scale 地图缩放大小
	Scale int `xml:"Scale"`
	// Label 地理位置信息
	Label string `xml:"Label"`
	// AppType app类型，在企业微信固定返回wxwork，在微信不返回该字段
	AppType string `xml:"AppType"`
}

// rxLinkMessageSpecifics 接收的链接消息，特有字段
type rxLinkMessageSpecifics struct {
	// Title 标题
	Title string `xml:"Title"`
	// Description 描述
	Description string `xml:"Description"`
	// URL 链接跳转的url
	URL string `xml:"Url"`
	// PicURL 封面缩略图的url
	PicURL string `xml:"PicUrl"`
}