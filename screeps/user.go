package screeps

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"strings"
)

// GetMemoryResponse Response for #GetMemory
type GetMemoryResponse struct {
	BaseResponse
	Data string
}

// Decompress - Decompresses gzip compressed data
func (mr *GetMemoryResponse) Decompress() error {
	if !strings.HasPrefix(mr.Data, "gz:") {
		return nil
	}
	sr := strings.NewReader(mr.Data)
	gr, err := gzip.NewReader(sr)
	if err != nil {
		return err
	}
	ret, err := ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	mr.Data = string(ret)
	return nil
}

// Parse - Parses JSON data into struct
func (mr *GetMemoryResponse) Parse(tgt *interface{}) error {
	err := mr.Decompress()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(mr.Data), tgt)
	if err != nil {
		return err
	}
	return nil
}

// GetMemory GET /api/user/memory
func (c *Client) GetMemory(path string, shard string) (*GetMemoryResponse, error) {
	if shard == "" {
		shard = c.DefaultShard
	}
	resp, err := c.r.R().
		SetResult(GetMemoryResponse{}).
		Get("/api/user/memory")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*GetMemoryResponse), nil
}

// GetMemorySegment GET /api/user/memory-segment
func (c *Client) GetMemorySegment(path string, shard string) (*GetMemoryResponse, error) {
	if shard == "" {
		shard = c.DefaultShard
	}
	resp, err := c.r.R().
		SetResult(GetMemoryResponse{}).
		Get("/api/user/memory-segment")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*GetMemoryResponse), nil
}
