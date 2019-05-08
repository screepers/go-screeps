package screeps

import (
	"fmt"
)

// AuthSigninRequest AuthSigninRequest
type AuthSigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthSigninResponse AuthSigninResponse
type AuthSigninResponse struct {
	BaseResponse
	Token string
}

// Badge Screeps Badge format
// BUG(): Doesn't handle custom badges
type Badge struct {
	Type   int8
	Color1 string
	Color2 string
	Color3 string
	Param  int
	Flip   bool
}

// AuthMeResponse AuthMeResponse
type AuthMeResponse struct {
	BaseResponse    `json:",inline"`
	ID              string `json:"_id"`
	Email           string
	Username        string
	CPU             int
	Badge           Badge
	Password        bool
	LastRespawnDate int64
	NotifyPrefs     struct {
		ErrorsInterval int
	}
	GCL                  int64
	Credits              int64
	Subscription         bool
	LifetimeSubscription bool
	Power                int64
	Money                float64
	SubscriptionTokens   int
	CPUShard             map[string]int
	CPUShardUpdatedTime  int64
	Runtime              struct {
		IVM bool
	}
	PowerExperimentations    int
	PowerExperimentationTime int64
	Github                   struct {
		ID       string
		Username string
	}
	Steam struct {
		ID          string
		DisplayName string
		Ownership   []int
	}
}

// Token Screeps Token Info
// TODO: Add limited token fields
type Token struct {
	ID               string `json:"_id"`
	Full             bool
	Token            string
	NoRateLimitUntil int64
}

// QueryTokenResponse QueryTokenResponse
type QueryTokenResponse struct {
	BaseResponse
	Token Token
}

// AuthSignin POST /api/auth/signin
func (c *Client) AuthSignin() (*AuthSigninResponse, error) {
	if c.IsOfficial() {
		return nil, fmt.Errorf("Use a token for official")
	}
	resp, err := c.r.R().
		SetBody(AuthSigninRequest{
			Email:    c.config.Username,
			Password: c.config.Password,
		}).
		SetResult(AuthSigninResponse{}).
		Post("/api/auth/signin")
	if err != nil {
		return nil, err
	}
	if resp.String() == "Unauthorized" {
		return nil, fmt.Errorf("Unauthorized")
	}
	return resp.Result().(*AuthSigninResponse), nil
}

// AuthMe GET /api/auth/me
func (c *Client) AuthMe() (*AuthMeResponse, error) {
	resp, err := c.r.R().
		SetResult(AuthMeResponse{}).
		Get("/api/auth/me")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*AuthMeResponse), nil
}

// QueryToken GET /api/auth/query-token
func (c *Client) QueryToken(token string) (*QueryTokenResponse, error) {
	resp, err := c.r.R().
		SetQueryParams(map[string]string{
			"token": token,
		}).
		SetResult(QueryTokenResponse{}).
		Get("/api/auth/query-token")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*QueryTokenResponse), nil
}
