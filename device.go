package jpush

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ==================== 设备载荷 ====================

// DevicePayload 设备管理载荷
type DevicePayload struct {
	client *Client
}

// DeviceInfo 设备信息
type DeviceInfo struct {
	Tags      []string `json:"tags,omitempty"`
	Alias     string   `json:"alias,omitempty"`
	Mobile    string   `json:"mobile,omitempty"`
	Mark      string   `json:"mark,omitempty"`
}

// TagOperation 标签操作
type TagOperation struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

// AliasOperation 别名操作
type AliasOperation struct {
	Alias string `json:"alias,omitempty"`
}

// ==================== Device API 方法 ====================

// GetDeviceInfo 获取设备信息
func (d *DevicePayload) GetDeviceInfo(registrationID string) (*DeviceInfoResponse, error) {
	url := GetURL("DEVICE", d.client.Zone) + registrationID
	respBody, statusCode, err := d.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
	})
	if err != nil {
		return nil, err
	}

	var result DeviceInfoResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// UpdateDeviceTags 更新设备标签
func (d *DevicePayload) UpdateDeviceTags(registrationID string, tags *TagOperation) (*DeviceResponse, error) {
	data, err := json.Marshal(map[string]interface{}{
		"tags": tags,
	})
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("DEVICE", d.client.Zone) + registrationID
	respBody, statusCode, err := d.client.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result DeviceResponse
	result.StatusCode = statusCode
	result.Response = string(respBody)
	return &result, nil
}

// UpdateDeviceAlias 更新设备别名
func (d *DevicePayload) UpdateDeviceAlias(registrationID, alias string) (*DeviceResponse, error) {
	data, err := json.Marshal(map[string]interface{}{
		"alias": alias,
	})
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("DEVICE", d.client.Zone) + registrationID
	respBody, statusCode, err := d.client.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result DeviceResponse
	result.StatusCode = statusCode
	result.Response = string(respBody)
	return &result, nil
}

// UpdateDeviceMobile 更新设备手机号
func (d *DevicePayload) UpdateDeviceMobile(registrationID, mobile string) (*DeviceResponse, error) {
	data, err := json.Marshal(map[string]interface{}{
		"mobile": mobile,
	})
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("DEVICE", d.client.Zone) + registrationID
	respBody, statusCode, err := d.client.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result DeviceResponse
	result.StatusCode = statusCode
	result.Response = string(respBody)
	return &result, nil
}

// GetDeviceStatus 获取设备在线状态
func (d *DevicePayload) GetDeviceStatus(registrationIDs []string) (*DeviceStatusResponse, error) {
	data, err := json.Marshal(map[string]interface{}{
		"registration_ids": registrationIDs,
	})
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("DEVICE", d.client.Zone) + "status"
	respBody, statusCode, err := d.client.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result DeviceStatusResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// ==================== Tag API 方法 ====================

// GetTagList 获取标签列表
func (d *DevicePayload) GetTagList() (*TagListResponse, error) {
	url := GetURL("TAG", d.client.Zone)
	respBody, statusCode, err := d.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
	})
	if err != nil {
		return nil, err
	}

	var result TagListResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// DeleteTag 删除标签
func (d *DevicePayload) DeleteTag(tag, platform string) (*DeviceResponse, error) {
	url := GetURL("TAG", d.client.Zone) + tag
	if platform != "" {
		url += "?platform=" + platform
	}

	respBody, statusCode, err := d.client.doRequest(RequestOptions{
		Method: http.MethodDelete,
		URL:    url,
	})
	if err != nil {
		return nil, err
	}

	var result DeviceResponse
	result.StatusCode = statusCode
	result.Response = string(respBody)
	return &result, nil
}

// UpdateTagUsers 更新标签用户
func (d *DevicePayload) UpdateTagUsers(tag string, operation *TagOperation) (*DeviceResponse, error) {
	data, err := json.Marshal(map[string]interface{}{
		"registration_ids": operation,
	})
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("TAG", d.client.Zone) + tag
	respBody, statusCode, err := d.client.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result DeviceResponse
	result.StatusCode = statusCode
	result.Response = string(respBody)
	return &result, nil
}

// CheckTagUserExist 检查用户是否在标签中
func (d *DevicePayload) CheckTagUserExist(tag, registrationID string) (*TagUserCheckResponse, error) {
	url := GetURL("TAG", d.client.Zone) + tag + "/registration_ids/" + registrationID
	respBody, statusCode, err := d.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
	})
	if err != nil {
		return nil, err
	}

	var result TagUserCheckResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// ==================== Alias API 方法 ====================

// GetAliasUsers 获取别名用户列表
func (d *DevicePayload) GetAliasUsers(alias, platform string) (*AliasUsersResponse, error) {
	url := GetURL("ALIAS", d.client.Zone) + alias
	params := map[string]string{}
	if platform != "" {
		params["platform"] = platform
	}

	respBody, statusCode, err := d.client.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
		Params: params,
	})
	if err != nil {
		return nil, err
	}

	var result AliasUsersResponse
	result.StatusCode = statusCode
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// DeleteAlias 删除别名
func (d *DevicePayload) DeleteAlias(alias, platform string) (*DeviceResponse, error) {
	url := GetURL("ALIAS", d.client.Zone) + alias
	if platform != "" {
		url += "?platform=" + platform
	}

	respBody, statusCode, err := d.client.doRequest(RequestOptions{
		Method: http.MethodDelete,
		URL:    url,
	})
	if err != nil {
		return nil, err
	}

	var result DeviceResponse
	result.StatusCode = statusCode
	result.Response = string(respBody)
	return &result, nil
}

// ==================== Device 响应 ====================

// DeviceResponse 设备基础响应
type DeviceResponse struct {
	StatusCode int    `json:"-"`
	Response   string `json:"-"`
}

// IsSuccess 检查是否成功
func (r *DeviceResponse) IsSuccess() bool {
	return r.StatusCode == http.StatusOK
}

// DeviceInfoResponse 设备信息响应
type DeviceInfoResponse struct {
	StatusCode int      `json:"-"`
	Tags       []string `json:"tags,omitempty"`
	Alias      string   `json:"alias,omitempty"`
	Mobile     string   `json:"mobile,omitempty"`
}

// DeviceStatusResponse 设备状态响应
type DeviceStatusResponse struct {
	StatusCode int                  `json:"-"`
	Devices    []DeviceStatusDetail `json:"devices,omitempty"`
}

// DeviceStatusDetail 设备状态详情
type DeviceStatusDetail struct {
	RegistrationID string `json:"registration_id"`
	Online         bool   `json:"online"`
	LastOnlineTime string `json:"last_online_time,omitempty"`
}

// TagListResponse 标签列表响应
type TagListResponse struct {
	StatusCode int      `json:"-"`
	Tags       []string `json:"tags,omitempty"`
}

// TagUserCheckResponse 标签用户检查响应
type TagUserCheckResponse struct {
	StatusCode int  `json:"-"`
	Result     bool `json:"result"` // true: 存在, false: 不存在
}

// AliasUsersResponse 别名用户响应
type AliasUsersResponse struct {
	StatusCode    int      `json:"-"`
	RegistrationIDs []string `json:"registration_ids,omitempty"`
	Alias         string   `json:"alias,omitempty"`
}

// ==================== Device 构建器辅助函数 ====================

// NewTagOperation 创建标签操作
func NewTagOperation(add, remove []string) *TagOperation {
	return &TagOperation{
		Add:    add,
		Remove: remove,
	}
}

// AddTags 添加标签
func (t *TagOperation) AddTags(tags ...string) *TagOperation {
	t.Add = append(t.Add, tags...)
	return t
}

// RemoveTags 移除标签
func (t *TagOperation) RemoveTags(tags ...string) *TagOperation {
	t.Remove = append(t.Remove, tags...)
	return t
}

// ==================== 便捷方法 ====================

// AddDeviceTag 添加设备标签
func (d *DevicePayload) AddDeviceTag(registrationID string, tags ...string) (*DeviceResponse, error) {
	return d.UpdateDeviceTags(registrationID, NewTagOperation(tags, nil))
}

// RemoveDeviceTag 移除设备标签
func (d *DevicePayload) RemoveDeviceTag(registrationID string, tags ...string) (*DeviceResponse, error) {
	return d.UpdateDeviceTags(registrationID, NewTagOperation(nil, tags))
}

// AddTagUsers 添加标签用户
func (d *DevicePayload) AddTagUsers(tag string, regIDs ...string) (*DeviceResponse, error) {
	return d.UpdateTagUsers(tag, NewTagOperation(regIDs, nil))
}

// RemoveTagUsers 移除标签用户
func (d *DevicePayload) RemoveTagUsers(tag string, regIDs ...string) (*DeviceResponse, error) {
	return d.UpdateTagUsers(tag, NewTagOperation(nil, regIDs))
}

// ==================== 批量操作 ====================

// BatchUpdateDeviceTags 批量更新设备标签
func (d *DevicePayload) BatchUpdateDeviceTags(registrationIDs []string, tags *TagOperation) (map[string]*DeviceResponse, error) {
	results := make(map[string]*DeviceResponse)

	for _, regID := range registrationIDs {
		resp, err := d.UpdateDeviceTags(regID, tags)
		if err != nil {
			results[regID] = &DeviceResponse{
				Response: err.Error(),
			}
		} else {
			results[regID] = resp
		}
	}

	return results, nil
}

// ==================== 字符串处理 ====================

// ParsePlatform 解析平台字符串
func ParsePlatform(platform string) []string {
	if platform == "" {
		return nil
	}
	return strings.Split(platform, ",")
}

// JoinPlatform 合并平台为字符串
func JoinPlatform(platforms []string) string {
	return strings.Join(platforms, ",")
}
