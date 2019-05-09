package screeps

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
)

// GetMemorySegmentsResponse GetMemorySegmentsResponse
type GetMemorySegmentsResponse struct {
	BaseResponse
	Data []string
}

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
	data, err := base64.StdEncoding.DecodeString(mr.Data[3:])
	if err != nil {
		return err
	}
	br := bytes.NewReader(data)
	gr, err := gzip.NewReader(br)
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
		SetQueryParams(map[string]string{
			"path":  path,
			"shard": shard,
		}).
		SetResult(GetMemoryResponse{}).
		Get("/api/user/memory")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*GetMemoryResponse), nil
}

// GetMemorySegment GET /api/user/memory-segment
func (c *Client) GetMemorySegment(segment int, shard string) (*GetMemoryResponse, error) {
	if shard == "" {
		shard = c.DefaultShard
	}
	resp, err := c.r.R().
		SetQueryParams(map[string]string{
			"segment": strconv.Itoa(segment),
			"shard":   shard,
		}).
		SetResult(GetMemoryResponse{}).
		Get("/api/user/memory-segment")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*GetMemoryResponse), nil
}

// GetMemorySegments GET /api/user/memory-segment
func (c *Client) GetMemorySegments(segments []int, shard string) ([]GetMemoryResponse, error) {
	if shard == "" {
		shard = c.DefaultShard
	}
	ret := make([]GetMemoryResponse, len(segments))
	if c.IsOfficial() {
		segStrings := make([]string, len(segments))
		for i, seg := range segments {
			segStrings[i] = strconv.Itoa(seg)
		}
		resp, err := c.r.R().
			SetQueryParams(map[string]string{
				"segment": strings.Join(segStrings, ","),
				"shard":   shard,
			}).
			SetResult(GetMemorySegmentsResponse{}).
			Get("/api/user/memory-segment")
		if err != nil {
			return nil, err
		}
		res := resp.Result().(*GetMemorySegmentsResponse)
		for i, s := range res.Data {
			ret[i] = GetMemoryResponse{
				BaseResponse: res.BaseResponse,
				Data:         s,
			}
		}
	} else {
		for _, seg := range segments {
			mem, err := c.GetMemorySegment(seg, shard)
			if err != nil {
				return nil, err
			}
			ret = append(ret, *mem)
		}
	}
	return ret, nil
}
