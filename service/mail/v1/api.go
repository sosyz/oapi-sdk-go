// Package mail code generated by oapi sdk gen
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

package larkmail

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *MailService {
	m := &MailService{config: config}
	m.Mailgroup = &mailgroup{service: m}
	m.MailgroupAlias = &mailgroupAlias{service: m}
	m.MailgroupMember = &mailgroupMember{service: m}
	m.MailgroupPermissionMember = &mailgroupPermissionMember{service: m}
	m.PublicMailbox = &publicMailbox{service: m}
	m.PublicMailboxAlias = &publicMailboxAlias{service: m}
	m.PublicMailboxMember = &publicMailboxMember{service: m}
	m.User = &user{service: m}
	m.UserMailbox = &userMailbox{service: m}
	m.UserMailboxAlias = &userMailboxAlias{service: m}
	return m
}

// 业务域服务定义
type MailService struct {
	config                    *larkcore.Config
	Mailgroup                 *mailgroup
	MailgroupAlias            *mailgroupAlias
	MailgroupMember           *mailgroupMember
	MailgroupPermissionMember *mailgroupPermissionMember
	PublicMailbox             *publicMailbox
	PublicMailboxAlias        *publicMailboxAlias
	PublicMailboxMember       *publicMailboxMember
	User                      *user
	UserMailbox               *userMailbox
	UserMailboxAlias          *userMailboxAlias
}

// 资源服务定义
type mailgroup struct {
	service *MailService
}
type mailgroupAlias struct {
	service *MailService
}
type mailgroupMember struct {
	service *MailService
}
type mailgroupPermissionMember struct {
	service *MailService
}
type publicMailbox struct {
	service *MailService
}
type publicMailboxAlias struct {
	service *MailService
}
type publicMailboxMember struct {
	service *MailService
}
type user struct {
	service *MailService
}
type userMailbox struct {
	service *MailService
}
type userMailboxAlias struct {
	service *MailService
}

// 资源服务方法定义
func (m *mailgroup) Create(ctx context.Context, req *CreateMailgroupReq, options ...larkcore.RequestOptionFunc) (*CreateMailgroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateMailgroupResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroup) Delete(ctx context.Context, req *DeleteMailgroupReq, options ...larkcore.RequestOptionFunc) (*DeleteMailgroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteMailgroupResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroup) Get(ctx context.Context, req *GetMailgroupReq, options ...larkcore.RequestOptionFunc) (*GetMailgroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetMailgroupResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroup) List(ctx context.Context, req *ListMailgroupReq, options ...larkcore.RequestOptionFunc) (*ListMailgroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListMailgroupResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroup) ListByIterator(ctx context.Context, req *ListMailgroupReq, options ...larkcore.RequestOptionFunc) (*ListMailgroupIterator, error) {
	return &ListMailgroupIterator{
		ctx:      ctx,
		req:      req,
		listFunc: m.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (m *mailgroup) Patch(ctx context.Context, req *PatchMailgroupReq, options ...larkcore.RequestOptionFunc) (*PatchMailgroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchMailgroupResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroup) Update(ctx context.Context, req *UpdateMailgroupReq, options ...larkcore.RequestOptionFunc) (*UpdateMailgroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id"
	apiReq.HttpMethod = http.MethodPut
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateMailgroupResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupAlias) Create(ctx context.Context, req *CreateMailgroupAliasReq, options ...larkcore.RequestOptionFunc) (*CreateMailgroupAliasResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateMailgroupAliasResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupAlias) Delete(ctx context.Context, req *DeleteMailgroupAliasReq, options ...larkcore.RequestOptionFunc) (*DeleteMailgroupAliasResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases/:alias_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteMailgroupAliasResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupAlias) List(ctx context.Context, req *ListMailgroupAliasReq, options ...larkcore.RequestOptionFunc) (*ListMailgroupAliasResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListMailgroupAliasResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupMember) Create(ctx context.Context, req *CreateMailgroupMemberReq, options ...larkcore.RequestOptionFunc) (*CreateMailgroupMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id/members"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateMailgroupMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupMember) Delete(ctx context.Context, req *DeleteMailgroupMemberReq, options ...larkcore.RequestOptionFunc) (*DeleteMailgroupMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteMailgroupMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupMember) Get(ctx context.Context, req *GetMailgroupMemberReq, options ...larkcore.RequestOptionFunc) (*GetMailgroupMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetMailgroupMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupMember) List(ctx context.Context, req *ListMailgroupMemberReq, options ...larkcore.RequestOptionFunc) (*ListMailgroupMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id/members"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListMailgroupMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupMember) ListByIterator(ctx context.Context, req *ListMailgroupMemberReq, options ...larkcore.RequestOptionFunc) (*ListMailgroupMemberIterator, error) {
	return &ListMailgroupMemberIterator{
		ctx:      ctx,
		req:      req,
		listFunc: m.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (m *mailgroupPermissionMember) Create(ctx context.Context, req *CreateMailgroupPermissionMemberReq, options ...larkcore.RequestOptionFunc) (*CreateMailgroupPermissionMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateMailgroupPermissionMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupPermissionMember) Delete(ctx context.Context, req *DeleteMailgroupPermissionMemberReq, options ...larkcore.RequestOptionFunc) (*DeleteMailgroupPermissionMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/:permission_member_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteMailgroupPermissionMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupPermissionMember) Get(ctx context.Context, req *GetMailgroupPermissionMemberReq, options ...larkcore.RequestOptionFunc) (*GetMailgroupPermissionMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/:permission_member_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetMailgroupPermissionMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupPermissionMember) List(ctx context.Context, req *ListMailgroupPermissionMemberReq, options ...larkcore.RequestOptionFunc) (*ListMailgroupPermissionMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListMailgroupPermissionMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *mailgroupPermissionMember) ListByIterator(ctx context.Context, req *ListMailgroupPermissionMemberReq, options ...larkcore.RequestOptionFunc) (*ListMailgroupPermissionMemberIterator, error) {
	return &ListMailgroupPermissionMemberIterator{
		ctx:      ctx,
		req:      req,
		listFunc: m.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (p *publicMailbox) Create(ctx context.Context, req *CreatePublicMailboxReq, options ...larkcore.RequestOptionFunc) (*CreatePublicMailboxResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreatePublicMailboxResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailbox) Delete(ctx context.Context, req *DeletePublicMailboxReq, options ...larkcore.RequestOptionFunc) (*DeletePublicMailboxResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeletePublicMailboxResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailbox) Get(ctx context.Context, req *GetPublicMailboxReq, options ...larkcore.RequestOptionFunc) (*GetPublicMailboxResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetPublicMailboxResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailbox) List(ctx context.Context, req *ListPublicMailboxReq, options ...larkcore.RequestOptionFunc) (*ListPublicMailboxResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListPublicMailboxResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailbox) ListByIterator(ctx context.Context, req *ListPublicMailboxReq, options ...larkcore.RequestOptionFunc) (*ListPublicMailboxIterator, error) {
	return &ListPublicMailboxIterator{
		ctx:      ctx,
		req:      req,
		listFunc: p.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (p *publicMailbox) Patch(ctx context.Context, req *PatchPublicMailboxReq, options ...larkcore.RequestOptionFunc) (*PatchPublicMailboxResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchPublicMailboxResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailbox) Update(ctx context.Context, req *UpdatePublicMailboxReq, options ...larkcore.RequestOptionFunc) (*UpdatePublicMailboxResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id"
	apiReq.HttpMethod = http.MethodPut
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdatePublicMailboxResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailboxAlias) Create(ctx context.Context, req *CreatePublicMailboxAliasReq, options ...larkcore.RequestOptionFunc) (*CreatePublicMailboxAliasResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreatePublicMailboxAliasResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailboxAlias) Delete(ctx context.Context, req *DeletePublicMailboxAliasReq, options ...larkcore.RequestOptionFunc) (*DeletePublicMailboxAliasResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases/:alias_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeletePublicMailboxAliasResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailboxAlias) List(ctx context.Context, req *ListPublicMailboxAliasReq, options ...larkcore.RequestOptionFunc) (*ListPublicMailboxAliasResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListPublicMailboxAliasResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailboxMember) Clear(ctx context.Context, req *ClearPublicMailboxMemberReq, options ...larkcore.RequestOptionFunc) (*ClearPublicMailboxMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/clear"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ClearPublicMailboxMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailboxMember) Create(ctx context.Context, req *CreatePublicMailboxMemberReq, options ...larkcore.RequestOptionFunc) (*CreatePublicMailboxMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreatePublicMailboxMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailboxMember) Delete(ctx context.Context, req *DeletePublicMailboxMemberReq, options ...larkcore.RequestOptionFunc) (*DeletePublicMailboxMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeletePublicMailboxMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailboxMember) Get(ctx context.Context, req *GetPublicMailboxMemberReq, options ...larkcore.RequestOptionFunc) (*GetPublicMailboxMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetPublicMailboxMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailboxMember) List(ctx context.Context, req *ListPublicMailboxMemberReq, options ...larkcore.RequestOptionFunc) (*ListPublicMailboxMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListPublicMailboxMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *publicMailboxMember) ListByIterator(ctx context.Context, req *ListPublicMailboxMemberReq, options ...larkcore.RequestOptionFunc) (*ListPublicMailboxMemberIterator, error) {
	return &ListPublicMailboxMemberIterator{
		ctx:      ctx,
		req:      req,
		listFunc: p.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (u *user) Query(ctx context.Context, req *QueryUserReq, options ...larkcore.RequestOptionFunc) (*QueryUserResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/users/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userMailbox) Delete(ctx context.Context, req *DeleteUserMailboxReq, options ...larkcore.RequestOptionFunc) (*DeleteUserMailboxResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/user_mailboxes/:user_mailbox_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteUserMailboxResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userMailboxAlias) Create(ctx context.Context, req *CreateUserMailboxAliasReq, options ...larkcore.RequestOptionFunc) (*CreateUserMailboxAliasResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateUserMailboxAliasResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userMailboxAlias) Delete(ctx context.Context, req *DeleteUserMailboxAliasReq, options ...larkcore.RequestOptionFunc) (*DeleteUserMailboxAliasResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases/:alias_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteUserMailboxAliasResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userMailboxAlias) List(ctx context.Context, req *ListUserMailboxAliasReq, options ...larkcore.RequestOptionFunc) (*ListUserMailboxAliasResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListUserMailboxAliasResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
