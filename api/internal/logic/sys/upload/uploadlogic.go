package upload

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(r *http.Request, ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadLogic) Upload() (resp *types.UploadResp, err error) {
	var AccessKey = "dXjfJ47I2-qY5SCiQ2KvWENU8BXsTXKiNpvocA9I"
	var SerectKey = "VnWaBZkO9_AUcsEjr7iZd8-XqYn7nEUlKLMx0fFO"
	var Bucket = "mallupload"
	var ImgUrl = "rz9o6lxu1.hd-bkt.clouddn.com"

	file, handler, _ := l.r.FormFile("file")
	if err != nil {
		// 处理错误
		return nil, err
	}
	defer file.Close()

	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	key := generateTimestampKey()
	err = formUploader.Put(context.Background(), &ret, upToken, key, file, handler.Size, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	imageURL := "http://" + ImgUrl + "/" + ret.Key
	return &types.UploadResp{
		Data: imageURL,
	}, nil
}
func generateTimestampKey() string {
	timestamp := time.Now().Unix()
	key := fmt.Sprintf("%d", timestamp)
	return key
}
