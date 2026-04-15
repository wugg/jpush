package jpush

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ==================== 定时任务载荷 ====================

// SchedulePayload 定时任务载荷
type SchedulePayload struct {
	client     *Client
	Name       string                 `json:"name,omitempty"`
	Enabled    bool                   `json:"enabled"`
	Trigger    *ScheduleTrigger       `json:"trigger,omitempty"`
	Push       map[string]interface{} `json:"push,omitempty"`
}

// ScheduleTrigger 触发器
type ScheduleTrigger struct {
	Single      *SingleTrigger      `json:"single,omitempty"`
	Periodical  *PeriodicalTrigger   `json:"periodical,omitempty"`
}

// SingleTrigger 单次触发
type SingleTrigger struct {
	Time string `json:"time"` // 格式: "yyyy-MM-dd HH:mm:ss"
}

// PeriodicalTrigger 周期触发
type PeriodicalTrigger struct {
	Start     string   `json:"start,omitempty"`      // 格式: "yyyy-MM-dd HH:mm:ss"
	End       string   `json:"end,omitempty"`         // 格式: "yyyy-MM-dd HH:mm:ss"
	Time      string   `json:"time,omitempty"`        // 格式: "HH:mm:ss"
	TimeUnit  string   `json:"time_unit,omitempty"`   // day, week, month
	Frequency int      `json:"frequency,omitempty"`   // 周期频率
	Point     []string `json:"point,omitempty"`      // 执行时间点
}

// ==================== Schedule 创建方法 ====================

// SetName 设置任务名称
func (s *SchedulePayload) SetName(name string) *SchedulePayload {
	s.Name = name
	return s
}

// SetEnabled 设置任务状态
func (s *SchedulePayload) SetEnabled(enabled bool) *SchedulePayload {
	s.Enabled = enabled
	return s
}

// SetSingleTrigger 设置单次触发
func (s *SchedulePayload) SetSingleTrigger(time string) *SchedulePayload {
	s.Trigger = &ScheduleTrigger{
		Single: &SingleTrigger{
			Time: time,
		},
	}
	return s
}

// SetPeriodicalTrigger 设置周期触发
func (s *SchedulePayload) SetPeriodicalTrigger(trigger *PeriodicalTrigger) *SchedulePayload {
	s.Trigger = &ScheduleTrigger{
		Periodical: trigger,
	}
	return s
}

// SetPush 设置推送内容
func (s *SchedulePayload) SetPush(push map[string]interface{}) *SchedulePayload {
	s.Push = push
	return s
}

// ==================== Schedule API 方法 ====================

// Create 创建定时任务
func (s *SchedulePayload) Create() (*ScheduleResponse, error) {
	if s.Trigger == nil {
		return nil, fmt.Errorf("触发器不能为空")
	}
	if s.Push == nil {
		return nil, fmt.Errorf("推送内容不能为空")
	}

	data, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("SCHEDULE", s.client.Zone)
	respBody, statusCode, err := s.client.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result ScheduleResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// Update 更新定时任务
func (s *SchedulePayload) Update(scheduleID string) (*ScheduleResponse, error) {
	if s.Trigger == nil {
		return nil, fmt.Errorf("触发器不能为空")
	}
	if s.Push == nil {
		return nil, fmt.Errorf("推送内容不能为空")
	}

	data, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("SCHEDULE", s.client.Zone) + scheduleID
	respBody, statusCode, err := s.client.doRequest(RequestOptions{
		Method: http.MethodPut,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result ScheduleResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// Delete 删除定时任务
func (s *SchedulePayload) Delete(scheduleID string) (*ScheduleResponse, error) {
	url := GetURL("SCHEDULE", s.client.Zone) + scheduleID
	respBody, statusCode, err := s.client.doRequest(RequestOptions{
		Method: http.MethodDelete,
		URL:    url,
	})
	if err != nil {
		return nil, err
	}

	var result ScheduleResponse
	result.StatusCode = statusCode
	if len(respBody) > 0 {
		if err := json.Unmarshal(respBody, &result); err != nil {
			result.Message = string(respBody)
		}
	}

	return &result, nil
}

// GetByID 获取定时任务详情
func (s *SchedulePayload) GetByID(scheduleID string) (*ScheduleDetailResponse, error) {
	url := GetURL("SCHEDULE", s.client.Zone) + scheduleID
	respBody, statusCode, err := s.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
	})
	if err != nil {
		return nil, err
	}

	var result ScheduleDetailResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// GetList 获取定时任务列表
func (s *SchedulePayload) GetList(page int) (*ScheduleListResponse, error) {
	url := GetURL("SCHEDULE", s.client.Zone)
	respBody, statusCode, err := s.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
		Params: map[string]string{
			"page": strconv.Itoa(page),
		},
	})
	if err != nil {
		return nil, err
	}

	var result ScheduleListResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// GetMsgIDs 获取任务消息 ID
func (s *SchedulePayload) GetMsgIDs(scheduleID string) (*ScheduleMsgIDsResponse, error) {
	url := GetURL("SCHEDULE", s.client.Zone) + scheduleID + "/msg_ids"
	respBody, statusCode, err := s.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
	})
	if err != nil {
		return nil, err
	}

	var result ScheduleMsgIDsResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// ==================== Schedule 响应 ====================

// ScheduleResponse 定时任务基础响应
type ScheduleResponse struct {
	StatusCode   int    `json:"-"`
	ScheduleID   string `json:"schedule_id,omitempty"`
	Message      string `json:"message,omitempty"`
	ErrorCode    int    `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// IsSuccess 检查是否成功
func (r *ScheduleResponse) IsSuccess() bool {
	return r.ScheduleID != "" || r.Message == "success"
}

// ScheduleDetailResponse 定时任务详情响应
type ScheduleDetailResponse struct {
	StatusCode    int               `json:"-"`
	ScheduleID    string            `json:"schedule_id,omitempty"`
	Name          string            `json:"name,omitempty"`
	Enabled       bool              `json:"enabled"`
	Trigger       *ScheduleTrigger  `json:"trigger,omitempty"`
	Push          interface{}       `json:"push,omitempty"`
	CreatedAt     string            `json:"created_at,omitempty"`
	UpdatedAt     string            `json:"updated_at,omitempty"`
	ErrorCode     int               `json:"error_code,omitempty"`
	ErrorMessage  string            `json:"error_message,omitempty"`
}

// ScheduleListResponse 定时任务列表响应
type ScheduleListResponse struct {
	StatusCode int                    `json:"-"`
	Total      int                    `json:"total"`
	Page       int                    `json:"page"`
	Schedules  []ScheduleListItem    `json:"schedules,omitempty"`
	ErrorCode  int                    `json:"error_code,omitempty"`
	ErrorMessage string               `json:"error_message,omitempty"`
}

// ScheduleListItem 定时任务列表项
type ScheduleListItem struct {
	ScheduleID   string `json:"schedule_id"`
	Name         string `json:"name"`
	Enabled      bool   `json:"enabled"`
	TriggerType  string `json:"trigger_type,omitempty"`  // single, periodical
	NextFireTime string `json:"next_fire_time,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
}

// ScheduleMsgIDsResponse 消息 ID 响应
type ScheduleMsgIDsResponse struct {
	StatusCode int      `json:"-"`
	MsgIDs     []string `json:"msg_ids,omitempty"`
	ErrorCode  int      `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// ==================== Schedule 构建器辅助函数 ====================

// NewSingleTrigger 创建单次触发器
func NewSingleTrigger(time string) *SingleTrigger {
	return &SingleTrigger{
		Time: time,
	}
}

// NewPeriodicalTrigger 创建周期触发器
func NewPeriodicalTrigger(start, time, timeUnit string, frequency int) *PeriodicalTrigger {
	return &PeriodicalTrigger{
		Start:     start,
		Time:      time,
		TimeUnit:  timeUnit,
		Frequency: frequency,
	}
}

// SetEnd 设置结束时间
func (t *PeriodicalTrigger) SetEnd(end string) *PeriodicalTrigger {
	t.End = end
	return t
}

// SetPoint 设置执行时间点
func (t *PeriodicalTrigger) SetPoint(point []string) *PeriodicalTrigger {
	t.Point = point
	return t
}

// NewDailyTrigger 创建每天触发器
func NewDailyTrigger(start, time string, frequency int) *PeriodicalTrigger {
	return &PeriodicalTrigger{
		Start:     start,
		Time:      time,
		TimeUnit:  TimeUnitDay,
		Frequency: frequency,
	}
}

// NewWeeklyTrigger 创建每周触发器
func NewWeeklyTrigger(start, time string, point []string, frequency int) *PeriodicalTrigger {
	return &PeriodicalTrigger{
		Start:     start,
		Time:      time,
		TimeUnit:  TimeUnitWeek,
		Frequency: frequency,
		Point:     point,
	}
}

// NewMonthlyTrigger 创建每月触发器
func NewMonthlyTrigger(start, time string, point []string, frequency int) *PeriodicalTrigger {
	return &PeriodicalTrigger{
		Start:     start,
		Time:      time,
		TimeUnit:  TimeUnitMonth,
		Frequency: frequency,
		Point:     point,
	}
}

// NewSchedulePayload 创建定时任务载荷
func NewSchedulePayload(name string, enabled bool, trigger *ScheduleTrigger, push map[string]interface{}) *SchedulePayload {
	return &SchedulePayload{
		Name:    name,
		Enabled: enabled,
		Trigger: trigger,
		Push:    push,
	}
}
