package jpush

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ==================== 推送载荷 ====================

// PushPayload 推送载荷
type PushPayload struct {
	client       *Client
	Platform     interface{} `json:"platform,omitempty"`
	Audience     interface{} `json:"audience,omitempty"`
	Notification interface{} `json:"notification,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	SmsMessage   interface{} `json:"sms_message,omitempty"`
	Options      interface{} `json:"options,omitempty"`
	InappMessage interface{} `json:"inapp_message,omitempty"`
	CID          string      `json:"cid,omitempty"`
}

// ==================== 平台设置 ====================

// SetPlatform 设置推送平台
func (p *PushPayload) SetPlatform(platform interface{}) *PushPayload {
	p.Platform = platform
	return p
}

// SetPlatformAll 设置推送平台为全部
func (p *PushPayload) SetPlatformAll() *PushPayload {
	p.Platform = PlatformAll
	return p
}

// SetPlatforms 设置多个推送平台
func (p *PushPayload) SetPlatforms(platforms ...string) *PushPayload {
	if len(platforms) == 1 && platforms[0] == PlatformAll {
		p.Platform = PlatformAll
	} else {
		p.Platform = platforms
	}
	return p
}

// ==================== 受众设置 ====================

// SetAudience 设置推送受众
func (p *PushPayload) SetAudience(audience interface{}) *PushPayload {
	p.Audience = audience
	return p
}

// SetAudienceAll 设置受众为全部
func (p *PushPayload) SetAudienceAll() *PushPayload {
	p.Audience = PlatformAll
	return p
}

// SetAudienceWithMap 设置受众（使用 Map）
func (p *PushPayload) SetAudienceWithMap(audience map[string][]string) *PushPayload {
	p.Audience = audience
	return p
}

// AddTag 添加标签
func (p *PushPayload) AddTag(tags ...string) *PushPayload {
	if p.Audience == nil {
		p.Audience = make(map[string][]string)
	}
	if m, ok := p.Audience.(map[string][]string); ok {
		m[AudienceTag] = append(m[AudienceTag], tags...)
	}
	return p
}

// AddTagAnd 添加标签与
func (p *PushPayload) AddTagAnd(tags ...string) *PushPayload {
	if p.Audience == nil {
		p.Audience = make(map[string][]string)
	}
	if m, ok := p.Audience.(map[string][]string); ok {
		m[AudienceTagAnd] = append(m[AudienceTagAnd], tags...)
	}
	return p
}

// AddTagNot 添加标签非
func (p *PushPayload) AddTagNot(tags ...string) *PushPayload {
	if p.Audience == nil {
		p.Audience = make(map[string][]string)
	}
	if m, ok := p.Audience.(map[string][]string); ok {
		m[AudienceTagNot] = append(m[AudienceTagNot], tags...)
	}
	return p
}

// AddAlias 添加别名
func (p *PushPayload) AddAlias(aliases ...string) *PushPayload {
	if p.Audience == nil {
		p.Audience = make(map[string][]string)
	}
	if m, ok := p.Audience.(map[string][]string); ok {
		m[AudienceAlias] = append(m[AudienceAlias], aliases...)
	}
	return p
}

// AddRegistrationID 添加设备 ID
func (p *PushPayload) AddRegistrationID(regIDs ...string) *PushPayload {
	if p.Audience == nil {
		p.Audience = make(map[string][]string)
	}
	if m, ok := p.Audience.(map[string][]string); ok {
		m[AudienceRegistrationID] = append(m[AudienceRegistrationID], regIDs...)
	}
	return p
}

// AddSegment 添加用户分群
func (p *PushPayload) AddSegment(segments ...string) *PushPayload {
	if p.Audience == nil {
		p.Audience = make(map[string][]string)
	}
	if m, ok := p.Audience.(map[string][]string); ok {
		m[AudienceSegment] = append(m[AudienceSegment], segments...)
	}
	return p
}

// AddAbtest 添加 A/B 测试
func (p *PushPayload) AddAbtest(abtests ...string) *PushPayload {
	if p.Audience == nil {
		p.Audience = make(map[string][]string)
	}
	if m, ok := p.Audience.(map[string][]string); ok {
		m[AudienceAbtest] = append(m[AudienceAbtest], abtests...)
	}
	return p
}

// ==================== 通知设置 ====================

// SetNotification 设置通知
func (p *PushPayload) SetNotification(notification interface{}) *PushPayload {
	p.Notification = notification
	return p
}

// SetNotificationAlert 设置通知内容
func (p *PushPayload) SetNotificationAlert(alert string) *PushPayload {
	p.Notification = map[string]interface{}{
		"alert": alert,
	}
	return p
}

// SetAndroidNotification 设置 Android 通知
func (p *PushPayload) SetAndroidNotification(notification *AndroidNotification) *PushPayload {
	if p.Notification == nil {
		p.Notification = make(map[string]interface{})
	}
	if n, ok := p.Notification.(map[string]interface{}); ok {
		n["android"] = notification.ToMap()
	}
	return p
}

// SetIosNotification 设置 iOS 通知
func (p *PushPayload) SetIosNotification(notification *IOSNotification) *PushPayload {
	if p.Notification == nil {
		p.Notification = make(map[string]interface{})
	}
	if n, ok := p.Notification.(map[string]interface{}); ok {
		n["ios"] = notification.ToMap()
	}
	return p
}

// SetHmosNotification 设置鸿蒙通知
func (p *PushPayload) SetHmosNotification(notification *HmosNotification) *PushPayload {
	if p.Notification == nil {
		p.Notification = make(map[string]interface{})
	}
	if n, ok := p.Notification.(map[string]interface{}); ok {
		n["hmos"] = notification.ToMap()
	}
	return p
}

// SetQuickAppNotification 设置快应用通知
func (p *PushPayload) SetQuickAppNotification(notification *QuickAppNotification) *PushPayload {
	if p.Notification == nil {
		p.Notification = make(map[string]interface{})
	}
	if n, ok := p.Notification.(map[string]interface{}); ok {
		n["quickapp"] = notification.ToMap()
	}
	return p
}

// ==================== 自定义消息设置 ====================

// SetMessage 设置自定义消息
func (p *PushPayload) SetMessage(message *Message) *PushPayload {
	p.Message = message.ToMap()
	return p
}

// ==================== 短信消息设置 ====================

// SetSmsMessage 设置短信消息
func (p *PushPayload) SetSmsMessage(message *SmsMessage) *PushPayload {
	p.SmsMessage = message.ToMap()
	return p
}

// ==================== 应用内增强提醒 ====================

// SetInappMessage 设置应用内增强提醒
func (p *PushPayload) SetInappMessage(enabled bool) *PushPayload {
	p.InappMessage = map[string]bool{
		"inapp_message": enabled,
	}
	return p
}

// ==================== 推送选项设置 ====================

// SetOptions 设置推送选项
func (p *PushPayload) SetOptions(options *Options) *PushPayload {
	p.Options = options.ToMap()
	return p
}

// SetApnsProduction 设置 APNs 生产环境
func (p *PushPayload) SetApnsProduction(isProd bool) *PushPayload {
	if p.Options == nil {
		p.Options = &Options{}
	}
	if opt, ok := p.Options.(*Options); ok {
		opt.ApnsProduction = isProd
	}
	return p
}

// SetTimeToLive 设置离线消息保留时间
func (p *PushPayload) SetTimeToLive(ttl int) *PushPayload {
	if p.Options == nil {
		p.Options = &Options{}
	}
	if opt, ok := p.Options.(*Options); ok {
		opt.TimeToLive = ttl
	}
	return p
}

// SetSendNo 设置发送序号
func (p *PushPayload) SetSendNo(sendno int) *PushPayload {
	if p.Options == nil {
		p.Options = &Options{}
	}
	if opt, ok := p.Options.(*Options); ok {
		opt.SendNo = sendno
	}
	return p
}

// SetOverrideMsgID 设置覆盖消息 ID
func (p *PushPayload) SetOverrideMsgID(msgID string) *PushPayload {
	if p.Options == nil {
		p.Options = &Options{}
	}
	if opt, ok := p.Options.(*Options); ok {
		opt.OverrideMsgID = msgID
	}
	return p
}

// SetBigPushDuration 设置定速推送时长
func (p *PushPayload) SetBigPushDuration(duration int) *PushPayload {
	if p.Options == nil {
		p.Options = &Options{}
	}
	if opt, ok := p.Options.(*Options); ok {
		opt.BigPushDuration = duration
	}
	return p
}

// SetApnsCollapseID 设置 APNs 折叠 ID
func (p *PushPayload) SetApnsCollapseID(id string) *PushPayload {
	if p.Options == nil {
		p.Options = &Options{}
	}
	if opt, ok := p.Options.(*Options); ok {
		opt.ApnsCollapseID = id
	}
	return p
}

// SetTestMessage 设置测试消息标识
func (p *PushPayload) SetTestMessage(test bool) *PushPayload {
	if p.Options == nil {
		p.Options = &Options{}
	}
	if opt, ok := p.Options.(*Options); ok {
		opt.TestMessage = test
	}
	return p
}

// SetCID 设置 CID
func (p *PushPayload) SetCID(cid string) *PushPayload {
	p.CID = cid
	return p
}

// ==================== 发送方法 ====================

// Send 发送推送
func (p *PushPayload) Send() (*PushResponse, error) {
	// 验证必填字段
	if p.Platform == nil {
		return nil, fmt.Errorf("推送平台不能为空")
	}
	if p.Audience == nil {
		return nil, fmt.Errorf("推送受众不能为空")
	}
	if p.Notification == nil && p.Message == nil {
		return nil, fmt.Errorf("通知和消息不能同时为空")
	}

	data, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("PUSH", p.client.Zone)
	respBody, statusCode, err := p.client.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result PushResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// SendValidate 验证推送（不实际发送）
func (p *PushPayload) SendValidate() (*PushResponse, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("PUSH", p.client.Zone) + "/validate"
	respBody, statusCode, err := p.client.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result PushResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// ==================== 通知对象定义 ====================

// AndroidNotification Android 通知
type AndroidNotification struct {
	Alert        string                 `json:"alert,omitempty"`
	Title        string                 `json:"title,omitempty"`
	BuilderID    int                    `json:"builder_id,omitempty"`
	ChannelID    string                 `json:"channel_id,omitempty"`
	Category     string                 `json:"category,omitempty"`
	Priority     int                    `json:"priority,omitempty"`
	Style        int                    `json:"style,omitempty"`
	AlertType    int                    `json:"alert_type,omitempty"`
	BigText      string                 `json:"big_text,omitempty"`
	Inbox        map[string]interface{} `json:"inbox,omitempty"`
	BigPicPath   string                 `json:"big_pic_path,omitempty"`
	LargeIcon    string                 `json:"large_icon,omitempty"`
	SmallIconURI string                 `json:"small_icon_uri,omitempty"`
	IconBgColor  string                 `json:"icon_bg_color,omitempty"`
	Intent       *Intent                `json:"intent,omitempty"`
	Extras       map[string]interface{} `json:"extras,omitempty"`
}

// ToMap 转换为 Map
func (n *AndroidNotification) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if n.Alert != "" {
		m["alert"] = n.Alert
	}
	if n.Title != "" {
		m["title"] = n.Title
	}
	if n.BuilderID != 0 {
		m["builder_id"] = n.BuilderID
	}
	if n.ChannelID != "" {
		m["channel_id"] = n.ChannelID
	}
	if n.Category != "" {
		m["category"] = n.Category
	}
	if n.Priority != 0 {
		m["priority"] = n.Priority
	}
	if n.Style != 0 {
		m["style"] = n.Style
	}
	if n.AlertType != 0 {
		m["alert_type"] = n.AlertType
	}
	if n.BigText != "" {
		m["big_text"] = n.BigText
	}
	if n.Inbox != nil {
		m["inbox"] = n.Inbox
	}
	if n.BigPicPath != "" {
		m["big_pic_path"] = n.BigPicPath
	}
	if n.LargeIcon != "" {
		m["large_icon"] = n.LargeIcon
	}
	if n.SmallIconURI != "" {
		m["small_icon_uri"] = n.SmallIconURI
	}
	if n.IconBgColor != "" {
		m["icon_bg_color"] = n.IconBgColor
	}
	if n.Intent != nil {
		m["intent"] = n.Intent.ToMap()
	}
	if n.Extras != nil {
		m["extras"] = n.Extras
	}
	return m
}

// Intent Android 跳转 Intent
type Intent struct {
	Action    string `json:"action,omitempty"`
	Component string `json:"component,omitempty"`
}

// ToMap 转换为 Map
func (i *Intent) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if i.Action != "" {
		m["action"] = i.Action
	}
	if i.Component != "" {
		m["component"] = i.Component
	}
	return m
}

// IOSNotification iOS 通知
type IOSNotification struct {
	Alert             interface{}            `json:"alert,omitempty"`
	Sound             string                 `json:"sound,omitempty"`
	Badge             interface{}            `json:"badge,omitempty"`
	ContentAvailable  int                    `json:"content-available,omitempty"`
	MutableContent    int                    `json:"mutable-content,omitempty"`
	Category          string                 `json:"category,omitempty"`
	ThreadID          string                 `json:"thread-id,omitempty"`
	InterruptionLevel string                 `json:"interruption-level,omitempty"`
	Extras            map[string]interface{} `json:"extras,omitempty"`
}

// ToMap 转换为 Map
func (n *IOSNotification) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if n.Alert != nil {
		m["alert"] = n.Alert
	}
	if n.Sound != "" {
		m["sound"] = n.Sound
	}
	if n.Badge != nil {
		m["badge"] = n.Badge
	}
	if n.ContentAvailable != 0 {
		m["content-available"] = n.ContentAvailable
	}
	if n.MutableContent != 0 {
		m["mutable-content"] = n.MutableContent
	}
	if n.Category != "" {
		m["category"] = n.Category
	}
	if n.ThreadID != "" {
		m["thread-id"] = n.ThreadID
	}
	if n.InterruptionLevel != "" {
		m["interruption-level"] = n.InterruptionLevel
	}
	if n.Extras != nil {
		m["extras"] = n.Extras
	}
	return m
}

// HmosNotification 鸿蒙通知
type HmosNotification struct {
	Alert       string                 `json:"alert,omitempty"`
	Title       string                 `json:"title,omitempty"`
	Category    string                 `json:"category,omitempty"`
	LargeIcon   string                 `json:"large_icon,omitempty"`
	Intent      *HmosIntent            `json:"intent,omitempty"`
	BadgeAddNum int                    `json:"badge_add_num,omitempty"`
	BadgeSetNum int                    `json:"badge_set_num,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
	Style       int                    `json:"style,omitempty"`
	Inbox       map[string]interface{} `json:"inbox,omitempty"`
}

// ToMap 转换为 Map
func (n *HmosNotification) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if n.Alert != "" {
		m["alert"] = n.Alert
	}
	if n.Title != "" {
		m["title"] = n.Title
	}
	if n.Category != "" {
		m["category"] = n.Category
	}
	if n.LargeIcon != "" {
		m["large_icon"] = n.LargeIcon
	}
	if n.Intent != nil {
		m["intent"] = n.Intent.ToMap()
	}
	if n.BadgeAddNum != 0 {
		m["badge_add_num"] = n.BadgeAddNum
	}
	if n.BadgeSetNum != 0 {
		m["badge_set_num"] = n.BadgeSetNum
	}
	if n.Extras != nil {
		m["extras"] = n.Extras
	}
	if n.Style != 0 {
		m["style"] = n.Style
	}
	if n.Inbox != nil {
		m["inbox"] = n.Inbox
	}
	return m
}

// HmosIntent 鸿蒙 Intent
type HmosIntent struct {
	Action string `json:"action,omitempty"`
}

// ToMap 转换为 Map
func (i *HmosIntent) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if i.Action != "" {
		m["action"] = i.Action
	}
	return m
}

// QuickAppNotification 快应用通知
type QuickAppNotification struct {
	Alert  string                 `json:"alert,omitempty"`
	Title  string                 `json:"title,omitempty"`
	Page   string                 `json:"page,omitempty"`
	Extras map[string]interface{} `json:"extras,omitempty"`
}

// ToMap 转换为 Map
func (n *QuickAppNotification) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if n.Alert != "" {
		m["alert"] = n.Alert
	}
	if n.Title != "" {
		m["title"] = n.Title
	}
	if n.Page != "" {
		m["page"] = n.Page
	}
	if n.Extras != nil {
		m["extras"] = n.Extras
	}
	return m
}

// ==================== 自定义消息定义 ====================

// Message 自定义消息
type Message struct {
	MsgContent  string                 `json:"msg_content,omitempty"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

// ToMap 转换为 Map
func (m *Message) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	if m.MsgContent != "" {
		result["msg_content"] = m.MsgContent
	}
	if m.Title != "" {
		result["title"] = m.Title
	}
	if m.ContentType != "" {
		result["content_type"] = m.ContentType
	}
	if m.Extras != nil {
		result["extras"] = m.Extras
	}
	return result
}

// ==================== 短信消息定义 ====================

// SmsMessage 短信消息
type SmsMessage struct {
	DelayTime    int                    `json:"delay_time,omitempty"`
	TempID       int64                  `json:"temp_id,omitempty"`
	SignID       int64                  `json:"signid,omitempty"`
	TempPara     map[string]interface{} `json:"temp_para,omitempty"`
	ActiveFilter bool                   `json:"active_filter,omitempty"`
}

// ToMap 转换为 Map
func (s *SmsMessage) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if s.DelayTime != 0 {
		m["delay_time"] = s.DelayTime
	}
	if s.TempID != 0 {
		m["temp_id"] = s.TempID
	}
	if s.SignID != 0 {
		m["signid"] = s.SignID
	}
	if s.TempPara != nil {
		m["temp_para"] = s.TempPara
	}
	if !s.ActiveFilter {
		m["active_filter"] = false
	}
	return m
}

// ==================== 推送选项定义 ====================

// Options 推送选项
type Options struct {
	SendNo             int    `json:"sendno,omitempty"`
	TimeToLive         int    `json:"time_to_live,omitempty"`
	OverrideMsgID      string `json:"override_msg_id,omitempty"`
	ApnsProduction     bool   `json:"apns_production,omitempty"`
	ApnsCollapseID     string `json:"apns_collapse_id,omitempty"`
	BigPushDuration    int    `json:"big_push_duration,omitempty"`
	Classification     int    `json:"classification,omitempty"`
	TestMessage        bool   `json:"test_message,omitempty"`
	ReceiptID          string `json:"receipt_id,omitempty"`
	ActivePush         bool   `json:"active_push,omitempty"`
	NeedBackup         bool   `json:"need_backup,omitempty"`
	TestModel          bool   `json:"test_model,omitempty"`
	Notification3rdVer string `json:"notification_3rd_ver,omitempty"`
	AutoTruncation     bool   `json:"auto_truncation,omitempty"`
	MktEnable          bool   `json:"mkt_enable,omitempty"`
}

// ToMap 转换为 Map
func (o *Options) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if o.SendNo != 0 {
		m["sendno"] = o.SendNo
	}
	if o.TimeToLive != 0 {
		m["time_to_live"] = o.TimeToLive
	}
	if o.OverrideMsgID != "" {
		m["override_msg_id"] = o.OverrideMsgID
	}
	m["apns_production"] = o.ApnsProduction
	if o.ApnsCollapseID != "" {
		m["apns_collapse_id"] = o.ApnsCollapseID
	}
	if o.BigPushDuration != 0 {
		m["big_push_duration"] = o.BigPushDuration
	}
	if o.Classification != 0 {
		m["classification"] = o.Classification
	}
	if o.TestMessage {
		m["test_message"] = o.TestMessage
	}
	if o.ReceiptID != "" {
		m["receipt_id"] = o.ReceiptID
	}
	if o.ActivePush {
		m["active_push"] = o.ActivePush
	}
	if o.NeedBackup {
		m["need_backup"] = o.NeedBackup
	}
	if o.TestModel {
		m["test_model"] = o.TestModel
	}
	if o.Notification3rdVer != "" {
		m["notification_3rd_ver"] = o.Notification3rdVer
	}
	if o.AutoTruncation {
		m["auto_truncation"] = o.AutoTruncation
	}
	if o.MktEnable {
		m["mkt_enable"] = o.MktEnable
	}
	return m
}

// ==================== 批量推送 ====================

// BatchPushByRegID 批量按 RegistrationID 推送
func (c *Client) BatchPushByRegID(pushList []map[string]interface{}) (*PushResponse, error) {
	cidList, err := c.GetCID(len(pushList), "push")
	if err != nil {
		return nil, err
	}

	batchPayload := map[string]map[string]interface{}{
		"pushlist": {},
	}
	for i, payload := range pushList {
		batchPayload["pushlist"][cidList[i]] = payload
	}

	data, err := json.Marshal(batchPayload)
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("PUSH", c.Zone) + "/batch/regid/single"
	respBody, statusCode, err := c.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result PushResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// BatchPushByAlias 批量按别名推送
func (c *Client) BatchPushByAlias(pushList []map[string]interface{}) (*PushResponse, error) {
	cidList, err := c.GetCID(len(pushList), "push")
	if err != nil {
		return nil, err
	}

	batchPayload := map[string]map[string]interface{}{
		"pushlist": {},
	}
	for i, payload := range pushList {
		batchPayload["pushlist"][cidList[i]] = payload
	}

	data, err := json.Marshal(batchPayload)
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("PUSH", c.Zone) + "/batch/alias/single"
	respBody, statusCode, err := c.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result PushResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// ==================== 推送撤销 ====================

// RevokePush 撤销推送消息
// 功能说明:
//   - Android 消息（排队中/发送中状态）可从服务端撤销
//   - iOS 消息（排队中状态）可从服务端撤销
//   - 针对 JPush SDK v3.5.0+(Android) 和 v3.2.8+(iOS)，会尝试从设备端撤销已展示但未被点击的消息
//   - 目前支持设备端撤销的厂商：小米、vivo
//
// 调用地址: DELETE https://api.jpush.cn/v3/push/{msgid}
// 参数: msgID - 要撤销的消息ID
func (c *Client) RevokePush(msgID string) (*PushResponse, error) {
	url := GetURL("PUSH", c.Zone) + "/" + msgID
	respBody, statusCode, err := c.doRequest(RequestOptions{
		Method: http.MethodDelete,
		URL:    url,
	})
	if err != nil {
		return nil, err
	}

	var result PushResponse
	result.StatusCode = statusCode
	if len(respBody) > 0 {
		if err := json.Unmarshal(respBody, &result); err != nil {
			result.ErrorMessage = string(respBody)
		}
	}

	return &result, nil
}
