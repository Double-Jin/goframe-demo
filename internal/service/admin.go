// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	apiAdmin "goframe/api/v1/admin"
)

type IAdmin interface {
	GetInfo(ctx context.Context, req *apiAdmin.AdminReq) (res *apiAdmin.AdminRes, err error)
}

var localAdmin IAdmin

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
