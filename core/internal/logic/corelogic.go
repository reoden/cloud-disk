package logic

import (
	"bytes"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreLogic) Core(req *types.Request) (resp *types.Response, err error) {
	data := make([]*models.UserBasic, 0)
	err = models.Engine.Find(&data)
	if err != nil {
		log.Println("Get UserBasic Error", err)
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		log.Println("Marshal Error", err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, marshal, "", "  ")
	if err != nil {
		log.Println("JSON Indent Error", err)
	}

	fmt.Println(dst.String())
	resp = new(types.Response)
	resp.Message = dst.String()
	return
}
