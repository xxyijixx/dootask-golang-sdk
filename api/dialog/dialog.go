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

// GetDialogList 01.对话列表
func (s *Service) GetDialogList(req *types.DialogListsRequest) (*types.DialogListsResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/lists", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DialogListsResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SearchDialog 2.搜索会话
func (s *Service) SearchDialog(req *types.SearchDialogRequest) (*types.SearchDialogResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/search", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.SearchDialogResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetDialogDetail 3.获取单个会话信息
func (s *Service) GetDialogDetail(req *types.DialogOneRequest) (*types.DialogOneResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/one", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DialogOneResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetDialogMembers 4.获取会话成员
func (s *Service) GetDialogMembers(req *types.DialogUserRequest) (*types.DialogUserResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/user", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DialogUserResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetDialogTodo 5.获取会话待办
func (s *Service) GetDialogTodo(req *types.DialogTodoRequest) (*types.DialogTodoResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/todo", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DialogTodoResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// TopDialog 6.会话置顶
func (s *Service) TopDialog(req *types.DialogTopRequest) (*types.DialogTopResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/top", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DialogTopResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetDialogTel 7.获取对方联系电话
func (s *Service) GetDialogTel(req *types.DialogTelRequest) (*types.DialogTelResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/tel", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DialogTelResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// OpenDialog 8.打开会话
func (s *Service) OpenDialog(req *types.OpenDialogRequest) (*types.OpenDialogResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/open/user", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.OpenDialogResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetMessageList 9.获取消息列表
func (s *Service) GetMessageList(req *types.MessageListRequest) (*types.MessageListResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/list", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.MessageListResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SearchMessage 10.搜索消息位置
func (s *Service) SearchMessage(req *types.SearchMessageRequest) (*types.SearchMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/search", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.SearchMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetMessageOne 11.获取单条消息
func (s *Service) GetMessageOne(req *types.MessageOneRequest) (*types.MessageOneResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/one", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.MessageOneResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ReadMessage 12.已读聊天消息
func (s *Service) ReadMessage(req *types.ReadMessageRequest) (*types.ReadMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/read", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReadMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetUnreadMessage 13.获取未读消息数据
func (s *Service) GetUnreadMessage(req *types.UnreadMessageRequest) (*types.UnreadMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/unread", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.UnreadMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// StreamMessage 14.通知成员监听消息
func (s *Service) StreamMessage(req *types.StreamMessageRequest) (*types.StreamMessageResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/dialog/msg/stream", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.StreamMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SendMessage 15.发送消息
func (s *Service) SendMessage(req *types.SendMessageRequest) (*types.SendMessageResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/dialog/msg/sendtext", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.SendMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SendRecord 16.发送语音
func (s *Service) SendRecord(req *types.SendRecordRequest) (*types.SendRecordResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/dialog/msg/sendrecord", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.SendRecordResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SendFile 17.文件上传
func (s *Service) SendFile(req *types.SendFileRequest) (*types.SendFileResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/dialog/msg/sendfile", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.SendFileResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SendFiles 18.群发文件上传
func (s *Service) SendFiles(req *types.SendFilesRequest) (*types.SendFilesResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/dialog/msg/sendfiles", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.SendFilesResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SendFileId 19.通过文件ID发送文件
func (s *Service) SendFileId(req *types.SendFileIDRequest) (*types.SendFileIDResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/sendfileid", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.SendFileIDResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SendAnonMessage 20.发送匿名消息
func (s *Service) SendAnonMessage(req *types.SendAnonRequest) (*types.SendAnonResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/dialog/msg/sendanon", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.SendAnonResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetMessageReadList 21.获取消息阅读情况
func (s *Service) GetMessageReadList(req *types.MessageReadListRequest) (*types.MessageReadListResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/readlist", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.MessageReadListResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetMessageDetail 22.消息详情
func (s *Service) GetMessageDetail(req *types.MessageDetailRequest) (*types.MessageDetailResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/detail", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.MessageDetailResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DownloadFile 23.文件下载
func (s *Service) DownloadFile(req *types.DownloadFileRequest) (*types.DownloadFileResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/download", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DownloadFileResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// WithdrawMessage 24.聊天消息撤回
func (s *Service) WithdrawMessage(req *types.WithdrawMessageRequest) (*types.WithdrawMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/withdraw", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.WithdrawMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// MarkMessage 25.消息标记操作
func (s *Service) MarkMessage(req *types.MarkMessageRequest) (*types.MarkMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/mark", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.MarkMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SilenceMessage 26.消息免打扰
func (s *Service) SilenceMessage(req *types.SilenceMessageRequest) (*types.SilenceMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/silence", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.SilenceMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ForwardMessage 27.转发消息给
func (s *Service) ForwardMessage(req *types.ForwardMessageRequest) (*types.ForwardMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/forward", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ForwardMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// EmojiMessage 28.emoji回复
func (s *Service) EmojiMessage(req *types.EmojiMessageRequest) (*types.EmojiMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/emoji", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.EmojiMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// TagMessage 29.标注/取消标注
func (s *Service) TagMessage(req *types.TagMessageRequest) (*types.TagMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/tag", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.TagMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// TodoMessage 30.设待办/取消待办
func (s *Service) TodoMessage(req *types.TodoMessageRequest) (*types.TodoMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/todo", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.TodoMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetTodoListMessage 31.获取消息待办情况
func (s *Service) GetTodoListMessage(req *types.TodoListMessageRequest) (*types.TodoListMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/todolist", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.TodoListMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DoneTodo 32.完成待办
func (s *Service) DoneTodo(req *types.DoneTodoRequest) (*types.DoneTodoResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/done", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DoneTodoResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ColorMessage 33.设置颜色
func (s *Service) ColorMessage(req *types.ColorMessageRequest) (*types.ColorMessageResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/msg/color", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ColorMessageResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// CreateGroup 34.新增群组
func (s *Service) CreateGroup(req *types.CreateGroupRequest) (*types.CreateGroupResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/group/add", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.CreateGroupResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// EditGroup 35.修改群组
func (s *Service) EditGroup(req *types.EditGroupRequest) (*types.EditGroupResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/group/edit", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.EditGroupResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// AddGroupUser 36.添加群成员
func (s *Service) AddGroupUser(req *types.AddGroupUserRequest) (*types.AddGroupUserResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/group/adduser", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.AddGroupUserResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DelGroupUser 37.移出（退出）群成员
func (s *Service) DelGroupUser(req *types.DelGroupUserRequest) (*types.DelGroupUserResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/group/deluser", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DelGroupUserResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// TransferGroup 38.转让群组
func (s *Service) TransferGroup(req *types.TransferGroupRequest) (*types.TransferGroupResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/group/transfer", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.TransferGroupResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DisbandGroup 39.解散群组
func (s *Service) DisbandGroup(req *types.DisbandGroupRequest) (*types.DisbandGroupResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/group/disband", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.DisbandGroupResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SearchGroupUser 40.搜索个人群（仅限管理员）
func (s *Service) SearchGroupUser(req *types.SearchGroupUserRequest) (*types.SearchGroupUserResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/dialog/group/searchuser", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.SearchGroupUserResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// CreateOkrDialog 41.创建OKR评论会话
func (s *Service) CreateOkrDialog(req *types.CreateOkrDialogRequest) (*types.CreateOkrDialogResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/dialog/okr/add", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.CreateOkrDialogResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// PushOkrInfo 42.推送OKR相关信息
func (s *Service) PushOkrInfo(req *types.PushOkrInfoRequest) (*types.PushOkrInfoResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/dialog/okr/push", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.PushOkrInfoResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}
