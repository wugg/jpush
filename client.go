package jpush

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// ==================== 客户端定义 ====================

// Client JPush 客户端
type Client struct {
	AppKey       string        // 应用 Key
	MasterSecret string        // Master Secret
	Zone         Zone          // 区域配置
	HttpClient   *http.Client  // HTTP 客户端
	timeout      time.Duration // 请求超时时间
}



// ClientOption 客户端配置选项
type ClientOption func(*Client)

// WithZone 设置区域
func WithZone(zone Zone) ClientOption {
	return func(c *Client) {
		c.Zone = zone
	}
}

// WithTimeout 设置超时时间
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.timeout = timeout
		c.HttpClient.Timeout = timeout
	}
}

// NewClient 创建 JPush 客户端
func NewClient(appKey, masterSecret string, opts ...ClientOption) *Client {
	c := &Client{
		AppKey:       appKey,
		MasterSecret: masterSecret,
		Zone:         ZoneDefault,
		timeout:      30 * time.Second,
		HttpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// ==================== 请求方法 ====================

// RequestOptions 请求配置
type RequestOptions struct {
	Method      string
	URL         string
	Body        []byte
	ContentType string
	Params      map[string]string
}

// doRequest 发起 HTTP 请求
func (c *Client) doRequest(opts RequestOptions) ([]byte, int, error) {
	var bodyReader io.Reader
	if opts.Body != nil {
		bodyReader = bytes.NewReader(opts.Body)
	}

	// 构建 URL（添加查询参数）
	url := opts.URL
	if len(opts.Params) > 0 {
		url = addQueryParams(opts.URL, opts.Params)
	}

	req, err := http.NewRequest(opts.Method, url, bodyReader)
	if err != nil {
		return nil, 0, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", c.getAuthHeader())
	req.Header.Set("User-Agent", "jpush-api-go-client")
	req.Header.Set("Connection", "keep-alive")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, 0, NewAPIConnectionException(fmt.Sprintf("连接失败: %v", err))
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, NewAPIRequestException(fmt.Sprintf("读取响应失败: %v", err))
	}

	// 处理错误响应
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, NewUnauthorized("认证失败，请检查 AppKey 和 MasterSecret")
	}

	if resp.StatusCode >= 400 {
		var errResp struct {
			Error struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			} `json:"error"`
		}
		if err := json.Unmarshal(respBody, &errResp); err == nil && errResp.Error.Code != 0 {
			return nil, resp.StatusCode, NewJPushFailure(errResp.Error.Code, errResp.Error.Message, string(respBody))
		}
		return nil, resp.StatusCode, fmt.Errorf("请求失败: HTTP %d, Body: %s", resp.StatusCode, string(respBody))
	}

	return respBody, resp.StatusCode, nil
}

// getAuthHeader 获取认证头
func (c *Client) getAuthHeader() string {
	auth := c.AppKey + ":" + c.MasterSecret
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

// ==================== Push API ====================

// Push 创建推送载荷构建器
func (c *Client) Push() *PushPayload {
	return &PushPayload{client: c}
}

// ==================== Schedule API ====================

// Schedule 创建定时任务构建器
func (c *Client) Schedule() *SchedulePayload {
	return &SchedulePayload{client: c}
}

// ==================== Device API ====================

// Device 创建设备管理构建器
func (c *Client) Device() *DevicePayload {
	return &DevicePayload{client: c}
}

// ==================== Report API ====================

// Report 创建报告构建器
func (c *Client) Report() *ReportPayload {
	return &ReportPayload{client: c}
}

// ==================== 工具方法 ====================

// ValidatePush 验证推送载荷
func (c *Client) ValidatePush(payload *PushPayload) (*PushResponse, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	url := GetURL("PUSH", c.Zone) + "/validate"
	respBody, _, err := c.doRequest(RequestOptions{
		Method: http.MethodPost,
		URL:    url,
		Body:   data,
	})
	if err != nil {
		return nil, err
	}

	var result PushResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// GetCID 获取推送 CID
func (c *Client) GetCID(count int, pushType string) ([]string, error) {
	url := GetURL("PUSH", c.Zone) + "/cid" // 注意：PUSH URL 已包含 /v3/push，这里只加 /cid
	respBody, _, err := c.doRequest(RequestOptions{
		Method: http.MethodGet,
		URL:    url,
		Params: map[string]string{
			"count": fmt.Sprintf("%d", count),
			"type":  pushType,
		},
	})
	if err != nil {
		return nil, err
	}

	var result struct {
		CIDList []string `json:"cidlist"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return result.CIDList, nil
}

// addQueryParams 添加查询参数到 URL
func addQueryParams(rawURL string, params map[string]string) string {
	if len(params) == 0 {
		return rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String()
}
