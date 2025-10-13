package dialog

import (
	"fmt"

	"github.com/xxyijixx/dootask-golang-sdk/internal/core"
	ihttp "github.com/xxyijixx/dootask-golang-sdk/internal/http"
	"github.com/xxyijixx/dootask-golang-sdk/types"
)

// Service 对话服务
type Service struct {
	client core.HTTPDoer
}

// New 创建对话服务实例
func New(client core.HTTPDoer) *Service {
	return &Service{
		client: client,
	}
}

// CreateDialog 创建对话
func (s *Service) CreateDialog(req *types.CreateDialogRequest) (*types.Dialog, error) {
	resp, err := s.client.DoRequest("POST", "/api/dialog/create", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.Dialog
	err = ihttp.ParseAPIResponse(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetDialogList 获取对话列表
func (s *Service) GetDialogList(req *types.DialogListRequest) (*types.DialogListResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/lists", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DialogListResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetDialogDetail 获取对话详情
func (s *Service) GetDialogDetail(dialogID int) (*types.Dialog, error) {
	endpoint := fmt.Sprintf("/api/dialog/detail/%d", dialogID)
	resp, err := s.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.Dialog
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// UpdateDialog 更新对话
func (s *Service) UpdateDialog(dialogID int, req *types.UpdateDialogRequest) (*types.Dialog, error) {
	endpoint := fmt.Sprintf("/api/dialog/update/%d", dialogID)
	resp, err := s.client.DoRequest("PUT", endpoint, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.Dialog
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DeleteDialog 删除对话
func (s *Service) DeleteDialog(dialogID int) error {
	endpoint := fmt.Sprintf("/api/dialog/delete/%d", dialogID)
	resp, err := s.client.DoRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var ok struct{}
	if err := ihttp.ParseAPIResponse(resp, &ok); err != nil {
		return fmt.Errorf("API error: %s", err.Error())
	}

	return nil
}

// GetDialogMembers 获取对话成员
func (s *Service) GetDialogMembers(dialogID int) (*types.DialogMembersResponse, error) {
	endpoint := fmt.Sprintf("/api/dialog/members/%d", dialogID)
	resp, err := s.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DialogMembersResponse
	if err := ihttp.ParseAPIResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// AddDialogMembers 添加对话成员
func (s *Service) AddDialogMembers(dialogID int, req *types.AddMemberRequest) (*types.DialogMembersResponse, error) {
	endpoint := fmt.Sprintf("/api/dialog/members/%d/add", dialogID)
	resp, err := s.client.DoRequest("POST", endpoint, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DialogMembersResponse
	if err := ihttp.ParseAPIResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// RemoveDialogMembers 移除对话成员
func (s *Service) RemoveDialogMembers(dialogID int, req *types.RemoveMemberRequest) error {
	endpoint := fmt.Sprintf("/api/dialog/members/%d/remove", dialogID)
	resp, err := s.client.DoRequest("POST", endpoint, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var ok struct{}
	if err := ihttp.ParseAPIResponse(resp, &ok); err != nil {
		return fmt.Errorf("API error: %s", err.Error())
	}

	return nil
}

// SendMessage 发送消息
func (s *Service) SendMessage(req *types.SendMessageRequest) (*types.DialogMessageResponse, error) {
	endpoint := "/api/dialog/msg/sendtext"
	resp, err := s.client.DoRequest("POST", endpoint, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DialogMessageResponse
	if err := ihttp.ParseAPIResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetDialogMessages 获取对话消息
func (s *Service) GetDialogMessages(dialogID int, req *types.DialogMessagesRequest) (*types.DialogMessagesResponse, error) {
	endpoint := fmt.Sprintf("/api/dialog/%d/messages", dialogID)
	resp, err := s.client.DoRequest("GET", endpoint, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DialogMessagesResponse
	if err := ihttp.ParseAPIResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// MarkMessagesAsRead 标记消息为已读
func (s *Service) MarkMessagesAsRead(dialogID int, messageIDs []int) error {
	req := struct {
		MessageIDs []int `json:"message_ids"`
	}{
		MessageIDs: messageIDs,
	}

	endpoint := fmt.Sprintf("/api/dialog/%d/messages/read", dialogID)
	resp, err := s.client.DoRequest("POST", endpoint, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var ok struct{}
	if err := ihttp.ParseAPIResponse(resp, &ok); err != nil {
		return fmt.Errorf("API error: %s", err.Error())
	}

	return nil
}

// DeleteMessage 删除消息
func (s *Service) DeleteMessage(dialogID, messageID int) error {
	endpoint := fmt.Sprintf("/api/dialog/%d/message/%d/delete", dialogID, messageID)
	resp, err := s.client.DoRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var ok struct{}
	if err := ihttp.ParseAPIResponse(resp, &ok); err != nil {
		return fmt.Errorf("API error: %s", err.Error())
	}

	return nil
}

// SetDialogTop 设置对话置顶
func (s *Service) SetDialogTop(dialogID int, isTop bool) error {
	req := struct {
		IsTop int `json:"is_top"`
	}{
		IsTop: 0,
	}
	if isTop {
		req.IsTop = 1
	}

	endpoint := fmt.Sprintf("/api/dialog/%d/top", dialogID)
	resp, err := s.client.DoRequest("POST", endpoint, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var ok struct{}
	if err := ihttp.ParseAPIResponse(resp, &ok); err != nil {
		return fmt.Errorf("API error: %s", err.Error())
	}

	return nil
}

// SetDialogMute 设置对话静音
func (s *Service) SetDialogMute(dialogID int, isMute bool) error {
	req := struct {
		IsMute int `json:"is_mute"`
	}{
		IsMute: 0,
	}
	if isMute {
		req.IsMute = 1
	}

	endpoint := fmt.Sprintf("/api/dialog/%d/mute", dialogID)
	resp, err := s.client.DoRequest("POST", endpoint, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var ok struct{}
	if err := ihttp.ParseAPIResponse(resp, &ok); err != nil {
		return fmt.Errorf("API error: %s", err.Error())
	}

	return nil
}

// ArchiveDialog 归档对话
func (s *Service) ArchiveDialog(dialogID int, isArchive bool) error {
	req := struct {
		IsArchive int `json:"is_archive"`
	}{
		IsArchive: 0,
	}
	if isArchive {
		req.IsArchive = 1
	}

	endpoint := fmt.Sprintf("/api/dialog/%d/archive", dialogID)
	resp, err := s.client.DoRequest("POST", endpoint, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var ok struct{}
	if err := ihttp.ParseAPIResponse(resp, &ok); err != nil {
		return fmt.Errorf("API error: %s", err.Error())
	}

	return nil
}
