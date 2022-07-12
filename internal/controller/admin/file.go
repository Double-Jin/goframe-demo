package admin

import (
	"context"
	"goframe/api/v1/admin"
	"goframe/internal/service"
)

var (
	File = cFile{}
)

type cFile struct{}

func (c *cAdmin) UploadFile(ctx context.Context, req *admin.UploadFileReq) (res *admin.UploadFileRes, err error) {
	//res, err = service.File().UploadLocal(ctx, req)
	res, err = service.File().UploadOss(ctx, req)
	return
}

