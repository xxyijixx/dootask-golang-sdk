package file

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/xxyijixx/dootask-golang-sdk/internal/http"
	"github.com/xxyijixx/dootask-golang-sdk/types"
)

// Content 获取文件内容
// id: 文件ID(int) 或 链接码(string)
// onlyUpdateAt: 仅获取update_at字段 (yes/no, 可选)
// down: 下载模式 (no/yes/preview, 可选)
// historyID: 读取历史记录ID (可选)
func (s *Service) Content(id interface{}, onlyUpdateAt, down *string, historyID *int) (interface{}, error) {
	params := url.Values{}

	switch v := id.(type) {
	case int:
		params.Set("id", strconv.Itoa(v))
	case string:
		params.Set("id", v)
	default:
		return nil, fmt.Errorf("invalid id type, must be int or string")
	}

	if onlyUpdateAt != nil {
		params.Set("only_update_at", *onlyUpdateAt)
	}
	if down != nil {
		params.Set("down", *down)
	}
	if historyID != nil {
		params.Set("history_id", strconv.Itoa(*historyID))
	}

	resp, err := s.client.DoRequest("GET", "/api/file/content?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 如果是预览或下载，直接返回响应
	if down != nil && (*down == "yes" || *down == "preview") {
		return resp, nil
	}

	// 仅获取更新时间
	if onlyUpdateAt != nil && *onlyUpdateAt == "yes" {
		var result struct {
			Ret  int    `json:"ret"`
			Msg  string `json:"msg"`
			Data struct {
				ID       int    `json:"id"`
				UpdateAt string `json:"update_at"`
			} `json:"data"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, err
		}

		if result.Ret != 1 {
			return nil, fmt.Errorf("API error: %s", result.Msg)
		}

		return result.Data, nil
	}

	// 返回文件内容（可能是文件数据、文本内容等）
	// 这里返回原始响应，让调用者自己处理
	return resp, nil
}

// ContentSave 保存文件内容
// id: 文件ID
// content: 内容数据
func (s *Service) ContentSave(id int, content interface{}) (*types.FileContent, error) {
	req := types.FileContentSaveRequest{
		ID:      id,
		Content: content,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/content/save", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var fileContent types.FileContent
	err = http.ParseAPIResponse(resp, &fileContent)
	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &fileContent, nil
}

// ContentUpload 上传文件
// pid: 父级ID (可选)
// cover: 覆盖已存在的文件 0不覆盖 1覆盖 (可选)
// webkitRelativePath: 相对路径 (可选)
func (s *Service) ContentUpload(pid, cover *int, webkitRelativePath *string) (map[string]interface{}, error) {
	req := types.FileContentUploadRequest{
		PID:                pid,
		Cover:              cover,
		WebkitRelativePath: webkitRelativePath,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/content/upload", req)
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

// ContentHistory 获取文件内容历史
// id: 文件ID
// page: 当前页，默认1 (可选)
// pageSize: 每页显示数量，默认20，最大100 (可选)
func (s *Service) ContentHistory(id int, page, pageSize *int) (map[string]interface{}, error) {
	params := url.Values{}
	params.Set("id", strconv.Itoa(id))

	if page != nil {
		params.Set("page", strconv.Itoa(*page))
	}
	if pageSize != nil {
		params.Set("pagesize", strconv.Itoa(*pageSize))
	}

	resp, err := s.client.DoRequest("GET", "/api/file/content/history?"+params.Encode(), nil)
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

// ContentRestore 恢复文件历史版本
// id: 文件ID
// historyID: 历史数据ID
func (s *Service) ContentRestore(id, historyID int) error {
	req := types.FileContentRestoreRequest{
		ID:        id,
		HistoryID: historyID,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/content/restore", req)
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
