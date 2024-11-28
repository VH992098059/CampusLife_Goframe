package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/model"
	"demo3/internal/service"
	"github.com/google/uuid"
)

var ActivityInfo = cActivityInfo{}

type cActivityInfo struct{}

func (i *cActivityInfo) Create(ctx context.Context, req *backapi.ActivityInfoAddReq) (res *backapi.ActivityInfoAddRes, err error) {
	_, err = service.ActivityInfo().Create(ctx, model.ActivityInfoModelAddInput{
		ActivityPosters:       req.ActivityPosters,
		ActivityTitle:         req.ActivityTitle,
		ActivityNumber:        uuid.New().String(),
		ActivityTypeID:        req.ActivityTypeID,
		Keywords:              req.Keywords,
		ReleaseTime:           req.ReleaseTime,
		CheckStatus:           req.CheckStatus,
		CheckTime:             req.CheckTime,
		CheckPersonContact:    req.CheckPersonContact,
		CheckRemark:           req.CheckRemark,
		RegistrationStartTime: req.RegistrationStartTime,
		ActivityStartTime:     req.ActivityStartTime,
		Addr:                  req.Addr,
		PersonLimit:           req.PersonLimit,
		RegistrationFee:       req.RegistrationFee,
		WXCustomerCode:        req.WXCustomerCode,
		PaymentMethod:         req.PaymentMethod,
		RegistrationEndTime:   req.RegistrationEndTime,
		ActivityEndTime:       req.ActivityEndTime,
		CheckNeed:             req.CheckNeed,
		EnterActivistsID:      req.EnterActivistsID,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.ActivityInfoAddRes{}, nil
}
