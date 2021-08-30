// Code generated by lark suite oapi sdk gen
package v2

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
)

type Avatar struct {
	AvatarOrigin    string   `json:"avatar_origin,omitempty"`
	Avatar72        string   `json:"avatar_72,omitempty"`
	Avatar240       string   `json:"avatar_240,omitempty"`
	Avatar640       string   `json:"avatar_640,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Avatar) MarshalJSON() ([]byte, error) {
	type cp Avatar
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Tenant struct {
	Name            string   `json:"name,omitempty"`
	DisplayId       string   `json:"display_id,omitempty"`
	TenantTag       int      `json:"tenant_tag,omitempty"`
	TenantKey       string   `json:"tenant_key,omitempty"`
	Avatar          *Avatar  `json:"avatar,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Tenant) MarshalJSON() ([]byte, error) {
	type cp Tenant
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type TenantQueryResult struct {
	Tenant *Tenant `json:"tenant,omitempty"`
}
