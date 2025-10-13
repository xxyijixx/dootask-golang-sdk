package file

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/xxyijixx/dootask-golang-sdk/types"
)

// DownloadPack 打包下载文件
// ids: 文件ID列表
// name: 下载文件名 (可选)
func (s *Service) DownloadPack(ids []int, name *string) (map[string]interface{}, error) {
	req := types.FileDownloadPackRequest{
		IDs: ids,
	}

	if name != nil {
		req.Name = *name
	}

	resp, err := s.client.DoRequest("GET", "/api/file/download/pack", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Ret  int                    `json:"ret"`
		Msg  string                 `json:"msg"`
		Data map[string]interface{} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Ret != 1 {
		return nil, fmt.Errorf("API error: %s", result.Msg)
	}

	return result.Data, nil
}

// DownloadPackFile 下载打包文件 (通过key)
// key: 下载密钥
func (s *Service) DownloadPackFile(key string) (*http.Response, error) {
	params := url.Values{}
	params.Set("key", key)

	resp, err := s.client.DoRequest("GET", "/api/file/download/pack?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
