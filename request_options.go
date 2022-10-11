package zadig

import "github.com/hashicorp/go-retryablehttp"

// RequestOptionFunc 自定义请求
type RequestOptionFunc func(*retryablehttp.Request) error

// WithToken 带token请求
func WithToken(authType AuthType, token string) RequestOptionFunc {
	return func(req *retryablehttp.Request) error {
		switch authType {
		case OAuthToken:
			req.Header.Set("Authorization", "Bearer "+token)
		case PrivateToken:
			req.Header.Set("PRIVATE-TOKEN", token)
		}
		return nil
	}
}
