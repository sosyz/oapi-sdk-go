// Package dispatcher code generated by oapi sdk gen
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

package dispatcher

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

func (dispatcher *EventDispatcher) OnP2CustomAttrEventUpdatedV3(handler func(ctx context.Context, event *larkcontact.P2CustomAttrEventUpdatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.custom_attr_event.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.custom_attr_event.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.custom_attr_event.updated_v3"] = larkcontact.NewP2CustomAttrEventUpdatedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2DepartmentCreatedV3(handler func(ctx context.Context, event *larkcontact.P2DepartmentCreatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.department.created_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.department.created_v3")
	}
	dispatcher.eventType2EventHandler["contact.department.created_v3"] = larkcontact.NewP2DepartmentCreatedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2DepartmentDeletedV3(handler func(ctx context.Context, event *larkcontact.P2DepartmentDeletedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.department.deleted_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.department.deleted_v3")
	}
	dispatcher.eventType2EventHandler["contact.department.deleted_v3"] = larkcontact.NewP2DepartmentDeletedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2DepartmentUpdatedV3(handler func(ctx context.Context, event *larkcontact.P2DepartmentUpdatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.department.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.department.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.department.updated_v3"] = larkcontact.NewP2DepartmentUpdatedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2EmployeeTypeEnumActivedV3(handler func(ctx context.Context, event *larkcontact.P2EmployeeTypeEnumActivedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.actived_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.actived_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.actived_v3"] = larkcontact.NewP2EmployeeTypeEnumActivedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2EmployeeTypeEnumCreatedV3(handler func(ctx context.Context, event *larkcontact.P2EmployeeTypeEnumCreatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.created_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.created_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.created_v3"] = larkcontact.NewP2EmployeeTypeEnumCreatedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2EmployeeTypeEnumDeactivatedV3(handler func(ctx context.Context, event *larkcontact.P2EmployeeTypeEnumDeactivatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.deactivated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.deactivated_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.deactivated_v3"] = larkcontact.NewP2EmployeeTypeEnumDeactivatedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2EmployeeTypeEnumDeletedV3(handler func(ctx context.Context, event *larkcontact.P2EmployeeTypeEnumDeletedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.deleted_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.deleted_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.deleted_v3"] = larkcontact.NewP2EmployeeTypeEnumDeletedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2EmployeeTypeEnumUpdatedV3(handler func(ctx context.Context, event *larkcontact.P2EmployeeTypeEnumUpdatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.updated_v3"] = larkcontact.NewP2EmployeeTypeEnumUpdatedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2ScopeUpdatedV3(handler func(ctx context.Context, event *larkcontact.P2ScopeUpdatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.scope.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.scope.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.scope.updated_v3"] = larkcontact.NewP2ScopeUpdatedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2UserCreatedV3(handler func(ctx context.Context, event *larkcontact.P2UserCreatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.user.created_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.user.created_v3")
	}
	dispatcher.eventType2EventHandler["contact.user.created_v3"] = larkcontact.NewP2UserCreatedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2UserDeletedV3(handler func(ctx context.Context, event *larkcontact.P2UserDeletedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.user.deleted_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.user.deleted_v3")
	}
	dispatcher.eventType2EventHandler["contact.user.deleted_v3"] = larkcontact.NewP2UserDeletedV3Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2UserUpdatedV3(handler func(ctx context.Context, event *larkcontact.P2UserUpdatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.user.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.user.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.user.updated_v3"] = larkcontact.NewP2UserUpdatedV3Handler(handler)
	return dispatcher
}
