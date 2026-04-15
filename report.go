package jpush

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ==================== 报告载荷 ====================

// ReportPayload 报告载荷
type ReportPayload struct {
	client *Client
}

// ==================== Report API 方法 ====================

// GetReceived 获取送达统计
func (r *ReportPayload) GetReceived(msgIDs string) (*ReceivedResponse, error) {
	url := GetURL("REPORT", r.client.Zone) + "received"
	respBody, statusCode, err := r.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
		Params: map[string]string{
			"msg_ids": msgIDs,
		},
	})
	if err != nil {
		return nil, err
	}

	var result ReceivedResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// GetReceivedDetail 获取送达详情
func (r *ReportPayload) GetReceivedDetail(msgIDs string) (*ReceivedDetailResponse, error) {
	url := GetURL("REPORT", r.client.Zone) + "received/detail"
	respBody, statusCode, err := r.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
		Params: map[string]string{
			"msg_ids": msgIDs,
		},
	})
	if err != nil {
		return nil, err
	}

	var result ReceivedDetailResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// GetMessages 获取消息统计
func (r *ReportPayload) GetMessages(msgIDs string) (*MessagesResponse, error) {
	url := GetURL("REPORT", r.client.Zone) + "messages"
	respBody, statusCode, err := r.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
		Params: map[string]string{
			"msg_ids": msgIDs,
		},
	})
	if err != nil {
		return nil, err
	}

	var result MessagesResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// GetMessagesDetail 获取消息详情
func (r *ReportPayload) GetMessagesDetail(msgIDs string) (*MessagesDetailResponse, error) {
	url := GetURL("REPORT", r.client.Zone) + "messages/detail"
	respBody, statusCode, err := r.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
		Params: map[string]string{
			"msg_ids": msgIDs,
		},
	})
	if err != nil {
		return nil, err
	}

	var result MessagesDetailResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// GetStatusMessage 获取消息状态
func (r *ReportPayload) GetStatusMessage(msgID string, regIDs []string, date string) (*StatusMessageResponse, error) {
	body := map[string]interface{}{
		"msg_id":           msgID,
		"registration_ids": regIDs,
	}
	if date != "" {
		body["date"] = date
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("REPORT", r.client.Zone) + "status/message"
	respBody, statusCode, err := r.client.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result StatusMessageResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// GetUsers 获取用户统计
func (r *ReportPayload) GetUsers(timeUnit, start string, duration int) (*UsersResponse, error) {
	url := GetURL("REPORT", r.client.Zone) + "users"
	respBody, statusCode, err := r.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
		Params: map[string]string{
			"time_unit": timeUnit,
			"start":    start,
			"duration": fmt.Sprintf("%d", duration),
		},
	})
	if err != nil {
		return nil, err
	}

	var result UsersResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// ==================== Report 响应 ====================

// ReceivedResponse 送达统计响应
type ReceivedResponse struct {
	StatusCode int                 `json:"-"`
	ReceivedList []ReceivedDetail  `json:"received_list,omitempty"`
	ErrorCode  int                 `json:"error_code,omitempty"`
	ErrorMessage string           `json:"error_message,omitempty"`
}

// ReceivedDetail 送达详情
type ReceivedDetail struct {
	MsgID     string `json:"msg_id"`
	Android   int    `json:"android"`
	Ios       int    `json:"ios"`
	Hmos      int    `json:"hmos"`
	OnlinePush int   `json:"online_push,omitempty"`
}

// ReceivedDetailResponse 送达详情响应
type ReceivedDetailResponse struct {
	StatusCode int                    `json:"-"`
	ReceivedList []ReceivedDetail     `json:"received_list,omitempty"`
	ErrorCode  int                    `json:"error_code,omitempty"`
	ErrorMessage string               `json:"error_message,omitempty"`
}

// MessagesResponse 消息统计响应
type MessagesResponse struct {
	StatusCode int                `json:"-"`
	MessageList []MessageDetail   `json:"message_list,omitempty"`
	ErrorCode  int                 `json:"error_code,omitempty"`
	ErrorMessage string            `json:"error_message,omitempty"`
}

// MessageDetail 消息详情
type MessageDetail struct {
	MsgID       string        `json:"msg_id"`
	Android     int           `json:"android"`
	Ios         int           `json:"ios"`
	Hmos        int           `json:"hmos"`
	QuickApp    int           `json:"quickapp,omitempty"`
	OnlinePush  int           `json:"online_push,omitempty"`
}

// MessagesDetailResponse 消息详情响应
type MessagesDetailResponse struct {
	StatusCode int                   `json:"-"`
	Messages   []MessageDetail       `json:"messages,omitempty"`
	ErrorCode  int                   `json:"error_code,omitempty"`
	ErrorMessage string              `json:"error_message,omitempty"`
}

// StatusMessageResponse 消息状态响应
type StatusMessageResponse struct {
	StatusCode int                    `json:"-"`
	List      []StatusMessageDetail   `json:"list,omitempty"`
	ErrorCode int                    `json:"error_code,omitempty"`
	ErrorMessage string               `json:"error_message,omitempty"`
}

// StatusMessageDetail 消息状态详情
type StatusMessageDetail struct {
	RegistrationID string `json:"registration_id"`
	Status        string `json:"status,omitempty"` // received, offline, not_received
	ReceivedTime  int64  `json:"received_time,omitempty"`
}

// UsersResponse 用户统计响应
type UsersResponse struct {
	StatusCode int          `json:"-"`
	TimeUnit   string       `json:"time_unit,omitempty"`
	Start      string       `json:"start,omitempty"`
	Duration   int          `json:"duration,omitempty"`
	List       []UserDetail `json:"list,omitempty"`
	ErrorCode  int          `json:"error_code,omitempty"`
	ErrorMessage string     `json:"error_message,omitempty"`
}

// UserDetail 用户详情
type UserDetail struct {
	Time   string `json:"time,omitempty"`
	New    int    `json:"new,omitempty"`
	Online int    `json:"online,omitempty"`
	Total  int64  `json:"total,omitempty"`
}

// ==================== 便捷方法 ====================

// GetReceivedByIDs 通过消息 ID 列表获取送达统计
func (r *ReportPayload) GetReceivedByIDs(msgIDs []int64) (*ReceivedResponse, error) {
	ids := make([]string, len(msgIDs))
	for i, id := range msgIDs {
		ids[i] = fmt.Sprintf("%d", id)
	}
	return r.GetReceived(strings.Join(ids, ","))
}

// GetMessagesByIDs 通过消息 ID 列表获取消息统计
func (r *ReportPayload) GetMessagesByIDs(msgIDs []int64) (*MessagesResponse, error) {
	ids := make([]string, len(msgIDs))
	for i, id := range msgIDs {
		ids[i] = fmt.Sprintf("%d", id)
	}
	return r.GetMessages(strings.Join(ids, ","))
}

// GetUsersDaily 获取每日用户统计
func (r *ReportPayload) GetUsersDaily(start string, duration int) (*UsersResponse, error) {
	return r.GetUsers("DAY", start, duration)
}

// GetUsersHourly 获取每小时用户统计
func (r *ReportPayload) GetUsersHourly(start string, duration int) (*UsersResponse, error) {
	return r.GetUsers("HOUR", start, duration)
}

// GetUsersMonthly 获取每月用户统计
func (r *ReportPayload) GetUsersMonthly(start string, duration int) (*UsersResponse, error) {
	return r.GetUsers("MONTH", start, duration)
}
