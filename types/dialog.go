package types

// Dialog 对话结构
type Dialog struct {
	// ID          int      `json:"id"`
	// Name        string   `json:"name"`
	// Description string   `json:"description"`
	// Type        int      `json:"type"`
	// Avatar      string   `json:"avatar"`
	// CreatedAt   DateTime `json:"created_at"`
	// UpdatedAt   DateTime `json:"updated_at"`
	// Creator     int      `json:"creator"`
	// Status      int      `json:"status"`
	// IsTop       int      `json:"is_top"`
	// IsMute      int      `json:"is_mute"`
	// IsArchive   int      `json:"is_archive"`
	// LastMsgID   int      `json:"last_msg_id"`
	// LastMsgTime DateTime `json:"last_msg_time"`
	// UnreadCount int      `json:"unread_count"`
	// MemberCount int      `json:"member_count"`

	ID         int           `json:"id"`
	Type       string        `json:"type"`
	GroupType  string        `json:"group_type"`
	SessionID  int           `json:"session_id"`
	Name       string        `json:"name"`
	Avatar     string        `json:"avatar"`
	OwnerID    int           `json:"owner_id"`
	LinkID     int           `json:"link_id"`
	TopUserid  int           `json:"top_userid"`
	TopMsgID   int           `json:"top_msg_id"`
	CreatedAt  DateTime      `json:"created_at"`
	UpdatedAt  DateTime      `json:"updated_at"`
	DeletedAt  *DateTime     `json:"deleted_at"`
	TopAt      *DateTime     `json:"top_at"`
	LastAt     *DateTime     `json:"last_at"`
	MarkUnread int           `json:"mark_unread"`
	Silence    int           `json:"silence"`
	Hide       int           `json:"hide"`
	Color      string        `json:"color"`
	UserAt     string        `json:"user_at"`
	UserMs     int64         `json:"user_ms"`
	Unread     int           `json:"unread"`
	UnreadOne  int           `json:"unread_one"`
	Mention    int           `json:"mention"`
	MentionIds []interface{} `json:"mention_ids"`
	People     int           `json:"people"`
	PeopleUser int           `json:"people_user"`
	PeopleBot  int           `json:"people_bot"`
	TodoNum    int           `json:"todo_num"`
	LastMsg    struct {
		ID     int    `json:"id"`
		Userid int    `json:"userid"`
		Type   string `json:"type"`
		Msg    struct {
			Preview string `json:"preview"`
		} `json:"msg"`
		Emoji      []interface{} `json:"emoji"`
		CreatedAt  string        `json:"created_at"`
		Percentage int           `json:"percentage"`
	} `json:"last_msg"`
	Pinyin       string        `json:"pinyin"`
	QuickMsgs    []interface{} `json:"quick_msgs"`
	DialogUser   interface{}   `json:"dialog_user"`
	GroupInfo    interface{}   `json:"group_info"`
	Bot          int           `json:"bot"`
	DialogMute   string        `json:"dialog_mute,omitempty"`
	HasTag       bool          `json:"has_tag,omitempty"`
	HasTodo      bool          `json:"has_todo,omitempty"`
	HasImage     bool          `json:"has_image,omitempty"`
	HasFile      bool          `json:"has_file,omitempty"`
	HasLink      bool          `json:"has_link,omitempty"`
	Email        string        `json:"email,omitempty"`
	Userimg      string        `json:"userimg,omitempty"`
	IsDisable    bool          `json:"is_disable,omitempty"`
	DialogDelete int           `json:"dialog_delete,omitempty"`
}

// DialogMember 对话成员结构
type DialogMember struct {
	ID        int      `json:"id"`
	DialogID  int      `json:"dialog_id"`
	UserID    int      `json:"user_id"`
	Role      int      `json:"role"`
	Status    int      `json:"status"`
	CreatedAt DateTime `json:"created_at"`
	UpdatedAt DateTime `json:"updated_at"`
	JoinedAt  DateTime `json:"joined_at"`
	LeftAt    DateTime `json:"left_at"`
}

// DialogMessage 对话消息结构
type DialogMessage struct {
	ID        int      `json:"id"`
	DialogID  int      `json:"dialog_id"`
	UserID    int      `json:"user_id"`
	Type      int      `json:"type"`
	Content   string   `json:"content"`
	Extra     string   `json:"extra"`
	ReplyID   int      `json:"reply_id"`
	CreatedAt DateTime `json:"created_at"`
	UpdatedAt DateTime `json:"updated_at"`
	Status    int      `json:"status"`
	IsRead    int      `json:"is_read"`
	IsDelete  int      `json:"is_delete"`
}

// CreateDialogRequest 创建对话请求
type CreateDialogRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int    `json:"type"`
	Avatar      string `json:"avatar"`
	Members     []int  `json:"members"`
}

// UpdateDialogRequest 更新对话请求
type UpdateDialogRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	IsTop       *int   `json:"is_top,omitempty"`
	IsMute      *int   `json:"is_mute,omitempty"`
	IsArchive   *int   `json:"is_archive,omitempty"`
}

// AddMemberRequest 添加成员请求
type AddMemberRequest struct {
	UserIDs []int `json:"user_ids"`
}

// RemoveMemberRequest 移除成员请求
type RemoveMemberRequest struct {
	UserIDs []int `json:"user_ids"`
}

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	DialogID   int    `json:"dialog_id"`
	Text       string `json:"text"`
	TextType   int    `json:"text_type"`
	UpdateID   int    `json:"update_id"`
	UpdateMark string `json:"update_mark"`
	ReplyID    int    `json:"reply_id,omitempty"`
	Silence    string `json:"silence"`
}

// DialogListRequest 对话列表请求
type DialogListRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Type     int `json:"type,omitempty"`
	Status   int `json:"status,omitempty"`
}

// DialogListResponse 对话列表响应
type DialogListResponse struct {
	CurrentPage  int      `json:"current_page"`
	Data         []Dialog `json:"data"`
	FirstPageURL string   `json:"first_page_url"`
	From         int      `json:"from"`
	LastPage     int      `json:"last_page"`
	LastPageURL  string   `json:"last_page_url"`
	Links        []struct {
		URL    interface{} `json:"url"`
		Label  string      `json:"label"`
		Active bool        `json:"active"`
	} `json:"links"`
	NextPageURL interface{}   `json:"next_page_url"`
	Path        string        `json:"path"`
	PerPage     int           `json:"per_page"`
	PrevPageURL interface{}   `json:"prev_page_url"`
	To          int           `json:"to"`
	Total       int           `json:"total"`
	DeletedID   []interface{} `json:"deleted_id"`
}

// DialogDetailResponse 对话详情响应
type DialogDetailResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Dialog `json:"data"`
}

// DialogMembersResponse 对话成员响应
type DialogMembersResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    []DialogMember `json:"data"`
}

// DialogMessagesRequest 对话消息请求
type DialogMessagesRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	LastID   int `json:"last_id,omitempty"`
}

// DialogMessagesResponse 对话消息响应
type DialogMessagesResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []DialogMessage `json:"data"`
	Total   int             `json:"total"`
}

// DialogMessageResponse 对话消息响应
type DialogMessageResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    DialogMessage `json:"data"`
}
