package report

import (
	"fmt"

	"github.com/xxyijixx/dootask-golang-sdk/internal/core"
	ihttp "github.com/xxyijixx/dootask-golang-sdk/internal/http"
	"github.com/xxyijixx/dootask-golang-sdk/types"
)

// Service 报告服务
type Service struct {
	client core.HTTPDoer
}

// New 创建报告服务实例
func New(client core.HTTPDoer) *Service {
	return &Service{
		client: client,
	}
}

// ==================== 报告查看 ====================

// GetMyReports 01. 我发送的汇报
// 查看我发送的所有汇报，支持按类型和时间搜索
func (s *Service) GetMyReports(req *types.ReportMyRequest) (*types.ReportMyResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/report/my", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReportMyResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetReceiveReports 02. 我接收的汇报
// 查看我接收到的汇报，支持关键词、类型和时间搜索
func (s *Service) GetReceiveReports(req *types.ReportReceiveRequest) (*types.ReportReceiveResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/report/receive", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReportReceiveResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetReportDetail 05. 报告详情
// 查看具体报告的详细信息，包括阅读状态等
func (s *Service) GetReportDetail(req *types.ReportDetailRequest) (*types.ReportDetailResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/report/detail", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReportDetailResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 报告创建与发送 ====================

// StoreReport 03. 保存并发送工作汇报
// 创建新报告或修改现有报告，支持设置标题、类型、内容和接收人
func (s *Service) StoreReport(req *types.ReportStoreRequest) (*types.ReportStoreResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/report/store", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReportStoreResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GenerateReportTemplate 04. 生成汇报模板
// 自动生成报告模板，支持周报、日报类型，会自动填充相关任务
func (s *Service) GenerateReportTemplate(req *types.ReportTemplateRequest) (*types.ReportTemplateResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/report/template", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReportTemplateResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 阅读状态管理 ====================

// MarkReport 06. 标记已读/未读
// 单个报告阅读状态标记，支持已读/未读切换
func (s *Service) MarkReport(req *types.ReportMarkRequest) (*types.ReportMarkResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/report/mark", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReportMarkResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// MarkReportsRead 09. 批量标记汇报已读
// 批量标记多个报告为已读状态
func (s *Service) MarkReportsRead(req *types.ReportReadRequest) (*types.ReportReadResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/report/read", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReportReadResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// GetUnreadReports 08. 获取未读
// 获取未读报告统计，可按类型分组查看
func (s *Service) GetUnreadReports(req *types.ReportUnreadRequest) (*types.ReportUnreadResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/report/unread", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReportUnreadResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}

// ==================== 辅助功能 ====================

// GetLastSubmitter 07. 获取最后一次提交的接收人
// 获取上次提交的接收人信息，便于快速填写新报告的接收人
func (s *Service) GetLastSubmitter(req *types.ReportLastSubmitterRequest) (*types.ReportLastSubmitterResponse, error) {
	resp, err := s.client.DoRequest("GET", "/api/report/last_submitter", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result types.ReportLastSubmitterResponse
	err = ihttp.ParseAPIResponse(resp, &result)

	if err != nil {
		return nil, fmt.Errorf("API error: %s", err.Error())
	}

	return &result, nil
}
