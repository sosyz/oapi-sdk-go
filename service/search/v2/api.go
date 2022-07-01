// Package search code generated by oapi sdk gen
package larksearch

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/core"
)

// 构建业务域服务实例
func NewService(httpClient *http.Client, config *core.Config) *SearchService {
	s := &SearchService{httpClient: httpClient, config: config}
	s.DataSource = &dataSource{service: s}
	s.DataSourceItem = &dataSourceItem{service: s}
	return s
}

// 业务域服务定义
type SearchService struct {
	httpClient     *http.Client
	config         *core.Config
	DataSource     *dataSource
	DataSourceItem *dataSourceItem
}

// 资源服务定义
type dataSource struct {
	service *SearchService
}
type dataSourceItem struct {
	service *SearchService
}

// 资源服务方法定义
func (d *dataSource) Create(ctx context.Context, req *CreateDataSourceReq, options ...core.RequestOptionFunc) (*CreateDataSourceResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, d.service.config, http.MethodPost,
		"/open-apis/search/v2/data_sources", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDataSourceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSource) Delete(ctx context.Context, req *DeleteDataSourceReq, options ...core.RequestOptionFunc) (*DeleteDataSourceResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, d.service.config, http.MethodDelete,
		"/open-apis/search/v2/data_sources/:data_source_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteDataSourceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSource) Get(ctx context.Context, req *GetDataSourceReq, options ...core.RequestOptionFunc) (*GetDataSourceResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/search/v2/data_sources/:data_source_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDataSourceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSource) List(ctx context.Context, req *ListDataSourceReq, options ...core.RequestOptionFunc) (*ListDataSourceResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/search/v2/data_sources", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListDataSourceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSource) ListDataSource(ctx context.Context, req *ListDataSourceReq, options ...core.RequestOptionFunc) (*ListDataSourceIterator, error) {
	return &ListDataSourceIterator{
		ctx:      ctx,
		req:      req,
		listFunc: d.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (d *dataSource) Patch(ctx context.Context, req *PatchDataSourceReq, options ...core.RequestOptionFunc) (*PatchDataSourceResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, d.service.config, http.MethodPatch,
		"/open-apis/search/v2/data_sources/:data_source_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchDataSourceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSourceItem) Create(ctx context.Context, req *CreateDataSourceItemReq, options ...core.RequestOptionFunc) (*CreateDataSourceItemResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, d.service.config, http.MethodPost,
		"/open-apis/search/v2/data_sources/:data_source_id/items", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDataSourceItemResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSourceItem) Delete(ctx context.Context, req *DeleteDataSourceItemReq, options ...core.RequestOptionFunc) (*DeleteDataSourceItemResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, d.service.config, http.MethodDelete,
		"/open-apis/search/v2/data_sources/:data_source_id/items/:item_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteDataSourceItemResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSourceItem) Get(ctx context.Context, req *GetDataSourceItemReq, options ...core.RequestOptionFunc) (*GetDataSourceItemResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/search/v2/data_sources/:data_source_id/items/:item_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDataSourceItemResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
