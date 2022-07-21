// Package docx code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkdocx

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *DocxService {
	d := &DocxService{config: config}
	d.Document = &document{service: d}
	d.DocumentBlock = &documentBlock{service: d}
	d.DocumentBlockChildren = &documentBlockChildren{service: d}
	return d
}

// 业务域服务定义
type DocxService struct {
	config                *larkcore.Config
	Document              *document
	DocumentBlock         *documentBlock
	DocumentBlockChildren *documentBlockChildren
}

// 资源服务定义
type document struct {
	service *DocxService
}
type documentBlock struct {
	service *DocxService
}
type documentBlockChildren struct {
	service *DocxService
}

// 资源服务方法定义
func (d *document) Create(ctx context.Context, req *CreateDocumentReq, options ...larkcore.RequestOptionFunc) (*CreateDocumentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/docx/v1/documents"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDocumentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *document) Get(ctx context.Context, req *GetDocumentReq, options ...larkcore.RequestOptionFunc) (*GetDocumentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/docx/v1/documents/:document_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDocumentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *document) RawContent(ctx context.Context, req *RawContentDocumentReq, options ...larkcore.RequestOptionFunc) (*RawContentDocumentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/docx/v1/documents/:document_id/raw_content"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RawContentDocumentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlock) BatchUpdate(ctx context.Context, req *BatchUpdateDocumentBlockReq, options ...larkcore.RequestOptionFunc) (*BatchUpdateDocumentBlockResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/docx/v1/documents/:document_id/blocks/batch_update"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchUpdateDocumentBlockResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlock) Get(ctx context.Context, req *GetDocumentBlockReq, options ...larkcore.RequestOptionFunc) (*GetDocumentBlockResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/docx/v1/documents/:document_id/blocks/:block_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDocumentBlockResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlock) List(ctx context.Context, req *ListDocumentBlockReq, options ...larkcore.RequestOptionFunc) (*ListDocumentBlockResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/docx/v1/documents/:document_id/blocks"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListDocumentBlockResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlock) ListByIterator(ctx context.Context, req *ListDocumentBlockReq, options ...larkcore.RequestOptionFunc) (*ListDocumentBlockIterator, error) {
	return &ListDocumentBlockIterator{
		ctx:      ctx,
		req:      req,
		listFunc: d.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (d *documentBlock) Patch(ctx context.Context, req *PatchDocumentBlockReq, options ...larkcore.RequestOptionFunc) (*PatchDocumentBlockResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/docx/v1/documents/:document_id/blocks/:block_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchDocumentBlockResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlockChildren) BatchDelete(ctx context.Context, req *BatchDeleteDocumentBlockChildrenReq, options ...larkcore.RequestOptionFunc) (*BatchDeleteDocumentBlockChildrenResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children/batch_delete"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchDeleteDocumentBlockChildrenResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlockChildren) Create(ctx context.Context, req *CreateDocumentBlockChildrenReq, options ...larkcore.RequestOptionFunc) (*CreateDocumentBlockChildrenResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDocumentBlockChildrenResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlockChildren) Get(ctx context.Context, req *GetDocumentBlockChildrenReq, options ...larkcore.RequestOptionFunc) (*GetDocumentBlockChildrenResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDocumentBlockChildrenResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *documentBlockChildren) GetByIterator(ctx context.Context, req *GetDocumentBlockChildrenReq, options ...larkcore.RequestOptionFunc) (*GetDocumentBlockChildrenIterator, error) {
	return &GetDocumentBlockChildrenIterator{
		ctx:      ctx,
		req:      req,
		listFunc: d.Get,
		options:  options,
		limit:    req.Limit}, nil
}
