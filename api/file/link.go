package file

import (
	"fmt"

	"github.com/xxyijixx/dootask-golang-sdk/internal/http"
	"github.com/xxyijixx/dootask-golang-sdk/types"
)

// Link 获取文件链接
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
