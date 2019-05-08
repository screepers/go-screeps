package screeps

// ServersListResponse Response for #ServersList
type ServersListResponse struct {
	BaseResponse
	Likes   []interface{}
	Servers []struct {
		ID        string `json:"_id"`
		LikeCount int
		Name      string
		Settings  struct {
			Host string
			Pass string
			Port string
		}
		Status string
	}
}

// ServersList POST /api/servers/list
func (c *Client) ServersList() (*ServersListResponse, error) {
	resp, err := c.r.R().
		SetResult(ServersListResponse{}).
		Post("/api/servers/list")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*ServersListResponse), nil
}
