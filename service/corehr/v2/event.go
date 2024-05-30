// Package corehr code generated by oapi sdk gen
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

package larkcorehr

import (
	"context"
)

// 消息处理器定义
type P2OffboardingChecklistUpdatedV2Handler struct {
	handler func(context.Context, *P2OffboardingChecklistUpdatedV2) error
}

func NewP2OffboardingChecklistUpdatedV2Handler(handler func(context.Context, *P2OffboardingChecklistUpdatedV2) error) *P2OffboardingChecklistUpdatedV2Handler {
	h := &P2OffboardingChecklistUpdatedV2Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2OffboardingChecklistUpdatedV2Handler) Event() interface{} {
	return &P2OffboardingChecklistUpdatedV2{}
}

// 回调开发者注册的handle
func (h *P2OffboardingChecklistUpdatedV2Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2OffboardingChecklistUpdatedV2))
}

// 消息处理器定义
type P2OffboardingStatusUpdatedV2Handler struct {
	handler func(context.Context, *P2OffboardingStatusUpdatedV2) error
}

func NewP2OffboardingStatusUpdatedV2Handler(handler func(context.Context, *P2OffboardingStatusUpdatedV2) error) *P2OffboardingStatusUpdatedV2Handler {
	h := &P2OffboardingStatusUpdatedV2Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2OffboardingStatusUpdatedV2Handler) Event() interface{} {
	return &P2OffboardingStatusUpdatedV2{}
}

// 回调开发者注册的handle
func (h *P2OffboardingStatusUpdatedV2Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2OffboardingStatusUpdatedV2))
}

// 消息处理器定义
type P2OffboardingUpdatedV2Handler struct {
	handler func(context.Context, *P2OffboardingUpdatedV2) error
}

func NewP2OffboardingUpdatedV2Handler(handler func(context.Context, *P2OffboardingUpdatedV2) error) *P2OffboardingUpdatedV2Handler {
	h := &P2OffboardingUpdatedV2Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2OffboardingUpdatedV2Handler) Event() interface{} {
	return &P2OffboardingUpdatedV2{}
}

// 回调开发者注册的handle
func (h *P2OffboardingUpdatedV2Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2OffboardingUpdatedV2))
}

// 消息处理器定义
type P2ProbationUpdatedV2Handler struct {
	handler func(context.Context, *P2ProbationUpdatedV2) error
}

func NewP2ProbationUpdatedV2Handler(handler func(context.Context, *P2ProbationUpdatedV2) error) *P2ProbationUpdatedV2Handler {
	h := &P2ProbationUpdatedV2Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ProbationUpdatedV2Handler) Event() interface{} {
	return &P2ProbationUpdatedV2{}
}

// 回调开发者注册的handle
func (h *P2ProbationUpdatedV2Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ProbationUpdatedV2))
}

// 消息处理器定义
type P2ProcessUpdatedV2Handler struct {
	handler func(context.Context, *P2ProcessUpdatedV2) error
}

func NewP2ProcessUpdatedV2Handler(handler func(context.Context, *P2ProcessUpdatedV2) error) *P2ProcessUpdatedV2Handler {
	h := &P2ProcessUpdatedV2Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ProcessUpdatedV2Handler) Event() interface{} {
	return &P2ProcessUpdatedV2{}
}

// 回调开发者注册的handle
func (h *P2ProcessUpdatedV2Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ProcessUpdatedV2))
}

// 消息处理器定义
type P2ProcessApproverUpdatedV2Handler struct {
	handler func(context.Context, *P2ProcessApproverUpdatedV2) error
}

func NewP2ProcessApproverUpdatedV2Handler(handler func(context.Context, *P2ProcessApproverUpdatedV2) error) *P2ProcessApproverUpdatedV2Handler {
	h := &P2ProcessApproverUpdatedV2Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ProcessApproverUpdatedV2Handler) Event() interface{} {
	return &P2ProcessApproverUpdatedV2{}
}

// 回调开发者注册的handle
func (h *P2ProcessApproverUpdatedV2Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ProcessApproverUpdatedV2))
}

// 消息处理器定义
type P2ProcessCcUpdatedV2Handler struct {
	handler func(context.Context, *P2ProcessCcUpdatedV2) error
}

func NewP2ProcessCcUpdatedV2Handler(handler func(context.Context, *P2ProcessCcUpdatedV2) error) *P2ProcessCcUpdatedV2Handler {
	h := &P2ProcessCcUpdatedV2Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ProcessCcUpdatedV2Handler) Event() interface{} {
	return &P2ProcessCcUpdatedV2{}
}

// 回调开发者注册的handle
func (h *P2ProcessCcUpdatedV2Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ProcessCcUpdatedV2))
}

// 消息处理器定义
type P2ProcessNodeUpdatedV2Handler struct {
	handler func(context.Context, *P2ProcessNodeUpdatedV2) error
}

func NewP2ProcessNodeUpdatedV2Handler(handler func(context.Context, *P2ProcessNodeUpdatedV2) error) *P2ProcessNodeUpdatedV2Handler {
	h := &P2ProcessNodeUpdatedV2Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ProcessNodeUpdatedV2Handler) Event() interface{} {
	return &P2ProcessNodeUpdatedV2{}
}

// 回调开发者注册的handle
func (h *P2ProcessNodeUpdatedV2Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ProcessNodeUpdatedV2))
}
