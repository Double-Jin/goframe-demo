package file

import (
	"context"
	"fmt"
	apiAdmin "goframe/api/v1/admin"
	"goframe/internal/service"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

type sFile struct {}

func init()  {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}


func (s *sFile) UploadLocal(ctx context.Context, req *apiAdmin.UploadFileReq) (res *apiAdmin.UploadFileRes, err error) {
	files := req.File
	names, err := files.Save("public/upload/")

	if err != nil {
		err = gerror.New("文件上传失败")
		return
	}

	res = &apiAdmin.UploadFileRes{
		Name: names,
		Url:  "public/upload/" + names,
	}

	return
}

func (s *sFile) UploadOss(ctx context.Context, req *apiAdmin.UploadFileReq) (res *apiAdmin.UploadFileRes, err error) {
	files := req.File
	names, err := files.Save("public/upload/")
	if err != nil {
		err = gerror.New("文件上传失败")
		return
	}

	client, err := oss.New(
		g.Cfg().MustGet(ctx, "oss.endpoint").String(),
		g.Cfg().MustGet(ctx, "oss.accessKeyId").String(),
		g.Cfg().MustGet(ctx, "oss.accessKeySecret").String())

	if err != nil {
		return
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(g.Cfg().MustGet(ctx, "oss.bucket").String())
	if err != nil {
		return
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile("public/upload/"+names, gfile.Pwd()+"/public/upload/"+names)
	if err != nil {
		fmt.Println("Error:", err)
	}

	res = &apiAdmin.UploadFileRes{
		Name: names,
		Url:  g.Cfg().MustGet(ctx, "oss.oss_url").String() + "/public/upload/" + names,
	}

	return
}
