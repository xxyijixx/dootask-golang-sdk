package types

import (
	"encoding/json"
	"fmt"
)

// ==================== 用户相关类型定义 ====================

// 34. 年度报告
type UserAnnualReportRequest struct {
	Year int `json:"year,omitempty"` // 年份，不传则当前年份
}

type UserAnnualReportResponse struct {
	Data interface{} `json:"data"` // 年度报告数据
}

// 35. 获取个人应用排序
type UserAppSortRequest struct{}

type UserAppSortResponse struct {
	Sorts map[string]interface{} `json:"sorts"` // 排序配置
}

// 36. 保存个人应用排序
type UserAppSortSaveRequest struct {
	Sorts map[string]interface{} `json:"sorts"` // 排序配置
}

type UserAppSortSaveResponse struct {
	Sorts map[string]interface{} `json:"sorts"` // 保存后的排序配置
}

// 37. 删除机器人
type UserBotDeleteRequest struct {
	BotID int `json:"bot_id"` // 机器人ID
}

type UserBotDeleteResponse struct {
	Success bool `json:"success"`
}

// 38. 机器人列表
type UserBotListRequest struct {
	Page     int `json:"page,omitempty"`     // 页码，默认1
	PageSize int `json:"pagesize,omitempty"` // 每页数量，默认20
}

type UserBotListResponse struct {
	CurrentPage  int           `json:"current_page"`
	Data         []UserBotItem `json:"data"`
	FirstPageUrl *string       `json:"first_page_url"`
	From         *int          `json:"from"`
	LastPage     int           `json:"last_page"`
	LastPageUrl  *string       `json:"last_page_url"`
	Links        []interface{} `json:"links"`
	NextPageUrl  *string       `json:"next_page_url"`
	Path         string        `json:"path"`
	PerPage      int           `json:"per_page"`
	PrevPageUrl  *string       `json:"prev_page_url"`
	To           *int          `json:"to"`
	Total        int           `json:"total"`
}

type UserBotItem struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userid"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
}

// 39. 同步部门成员（限管理员）
type UserDepartmentSyncRequest struct{}

type UserDepartmentSyncResponse struct {
	Success bool `json:"success"`
}

// 40. 编辑设备
type UserDeviceEditRequest struct {
	DeviceID int    `json:"device_id"` // 设备ID
	Name     string `json:"name"`      // 设备名称
}

type UserDeviceEditResponse struct {
	Success bool `json:"success"`
}

// 41. 获取设备列表
type UserDeviceListRequest struct{}

// UserDeviceListResponse 支持数组和对象两种格式
type UserDeviceListResponse []UserDeviceItem

// UnmarshalJSON 自定义解析，支持数组和对象格式
func (u *UserDeviceListResponse) UnmarshalJSON(data []byte) error {
	// 尝试解析为数组
	var arr []UserDeviceItem
	if err := json.Unmarshal(data, &arr); err == nil {
		*u = arr
		return nil
	}
	// 尝试解析为对象（单个元素）
	var obj UserDeviceItem
	if err := json.Unmarshal(data, &obj); err == nil {
		*u = []UserDeviceItem{obj}
		return nil
	}
	return fmt.Errorf("cannot unmarshal UserDeviceListResponse")
}

type UserDeviceItem struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userid"`
	Name      string `json:"name"`
	Device    string `json:"device"`
	IP        string `json:"ip"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// 42. 登出设备（删除设备）
type UserDeviceLogoutRequest struct {
	DeviceID int `json:"device_id"` // 设备ID
}

type UserDeviceLogoutResponse struct {
	Success bool `json:"success"`
}

// 43. 获取会员扩展信息
type UserExtraRequest struct {
	UserID int `json:"user_id,omitempty"` // 用户ID，不传则当前用户
}

type UserExtraResponse interface{} // 扩展信息，可能是任意结构

// 44. 检查收藏状态
type UserFavoriteCheckRequest struct {
	Type string `json:"type"` // 收藏类型 (task/project/file/message)
	ID   int    `json:"id"`   // 收藏对象ID
}

type UserFavoriteCheckResponse struct {
	IsFavorite bool `json:"is_favorite"` // 是否已收藏
}

// 45. 修改收藏备注
type UserFavoriteRemarkRequest struct {
	Type   string `json:"type"`   // 收藏类型 (task/project/file/message)
	ID     int    `json:"id"`     // 收藏对象ID
	Remark string `json:"remark"` // 备注
}

type UserFavoriteRemarkResponse struct {
	Success bool `json:"success"`
}

// 46. 切换收藏状态
type UserFavoriteToggleRequest struct {
	Type string `json:"type"` // 收藏类型 (task/project/file/message)
	ID   int    `json:"id"`   // 收藏对象ID
}

type UserFavoriteToggleResponse struct {
	IsFavorite bool `json:"is_favorite"` // 切换后的收藏状态
}

// 47. 获取用户收藏列表
type UserFavoritesRequest struct {
	Type     string `json:"type,omitempty"`     // 收藏类型过滤 (task/project/file/message)
	Page     int    `json:"page,omitempty"`     // 页码，默认1
	PageSize int    `json:"pagesize,omitempty"` // 每页数量，默认20
}

type UserFavoritesResponse struct {
	CurrentPage  int                `json:"current_page"`
	Data         []UserFavoriteItem `json:"data"`
	FirstPageUrl *string            `json:"first_page_url"`
	From         *int               `json:"from"`
	LastPage     int                `json:"last_page"`
	LastPageUrl  *string            `json:"last_page_url"`
	Links        []interface{}      `json:"links"`
	NextPageUrl  *string            `json:"next_page_url"`
	Path         string             `json:"path"`
	PerPage      int                `json:"per_page"`
	PrevPageUrl  *string            `json:"prev_page_url"`
	To           *int               `json:"to"`
	Total        int                `json:"total"`
}

type UserFavoriteItem struct {
	ID        int     `json:"id"`
	Type      string  `json:"type"`
	TargetID  int     `json:"target_id"`
	Remark    *string `json:"remark"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	// 根据类型不同，可能包含不同的关联数据
	Task    interface{} `json:"task,omitempty"`
	Project interface{} `json:"project,omitempty"`
	File    interface{} `json:"file,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

// 48. 清理用户收藏
type UserFavoritesCleanRequest struct {
	Type string `json:"type,omitempty"` // 收藏类型，不传则清理所有
}

type UserFavoritesCleanResponse struct {
	Success bool `json:"success"`
}

// 49. 获取我的部门列表
type UserInfoDepartmentsRequest struct{}

type UserInfoDepartmentsResponse []UserDepartmentItem

type UserDepartmentItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id"`
}

// 50. 退出登录
type UserLogoutRequest struct{}

type UserLogoutResponse struct {
	Success bool        `json:"success,omitempty"`
	Data    interface{} `json:"data,omitempty"` // 可能是字符串或对象
}

// 51. 获取最近访问记录
type UserRecentBrowseRequest struct {
	Type     string `json:"type,omitempty"`     // 类型过滤 (task/project/file)
	Page     int    `json:"page,omitempty"`     // 页码，默认1
	PageSize int    `json:"pagesize,omitempty"` // 每页数量，默认20
}

type UserRecentBrowseResponse struct {
	CurrentPage  int              `json:"current_page"`
	Data         []UserRecentItem `json:"data"`
	FirstPageUrl *string          `json:"first_page_url"`
	From         *int             `json:"from"`
	LastPage     int              `json:"last_page"`
	LastPageUrl  *string          `json:"last_page_url"`
	Links        []interface{}    `json:"links"`
	NextPageUrl  *string          `json:"next_page_url"`
	Path         string           `json:"path"`
	PerPage      int              `json:"per_page"`
	PrevPageUrl  *string          `json:"prev_page_url"`
	To           *int             `json:"to"`
	Total        int              `json:"total"`
}

type UserRecentItem struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	TargetID  int    `json:"target_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	// 根据类型不同，可能包含不同的关联数据
	Task    interface{} `json:"task,omitempty"`
	Project interface{} `json:"project,omitempty"`
	File    interface{} `json:"file,omitempty"`
}

// 52. 删除最近访问记录
type UserRecentDeleteRequest struct {
	ID int `json:"id"` // 记录ID
}

type UserRecentDeleteResponse struct {
	Success bool `json:"success"`
}

// 53. 获取AI机器人
type UserSearchAIRequest struct{}

type UserSearchAIResponse []UserBotItem

// 54. 新增个性标签
type UserTagsAddRequest struct {
	Name string `json:"name"` // 标签名称
}

// UserTagsAddResponse 支持数组和对象两种格式
type UserTagsAddResponse struct {
	ID   int    `json:"id"`   // 标签ID
	Name string `json:"name"` // 标签名称
}

// UnmarshalJSON 自定义解析，支持数组和对象格式
func (u *UserTagsAddResponse) UnmarshalJSON(data []byte) error {
	// 尝试解析为数组（取第一个元素）
	var arr []UserTagsAddResponse
	if err := json.Unmarshal(data, &arr); err == nil && len(arr) > 0 {
		*u = arr[0]
		return nil
	}
	// 尝试解析为对象
	var obj UserTagsAddResponse
	if err := json.Unmarshal(data, &obj); err == nil {
		*u = obj
		return nil
	}
	return fmt.Errorf("cannot unmarshal UserTagsAddResponse")
}

// 55. 删除个性标签
type UserTagsDeleteRequest struct {
	TagID int `json:"tag_id"` // 标签ID
}

type UserTagsDeleteResponse struct {
	Success bool `json:"success"`
}

// 56. 获取个性标签列表
type UserTagsListsRequest struct {
	UserID int `json:"user_id,omitempty"` // 用户ID，不传则当前用户
}

// UserTagsListsResponse 支持数组和对象两种格式
type UserTagsListsResponse []UserTagItem

// UnmarshalJSON 自定义解析，支持数组和对象格式
func (u *UserTagsListsResponse) UnmarshalJSON(data []byte) error {
	// 尝试解析为数组
	var arr []UserTagItem
	if err := json.Unmarshal(data, &arr); err == nil {
		*u = arr
		return nil
	}
	// 尝试解析为对象（单个元素）
	var obj UserTagItem
	if err := json.Unmarshal(data, &obj); err == nil {
		*u = []UserTagItem{obj}
		return nil
	}
	return fmt.Errorf("cannot unmarshal UserTagsListsResponse")
}

type UserTagItem struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userid"`
	Name        string `json:"name"`
	Recognition int    `json:"recognition"` // 认可数
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// 57. 认可个性标签
type UserTagsRecognizeRequest struct {
	TagID int `json:"tag_id"` // 标签ID
}

type UserTagsRecognizeResponse struct {
	Success bool `json:"success"`
}

// 58. 修改个性标签
type UserTagsUpdateRequest struct {
	TagID int    `json:"tag_id"` // 标签ID
	Name  string `json:"name"`   // 新名称
}

type UserTagsUpdateResponse struct {
	ID   int    `json:"id"`   // 标签ID
	Name string `json:"name"` // 新名称
}

// 59. 获取任务浏览历史
type UserTaskBrowseRequest struct {
	Page     int `json:"page,omitempty"`     // 页码，默认1
	PageSize int `json:"pagesize,omitempty"` // 每页数量，默认20
}

// UserTaskBrowseResponse 支持分页对象和数组两种格式
type UserTaskBrowseResponse struct {
	CurrentPage  int                  `json:"current_page"`
	Data         []UserTaskBrowseItem `json:"data"`
	FirstPageUrl *string              `json:"first_page_url"`
	From         *int                 `json:"from"`
	LastPage     int                  `json:"last_page"`
	LastPageUrl  *string              `json:"last_page_url"`
	Links        []interface{}        `json:"links"`
	NextPageUrl  *string              `json:"next_page_url"`
	Path         string               `json:"path"`
	PerPage      int                  `json:"per_page"`
	PrevPageUrl  *string              `json:"prev_page_url"`
	To           *int                 `json:"to"`
	Total        int                  `json:"total"`
}

// UnmarshalJSON 自定义解析，支持分页对象和数组格式
func (u *UserTaskBrowseResponse) UnmarshalJSON(data []byte) error {
	// 尝试解析为分页对象
	var obj UserTaskBrowseResponse
	if err := json.Unmarshal(data, &obj); err == nil && obj.Data != nil {
		*u = obj
		return nil
	}
	// 尝试解析为数组（直接作为 data）
	var arr []UserTaskBrowseItem
	if err := json.Unmarshal(data, &arr); err == nil {
		*u = UserTaskBrowseResponse{
			Data: arr,
			Total: len(arr),
		}
		return nil
	}
	return fmt.Errorf("cannot unmarshal UserTaskBrowseResponse")
}

type UserTaskBrowseItem struct {
	ID        int         `json:"id"`
	TaskID    int         `json:"task_id"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	Task      interface{} `json:"task,omitempty"`
}

// 60. 清理任务浏览历史
type UserTaskBrowseCleanRequest struct{}

type UserTaskBrowseCleanResponse struct {
	Success bool `json:"success"`
}

// 61. 记录任务浏览历史
type UserTaskBrowseSaveRequest struct {
	TaskID int `json:"task_id"` // 任务ID
}

type UserTaskBrowseSaveResponse struct {
	Success bool `json:"success"`
}

// 62. 查询 token 过期时间
type UserTokenExpireRequest struct{}

type UserTokenExpireResponse struct {
	ExpireAt string      `json:"expire_at,omitempty"` // 过期时间
	Data     interface{} `json:"data,omitempty"`      // 可能是字符串或对象
}
