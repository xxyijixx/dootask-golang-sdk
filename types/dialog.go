package types

// DialogItem 对话项结构
type DialogItem struct {
	// 基础字段
	ID        int     `json:"id"`         // 对话ID
	Type      string  `json:"type"`       // 对话类型：user(私聊)或group(群聊)
	GroupType *string `json:"group_type"` // 群组类型：user(用户群)、project(项目群)、task(任务群)、all(全体成员)
	Name      string  `json:"name"`       // 对话名称
	Avatar    string  `json:"avatar"`     // 头像URL
	OwnerID   *int    `json:"owner_id"`   // 群主用户ID
	LinkID    *int    `json:"link_id"`    // 关联ID
	TopUserID *int    `json:"top_userid"` // 置顶的用户ID
	TopMsgID  *int    `json:"top_msg_id"` // 置顶的消息ID
	CreatedAt *string `json:"created_at"` // 创建时间
	UpdatedAt *string `json:"updated_at"` // 更新时间
	DeletedAt *string `json:"deleted_at"` // 删除时间

	// 用户关系字段
	TopAt      *string `json:"top_at"`      // 置顶时间
	LastAt     *string `json:"last_at"`     // 最后活动时间
	MarkUnread *int    `json:"mark_unread"` // 标记未读
	Silence    *int    `json:"silence"`     // 静音
	Hide       *int    `json:"hide"`        // 隐藏
	Color      *string `json:"color"`       // 颜色
	UserAt     *string `json:"user_at"`     // 用户更新时间
	UserMs     *int64  `json:"user_ms"`     // 用户更新时间戳

	// 未读消息统计
	Unread     int   `json:"unread"`      // 未读消息总数
	UnreadOne  int   `json:"unread_one"`  // 最早未读消息ID
	Mention    int   `json:"mention"`     // @我的消息数
	MentionIDs []int `json:"mention_ids"` // @我的消息ID列表

	// 对话人数统计
	People     int `json:"people"`      // 总人数
	PeopleUser int `json:"people_user"` // 用户数
	PeopleBot  int `json:"people_bot"`  // 机器人数

	// 其他信息
	TodoNum      int           `json:"todo_num"`      // 待办数量
	Pinyin       string        `json:"pinyin"`        // 拼音
	QuickMsgs    []interface{} `json:"quick_msgs"`    // 快捷消息
	DialogUser   *DialogUser   `json:"dialog_user"`   // 对话用户（私聊时使用）
	GroupInfo    *GroupInfo    `json:"group_info"`    // 群组信息（群聊时使用）
	Bot          int           `json:"bot"`           // 是否机器人
	DialogMute   *string       `json:"dialog_mute"`   // 对话禁言状态
	DialogDelete *int          `json:"dialog_delete"` // 是否删除

	// 用户聊天特有字段
	Email     *string `json:"email"`      // 用户邮箱（私聊时）
	Userimg   *string `json:"userimg"`    // 用户头像（私聊时）
	IsDisable *bool   `json:"is_disable"` // 是否禁用（私聊时）

	// 消息类型标记
	HasTag   *bool `json:"has_tag"`   // 是否有标签消息
	HasTodo  *bool `json:"has_todo"`  // 是否有待办消息
	HasImage *bool `json:"has_image"` // 是否有图片消息
	HasFile  *bool `json:"has_file"`  // 是否有文件消息
	HasLink  *bool `json:"has_link"`  // 是否有链接消息

	// 最后消息
	LastMsg *LastMsg `json:"last_msg"` // 最后一条消息
}

// DialogUser 对话用户结构
type DialogUser struct {
	ID        int     `json:"id"`
	DialogID  int     `json:"dialog_id"`
	UserID    int     `json:"userid"`
	Bot       int     `json:"bot"`
	Inviter   int     `json:"inviter"`
	Important int     `json:"important"`
	LastAt    *string `json:"last_at"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// GroupInfo 群组信息结构
type GroupInfo struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	ArchivedAt *string `json:"archived_at,omitempty"` // 项目群使用
	DeletedAt  *string `json:"deleted_at,omitempty"`  // 项目群使用
	CompleteAt *string `json:"complete_at,omitempty"` // 任务群使用
}

// LastMsg 最后消息结构
type LastMsg struct {
	ID         int                    `json:"id"`
	Type       string                 `json:"type"`
	Msg        map[string]interface{} `json:"msg"` // 消息内容，包含preview字段
	UserID     int                    `json:"userid"`
	Percentage *int                   `json:"percentage,omitempty"`
	Emoji      []interface{}          `json:"emoji"`
	CreatedAt  string                 `json:"created_at"`
}

// MessageItem 消息项结构
type MessageItem struct {
	ID         int                    `json:"id"`          // 消息ID
	DialogID   int                    `json:"dialog_id"`   // 对话ID
	DialogType string                 `json:"dialog_type"` // 对话类型
	SessionID  *int                   `json:"session_id"`  // 会话ID
	UserID     int                    `json:"userid"`      // 发送者ID
	Type       string                 `json:"type"`        // 消息类型
	Mtype      string                 `json:"mtype"`       // 消息类型（用于搜索）
	Msg        map[string]interface{} `json:"msg"`         // 消息内容
	Emoji      []interface{}          `json:"emoji"`       // emoji回复
	Read       int                    `json:"read"`        // 已阅数量
	Send       int                    `json:"send"`        // 发送数量
	Tag        int                    `json:"tag"`         // 标注会员ID
	Todo       int                    `json:"todo"`        // 设为待办会员ID
	Link       int                    `json:"link"`        // 是否存在链接
	Modify     int                    `json:"modify"`      // 是否编辑
	Bot        int                    `json:"bot"`         // 是否机器人的消息
	ReplyNum   int                    `json:"reply_num"`   // 回复数量
	ReplyID    *int                   `json:"reply_id"`    // 回复ID
	ForwardID  *int                   `json:"forward_id"`  // 转发ID
	ForwardNum int                    `json:"forward_num"` // 被转发次数
	CreatedAt  string                 `json:"created_at"`  // 创建时间
	DeletedAt  *string                `json:"deleted_at"`  // 删除时间

	// 扩展字段
	Mention    *int    `json:"mention,omitempty"` // 是否@我
	Dot        *int    `json:"dot,omitempty"`     // 是否有点标记
	ReadAt     *string `json:"read_at,omitempty"` // 阅读时间
	Percentage int     `json:"percentage"`        // 阅读占比
	TodoDone   bool    `json:"todo_done"`         // 待办是否完成
	NextID     int     `json:"next_id"`           // 下一条消息ID
	PrevID     int     `json:"prev_id"`           // 上一条消息ID
}

// DialogTodoItem 待办项
type DialogTodoItem struct {
	ID        int     `json:"id"`
	DialogID  int     `json:"dialog_id"`
	MsgID     int     `json:"msg_id"`
	UserID    int     `json:"userid"`
	DoneAt    *string `json:"done_at"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// DialogUserWithDetail 带用户详情的对话成员
type DialogUserWithDetail struct {
	ID        int     `json:"id"`
	DialogID  int     `json:"dialog_id"`
	UserID    int     `json:"userid"`
	Bot       int     `json:"bot"`
	Inviter   int     `json:"inviter"`
	Important int     `json:"important"`
	LastAt    *string `json:"last_at"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`

	// 详细用户信息
	UserID2   int     `json:"userid2,omitempty"`    // 实际用户ID
	Nickname  string  `json:"nickname,omitempty"`   // 昵称
	Email     string  `json:"email,omitempty"`      // 邮箱
	Userimg   string  `json:"userimg,omitempty"`    // 头像
	DisableAt *string `json:"disable_at,omitempty"` // 禁用时间
	Online    bool    `json:"online,omitempty"`     // 是否在线
}

// ============================ 对话管理接口 ============================

// 1. 对话列表
type DialogListsRequest struct {
	Timerange string `json:"timerange,omitempty"` // 时间范围，如："1678248944,1678248944"
	Page      int    `json:"page,omitempty"`      // 当前页，默认1
	PageSize  int    `json:"pagesize,omitempty"`  // 每页显示数量，默认50，最大100
}

type DialogListsResponse struct {
	CurrentPage  int           `json:"current_page"`
	Data         []DialogItem  `json:"data"`
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
	DeletedID    []int         `json:"deleted_id,omitempty"`
}

// 2. 搜索会话
type SearchDialogRequest struct {
	Key  string `json:"key"`            // 搜索关键词
	Take int    `json:"take,omitempty"` // 返回数量，默认20
}

type SearchDialogResponse []DialogItem

// 3. 获取单个会话信息
type DialogOneRequest struct {
	DialogID int `json:"dialog_id"` // 对话ID
}

type DialogOneResponse DialogItem

// 4. 获取会话成员
type DialogUserRequest struct {
	DialogID int `json:"dialog_id"`         // 会话ID
	Getuser  int `json:"getuser,omitempty"` // 获取会员详情：0(默认)或1
}

type DialogUserResponse []DialogUserWithDetail

// 5. 获取会话待办
type DialogTodoRequest struct {
	DialogID int `json:"dialog_id,omitempty"` // 会话ID（可选，不传则获取所有待办）
}

type DialogTodoResponse []DialogTodoItem

// 6. 会话置顶
type DialogTopRequest struct {
	DialogID int `json:"dialog_id"` // 会话ID
}

type DialogTopResponse struct {
	ID    int     `json:"id"`
	TopAt *string `json:"top_at"` // null表示取消置顶
}

// 7. 获取对方联系电话
type DialogTelRequest struct {
	DialogID int `json:"dialog_id"` // 会话ID
}

type DialogTelResponse struct {
	Tel string `json:"tel"` // 联系电话
}

// 8. 打开会话
type OpenDialogRequest struct {
	UserID int `json:"user_id"` // 用户ID
}

type OpenDialogResponse DialogItem

// ============================ 消息管理接口 ============================

// 9. 获取消息列表
type MessageListRequest struct {
	DialogID   int    `json:"dialog_id"`             // 对话ID
	MsgID      int    `json:"msg_id,omitempty"`      // 消息ID（获取回复）
	PositionID int    `json:"position_id,omitempty"` // 此消息ID前后的数据
	PrevID     int    `json:"prev_id,omitempty"`     // 此消息ID之前的数据
	NextID     int    `json:"next_id,omitempty"`     // 此消息ID之后的数据
	MsgType    string `json:"msg_type,omitempty"`    // 消息类型：tag、todo、link、text、image、file、record、meeting
	Take       int    `json:"take,omitempty"`        // 获取条数，默认50，最大100
}

type MessageListResponse struct {
	List   []MessageItem    `json:"list"`             // 消息列表
	Time   int64            `json:"time"`             // 时间戳
	Dialog *DialogItem      `json:"dialog,omitempty"` // 对话信息
	Todo   []DialogTodoItem `json:"todo,omitempty"`   // 待办列表
	Top    *MessageItem     `json:"top,omitempty"`    // 置顶消息
}

// 10. 搜索消息位置
type SearchMessageRequest struct {
	DialogID int    `json:"dialog_id,omitempty"` // 对话ID（可选）
	Key      string `json:"key"`                 // 搜索关键词
	Take     int    `json:"take,omitempty"`      // 返回数量，默认20
}

type SearchMessageResponse struct {
	List   []MessageItem `json:"list"`             // 消息列表
	Dialog *DialogItem   `json:"dialog,omitempty"` // 对话信息
}

// 11. 获取单条消息
type MessageOneRequest struct {
	DialogID int `json:"dialog_id"` // 对话ID
	MsgID    int `json:"msg_id"`    // 消息ID
}

type MessageOneResponse MessageItem

// 12. 已读聊天消息
type ReadMessageRequest struct {
	DialogID int `json:"dialog_id"` // 对话ID
}

type ReadMessageResponse struct {
	Read       int `json:"read"`       // 已读数量
	Send       int `json:"send"`       // 发送总数
	Percentage int `json:"percentage"` // 阅读百分比
}

// 13. 获取未读消息数据
type UnreadMessageRequest struct {
	DialogID int `json:"dialog_id"` // 对话ID
}

type UnreadMessageResponse struct {
	Unread     int   `json:"unread"`      // 未读消息总数
	UnreadOne  int   `json:"unread_one"`  // 最早未读消息ID
	Mention    int   `json:"mention"`     // @我的消息数
	MentionIDs []int `json:"mention_ids"` // @我的消息ID列表
}

// 14. 通知成员监听消息
type StreamMessageRequest struct {
	DialogID int   `json:"dialog_id"` // 对话ID
	UserIDs  []int `json:"user_ids"`  // 用户ID列表
}

type StreamMessageResponse struct {
	Success bool `json:"success"`
}

// 15. 发送消息
type SendMessageRequest struct {
	DialogID   int    `json:"dialog_id,omitempty"`   // 对话ID（存在dialog_ids时无效）
	DialogIDs  string `json:"dialog_ids,omitempty"`  // 对话ID列表，多个对话ID用逗号分隔
	Text       string `json:"text"`                  // 消息内容
	Key        string `json:"key,omitempty"`         // 搜索关键词
	TextType   string `json:"text_type,omitempty"`   // 消息类型：html(默认)或md
	UpdateID   int    `json:"update_id,omitempty"`   // 更新消息ID（优先大于reply_id）
	UpdateMark string `json:"update_mark,omitempty"` // 是否更新标记：no或yes(默认)
	ReplyID    int    `json:"reply_id,omitempty"`    // 回复ID
	ReplyCheck string `json:"reply_check,omitempty"` // 回复验证：no(默认)或yes
	Silence    string `json:"silence,omitempty"`     // 是否静默发送：no(默认)或yes
	ModelName  string `json:"model_name,omitempty"`  // 模型名称（仅AI机器人支持）
}

type SendMessageResponse MessageItem

// 16. 发送语音
type SendRecordRequest struct {
	DialogID int    `json:"dialog_id"`           // 对话ID
	File     string `json:"file"`                // 语音文件
	FileName string `json:"file_name,omitempty"` // 文件名称
	ReplyID  int    `json:"reply_id,omitempty"`  // 回复ID
	Silence  string `json:"silence,omitempty"`   // 是否静默发送
}

type SendRecordResponse MessageItem

// 17. 文件上传
type SendFileRequest struct {
	DialogID int    `json:"dialog_id"`           // 对话ID
	File     string `json:"file"`                // 文件
	FileName string `json:"file_name,omitempty"` // 文件名称
	ReplyID  int    `json:"reply_id,omitempty"`  // 回复ID
	Silence  string `json:"silence,omitempty"`   // 是否静默发送
}

type SendFileResponse MessageItem

// 18. 群发文件上传
type SendFilesRequest struct {
	DialogIDs string `json:"dialog_ids"`          // 对话ID（user_ids 二选一）
	UserIDs   string `json:"user_ids,omitempty"`  // 用户ID，多个用逗号分隔
	File      string `json:"file"`                // 文件
	FileName  string `json:"file_name,omitempty"` // 文件名称
	ReplyID   int    `json:"reply_id,omitempty"`  // 回复ID
	Silence   string `json:"silence,omitempty"`   // 是否静默发送
}

type SendFilesResponse []MessageItem

// 19. 通过文件ID发送文件
type SendFileIDRequest struct {
	DialogIDs []int `json:"dialogids"` // 转发给的对话ID
	FileID    int   `json:"file_id"`   // 文件ID
}

type SendFileIDResponse []MessageItem

// 20. 发送匿名消息
type SendAnonRequest struct {
	DialogID int    `json:"dialog_id"`          // 对话ID
	Text     string `json:"text"`               // 消息内容
	ReplyID  int    `json:"reply_id,omitempty"` // 回复ID
}

type SendAnonResponse MessageItem

// 21. 获取消息阅读情况
type MessageReadListRequest struct {
	DialogID int `json:"dialog_id"` // 对话ID
	MsgID    int `json:"msg_id"`    // 消息ID
}

type MessageReadListResponse []struct {
	UserID   int     `json:"user_id"`
	Nickname string  `json:"nickname"`
	Userimg  string  `json:"userimg"`
	ReadAt   *string `json:"read_at"`
}

// 22. 消息详情
type MessageDetailRequest struct {
	DialogID int `json:"dialog_id"` // 对话ID
	MsgID    int `json:"msg_id"`    // 消息ID
}

type MessageDetailResponse MessageItem

// 23. 文件下载
type DownloadFileRequest struct {
	DialogID int `json:"dialog_id"` // 对话ID
	MsgID    int `json:"msg_id"`    // 消息ID
}

type DownloadFileResponse struct {
	Url      string `json:"url"`      // 下载URL
	Filename string `json:"filename"` // 文件名
}

// 24. 聊天消息撤回
type WithdrawMessageRequest struct {
	DialogID int `json:"dialog_id"` // 对话ID
	MsgID    int `json:"msg_id"`    // 消息ID
}

type WithdrawMessageResponse MessageItem

// 25. 消息标记操作
type MarkMessageRequest struct {
	DialogID int `json:"dialog_id"` // 会话ID
	MsgID    int `json:"msg_id"`    // 消息ID
}

type MarkMessageResponse MessageItem

// 26. 消息免打扰
type SilenceMessageRequest struct {
	DialogID int `json:"dialog_id"` // 会话ID
}

type SilenceMessageResponse struct {
	ID      int `json:"id"`
	Silence int `json:"silence"` // 0=正常，1=免打扰
}

// 27. 转发消息给
type ForwardMessageRequest struct {
	DialogID  int   `json:"dialog_id"` // 源对话ID
	MsgID     int   `json:"msg_id"`    // 消息ID
	DialogIDs []int `json:"dialogids"` // 转发给的对话ID
}

type ForwardMessageResponse []MessageItem

// 28. emoji回复
type EmojiMessageRequest struct {
	DialogID int    `json:"dialog_id"` // 对话ID
	MsgID    int    `json:"msg_id"`    // 消息ID
	Emoji    string `json:"emoji"`     // emoji符号
}

type EmojiMessageResponse MessageItem

// 29. 标注/取消标注
type TagMessageRequest struct {
	DialogID int `json:"dialog_id"` // 会话ID
	MsgID    int `json:"msg_id"`    // 消息ID
}

type TagMessageResponse MessageItem

// 30. 设待办/取消待办
type TodoMessageRequest struct {
	DialogID int `json:"dialog_id"` // 会话ID
	MsgID    int `json:"msg_id"`    // 消息ID
}

type TodoMessageResponse MessageItem

// 31. 获取消息待办情况
type TodoListMessageRequest struct {
	DialogID int `json:"dialog_id"` // 对话ID
	MsgID    int `json:"msg_id"`    // 消息ID
}

type TodoListMessageResponse []DialogTodoItem

// 32. 完成待办
type DoneTodoRequest struct {
	DialogID int `json:"dialog_id"` // 会话ID
	MsgID    int `json:"msg_id"`    // 消息ID
}

type DoneTodoResponse DialogTodoItem

// 33. 设置颜色
type ColorMessageRequest struct {
	DialogID int    `json:"dialog_id"` // 会话ID
	Color    string `json:"color"`     // 颜色值
}

type ColorMessageResponse struct {
	ID    int    `json:"id"`
	Color string `json:"color"`
}

// ============================ 群组管理接口 ============================

// 34. 新增群组
type CreateGroupRequest struct {
	Avatar   string `json:"avatar,omitempty"`    // 群头像
	ChatName string `json:"chat_name,omitempty"` // 群名称
	UserIDs  []int  `json:"userids"`             // 群成员ID列表
}

type CreateGroupResponse DialogItem

// 35. 修改群组
type EditGroupRequest struct {
	DialogID int    `json:"dialog_id"`           // 会话ID
	Avatar   string `json:"avatar,omitempty"`    // 群头像
	ChatName string `json:"chat_name,omitempty"` // 群名称
}

type EditGroupResponse DialogItem

// 36. 添加群成员
type AddGroupUserRequest struct {
	DialogID int   `json:"dialog_id"` // 会话ID
	UserIDs  []int `json:"userids"`   // 要添加的用户ID列表
}

type AddGroupUserResponse struct {
	People     int `json:"people"`      // 群组总人数
	PeopleUser int `json:"people_user"` // 用户数
	PeopleBot  int `json:"people_bot"`  // 机器人数
}

// 37. 移出（退出）群成员
type DelGroupUserRequest struct {
	DialogID int    `json:"dialog_id"`         // 会话ID
	UserIDs  []int  `json:"userids,omitempty"` // 要移出的用户ID列表（不传则退出自己）
	Type     string `json:"type,omitempty"`    // 类型：exit(退出)或remove(移出)
}

type DelGroupUserResponse struct {
	People     int `json:"people"`      // 群组总人数
	PeopleUser int `json:"people_user"` // 用户数
	PeopleBot  int `json:"people_bot"`  // 机器人数
}

// 38. 转让群组
type TransferGroupRequest struct {
	DialogID int `json:"dialog_id"` // 会话ID
	UserID   int `json:"userid"`    // 新群主用户ID
}

type TransferGroupResponse DialogItem

// 39. 解散群组
type DisbandGroupRequest struct {
	DialogID int `json:"dialog_id"` // 会话ID
}

type DisbandGroupResponse struct {
	Success bool `json:"success"`
}

// 40. 搜索个人群（仅限管理员）
type SearchGroupUserRequest struct {
	Key string `json:"key"` // 搜索关键词
}

type SearchGroupUserResponse []struct {
	DialogID int    `json:"dialog_id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	OwnerID  int    `json:"owner_id"`
}

// ============================ OKR接口 ============================

// 41. 创建OKR评论会话
type CreateOkrDialogRequest struct {
	OkrID   int    `json:"okr_id"`            // OKR ID
	Content string `json:"content,omitempty"` // 评论内容
}

type CreateOkrDialogResponse DialogItem

// 42. 推送OKR相关信息
type PushOkrInfoRequest struct {
	OkrID    int    `json:"okr_id"`              // OKR ID
	DialogID int    `json:"dialog_id"`           // 对话ID
	Content  string `json:"content"`             // 推送内容
	PushType string `json:"push_type,omitempty"` // 推送类型
}

type PushOkrInfoResponse struct {
	Success bool `json:"success"`
}
