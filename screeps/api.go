package screeps

// BaseResponse BaseResponse
type BaseResponse struct {
	Ok    int8
	Error string
}

// VersionResponse VersionResponse
type VersionResponse struct {
	BaseResponse
	Package    int
	Protocol   int
	ServerData struct {
		HistoryChunkSize int
		Shards           []string
	}
	Users int
}

// AuthmodResponse AuthmodResponse
type AuthmodResponse struct {
	BaseResponse
	Name              string
	Version           string
	AllowRegistration bool
	Steam             bool
	Github            string
}

// Version GET /api/version
func (c *Client) Version() (*VersionResponse, error) {
	resp, err := c.r.R().
		SetResult(VersionResponse{}).
		Get("/api/version")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*VersionResponse), nil
}

// Authmod GET /api/authmod
func (c *Client) Authmod() (*AuthmodResponse, error) {
	if c.IsOfficial() {
		return &AuthmodResponse{
			BaseResponse:      BaseResponse{Ok: 1},
			Name:              "official",
			Steam:             true,
			AllowRegistration: true,
		}, nil
	}
	resp, err := c.r.R().
		SetResult(AuthmodResponse{}).
		Get("/api/authmod")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*AuthmodResponse), nil
}
