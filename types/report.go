package types

import (
	"encoding/json"
	"fmt"
)

// ==================== 报告管理相关 ====================

// ReportMyRequest 01. 我发送的汇报
type ReportMyRequest struct {
	Type   string `json:"type,omitempty"`   // 报告类型: day(日报), week(周报), month(月报)
	Search string `json:"search,omitempty"` // 搜索关键词
	Page   int    `json:"page,omitempty"`   // 页码
}

type ReportMyResponse struct {
	CurrentPage  int           `json:"current_page"`
	Data         []ReportInfo  `json:"data"`
	FirstPageUrl *string       `json:"first_page_url,omitempty"`
	From         *int          `json:"from,omitempty"`
	LastPage     int           `json:"last_page"`
	LastPageUrl  *string       `json:"last_page_url,omitempty"`
	Links        []interface{} `json:"links,omitempty"`
	NextPageUrl  *string       `json:"next_page_url,omitempty"`
	Path         string        `json:"path"`
	PerPage      int           `json:"per_page"`
	PrevPageUrl  *string       `json:"prev_page_url,omitempty"`
	To           *int          `json:"to,omitempty"`
	Total        int           `json:"total"`
}

// ReportReceiveRequest 02. 我接收的汇报
type ReportReceiveRequest struct {
	Type   string `json:"type,omitempty"`   // 报告类型: day(日报), week(周报), month(月报)
	Search string `json:"search,omitempty"` // 搜索关键词（标题、发送人）
	Time   string `json:"time,omitempty"`   // 时间范围筛选
	Page   int    `json:"page,omitempty"`   // 页码
}

type ReportReceiveResponse struct {
	CurrentPage  int           `json:"current_page"`
	Data         []ReportInfo  `json:"data"`
	FirstPageUrl *string       `json:"first_page_url,omitempty"`
	From         *int          `json:"from,omitempty"`
	LastPage     int           `json:"last_page"`
	LastPageUrl  *string       `json:"last_page_url,omitempty"`
	Links        []interface{} `json:"links,omitempty"`
	NextPageUrl  *string       `json:"next_page_url,omitempty"`
	Path         string        `json:"path"`
	PerPage      int           `json:"per_page"`
	PrevPageUrl  *string       `json:"prev_page_url,omitempty"`
	To           *int          `json:"to,omitempty"`
	Total        int           `json:"total"`
}

// ReportDetailRequest 05. 报告详情
type ReportDetailRequest struct {
	ReportID int `json:"report_id"` // 报告ID
}

// ReportDetailResponse 支持数组和对象两种格式
type ReportDetailResponse ReportDetail

// UnmarshalJSON 自定义解析，支持数组和对象格式
func (r *ReportDetailResponse) UnmarshalJSON(data []byte) error {
	// 尝试解析为数组（取第一个元素）
	var arr []ReportDetail
	if err := json.Unmarshal(data, &arr); err == nil && len(arr) > 0 {
		*r = ReportDetailResponse(arr[0])
		return nil
	}
	// 尝试解析为对象
	var obj ReportDetail
	if err := json.Unmarshal(data, &obj); err == nil {
		*r = ReportDetailResponse(obj)
		return nil
	}
	return fmt.Errorf("cannot unmarshal ReportDetailResponse")
}

// ReportStoreRequest 03. 保存并发送工作汇报
type ReportStoreRequest struct {
	ReportID  int    `json:"report_id,omitempty"` // 报告ID（修改时需要）
	Type      string `json:"type"`                // 报告类型: day(日报), week(周报), month(月报)
	Title     string `json:"title"`               // 报告标题
	Content   string `json:"content"`             // 报告内容（HTML格式）
	Recipient []int  `json:"recipient,omitempty"` // 接收人用户ID数组
	SendAt    string `json:"send_at,omitempty"`   // 发送时间，格式: YYYY-MM-DD HH:mm:ss
}

// ReportStoreResponse 支持数组和对象两种格式
type ReportStoreResponse ReportDetail

// UnmarshalJSON 自定义解析，支持数组和对象格式
func (r *ReportStoreResponse) UnmarshalJSON(data []byte) error {
	// 尝试解析为数组（取第一个元素）
	var arr []ReportDetail
	if err := json.Unmarshal(data, &arr); err == nil && len(arr) > 0 {
		*r = ReportStoreResponse(arr[0])
		return nil
	}
	// 尝试解析为对象
	var obj ReportDetail
	if err := json.Unmarshal(data, &obj); err == nil {
		*r = ReportStoreResponse(obj)
		return nil
	}
	return fmt.Errorf("cannot unmarshal ReportStoreResponse")
}

// ReportTemplateRequest 04. 生成汇报模板
type ReportTemplateRequest struct {
	Type string `json:"type"` // 报告类型: day(日报), week(周报), month(月报)
	Time string `json:"time"` // 时间，格式: YYYY-MM-DD
}

type ReportTemplateResponse ReportTemplate

// ReportMarkRequest 06. 标记已读/未读
type ReportMarkRequest struct {
	ReportID int    `json:"report_id"` // 报告ID
	Type     string `json:"type"`      // 操作类型: read(已读), unread(未读)
}

// ReportMarkResponse 支持数组和对象两种格式
type ReportMarkResponse struct {
	Data interface{} `json:"data"`
}

// UnmarshalJSON 自定义解析，支持数组和对象格式
func (r *ReportMarkResponse) UnmarshalJSON(data []byte) error {
	// 尝试解析为数组（取第一个元素）
	var arr []ReportMarkResponse
	if err := json.Unmarshal(data, &arr); err == nil && len(arr) > 0 {
		*r = arr[0]
		return nil
	}
	// 尝试解析为对象
	var obj ReportMarkResponse
	if err := json.Unmarshal(data, &obj); err == nil {
		*r = obj
		return nil
	}
	return fmt.Errorf("cannot unmarshal ReportMarkResponse")
}

// ReportReadRequest 09. 批量标记汇报已读
type ReportReadRequest struct {
	ReportIDs []int `json:"report_ids"` // 报告ID数组
}

// ReportReadResponse 支持数组和对象两种格式
type ReportReadResponse struct {
	Data interface{} `json:"data"`
}

// UnmarshalJSON 自定义解析，支持数组和对象格式
func (r *ReportReadResponse) UnmarshalJSON(data []byte) error {
	// 尝试解析为数组（取第一个元素）
	var arr []ReportReadResponse
	if err := json.Unmarshal(data, &arr); err == nil && len(arr) > 0 {
		*r = arr[0]
		return nil
	}
	// 尝试解析为对象
	var obj ReportReadResponse
	if err := json.Unmarshal(data, &obj); err == nil {
		*r = obj
		return nil
	}
	return fmt.Errorf("cannot unmarshal ReportReadResponse")
}

// ReportUnreadRequest 08. 获取未读
type ReportUnreadRequest struct {
	Type string `json:"type,omitempty"` // 报告类型: day(日报), week(周报), month(月报)，不传则全部
}

type ReportUnreadResponse UnreadInfo

// ReportLastSubmitterRequest 07. 获取最后一次提交的接收人
type ReportLastSubmitterRequest struct {
	Type string `json:"type"` // 报告类型: day(日报), week(周报), month(月报)
}

type ReportLastSubmitterResponse []User

// ==================== 报告相关数据结构 ====================

// ReportInfo 报告基本信息
type ReportInfo struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userid"`
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	SendAt      DateTime  `json:"send_at"`
	CreatedAt   DateTime  `json:"created_at"`
	UpdatedAt   DateTime  `json:"updated_at"`
	User        User      `json:"user"`         // 发送人信息
	Recipient   []User    `json:"recipient"`    // 接收人列表
	ReadStatus  int       `json:"read_status"`  // 阅读状态：0-未读，1-已读
	ReadAt      *DateTime `json:"read_at"`      // 阅读时间
	UnreadCount int       `json:"unread_count"` // 未读人数
}

// ReportDetail 报告详细信息
type ReportDetail struct {
	ID          int          `json:"id"`
	UserID      int          `json:"userid"`
	Type        string       `json:"type"`
	Title       string       `json:"title"`
	Content     string       `json:"content"`
	SendAt      DateTime     `json:"send_at"`
	CreatedAt   DateTime     `json:"created_at"`
	UpdatedAt   DateTime     `json:"updated_at"`
	User        User         `json:"user"`      // 发送人信息
	Recipient   []User       `json:"recipient"` // 接收人列表
	ReadStatus  int          `json:"read_status"`
	ReadAt      *DateTime    `json:"read_at"`
	ReadUsers   []ReadUser   `json:"read_users"`   // 已读用户列表
	UnreadUsers []UnreadUser `json:"unread_users"` // 未读用户列表
}

// ReportTemplate 报告模板
type ReportTemplate struct {
	Type      string     `json:"type"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Tasks     []TaskInfo `json:"tasks"`      // 相关任务列表（使用project中定义的TaskInfo）
	TimeRange TimeRange  `json:"time_range"` // 时间范围
	Recipient []User     `json:"recipient"`  // 推荐接收人
}

// TimeRange 时间范围
type TimeRange struct {
	Start string `json:"start"` // 开始时间 YYYY-MM-DD
	End   string `json:"end"`   // 结束时间 YYYY-MM-DD
}

// ReadUser 已读用户信息
type ReadUser struct {
	UserID int      `json:"userid"`
	ReadAt DateTime `json:"read_at"`
	User   User     `json:"user"`
}

// UnreadUser 未读用户信息
type UnreadUser struct {
	UserID int  `json:"userid"`
	User   User `json:"user"`
}

// UnreadInfo 未读统计信息
type UnreadInfo struct {
	Total    int            `json:"total"`     // 总未读数
	Day      int            `json:"day"`       // 日报未读数
	Week     int            `json:"week"`      // 周报未读数
	Month    int            `json:"month"`     // 月报未读数
	ByType   map[string]int `json:"by_type"`   // 按类型分组的未读数
	LastTime *DateTime      `json:"last_time"` // 最后一条未读时间
}

// User 用户信息（已在 common.go 中定义，这里用于类型引用）

// ==================== 新增接口类型定义 ====================

// ReportAnalysaveRequest 10. 保存工作汇报 AI 分析
type ReportAnalysaveRequest struct {
	ReportID int    `json:"report_id"` // 报告ID
	Content  string `json:"content"`   // AI分析内容
}

// ReportAnalysaveResponse 支持数组和对象两种格式
type ReportAnalysaveResponse struct {
	Success bool `json:"success"`
}

// UnmarshalJSON 自定义解析，支持数组和对象格式
func (r *ReportAnalysaveResponse) UnmarshalJSON(data []byte) error {
	// 尝试解析为数组（取第一个元素）
	var arr []ReportAnalysaveResponse
	if err := json.Unmarshal(data, &arr); err == nil && len(arr) > 0 {
		*r = arr[0]
		return nil
	}
	// 尝试解析为对象
	var obj ReportAnalysaveResponse
	if err := json.Unmarshal(data, &obj); err == nil {
		*r = obj
		return nil
	}
	return fmt.Errorf("cannot unmarshal ReportAnalysaveResponse")
}

// ReportShareRequest 11. 分享报告到消息
type ReportShareRequest struct {
	ReportID int `json:"report_id"` // 报告ID
	DialogID int `json:"dialog_id"` // 对话ID
}

// ReportShareResponse 支持数组和对象两种格式
type ReportShareResponse struct {
	Success bool `json:"success"`
}

// UnmarshalJSON 自定义解析，支持数组和对象格式
func (r *ReportShareResponse) UnmarshalJSON(data []byte) error {
	// 尝试解析为数组（取第一个元素）
	var arr []ReportShareResponse
	if err := json.Unmarshal(data, &arr); err == nil && len(arr) > 0 {
		*r = arr[0]
		return nil
	}
	// 尝试解析为对象
	var obj ReportShareResponse
	if err := json.Unmarshal(data, &obj); err == nil {
		*r = obj
		return nil
	}
	return fmt.Errorf("cannot unmarshal ReportShareResponse")
}
