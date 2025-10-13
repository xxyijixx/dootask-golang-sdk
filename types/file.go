package types

// File represents a file or folder in the system
type File struct {
	ID          int       `json:"id"`
	PID         *int      `json:"pid,omitempty"`          // 父级ID
	PIDs        *string   `json:"pids,omitempty"`         // 父级ID递归
	CID         *int      `json:"cid,omitempty"`          // 复制ID
	Name        string    `json:"name"`                   // 名称
	Type        string    `json:"type"`                   // 类型 (folder, document, mind, drawio, word, excel, ppt, etc.)
	Ext         *string   `json:"ext,omitempty"`          // 后缀名
	Size        *int64    `json:"size,omitempty"`         // 大小(字节)
	UserID      int       `json:"userid"`                 // 拥有者ID
	Share       *int      `json:"share,omitempty"`        // 是否共享
	GuestAccess *int      `json:"guest_access,omitempty"` // 是否允许游客访问
	PShare      *int      `json:"pshare,omitempty"`       // 所属分享ID
	CreatedID   int       `json:"created_id"`             // 创建者ID
	CreatedAt   DateTime  `json:"created_at"`
	UpdatedAt   DateTime  `json:"updated_at"`
	DeletedAt   *DateTime `json:"deleted_at,omitempty"`

	// 权限相关
	Permission *int `json:"permission,omitempty"` // 权限级别
}

// FileListRequest represents the request for file list
type FileListRequest struct {
	PID *int `json:"pid,omitempty"` // 父级ID
}

// FileOneRequest represents the request for single file
type FileOneRequest struct {
	ID interface{} `json:"id"` // 文件ID或链接码
}

// FileSearchRequest represents the search request
type FileSearchRequest struct {
	Link *string `json:"link,omitempty"` // 通过分享地址搜索
	Key  *string `json:"key,omitempty"`  // 关键词
	Take *int    `json:"take,omitempty"` // 获取数量，默认50，最大100
}

// FileAddRequest represents the add/modify file request
type FileAddRequest struct {
	Name string `json:"name"`          // 文件名称
	Type string `json:"type"`          // 文件类型
	ID   *int   `json:"id,omitempty"`  // 文件ID（修改时使用）
	PID  *int   `json:"pid,omitempty"` // 父级ID
}

// FileCopyRequest represents the copy file request
type FileCopyRequest struct {
	ID int `json:"id"` // 文件ID
}

// FileMoveRequest represents the move files request
type FileMoveRequest struct {
	IDs []int `json:"ids"` // 文件ID列表
	PID int   `json:"pid"` // 移动到的文件夹ID
}

// FileRemoveRequest represents the delete files request
type FileRemoveRequest struct {
	IDs []int `json:"ids"` // 文件ID列表
}

// FileContentRequest represents the get content request
type FileContentRequest struct {
	ID           interface{} `json:"id"`                       // 文件ID或链接码
	OnlyUpdateAt *string     `json:"only_update_at,omitempty"` // 仅获取update_at字段 (yes/no)
	Down         *string     `json:"down,omitempty"`           // 下载模式 (no/yes/preview)
	HistoryID    *int        `json:"history_id,omitempty"`     // 读取历史记录ID
}

// FileContentSaveRequest represents the save content request
type FileContentSaveRequest struct {
	ID      int         `json:"id"`      // 文件ID
	Content interface{} `json:"content"` // 内容
}

// FileContentUploadRequest represents the upload request
type FileContentUploadRequest struct {
	PID                *int    `json:"pid,omitempty"`                // 父级ID
	Cover              *int    `json:"cover,omitempty"`              // 覆盖已存在的文件 (0/1)
	WebkitRelativePath *string `json:"webkitRelativePath,omitempty"` // 相对路径
}

// FileContentHistoryRequest represents the history request
type FileContentHistoryRequest struct {
	ID       int  `json:"id"`                 // 文件ID
	Page     *int `json:"page,omitempty"`     // 当前页，默认1
	PageSize *int `json:"pagesize,omitempty"` // 每页显示数量，默认20，最大100
}

// FileContentRestoreRequest represents the restore request
type FileContentRestoreRequest struct {
	ID        int `json:"id"`         // 文件ID
	HistoryID int `json:"history_id"` // 历史数据ID
}

// FileShareRequest represents the get share info request
type FileShareRequest struct {
	ID int `json:"id"` // 文件ID
}

// FileShareUpdateRequest represents the update share request
type FileShareUpdateRequest struct {
	ID         int   `json:"id"`                // 文件ID
	UserIDs    []int `json:"userids,omitempty"` // 共享成员ID列表
	Permission int   `json:"permission"`        // 共享方式 (0只读/1读写/-1删除)
	Force      *int  `json:"force,omitempty"`   // 忽略提醒 (0/1)
}

// FileShareOutRequest represents the exit share request
type FileShareOutRequest struct {
	ID int `json:"id"` // 文件ID
}

// FileLinkRequest represents the get link request
type FileLinkRequest struct {
	ID          int    `json:"id"`           // 文件ID
	Refresh     string `json:"refresh"`      // 刷新链接 (no/yes)
	GuestAccess string `json:"guest_access"` // 是否允许游客访问 (no/yes)
}

// FileDownloadPackRequest represents the pack download request
type FileDownloadPackRequest struct {
	IDs  []int  `json:"ids"`            // 文件ID列表
	Name string `json:"name,omitempty"` // 下载文件名
}

// FileUser represents a shared user
type FileUser struct {
	FileID     int `json:"file_id"`
	UserID     int `json:"userid"`
	Permission int `json:"permission"` // 0只读, 1读写
}

// FileLink represents a file share link
type FileLink struct {
	Code        string `json:"code"`
	URL         string `json:"url"`
	GuestAccess int    `json:"guest_access"`
}

// FileContent represents file content data
type FileContent struct {
	ID        int         `json:"id"`
	FID       int         `json:"fid"`
	Content   interface{} `json:"content,omitempty"`
	Text      *string     `json:"text,omitempty"`
	Size      int64       `json:"size"`
	UserID    int         `json:"userid"`
	CreatedAt DateTime    `json:"created_at"`
}

// FileHistory represents file content history
type FileHistory struct {
	ID        int      `json:"id"`
	Size      int64    `json:"size"`
	UserID    int      `json:"userid"`
	CreatedAt DateTime `json:"created_at"`
}
