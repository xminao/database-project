package logic

import (
	"backend/app/university_goctl/models"
	"context"
	"errors"
	"strings"
	"time"

	"backend/app/university_goctl/api/internal/svc"
	"backend/app/university_goctl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddYearLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddYearLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddYearLogic {
	return &AddYearLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddYearLogic) AddYear(req *types.AddYearReq) (resp *types.AddYearResp, err error) {
	if len(strings.TrimSpace(req.YearName)) == 0 {
		return nil, errors.New("年级不能为空")
	}

	countBuilder := l.svcCtx.YearInfoModel.CountBuilder("*")
	count, err := l.svcCtx.YearInfoModel.FindCount(l.ctx, countBuilder)

	year := new(models.YearInfo)

	if count == 0 {
		year.YearId = 1
	} else {
		maxBuilder := l.svcCtx.YearInfoModel.MaxRowBuilder("created_at")
		LatestRecord, err := l.svcCtx.YearInfoModel.FindOneByQuery(l.ctx, maxBuilder)
		if err != nil {
			return nil, err
		}
		year.YearId = LatestRecord.YearId + 1
	}

	year.YearName = req.YearName
	year.CollegeId = req.CollegeId
	year.CreatedAt = time.Now()

	_, err = l.svcCtx.YearInfoModel.Insert(l.ctx, year)
	if err != nil {
		return nil, err
	}

	return &types.AddYearResp{
		Msg: "插入成功",
	}, nil
}
