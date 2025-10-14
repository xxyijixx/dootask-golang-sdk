### **dialog 模块接口**

| 序号 | 接口名称 | 接口地址 |
|------|----------|----------|
| 1 | 对话列表 | `GET /api/dialog/lists` |
| 2 | 搜索会话 | `GET /api/dialog/search` |
| 3 | 获取单个会话信息 | `GET /api/dialog/one` |
| 4 | 获取会话成员 | `GET /api/dialog/user` |
| 5 | 获取会话待办 | `GET /api/dialog/todo` |
| 6 | 会话置顶 | `GET /api/dialog/top` |
| 7 | 获取对方联系电话 | `GET /api/dialog/tel` |
| 8 | 打开会话 | `GET /api/dialog/open/user` |
| 9 | 获取消息列表 | `GET /api/dialog/msg/list` |
| 10 | 搜索消息位置 | `GET /api/dialog/msg/search` |
| 11 | 获取单条消息 | `GET /api/dialog/msg/one` |
| 12 | 已读聊天消息 | `GET /api/dialog/msg/read` |
| 13 | 获取未读消息数据 | `GET /api/dialog/msg/unread` |
| 14 | 通知成员监听消息 | `POST /api/dialog/msg/stream` |
| 15 | 发送消息 | `POST /api/dialog/msg/sendtext` |
| 16 | 发送语音 | `POST /api/dialog/msg/sendrecord` |
| 17 | 文件上传 | `POST /api/dialog/msg/sendfile` |
| 18 | 群发文件上传 | `POST /api/dialog/msg/sendfiles` |
| 19 | 通过文件ID发送文件 | `GET /api/dialog/msg/sendfileid` |
| 20 | 发送匿名消息 | `POST /api/dialog/msg/sendanon` |
| 21 | 获取消息阅读情况 | `GET /api/dialog/msg/readlist` |
| 22 | 消息详情 | `GET /api/dialog/msg/detail` |
| 23 | 文件下载 | `GET /api/dialog/msg/download` |
| 24 | 聊天消息撤回 | `GET /api/dialog/msg/withdraw` |
| 25 | 消息标记操作 | `GET /api/dialog/msg/mark` |
| 26 | 消息免打扰 | `GET /api/dialog/msg/silence` |
| 27 | 转发消息给 | `GET /api/dialog/msg/forward` |
| 28 | emoji回复 | `GET /api/dialog/msg/emoji` |
| 29 | 标注/取消标注 | `GET /api/dialog/msg/tag` |
| 30 | 设待办/取消待办 | `GET /api/dialog/msg/todo` |
| 31 | 获取消息待办情况 | `GET /api/dialog/msg/todolist` |
| 32 | 完成待办 | `GET /api/dialog/msg/done` |
| 33 | 设置颜色 | `GET /api/dialog/msg/color` |
| 34 | 新增群组 | `GET /api/dialog/group/add` |
| 35 | 修改群组 | `GET /api/dialog/group/edit` |
| 36 | 添加群成员 | `GET /api/dialog/group/adduser` |
| 37 | 移出（退出）群成员 | `GET /api/dialog/group/deluser` |
| 38 | 转让群组 | `GET /api/dialog/group/transfer` |
| 39 | 解散群组 | `GET /api/dialog/group/disband` |
| 40 | 搜索个人群（仅限管理员） | `GET /api/dialog/group/searchuser` |
| 41 | 创建OKR评论会话 | `POST /api/dialog/okr/add` |
| 42 | 推送OKR相关信息 | `POST /api/dialog/okr/push` |
