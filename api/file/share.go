package file

import (
	"fmt"

	"github.com/xxyijixx/dootask-golang-sdk/internal/http"
	"github.com/xxyijixx/dootask-golang-sdk/types"
)

// Share 获取共享信息
// id: 文件ID
func (s *Service) Share(id int) (map[string]interface{}, error) {
	req := types.FileShareRequest{
		ID: id,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/share", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = http.ParseAPIResponse(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return result, nil
}

// ShareUpdate 设置共享
// id: 文件ID
// userIDs: 共享成员ID列表 (可选)
// permission: 共享方式 0只读 1读写 -1删除
// force: 忽略提醒 0不忽略 1忽略 (可选)
func (s *Service) ShareUpdate(id int, userIDs []int, permission int, force *int) (*types.File, error) {
	req := types.FileShareUpdateRequest{
		ID:         id,
		UserIDs:    userIDs,
		Permission: permission,
		Force:      force,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/share/update", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var file types.File
	err = http.ParseAPIResponse(resp, &file)
	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &file, nil
}

// ShareOut 退出共享
// id: 文件ID
func (s *Service) ShareOut(id int) error {
	req := types.FileShareOutRequest{
		ID: id,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/share/out", req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Ret int    `json:"ret"`
		Msg string `json:"msg"`
	}

	err = http.ParseAPIResponse(resp, &result)
	if err != nil {
		return fmt.Errorf("API error: %s", err.Error())
	}

	return nil
}
