package userclient

import (
	"encoding/json"
	"github.com/rolandhe/daog/ttypes"
)

type UserInfo struct {
	Uid       int64                 `json:"uid"`
	Name      string                `json:"name"`
	Mobile    string                `json:"mobile,omitempty"`
	Email     string                `json:"email"`
	Nickname  string                `json:"nickname,omitempty"`
	Avatar    string                `json:"avatar,omitempty"`
	Domain    string                `json:"domain,omitempty"`
	DomainId  int64                 `json:"domainId,omitempty"`
	Role      string                `json:"role,omitempty"`
	GroupRole string                `json:"groupRole,omitempty"`
	Enabled   bool                  `json:"enabled"`
	Deleted   bool                  `json:"deleted"`
	Verified  bool                  `json:"verified"`
	CreatedAt ttypes.NormalDatetime `json:"createdAt,omitempty"`
}

func (u *UserInfo) String() string {
	js, err := json.Marshal(u)
	if err != nil {
		return err.Error()
	}
	return string(js)

}
