package file

import (
	"encoding/json"
	"fmt"
	"io"
	nethtp "net/http"
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

// ============================================================
// 📁 文件管理核心接口
// ============================================================

// ============================================================
// 文件浏览与搜索
// ============================================================

// Lists GET 01. 获取文件列表
// 按父级目录浏览文件
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

// One GET 02. 获取单条数据
// 支持文件ID或链接码获取
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

// Search GET 03. 搜索文件列表
// 支持关键词搜索和分享链接搜索
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

// ============================================================
// 文件操作
// ============================================================

// Add GET 04. 添加、修改文件(夹)
// 创建文件夹或重命名文件
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

// Copy GET 05. 复制文件(夹)
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

// Move GET 06. 移动文件(夹)
// 批量移动文件到指定文件夹
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

// Remove GET 07. 删除文件(夹)
// 批量删除文件
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

// ============================================================
// 📄 文件内容管理
// ============================================================

// ============================================================
// 内容读写
// ============================================================

// Content GET 08. 获取文件内容
// 支持下载、预览模式
// 支持历史版本读取
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
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		var result struct {
			Ret  int    `json:"ret"`
			Msg  string `json:"msg"`
			Data json.RawMessage `json:"data"`
		}

		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}

		if result.Ret != 1 {
			return nil, fmt.Errorf("API error: %s", result.Msg)
		}

		// 尝试解析为数组（取第一个元素）
		var arr []struct {
			ID       int    `json:"id"`
			UpdateAt string `json:"update_at"`
		}
		if err := json.Unmarshal(result.Data, &arr); err == nil && len(arr) > 0 {
			return arr[0], nil
		}

		// 尝试解析为对象
		var obj struct {
			ID       int    `json:"id"`
			UpdateAt string `json:"update_at"`
		}
		if err := json.Unmarshal(result.Data, &obj); err == nil {
			return obj, nil
		}

		return nil, fmt.Errorf("cannot unmarshal FileContent onlyUpdateAt response")
	}

	// 返回文件内容（可能是文件数据、文本内容等）
	// 这里返回原始响应，让调用者自己处理
	return resp, nil
}

// ContentSave GET 09. 保存文件内容
// 通过Request Payload提交内容
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

// ============================================================
// Office文档协作
// ============================================================

// OfficeToken GET 10. 获取token
// 获取Office协作token
// id: 文件ID
func (s *Service) OfficeToken(id int) (map[string]interface{}, error) {
	params := url.Values{}
	params.Set("id", strconv.Itoa(id))

	resp, err := s.client.DoRequest("GET", "/api/file/office/token?"+params.Encode(), nil)
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

// ContentOffice GET 11. 保存文件内容（office）
// Office文档保存
// id: 文件ID
// content: Office内容数据
func (s *Service) ContentOffice(id int, content interface{}) (map[string]interface{}, error) {
	req := map[string]interface{}{
		"id":      id,
		"content": content,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/content/office", req)
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

// ============================================================
// 文件上传
// ============================================================

// ContentUpload GET 12. 保存文件内容（上传文件）
// 文件上传功能
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

// ============================================================
// ⏱️ 版本历史管理
// ============================================================

// ContentHistory GET 13. 获取内容历史
// 查看文件修改历史记录
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

// ContentRestore GET 14. 恢复文件历史
// 恢复到指定历史版本
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

// ============================================================
// 🔗 共享与协作
// ============================================================

// ============================================================
// 文件共享管理
// ============================================================

// Share GET 15. 获取共享信息
// 查看文件共享状态
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

// ShareUpdate GET 16. 设置共享
// 设置共享成员和权限（只读/读写）
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

// ShareOut GET 17. 退出共享
// 退出他人共享的文件
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

// ============================================================
// 链接分享
// ============================================================

// Link GET 18. 获取链接
// 生成分享链接，支持刷新
// id: 文件ID
// refresh: 刷新链接 no不刷新 yes刷新 (默认no)
// guestAccess: 是否允许游客访问 no不允许 yes允许 (默认no)
func (s *Service) Link(id int, refresh, guestAccess string) (*types.FileLink, error) {
	req := types.FileLinkRequest{
		ID:          id,
		Refresh:     refresh,
		GuestAccess: guestAccess,
	}

	resp, err := s.client.DoRequest("GET", "/api/file/link", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var fileLink types.FileLink
	err = http.ParseAPIResponse(resp, &fileLink)
	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &fileLink, nil
}

// ============================================================
// 📦 下载管理
// ============================================================

// DownloadPack GET 19. 打包文件
// 批量文件打包下载
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

	var result map[string]interface{}
	err = http.ParseAPIResponse(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return result, nil
}

// DownloadConfirm GET 20. 确认下载
// 下载确认
// key: 下载密钥
func (s *Service) DownloadConfirm(key string) (*nethtp.Response, error) {
	params := url.Values{}
	params.Set("key", key)

	resp, err := s.client.DoRequest("GET", "/api/file/download/confirm?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
