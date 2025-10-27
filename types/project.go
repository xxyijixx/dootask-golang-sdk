package types

// ==================== 项目管理相关 ====================

// ProjectListsRequest 01. 获取项目列表
type ProjectListsRequest struct {
	Archived string `json:"archived,omitempty"` // 归档状态: no(未归档), yes(已归档), all(全部)，默认no
	Search   string `json:"search,omitempty"`   // 搜索关键词
	Page     int    `json:"page,omitempty"`     // 页码
	Pagesize int    `json:"pagesize,omitempty"` // 每页数量
}

type ProjectListsResponse struct {
	Data []ProjectInfo `json:"data"`
}

type ProjectInfo struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Desc          string    `json:"desc"`
	DialogID      int       `json:"dialog_id"`
	Owner         []int     `json:"owner"`
	OwnerUser     []User    `json:"owner_user"`
	CreatedAt     DateTime  `json:"created_at"`
	UpdatedAt     DateTime  `json:"updated_at"`
	ArchivedAt    *DateTime `json:"archived_at,omitempty"`
	TopAt         *DateTime `json:"top_at,omitempty"`
	Personal      int       `json:"personal"`
	Flow          string    `json:"flow"`
	ArchiveDays   int       `json:"archive_days"`
	ArchiveMethod string    `json:"archive_method"`
}

// ProjectOneRequest 02. 获取一个项目信息
type ProjectOneRequest struct {
	ProjectID int `json:"project_id"` // 项目ID
}

type ProjectOneResponse struct {
	Data ProjectDetail `json:"data"`
}

type ProjectDetail struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	Desc          string          `json:"desc"`
	DialogID      int             `json:"dialog_id"`
	Owner         []int           `json:"owner"`
	OwnerUser     []User          `json:"owner_user"`
	CreatedAt     DateTime        `json:"created_at"`
	UpdatedAt     DateTime        `json:"updated_at"`
	ArchivedAt    *DateTime       `json:"archived_at,omitempty"`
	TopAt         *DateTime       `json:"top_at,omitempty"`
	Personal      int             `json:"personal"`
	Flow          string          `json:"flow"`
	ArchiveDays   int             `json:"archive_days"`
	ArchiveMethod string          `json:"archive_method"`
	Columns       []ProjectColumn `json:"columns"`
	Members       []ProjectMember `json:"project_user"`
}

type ProjectColumn struct {
	ID        int      `json:"id"`
	ProjectID int      `json:"project_id"`
	Name      string   `json:"name"`
	Color     string   `json:"color"`
	Sort      int      `json:"sort"`
	CreatedAt DateTime `json:"created_at"`
}

type ProjectMember struct {
	ID        int      `json:"id"`
	ProjectID int      `json:"project_id"`
	UserID    int      `json:"userid"`
	Owner     int      `json:"owner"`
	CreatedAt DateTime `json:"created_at"`
	User      User     `json:"user"`
}

// ProjectAddRequest 03. 添加项目
type ProjectAddRequest struct {
	Name     string   `json:"name"`               // 项目名称
	Desc     string   `json:"desc,omitempty"`     // 项目描述
	Columns  []string `json:"columns,omitempty"`  // 初始化列名称数组
	Flow     string   `json:"flow,omitempty"`     // 是否开启流程: open/close
	Personal bool     `json:"personal,omitempty"` // 是否创建个人项目
}

type ProjectAddResponse struct {
	Data ProjectDetail `json:"data"`
}

// ProjectUpdateRequest 04. 修改项目
type ProjectUpdateRequest struct {
	ProjectID     int    `json:"project_id"`               // 项目ID
	Name          string `json:"name,omitempty"`           // 项目名称
	Desc          string `json:"desc,omitempty"`           // 项目描述
	ArchiveDays   int    `json:"archive_days,omitempty"`   // 自动归档天数
	ArchiveMethod string `json:"archive_method,omitempty"` // 归档方式
}

type ProjectUpdateResponse struct {
	Data ProjectDetail `json:"data"`
}

// ProjectTransferRequest 09. 移交项目
type ProjectTransferRequest struct {
	ProjectID int `json:"project_id"` // 项目ID
	UserID    int `json:"userid"`     // 新负责人用户ID
}

type ProjectTransferResponse struct {
	Data string `json:"data"`
}

// ProjectExitRequest 11. 退出项目
type ProjectExitRequest struct {
	ProjectID int `json:"project_id"` // 项目ID
}

type ProjectExitResponse struct {
	Data string `json:"data"`
}

// ProjectArchivedRequest 12. 归档项目
type ProjectArchivedRequest struct {
	ProjectID int    `json:"project_id"` // 项目ID
	Type      string `json:"type"`       // 操作类型: archive(归档), reduction(还原)
}

type ProjectArchivedResponse struct {
	Data string `json:"data"`
}

// ProjectRemoveRequest 13. 删除项目
type ProjectRemoveRequest struct {
	ProjectID int `json:"project_id"` // 项目ID
}

type ProjectRemoveResponse struct {
	Data string `json:"data"`
}

// ProjectTopRequest 43. 项目置顶
type ProjectTopRequest struct {
	ProjectID int    `json:"project_id"` // 项目ID
	Type      string `json:"type"`       // 操作类型: add(置顶), delete(取消置顶)
}

type ProjectTopResponse struct {
	Data string `json:"data"`
}

// ==================== 项目成员管理 ====================

// ProjectUserRequest 05. 修改项目成员
type ProjectUserRequest struct {
	ProjectID int    `json:"project_id"` // 项目ID
	Type      string `json:"type"`       // 操作类型: add(添加), delete(删除), owner(设置负责人)
	UserIDs   []int  `json:"userids"`    // 用户ID数组
}

type ProjectUserResponse struct {
	Data string `json:"data"`
}

// ProjectInviteRequest 06. 获取邀请链接
type ProjectInviteRequest struct {
	ProjectID int `json:"project_id"` // 项目ID
}

type ProjectInviteResponse struct {
	Data ProjectInviteData `json:"data"`
}

type ProjectInviteData struct {
	Code string `json:"code"`
	Link string `json:"link"`
}

// ProjectInviteInfoRequest 07. 通过邀请链接code获取项目信息
type ProjectInviteInfoRequest struct {
	Code string `json:"code"` // 邀请码
}

type ProjectInviteInfoResponse struct {
	Data ProjectInfo `json:"data"`
}

// ProjectInviteJoinRequest 08. 通过邀请链接code加入项目
type ProjectInviteJoinRequest struct {
	Code string `json:"code"` // 邀请码
}

type ProjectInviteJoinResponse struct {
	Data ProjectDetail `json:"data"`
}

// ==================== 任务列表（列）管理 ====================

// ProjectColumnListsRequest 14. 获取任务列表
type ProjectColumnListsRequest struct {
	ProjectID int `json:"project_id"` // 项目ID
}

type ProjectColumnListsResponse struct {
	Data []ProjectColumn `json:"data"`
}

// ProjectColumnAddRequest 15. 添加任务列表
type ProjectColumnAddRequest struct {
	ProjectID int    `json:"project_id"`      // 项目ID
	Name      string `json:"name"`            // 列名称
	Color     string `json:"color,omitempty"` // 颜色
}

type ProjectColumnAddResponse struct {
	Data ProjectColumn `json:"data"`
}

// ProjectColumnUpdateRequest 16. 修改任务列表
type ProjectColumnUpdateRequest struct {
	ColumnID int    `json:"column_id"`       // 列ID
	Name     string `json:"name,omitempty"`  // 列名称
	Color    string `json:"color,omitempty"` // 颜色
}

type ProjectColumnUpdateResponse struct {
	Data ProjectColumn `json:"data"`
}

// ProjectColumnRemoveRequest 17. 删除任务列表
type ProjectColumnRemoveRequest struct {
	ColumnID int `json:"column_id"` // 列ID
}

type ProjectColumnRemoveResponse struct {
	Data string `json:"data"`
}

// ProjectColumnOneRequest 18. 获取任务列详细
type ProjectColumnOneRequest struct {
	ColumnID int `json:"column_id"` // 列ID
}

type ProjectColumnOneResponse struct {
	Data ProjectColumn `json:"data"`
}

// ==================== 任务管理 ====================

// ProjectTaskListsRequest 19. 任务列表
type ProjectTaskListsRequest struct {
	ProjectID int    `json:"project_id,omitempty"` // 项目ID
	ColumnID  int    `json:"column_id,omitempty"`  // 列ID
	ParentID  int    `json:"parent_id,omitempty"`  // 主任务ID
	Status    string `json:"status,omitempty"`     // 状态: all, completed, uncompleted
	Search    string `json:"search,omitempty"`     // 搜索关键词
	Time      string `json:"time,omitempty"`       // 时间范围
	Page      int    `json:"page,omitempty"`       // 页码
	Pagesize  int    `json:"pagesize,omitempty"`   // 每页数量
}

type ProjectTaskListsResponse struct {
	Data []TaskInfo `json:"data"`
}

type TaskInfo struct {
	ID         int        `json:"id"`
	ProjectID  int        `json:"project_id"`
	ColumnID   int        `json:"column_id"`
	ParentID   int        `json:"parent_id"`
	Name       string     `json:"name"`
	Content    string     `json:"content,omitempty"`
	Color      string     `json:"color"`
	Owner      []int      `json:"owner"`
	Assist     []int      `json:"assist"`
	OwnerUser  []User     `json:"owner_user,omitempty"`
	AssistUser []User     `json:"assist_user,omitempty"`
	StartAt    *DateTime  `json:"start_at,omitempty"`
	EndAt      *DateTime  `json:"end_at,omitempty"`
	CompleteAt *DateTime  `json:"complete_at,omitempty"`
	ArchivedAt *DateTime  `json:"archived_at,omitempty"`
	CreatedAt  DateTime   `json:"created_at"`
	UpdatedAt  DateTime   `json:"updated_at"`
	DialogID   int        `json:"dialog_id"`
	Sort       int        `json:"sort"`
	SubTasks   []TaskInfo `json:"sub_tasks,omitempty"`
}

// ProjectTaskEasyListsRequest 20. 任务列表-简单的
type ProjectTaskEasyListsRequest struct {
	ProjectID int    `json:"project_id,omitempty"` // 项目ID
	Status    string `json:"status,omitempty"`     // 状态
}

type ProjectTaskEasyListsResponse struct {
	Data []TaskInfo `json:"data"`
}

// ProjectTaskOneRequest 24. 获取单个任务信息
type ProjectTaskOneRequest struct {
	TaskID int `json:"task_id"` // 任务ID
}

type ProjectTaskOneResponse struct {
	Data TaskInfo `json:"data"`
}

// ProjectTaskContentRequest 25. 获取任务详细描述
type ProjectTaskContentRequest struct {
	TaskID int `json:"task_id"` // 任务ID
}

type ProjectTaskContentResponse struct {
	Data TaskContent `json:"data"`
}

type TaskContent struct {
	Content string `json:"content"`
}

// ProjectTaskAddRequest 30. 添加任务
type ProjectTaskAddRequest struct {
	ProjectID int    `json:"project_id"`         // 项目ID
	ColumnID  int    `json:"column_id"`          // 列ID
	Name      string `json:"name"`               // 任务名称
	Content   string `json:"content,omitempty"`  // 任务内容
	Owner     []int  `json:"owner,omitempty"`    // 负责人用户ID数组
	Assist    []int  `json:"assist,omitempty"`   // 协助人员用户ID数组
	StartAt   string `json:"start_at,omitempty"` // 开始时间
	EndAt     string `json:"end_at,omitempty"`   // 结束时间
}

type ProjectTaskAddResponse struct {
	Data TaskInfo `json:"data"`
}

// ProjectTaskAddSubRequest 31. 添加子任务
type ProjectTaskAddSubRequest struct {
	TaskID int    `json:"task_id"` // 主任务ID
	Name   string `json:"name"`    // 子任务名称
}

type ProjectTaskAddSubResponse struct {
	Data TaskInfo `json:"data"`
}

// ProjectTaskUpdateRequest 32. 修改任务、子任务
type ProjectTaskUpdateRequest struct {
	TaskID     int         `json:"task_id"`               // 任务ID
	Name       string      `json:"name,omitempty"`        // 任务名称
	Content    string      `json:"content,omitempty"`     // 任务内容
	ColumnID   int         `json:"column_id,omitempty"`   // 移动到指定列ID
	Owner      []int       `json:"owner,omitempty"`       // 负责人用户ID数组
	Assist     []int       `json:"assist,omitempty"`      // 协助人员用户ID数组
	StartAt    string      `json:"start_at,omitempty"`    // 开始时间
	EndAt      string      `json:"end_at,omitempty"`      // 结束时间
	CompleteAt interface{} `json:"complete_at,omitempty"` // 完成时间，传时间字符串标记完成，传false标记未完成
}

type ProjectTaskUpdateResponse struct {
	Data TaskInfo `json:"data"`
}

// ProjectTaskDialogRequest 33. 创建/获取聊天室
type ProjectTaskDialogRequest struct {
	TaskID int `json:"task_id"` // 任务ID
}

type ProjectTaskDialogResponse struct {
	Data TaskDialogData `json:"data"`
}

type TaskDialogData struct {
	DialogID int `json:"dialog_id"`
}

// ProjectTaskArchivedRequest 34. 归档任务
type ProjectTaskArchivedRequest struct {
	TaskID int    `json:"task_id"` // 任务ID
	Type   string `json:"type"`    // 操作类型: archive(归档), reduction(还原)
}

type ProjectTaskArchivedResponse struct {
	Data string `json:"data"`
}

// ProjectTaskRemoveRequest 35. 删除任务
type ProjectTaskRemoveRequest struct {
	TaskID int `json:"task_id"` // 任务ID
}

type ProjectTaskRemoveResponse struct {
	Data string `json:"data"`
}

// ProjectTaskResetFromLogRequest 36. 根据日志重置任务
type ProjectTaskResetFromLogRequest struct {
	LogID int `json:"log_id"` // 日志ID
}

type ProjectTaskResetFromLogResponse struct {
	Data string `json:"data"`
}

// ProjectTaskMoveRequest 38. 任务移动
type ProjectTaskMoveRequest struct {
	TaskID    int `json:"task_id"`    // 任务ID
	ProjectID int `json:"project_id"` // 目标项目ID
	ColumnID  int `json:"column_id"`  // 目标列ID
}

type ProjectTaskMoveResponse struct {
	Data string `json:"data"`
}

// ProjectSortRequest 10. 排序任务
type ProjectSortRequest struct {
	ProjectID int   `json:"project_id"` // 项目ID
	TaskIDs   []int `json:"task_ids"`   // 任务ID数组，按顺序排列
}

type ProjectSortResponse struct {
	Data string `json:"data"`
}

// ==================== 文件管理 ====================

// ProjectTaskFilesRequest 26. 获取任务文件列表
type ProjectTaskFilesRequest struct {
	TaskID int `json:"task_id"` // 任务ID
}

type ProjectTaskFilesResponse struct {
	Data []TaskFile `json:"data"`
}

type TaskFile struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Size      int64    `json:"size"`
	Ext       string   `json:"ext"`
	Path      string   `json:"path"`
	Thumb     string   `json:"thumb"`
	UserID    int      `json:"userid"`
	CreatedAt DateTime `json:"created_at"`
}

// ProjectTaskFileDeleteRequest 27. 删除任务文件
type ProjectTaskFileDeleteRequest struct {
	FileID int `json:"file_id"` // 文件ID
}

type ProjectTaskFileDeleteResponse struct {
	Data string `json:"data"`
}

// ProjectTaskFileDetailRequest 28. 获取任务文件详情
type ProjectTaskFileDetailRequest struct {
	FileID int `json:"file_id"` // 文件ID
}

type ProjectTaskFileDetailResponse struct {
	Data TaskFile `json:"data"`
}

// ProjectTaskFileDownRequest 29. 下载任务文件
type ProjectTaskFileDownRequest struct {
	FileID int `json:"file_id"` // 文件ID
}

type ProjectTaskFileDownResponse struct {
	Data TaskFileDownData `json:"data"`
}

type TaskFileDownData struct {
	URL string `json:"url"`
}

// ==================== 工作流管理 ====================

// ProjectTaskFlowRequest 37. 任务工作流信息
type ProjectTaskFlowRequest struct {
	TaskID int `json:"task_id"` // 任务ID
}

type ProjectTaskFlowResponse struct {
	Data TaskFlowData `json:"data"`
}

type TaskFlowData struct {
	Flows []FlowItem `json:"flows"`
}

type FlowItem struct {
	ID        int      `json:"id"`
	ProjectID int      `json:"project_id"`
	Name      string   `json:"name"`
	Status    string   `json:"status"`
	Turns     []string `json:"turns"`
	CreatedAt DateTime `json:"created_at"`
	UpdatedAt DateTime `json:"updated_at"`
}

// ProjectFlowListRequest 39. 工作流列表
type ProjectFlowListRequest struct {
	ProjectID int `json:"project_id"` // 项目ID
}

type ProjectFlowListResponse struct {
	Data []FlowItem `json:"data"`
}

// ProjectFlowSaveRequest 40. 保存工作流
type ProjectFlowSaveRequest struct {
	ProjectID int      `json:"project_id"`        // 项目ID
	FlowID    int      `json:"flow_id,omitempty"` // 工作流ID（修改时传）
	Name      string   `json:"name"`              // 工作流名称
	Turns     []string `json:"turns"`             // 流转规则
}

type ProjectFlowSaveResponse struct {
	Data FlowItem `json:"data"`
}

// ProjectFlowDeleteRequest 41. 删除工作流
type ProjectFlowDeleteRequest struct {
	FlowID int `json:"flow_id"` // 工作流ID
}

type ProjectFlowDeleteResponse struct {
	Data string `json:"data"`
}

// ==================== 导出功能 ====================

// ProjectTaskExportRequest 21. 导出任务（限管理员）
type ProjectTaskExportRequest struct {
	ProjectID int    `json:"project_id,omitempty"` // 项目ID
	ColumnID  int    `json:"column_id,omitempty"`  // 列ID
	Status    string `json:"status,omitempty"`     // 状态
}

type ProjectTaskExportResponse struct {
	Data ExportData `json:"data"`
}

type ExportData struct {
	TaskID string `json:"task_id"`
	URL    string `json:"url"`
}

// ProjectTaskExportOverdueRequest 22. 导出超期任务（限管理员）
type ProjectTaskExportOverdueRequest struct {
	ProjectID int `json:"project_id,omitempty"` // 项目ID
}

type ProjectTaskExportOverdueResponse struct {
	Data ExportData `json:"data"`
}

// ProjectTaskDownRequest 23. 下载导出的任务
type ProjectTaskDownRequest struct {
	TaskID string `json:"task_id"` // 导出任务ID
}

type ProjectTaskDownResponse struct {
	Data DownloadData `json:"data"`
}

type DownloadData struct {
	URL string `json:"url"`
}

// ==================== 日志管理 ====================

// ProjectLogListsRequest 42. 获取项目、任务日志
type ProjectLogListsRequest struct {
	ProjectID int `json:"project_id,omitempty"` // 项目ID
	TaskID    int `json:"task_id,omitempty"`    // 任务ID
	Page      int `json:"page,omitempty"`       // 页码
	Pagesize  int `json:"pagesize,omitempty"`   // 每页数量
}

type ProjectLogListsResponse struct {
	Data []LogInfo `json:"data"`
}

type LogInfo struct {
	ID        int      `json:"id"`
	ProjectID int      `json:"project_id"`
	TaskID    int      `json:"task_id"`
	Type      string   `json:"type"`
	Content   string   `json:"content"`
	Record    string   `json:"record"`
	UserID    int      `json:"userid"`
	User      User     `json:"user"`
	CreatedAt DateTime `json:"created_at"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
