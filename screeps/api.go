package screeps

import (
	"fmt"
)

// BaseResponse BaseResponse
type BaseResponse struct {
	Ok    int8
	Error string
}

// VersionResponse VersionResponse
type VersionResponse struct {
	BaseResponse
	Package       int
	Protocol      int
	UseNativeAuth bool
	Users         int
	ServerData    struct {
		WelcomeText          string
		HistoryChunkSize     int
		Shards               []string
		SocketUpdateThrottle int
		Renderer             interface{}
		CustomObjectTypes    interface{}
	}
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

// RoomHistoryResponse RoomHistoryResponse
type RoomHistoryResponse struct {
	Timestamp int64
	Room      string
	Base      int64
	Ticks     map[string]map[string]RoomObject
}

// RoomObject - Screeps RoomObject
type RoomObject struct {
	ID    string `json:"_id"`
	Room  string
	Type  string
	X     int8
	Y     int8
	Props map[string]interface{} `json:",inline"`
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

// RoomHistory GET /api/room-history
// Returns empty RoomHistoryResponse if history doesn't exist
func (c *Client) RoomHistory(room string, tick int64, shard string) (*RoomHistoryResponse, error) {
	if shard == "" {
		shard = c.DefaultShard
	}
	interval := 20
	if c.IsOfficial() {
		interval = 100
	} else {
		res, err := c.GetVersion()
		if err != nil {
			return nil, err
		}
		if res.ServerData.HistoryChunkSize > 0 {
			interval = res.ServerData.HistoryChunkSize
		}
	}
	tick -= tick % int64(interval)
	if c.IsOfficial() {
		url := fmt.Sprintf("/room-history/%s/%s/%d.json", shard, room, tick)
		resp, err := c.r.R().
			SetResult(RoomHistoryResponse{}).
			Get(url)
		if err != nil {
			return nil, err
		}
		return resp.Result().(*RoomHistoryResponse), nil
	}
	resp, err := c.r.R().
		SetQueryParams(map[string]string{
			"room":  room,
			"shard": shard,
			"tick":  string(tick),
		}).
		SetResult(RoomHistoryResponse{}).
		Get("/room-history")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*RoomHistoryResponse), nil
}
