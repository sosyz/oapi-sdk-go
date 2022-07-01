// Package vc code generated by oapi sdk gen
package larkvc

import (
	"context"
)

// 消息处理器定义
type MeetingJoinMeetingEventHandler struct {
	handler func(context.Context, *MeetingJoinMeetingEvent) error
}

func NewMeetingJoinMeetingEventHandler(handler func(context.Context, *MeetingJoinMeetingEvent) error) *MeetingJoinMeetingEventHandler {
	h := &MeetingJoinMeetingEventHandler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *MeetingJoinMeetingEventHandler) Event() interface{} {
	return &MeetingJoinMeetingEvent{}
}

// 回调开发者注册的handle
func (h *MeetingJoinMeetingEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MeetingJoinMeetingEvent))
}

// 消息处理器定义
type MeetingLeaveMeetingEventHandler struct {
	handler func(context.Context, *MeetingLeaveMeetingEvent) error
}

func NewMeetingLeaveMeetingEventHandler(handler func(context.Context, *MeetingLeaveMeetingEvent) error) *MeetingLeaveMeetingEventHandler {
	h := &MeetingLeaveMeetingEventHandler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *MeetingLeaveMeetingEventHandler) Event() interface{} {
	return &MeetingLeaveMeetingEvent{}
}

// 回调开发者注册的handle
func (h *MeetingLeaveMeetingEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MeetingLeaveMeetingEvent))
}

// 消息处理器定义
type MeetingMeetingEndedEventHandler struct {
	handler func(context.Context, *MeetingMeetingEndedEvent) error
}

func NewMeetingMeetingEndedEventHandler(handler func(context.Context, *MeetingMeetingEndedEvent) error) *MeetingMeetingEndedEventHandler {
	h := &MeetingMeetingEndedEventHandler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *MeetingMeetingEndedEventHandler) Event() interface{} {
	return &MeetingMeetingEndedEvent{}
}

// 回调开发者注册的handle
func (h *MeetingMeetingEndedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MeetingMeetingEndedEvent))
}

// 消息处理器定义
type MeetingMeetingStartedEventHandler struct {
	handler func(context.Context, *MeetingMeetingStartedEvent) error
}

func NewMeetingMeetingStartedEventHandler(handler func(context.Context, *MeetingMeetingStartedEvent) error) *MeetingMeetingStartedEventHandler {
	h := &MeetingMeetingStartedEventHandler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *MeetingMeetingStartedEventHandler) Event() interface{} {
	return &MeetingMeetingStartedEvent{}
}

// 回调开发者注册的handle
func (h *MeetingMeetingStartedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MeetingMeetingStartedEvent))
}

// 消息处理器定义
type MeetingRecordingEndedEventHandler struct {
	handler func(context.Context, *MeetingRecordingEndedEvent) error
}

func NewMeetingRecordingEndedEventHandler(handler func(context.Context, *MeetingRecordingEndedEvent) error) *MeetingRecordingEndedEventHandler {
	h := &MeetingRecordingEndedEventHandler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *MeetingRecordingEndedEventHandler) Event() interface{} {
	return &MeetingRecordingEndedEvent{}
}

// 回调开发者注册的handle
func (h *MeetingRecordingEndedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MeetingRecordingEndedEvent))
}

// 消息处理器定义
type MeetingRecordingReadyEventHandler struct {
	handler func(context.Context, *MeetingRecordingReadyEvent) error
}

func NewMeetingRecordingReadyEventHandler(handler func(context.Context, *MeetingRecordingReadyEvent) error) *MeetingRecordingReadyEventHandler {
	h := &MeetingRecordingReadyEventHandler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *MeetingRecordingReadyEventHandler) Event() interface{} {
	return &MeetingRecordingReadyEvent{}
}

// 回调开发者注册的handle
func (h *MeetingRecordingReadyEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MeetingRecordingReadyEvent))
}

// 消息处理器定义
type MeetingRecordingStartedEventHandler struct {
	handler func(context.Context, *MeetingRecordingStartedEvent) error
}

func NewMeetingRecordingStartedEventHandler(handler func(context.Context, *MeetingRecordingStartedEvent) error) *MeetingRecordingStartedEventHandler {
	h := &MeetingRecordingStartedEventHandler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *MeetingRecordingStartedEventHandler) Event() interface{} {
	return &MeetingRecordingStartedEvent{}
}

// 回调开发者注册的handle
func (h *MeetingRecordingStartedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MeetingRecordingStartedEvent))
}

// 消息处理器定义
type MeetingShareEndedEventHandler struct {
	handler func(context.Context, *MeetingShareEndedEvent) error
}

func NewMeetingShareEndedEventHandler(handler func(context.Context, *MeetingShareEndedEvent) error) *MeetingShareEndedEventHandler {
	h := &MeetingShareEndedEventHandler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *MeetingShareEndedEventHandler) Event() interface{} {
	return &MeetingShareEndedEvent{}
}

// 回调开发者注册的handle
func (h *MeetingShareEndedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MeetingShareEndedEvent))
}

// 消息处理器定义
type MeetingShareStartedEventHandler struct {
	handler func(context.Context, *MeetingShareStartedEvent) error
}

func NewMeetingShareStartedEventHandler(handler func(context.Context, *MeetingShareStartedEvent) error) *MeetingShareStartedEventHandler {
	h := &MeetingShareStartedEventHandler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *MeetingShareStartedEventHandler) Event() interface{} {
	return &MeetingShareStartedEvent{}
}

// 回调开发者注册的handle
func (h *MeetingShareStartedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MeetingShareStartedEvent))
}
