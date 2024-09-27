package builder

import (
	"context"

	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/api/admin"
	inviteapi "github.com/aide-family/moon/api/admin/invite"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

var _ InviteModuleBuilder = (*inviteModuleBuilder)(nil)

type (
	InviteModuleBuilder interface {
		WithInviteUserRequest(*inviteapi.InviteUserRequest) ICreateInviteUserRequestBuilder
		WithUpdateInviteStatusRequest(*inviteapi.UpdateInviteStatusRequest) IUpdateInviteStatusRequestBuilder
		WithListInviteUserRequest(*inviteapi.ListUserInviteRequest) IListInviteUserRequestBuilder
		DoInviteBuilder(map[uint32]*model.SysTeam) IDoInviteBuilder
	}
	inviteModuleBuilder struct {
		ctx context.Context
	}

	ICreateInviteUserRequestBuilder interface {
		ToBo() *bo.InviteUserParams
	}

	createInviteUserRequestBuilder struct {
		ctx context.Context
		*inviteapi.InviteUserRequest
	}

	IUpdateInviteStatusRequestBuilder interface {
		ToBo() *bo.UpdateInviteStatusParams
	}
	updateInviteStatusRequestBuilder struct {
		ctx context.Context
		*inviteapi.UpdateInviteStatusRequest
	}

	IListInviteUserRequestBuilder interface {
		ToBo() *bo.QueryInviteListParams
	}

	listInviteUserRequestBuilder struct {
		ctx context.Context
		*inviteapi.ListUserInviteRequest
	}

	IDoInviteBuilder interface {
		ToAPI(*bizmodel.SysTeamInvite) *admin.InviteItem
		ToAPIs([]*bizmodel.SysTeamInvite) []*admin.InviteItem
	}

	doInviteBuilder struct {
		ctx     context.Context
		TeamMap map[uint32]*model.SysTeam
	}
)

func (d *doInviteBuilder) ToAPI(invite *bizmodel.SysTeamInvite) *admin.InviteItem {
	if types.IsNil(d) || types.IsNil(invite) {
		return nil
	}

	resItem := &admin.InviteItem{
		Id:         invite.ID,
		InviteType: api.InviteType(invite.InviteType),
		Roles: NewParamsBuild().
			RoleModuleBuilder().
			DoRoleBuilder().
			ToAPIs(invite.TeamRoles),
	}
	teamMap := d.TeamMap
	if !types.IsNil(teamMap) {
		team := teamMap[invite.TeamID]
		resItem.Team = NewParamsBuild().TeamModuleBuilder().DoTeamBuilder().ToAPI(team)
	}
	return resItem
}

func (d *doInviteBuilder) ToAPIs(invites []*bizmodel.SysTeamInvite) []*admin.InviteItem {
	if types.IsNil(d) || types.IsNil(invites) {
		return nil
	}
	return types.SliceTo(invites, func(invite *bizmodel.SysTeamInvite) *admin.InviteItem {
		return d.ToAPI(invite)
	})
}

func (i *inviteModuleBuilder) DoInviteBuilder(teamMap map[uint32]*model.SysTeam) IDoInviteBuilder {
	return &doInviteBuilder{
		ctx:     i.ctx,
		TeamMap: teamMap,
	}
}

func (i *listInviteUserRequestBuilder) ToBo() *bo.QueryInviteListParams {
	if types.IsNil(i) || types.IsNil(i.ListUserInviteRequest) {
		return nil
	}
	return &bo.QueryInviteListParams{
		InviteType: vobj.InviteType(i.GetType()),
		Keyword:    i.GetKeyword(),
		Page:       types.NewPagination(i.GetPagination()),
	}
}

func (i updateInviteStatusRequestBuilder) ToBo() *bo.UpdateInviteStatusParams {
	if types.IsNil(i) || types.IsNil(i.UpdateInviteStatusRequest) {
		return nil
	}
	return &bo.UpdateInviteStatusParams{
		InviteID:   i.GetId(),
		InviteType: vobj.InviteType(i.GetType()),
	}
}

func (i *createInviteUserRequestBuilder) ToBo() *bo.InviteUserParams {
	if types.IsNil(i) || types.IsNil(i.InviteUserRequest) {
		return nil
	}
	return &bo.InviteUserParams{
		TeamRoleIds: i.GetRoleId(),
		InviteCode:  i.GetInviteCode(),
	}
}

func (i *inviteModuleBuilder) WithInviteUserRequest(request *inviteapi.InviteUserRequest) ICreateInviteUserRequestBuilder {
	return &createInviteUserRequestBuilder{
		ctx:               i.ctx,
		InviteUserRequest: request,
	}
}

func (i *inviteModuleBuilder) WithUpdateInviteStatusRequest(request *inviteapi.UpdateInviteStatusRequest) IUpdateInviteStatusRequestBuilder {
	return &updateInviteStatusRequestBuilder{
		ctx:                       i.ctx,
		UpdateInviteStatusRequest: request,
	}
}

func (i *inviteModuleBuilder) WithListInviteUserRequest(request *inviteapi.ListUserInviteRequest) IListInviteUserRequestBuilder {
	return &listInviteUserRequestBuilder{
		ctx:                   i.ctx,
		ListUserInviteRequest: request,
	}
}
