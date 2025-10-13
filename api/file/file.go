package file

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/xxyijixx/dootask-golang-sdk/internal/core"
	"github.com/xxyijixx/dootask-golang-sdk/internal/http"

	"github.com/xxyijixx/dootask-golang-sdk/types"
)

// Service handles file-related API calls
type Service struct {
	client core.HTTPDoer
}

// New creates a new file service
func New(client core.HTTPDoer) *Service {
	return &Service{
		client: client,
	}
}

// Lists 获取文件列表
// pid: 父级ID (可选)
func (s *Service) Lists(pid *int) ([]types.File, error) {
	params := url.Values{}
	if pid != nil {
		params.Set("pid", strconv.Itoa(*pid))
	}

	resp, err := s.client.DoRequest("GET", "/api/file/lists?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var files []types.File
	err = http.ParseAPIResponse[[]types.File](resp, &files)
	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return files, nil
}

// One 获取单个文件信息
// id: 文件ID(int) 或 链接码(string)
func (s *Service) One(id interface{}) (*types.File, error) {
	params := url.Values{}

	switch v := id.(type) {
	case int:
		params.Set("id", strconv.Itoa(v))
	case string:
		params.Set("id", v)
	default:
		return nil, fmt.Errorf("invalid id type, must be int or string")
	}

	resp, err := s.client.DoRequest("GET", "/api/file/one?"+params.Encode(), nil)
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

// Search 搜索文件
// link: 通过分享地址搜索 (可选)
// key: 关键词 (可选)
// take: 获取数量，默认50，最大100 (可选)
func (s *Service) Search(link, key *string, take *int) ([]types.File, error) {
	params := url.Values{}
	if link != nil {
		params.Set("link", *link)
	}
	if key != nil {
		params.Set("key", *key)
	}
	if take != nil {
		params.Set("take", strconv.Itoa(*take))
	}

	resp, err := s.client.DoRequest("GET", "/api/file/search?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var files []types.File
	err = http.ParseAPIResponse(resp, &files)
	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return files, nil
}

// Add 添加或修改文件(夹)
// name: 文件名称
// fileType: 文件类型 (folder, document, mind, drawio, word, excel, ppt)
// id: 文件ID (修改时使用，可选)
// pid: 父级ID (可选)
func (s *Service) Add(name, fileType string, id, pid *int) (*types.File, error) {
	req := types.FileAddRequest{
		Name: name,
		Type: fileType,
		ID:   id,
		PID:  pid,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/add", req)
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

// Copy 复制文件(夹)
// id: 文件ID
func (s *Service) Copy(id int) (*types.File, error) {
	req := types.FileCopyRequest{
		ID: id,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/copy", req)
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

// Move 移动文件(夹)
// ids: 文件ID列表
// pid: 移动到的文件夹ID
func (s *Service) Move(ids []int, pid int) ([]types.File, error) {
	req := types.FileMoveRequest{
		IDs: ids,
		PID: pid,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/move", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var files []types.File
	err = http.ParseAPIResponse(resp, &files)
	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return files, nil
}

// Remove 删除文件(夹)
// ids: 文件ID列表
func (s *Service) Remove(ids []int) ([]types.File, error) {
	req := types.FileRemoveRequest{
		IDs: ids,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/remove", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var files []types.File
	err = http.ParseAPIResponse(resp, &files)
	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return files, nil
}
