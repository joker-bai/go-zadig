package zadig

// ClientOptionFunc 自定义配置
type ClientOptionFunc func(*Client) error

// WithBaseURL 设置baseURL
func WithBaseURL(urlStr string) ClientOptionFunc {
	return func(c *Client) error {
		return c.setBaseURL(urlStr)
	}
}
