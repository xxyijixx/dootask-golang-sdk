package users

import (
	"fmt"

	"github.com/xxyijixx/dootask-golang-sdk/internal/core"
	ihttp "github.com/xxyijixx/dootask-golang-sdk/internal/http"
	"github.com/xxyijixx/dootask-golang-sdk/types"
)

// Service 用户服务
type Service struct {
	client core.HTTPDoer
}

// New 创建用户服务实例
func New(client core.HTTPDoer) *Service {
	return &Service{
		client: client,
	}
}

// GetAnnualReport 34.年度报告
func (s *Service) GetAnnualReport(req *types.UserAnnualReportRequest) (*types.UserAnnualReportResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/annual/report", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserAnnualReportResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetAppSort 35.获取个人应用排序
func (s *Service) GetAppSort(req *types.UserAppSortRequest) (*types.UserAppSortResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/appsort", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserAppSortResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SaveAppSort 36.保存个人应用排序
func (s *Service) SaveAppSort(req *types.UserAppSortSaveRequest) (*types.UserAppSortSaveResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/users/appsort/save", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserAppSortSaveResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DeleteBot 37.删除机器人
func (s *Service) DeleteBot(req *types.UserBotDeleteRequest) (*types.UserBotDeleteResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/bot/delete", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserBotDeleteResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetBotList 38.机器人列表
func (s *Service) GetBotList(req *types.UserBotListRequest) (*types.UserBotListResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/bot/list", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserBotListResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SyncDepartment 39.同步部门成员（限管理员）
func (s *Service) SyncDepartment(req *types.UserDepartmentSyncRequest) (*types.UserDepartmentSyncResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/department/sync", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserDepartmentSyncResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// EditDevice 40.编辑设备
func (s *Service) EditDevice(req *types.UserDeviceEditRequest) (*types.UserDeviceEditResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/device/edit", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserDeviceEditResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetDeviceList 41.获取设备列表
func (s *Service) GetDeviceList(req *types.UserDeviceListRequest) (*types.UserDeviceListResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/device/list", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserDeviceListResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// LogoutDevice 42.登出设备（删除设备）
func (s *Service) LogoutDevice(req *types.UserDeviceLogoutRequest) (*types.UserDeviceLogoutResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/device/logout", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserDeviceLogoutResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetExtra 43.获取会员扩展信息
func (s *Service) GetExtra(req *types.UserExtraRequest) (*types.UserExtraResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/extra", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserExtraResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// CheckFavorite 44.检查收藏状态
func (s *Service) CheckFavorite(req *types.UserFavoriteCheckRequest) (*types.UserFavoriteCheckResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/favorite/check", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserFavoriteCheckResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// UpdateFavoriteRemark 45.修改收藏备注
func (s *Service) UpdateFavoriteRemark(req *types.UserFavoriteRemarkRequest) (*types.UserFavoriteRemarkResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/users/favorite/remark", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserFavoriteRemarkResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ToggleFavorite 46.切换收藏状态
func (s *Service) ToggleFavorite(req *types.UserFavoriteToggleRequest) (*types.UserFavoriteToggleResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/users/favorite/toggle", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserFavoriteToggleResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetFavorites 47.获取用户收藏列表
func (s *Service) GetFavorites(req *types.UserFavoritesRequest) (*types.UserFavoritesResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/favorites", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserFavoritesResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// CleanFavorites 48.清理用户收藏
func (s *Service) CleanFavorites(req *types.UserFavoritesCleanRequest) (*types.UserFavoritesCleanResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/users/favorites/clean", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserFavoritesCleanResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetInfoDepartments 49.获取我的部门列表
func (s *Service) GetInfoDepartments(req *types.UserInfoDepartmentsRequest) (*types.UserInfoDepartmentsResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/info/departments", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserInfoDepartmentsResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// Logout 50.退出登录
func (s *Service) Logout(req *types.UserLogoutRequest) (*types.UserLogoutResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/logout", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserLogoutResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetRecentBrowse 51.获取最近访问记录
func (s *Service) GetRecentBrowse(req *types.UserRecentBrowseRequest) (*types.UserRecentBrowseResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/recent/browse", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserRecentBrowseResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DeleteRecent 52.删除最近访问记录
func (s *Service) DeleteRecent(req *types.UserRecentDeleteRequest) (*types.UserRecentDeleteResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/users/recent/delete", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserRecentDeleteResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SearchAI 53.获取AI机器人
func (s *Service) SearchAI(req *types.UserSearchAIRequest) (*types.UserSearchAIResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/search/ai", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserSearchAIResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// AddTag 54.新增个性标签
func (s *Service) AddTag(req *types.UserTagsAddRequest) (*types.UserTagsAddResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/users/tags/add", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserTagsAddResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DeleteTag 55.删除个性标签
func (s *Service) DeleteTag(req *types.UserTagsDeleteRequest) (*types.UserTagsDeleteResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/users/tags/delete", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserTagsDeleteResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetTagsLists 56.获取个性标签列表
func (s *Service) GetTagsLists(req *types.UserTagsListsRequest) (*types.UserTagsListsResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/tags/lists", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserTagsListsResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// RecognizeTag 57.认可个性标签
func (s *Service) RecognizeTag(req *types.UserTagsRecognizeRequest) (*types.UserTagsRecognizeResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/users/tags/recognize", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserTagsRecognizeResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// UpdateTag 58.修改个性标签
func (s *Service) UpdateTag(req *types.UserTagsUpdateRequest) (*types.UserTagsUpdateResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/users/tags/update", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserTagsUpdateResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetTaskBrowse 59.获取任务浏览历史
func (s *Service) GetTaskBrowse(req *types.UserTaskBrowseRequest) (*types.UserTaskBrowseResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/task/browse", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserTaskBrowseResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// CleanTaskBrowse 60.清理任务浏览历史
func (s *Service) CleanTaskBrowse(req *types.UserTaskBrowseCleanRequest) (*types.UserTaskBrowseCleanResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/users/task/browse_clean", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserTaskBrowseCleanResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SaveTaskBrowse 61.记录任务浏览历史
func (s *Service) SaveTaskBrowse(req *types.UserTaskBrowseSaveRequest) (*types.UserTaskBrowseSaveResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/task/browse_save", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserTaskBrowseSaveResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetTokenExpire 62.查询 token 过期时间
func (s *Service) GetTokenExpire(req *types.UserTokenExpireRequest) (*types.UserTokenExpireResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/users/token/expire", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UserTokenExpireResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

