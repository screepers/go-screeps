package screeps

import (
	"fmt"
	"net/http"
	"screepers/go-screeps/v1/config"

	"gopkg.in/resty.v1"
)

// Client Screeps API client
type Client struct {
	r      *resty.Client
	config *config.ServerConfig
	token  string
}

// NewClient creates a new client
func NewClient(config config.ServerConfig) *Client {
	c := &Client{
		config: &config,
		r:      resty.New(),
	}
	if config.Token != "" {
		c.token = config.Token
	}
	c.r.SetRetryCount(3)
	c.r.AddRetryCondition(
		// Condition function will be provided with *resty.Response as a
		// parameter. It is expected to return (bool, error) pair. Resty will retry
		// in case condition returns true or non nil error.
		func(r *resty.Response) (bool, error) {
			if r.StatusCode() == http.StatusUnauthorized {
				if c.IsOfficial() {
					return false, nil
				}
				if r.String() == "Unauthorized" {
					return false, nil
				}
				resp, err := c.AuthSignin()
				if err != nil {
					return false, nil
				}
				c.token = resp.Token
				return true, nil
			}
			return r.StatusCode() == http.StatusTooManyRequests, nil
		},
	)
	c.r.OnBeforeRequest(func(rc *resty.Client, req *resty.Request) error {
		if c.token != "" {
			req.SetHeader("X-Token", c.token)
			req.SetHeader("X-Username", c.token)
		}
		return nil
	})
	c.r.OnAfterResponse(func(rc *resty.Client, res *resty.Response) error {
		if token := res.Header().Get("X-Token"); token != "" {
			c.token = token
		}
		return nil
	})
	if config.Host == "" {
		config.Host = "screeps.com"
		config.Port = 443
		config.Secure = true
	}
	proto := "http"
	if config.Secure || config.Host == "screeps.com" {
		proto = "https"
	}
	path := ""
	if config.PTR {
		path = "/ptr"
	}
	c.r.SetHostURL(fmt.Sprintf("%s://%s:%d%s", proto, config.Host, config.Port, path))
	c.r.SetDebug(true)
	return c
}

// IsOfficial Returns true if server is screeps.com
func (c *Client) IsOfficial() bool {
	return c.config.Host == "screeps.com"
}
