package zadig

import "github.com/hashicorp/go-retryablehttp"

// RequestOptionFunc can be passed to all API requests to customize the API request.
type RequestOptionFunc func(*retryablehttp.Request) error

// WithToken takes a token which is then used when making this one request.
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
