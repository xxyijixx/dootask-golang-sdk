package project

import (
	"fmt"

	"github.com/xxyijixx/dootask-golang-sdk/internal/core"
	ihttp "github.com/xxyijixx/dootask-golang-sdk/internal/http"
	"github.com/xxyijixx/dootask-golang-sdk/types"
)

// Service 项目服务
type Service struct {
	client core.HTTPDoer
}

// New 创建项目服务实例
func New(client core.HTTPDoer) *Service {
	return &Service{
		client: client,
	}
}

// ==================== 项目管理 ====================

// GetProjectLists 01. 获取项目列表
func (s *Service) GetProjectLists(req *types.ProjectListsRequest) (*types.ProjectListsResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/lists", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectListsResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetProjectOne 02. 获取一个项目信息
func (s *Service) GetProjectOne(req *types.ProjectOneRequest) (*types.ProjectOneResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/one", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectOneResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// AddProject 03. 添加项目
func (s *Service) AddProject(req *types.ProjectAddRequest) (*types.ProjectAddResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/add", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectAddResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// UpdateProject 04. 修改项目
func (s *Service) UpdateProject(req *types.ProjectUpdateRequest) (*types.ProjectUpdateResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/update", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectUpdateResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// TransferProject 09. 移交项目
func (s *Service) TransferProject(req *types.ProjectTransferRequest) (*types.ProjectTransferResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/transfer", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTransferResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ExitProject 11. 退出项目
func (s *Service) ExitProject(req *types.ProjectExitRequest) (*types.ProjectExitResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/exit", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectExitResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ArchivedProject 12. 归档项目
func (s *Service) ArchivedProject(req *types.ProjectArchivedRequest) (*types.ProjectArchivedResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/archived", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectArchivedResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// RemoveProject 13. 删除项目
func (s *Service) RemoveProject(req *types.ProjectRemoveRequest) (*types.ProjectRemoveResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/remove", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectRemoveResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// TopProject 43. 项目置顶
func (s *Service) TopProject(req *types.ProjectTopRequest) (*types.ProjectTopResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/top", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTopResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 项目成员管理 ====================

// ManageProjectUser 05. 修改项目成员
func (s *Service) ManageProjectUser(req *types.ProjectUserRequest) (*types.ProjectUserResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/user", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectUserResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetProjectInvite 06. 获取邀请链接
func (s *Service) GetProjectInvite(req *types.ProjectInviteRequest) (*types.ProjectInviteResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/invite", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectInviteResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetProjectInviteInfo 07. 通过邀请链接code获取项目信息
func (s *Service) GetProjectInviteInfo(req *types.ProjectInviteInfoRequest) (*types.ProjectInviteInfoResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/invite/info", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectInviteInfoResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// JoinProjectByInvite 08. 通过邀请链接code加入项目
func (s *Service) JoinProjectByInvite(req *types.ProjectInviteJoinRequest) (*types.ProjectInviteJoinResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/invite/join", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectInviteJoinResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 任务列表（列）管理 ====================

// GetColumnLists 14. 获取任务列表
func (s *Service) GetColumnLists(req *types.ProjectColumnListsRequest) (*types.ProjectColumnListsResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/column/lists", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectColumnListsResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// AddColumn 15. 添加任务列表
func (s *Service) AddColumn(req *types.ProjectColumnAddRequest) (*types.ProjectColumnAddResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/column/add", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectColumnAddResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// UpdateColumn 16. 修改任务列表
func (s *Service) UpdateColumn(req *types.ProjectColumnUpdateRequest) (*types.ProjectColumnUpdateResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/column/update", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectColumnUpdateResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// RemoveColumn 17. 删除任务列表
func (s *Service) RemoveColumn(req *types.ProjectColumnRemoveRequest) (*types.ProjectColumnRemoveResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/column/remove", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectColumnRemoveResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetColumnOne 18. 获取任务列详细
func (s *Service) GetColumnOne(req *types.ProjectColumnOneRequest) (*types.ProjectColumnOneResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/column/one", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectColumnOneResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 任务查询 ====================

// GetTaskLists 19. 任务列表
func (s *Service) GetTaskLists(req *types.ProjectTaskListsRequest) (*types.ProjectTaskListsResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/lists", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskListsResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetTaskEasyLists 20. 任务列表-简单的
func (s *Service) GetTaskEasyLists(req *types.ProjectTaskEasyListsRequest) (*types.ProjectTaskEasyListsResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/easylists", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskEasyListsResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetTaskOne 24. 获取单个任务信息
func (s *Service) GetTaskOne(req *types.ProjectTaskOneRequest) (*types.ProjectTaskOneResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/one", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskOneResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetTaskContent 25. 获取任务详细描述
func (s *Service) GetTaskContent(req *types.ProjectTaskContentRequest) (*types.ProjectTaskContentResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/content", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskContentResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 任务操作 ====================

// AddTask 30. 添加任务
func (s *Service) AddTask(req *types.ProjectTaskAddRequest) (*types.ProjectTaskAddResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/project/task/add", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskAddResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// AddSubTask 31. 添加子任务
func (s *Service) AddSubTask(req *types.ProjectTaskAddSubRequest) (*types.ProjectTaskAddSubResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/addsub", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskAddSubResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// UpdateTask 32. 修改任务、子任务
func (s *Service) UpdateTask(req *types.ProjectTaskUpdateRequest) (*types.ProjectTaskUpdateResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/project/task/update", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskUpdateResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetTaskDialog 33. 创建/获取聊天室
func (s *Service) GetTaskDialog(req *types.ProjectTaskDialogRequest) (*types.ProjectTaskDialogResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/dialog", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskDialogResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ArchivedTask 34. 归档任务
func (s *Service) ArchivedTask(req *types.ProjectTaskArchivedRequest) (*types.ProjectTaskArchivedResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/archived", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskArchivedResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// RemoveTask 35. 删除任务
func (s *Service) RemoveTask(req *types.ProjectTaskRemoveRequest) (*types.ProjectTaskRemoveResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/remove", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskRemoveResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ResetTaskFromLog 36. 根据日志重置任务
func (s *Service) ResetTaskFromLog(req *types.ProjectTaskResetFromLogRequest) (*types.ProjectTaskResetFromLogResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/resetfromlog", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskResetFromLogResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// MoveTask 38. 任务移动
func (s *Service) MoveTask(req *types.ProjectTaskMoveRequest) (*types.ProjectTaskMoveResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/move", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskMoveResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 任务排序 ====================

// SortTask 10. 排序任务
func (s *Service) SortTask(req *types.ProjectSortRequest) (*types.ProjectSortResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/project/sort", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectSortResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 文件管理 ====================

// GetTaskFiles 26. 获取任务文件列表
func (s *Service) GetTaskFiles(req *types.ProjectTaskFilesRequest) (*types.ProjectTaskFilesResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/files", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskFilesResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DeleteTaskFile 27. 删除任务文件
func (s *Service) DeleteTaskFile(req *types.ProjectTaskFileDeleteRequest) (*types.ProjectTaskFileDeleteResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/filedelete", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskFileDeleteResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetTaskFileDetail 28. 获取任务文件详情
func (s *Service) GetTaskFileDetail(req *types.ProjectTaskFileDetailRequest) (*types.ProjectTaskFileDetailResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/filedetail", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskFileDetailResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DownloadTaskFile 29. 下载任务文件
func (s *Service) DownloadTaskFile(req *types.ProjectTaskFileDownRequest) (*types.ProjectTaskFileDownResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/filedown", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskFileDownResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 工作流管理 ====================

// GetTaskFlow 37. 任务工作流信息
func (s *Service) GetTaskFlow(req *types.ProjectTaskFlowRequest) (*types.ProjectTaskFlowResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/flow", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskFlowResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetFlowList 39. 工作流列表
func (s *Service) GetFlowList(req *types.ProjectFlowListRequest) (*types.ProjectFlowListResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/flow/list", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectFlowListResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SaveFlow 40. 保存工作流
func (s *Service) SaveFlow(req *types.ProjectFlowSaveRequest) (*types.ProjectFlowSaveResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/project/flow/save", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectFlowSaveResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DeleteFlow 41. 删除工作流
func (s *Service) DeleteFlow(req *types.ProjectFlowDeleteRequest) (*types.ProjectFlowDeleteResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/flow/delete", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectFlowDeleteResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 导出功能 ====================

// ExportTask 21. 导出任务（限管理员）
func (s *Service) ExportTask(req *types.ProjectTaskExportRequest) (*types.ProjectTaskExportResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/export", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskExportResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ExportOverdueTask 22. 导出超期任务（限管理员）
func (s *Service) ExportOverdueTask(req *types.ProjectTaskExportOverdueRequest) (*types.ProjectTaskExportOverdueResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/exportoverdue", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskExportOverdueResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DownloadExportedTask 23. 下载导出的任务
func (s *Service) DownloadExportedTask(req *types.ProjectTaskDownRequest) (*types.ProjectTaskDownResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/down", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskDownResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 日志管理 ====================

// GetLogLists 42. 获取项目、任务日志
func (s *Service) GetLogLists(req *types.ProjectLogListsRequest) (*types.ProjectLogListsResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/log/lists", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectLogListsResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 项目权限管理 ====================

// GetProjectPermission 44. 获取项目权限设置
func (s *Service) GetProjectPermission(req *types.ProjectPermissionRequest) (*types.ProjectPermissionResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/permission", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectPermissionResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// UpdateProjectPermission 45. 项目权限设置
func (s *Service) UpdateProjectPermission(req *types.ProjectPermissionUpdateRequest) (*types.ProjectPermissionUpdateResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/permission/update", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectPermissionUpdateResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 标签管理 ====================

// GetProjectTagList 47. 标签列表
func (s *Service) GetProjectTagList(req *types.ProjectTagListRequest) (*types.ProjectTagListResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/tag/list", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTagListResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SaveProjectTag 48. 保存标签
func (s *Service) SaveProjectTag(req *types.ProjectTagSaveRequest) (*types.ProjectTagSaveResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/project/tag/save", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTagSaveResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SortProjectTag 49. 标签排序
func (s *Service) SortProjectTag(req *types.ProjectTagSortRequest) (*types.ProjectTagSortResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/project/tag/sort", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTagSortResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DeleteProjectTag 46. 删除标签
func (s *Service) DeleteProjectTag(req *types.ProjectTagDeleteRequest) (*types.ProjectTagDeleteResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/tag/delete", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTagDeleteResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 任务历史与关联 ====================

// GetTaskContentHistory 50. 获取任务详细历史描述
func (s *Service) GetTaskContentHistory(req *types.ProjectTaskContentHistoryRequest) (*types.ProjectTaskContentHistoryResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/content_history", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskContentHistoryResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// CopyTask 51. 复制任务
func (s *Service) CopyTask(req *types.ProjectTaskCopyRequest) (*types.ProjectTaskCopyResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/project/task/copy", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskCopyResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetTaskRelated 52. 获取任务关联任务列表
func (s *Service) GetTaskRelated(req *types.ProjectTaskRelatedRequest) (*types.ProjectTaskRelatedResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/related", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskRelatedResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetTaskSubdata 53. 获取子任务数据
func (s *Service) GetTaskSubdata(req *types.ProjectTaskSubdataRequest) (*types.ProjectTaskSubdataResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/subdata", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskSubdataResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 任务模板管理 ====================

// GetTaskTemplateList 56. 任务模板列表
func (s *Service) GetTaskTemplateList(req *types.ProjectTaskTemplateListRequest) (*types.ProjectTaskTemplateListResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/template_list", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskTemplateListResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SaveTaskTemplate 57. 保存任务模板
func (s *Service) SaveTaskTemplate(req *types.ProjectTaskTemplateSaveRequest) (*types.ProjectTaskTemplateSaveResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/project/task/template_save", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskTemplateSaveResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SortTaskTemplate 58. 排序任务模板
func (s *Service) SortTaskTemplate(req *types.ProjectTaskTemplateSortRequest) (*types.ProjectTaskTemplateSortResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/project/task/template_sort", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskTemplateSortResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// DeleteTaskTemplate 55. 删除任务模板
func (s *Service) DeleteTaskTemplate(req *types.ProjectTaskTemplateDeleteRequest) (*types.ProjectTaskTemplateDeleteResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/template_delete", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskTemplateDeleteResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// SetTaskTemplateDefault 54. 设置(取消)任务模板为默认
func (s *Service) SetTaskTemplateDefault(req *types.ProjectTaskTemplateDefaultRequest) (*types.ProjectTaskTemplateDefaultResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/template_default", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskTemplateDefaultResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 任务操作 ====================

// UpgradeTask 59. 子任务升级为主任务
func (s *Service) UpgradeTask(req *types.ProjectTaskUpgradeRequest) (*types.ProjectTaskUpgradeResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/project/task/upgrade", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectTaskUpgradeResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 项目排序 ====================

// SortProjectUser 60. 项目列表排序
func (s *Service) SortProjectUser(req *types.ProjectUserSortRequest) (*types.ProjectUserSortResponse, error) {
	resp, err := s.client.DoRequest("POST", "/api/project/user/sort", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ProjectUserSortResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}
