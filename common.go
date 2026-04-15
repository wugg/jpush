package jpush

import (
	"fmt"
	"net/http"
)

// ==================== 常量定义 ====================

// API 版本
const (
	APIVersion = "v3"
)

// Zone 配置区域
type Zone string

const (
	ZoneDefault = Zone("DEFAULT") // 默认区域
	ZoneBJ      = Zone("BJ")      // 北京机房
)

// API URL 配置
var URLs = map[Zone]map[string]string{
	ZoneDefault: {
		"PUSH":     "https://api.jpush.cn/v3/push",
		"REPORT":   "https://report.jpush.cn/v3/",
		"DEVICE":   "https://device.jpush.cn/v3/devices/",
		"ALIAS":    "https://device.jpush.cn/v3/aliases/",
		"TAG":      "https://device.jpush.cn/v3/tags/",
		"SCHEDULE": "https://api.jpush.cn/v3/schedules/",
		"ADMIN":    "https://admin.jpush.cn/v1/",
	},
	ZoneBJ: {
		"PUSH":     "https://bjapi.push.jiguang.cn/v3/push",
		"REPORT":   "https://bjapi.push.jiguang.cn/v3/report/",
		"DEVICE":   "https://bjapi.push.jiguang.cn/v3/devices/",
		"ALIAS":    "https://bjapi.push.jiguang.cn/v3/device/aliases/",
		"TAG":      "https://bjapi.push.jiguang.cn/v3/device/tags/",
		"SCHEDULE": "https://bjapi.push.jiguang.cn/v3/push/schedules/",
		"ADMIN":    "https://admin.jpush.cn/v1/",
	},
}

// 平台常量
const (
	PlatformAll      = "all"
	PlatformIos      = "ios"
	PlatformAndroid  = "android"
	PlatformHmos     = "hmos"
	PlatformQuickApp = "quickapp"
	PlatformWinPhone = "winphone"
)

// 受众常量
const (
	AudienceAll            = "all"
	AudienceTag            = "tag"
	AudienceTagAnd         = "tag_and"
	AudienceTagNot         = "tag_not"
	AudienceAlias          = "alias"
	AudienceRegistrationID = "registration_id"
	AudienceSegment        = "segment"
	AudienceAbtest         = "abtest"
)

// 时间单位
const (
	TimeUnitDay   = "day"
	TimeUnitWeek  = "week"
	TimeUnitMonth = "month"
)

// ==================== 异常定义 ====================

// Unauthorized 认证失败异常
type Unauthorized struct {
	Message string
}

func (e *Unauthorized) Error() string {
	return e.Message
}

// NewUnauthorized 创建认证失败异常
func NewUnauthorized(message string) *Unauthorized {
	return &Unauthorized{Message: message}
}

// JPushFailure API 错误响应异常
type JPushFailure struct {
	ErrorCode    int            `json:"error_code"`
	ErrorMessage string         `json:"error_message"`
	Details      string         `json:"details"`
	Response     *http.Response `json:"-"`
	ResponseBody []byte         `json:"-"`
}

func (e *JPushFailure) Error() string {
	return fmt.Sprintf("JPush API Error: Code=%d, Message=%s, Details=%s", e.ErrorCode, e.ErrorMessage, e.Details)
}

// NewJPushFailure 创建 API 错误异常
func NewJPushFailure(errorCode int, errorMessage, details string) *JPushFailure {
	return &JPushFailure{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		Details:      details,
	}
}

// APIConnectionException 连接异常
type APIConnectionException struct {
	Message string
}

func (e *APIConnectionException) Error() string {
	return e.Message
}

// NewAPIConnectionException 创建连接异常
func NewAPIConnectionException(message string) *APIConnectionException {
	return &APIConnectionException{Message: message}
}

// APIRequestException 请求异常
type APIRequestException struct {
	Message string
}

func (e *APIRequestException) Error() string {
	return e.Message
}

// NewAPIRequestException 创建请求异常
func NewAPIRequestException(message string) *APIRequestException {
	return &APIRequestException{Message: message}
}

// ==================== 响应定义 ====================

// Response 基础响应
type Response struct {
	StatusCode int
	Data       []byte
}

// PushResponse 推送响应
type PushResponse struct {
	StatusCode   int    `json:"-"`
	MsgID        string `json:"msg_id,omitempty"`
	SendNo       string `json:"sendno,omitempty"`
	CID          string `json:"cid,omitempty"`
	ErrorCode    int    `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// IsSuccess 检查推送是否成功
func (r *PushResponse) IsSuccess() bool {
	return r.ErrorCode == 0
}

// GetMsgID 获取消息 ID
func (r *PushResponse) GetMsgID() string {
	return r.MsgID
}

// GetCID 获取推送 CID
func (r *PushResponse) GetCID() string {
	return r.CID
}

// ==================== 工具函数 ====================

// GetURL 获取指定 zone 的 URL
func GetURL(key string, zone ...Zone) string {
	z := ZoneDefault
	if len(zone) > 0 {
		z = zone[0]
	}
	if urls, ok := URLs[z]; ok {
		if url, ok := urls[key]; ok {
			return url
		}
	}
	return URLs[ZoneDefault][key]
}
