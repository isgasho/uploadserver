package web

import (
	"context"

	"common"
	"github.com/logger"
	"github.com/httpsvr"
	"scheduler"
)

type ControlerGET struct{

}
type ControlerOPTIONS struct{

}
type ControlerPATCH struct{

}
type ControlerDELETE struct{

}

func (getc *ControlerGET) Do(ctx context.Context, req interface{}) (interface{}, httpsvr.APIErr) {
	logger.Info(ctx, logger.TAIHETagRequestIn, "Do")
	/*生成scheduler的信息并同步等待*/
	scheduleMsg, err := getc.genScheduleMsg(req)
	if  err !=  nil {
		return struct{}, common.GenErr(common.REQUESTFAILED, err)
	}
	res, err := scheduler.Get(scheduleMsg)
	if err !=  nil {
		return struct{}, common.GenErr(common.REQUESTFAILED, err)
	}
	return res, common.GenErr(common.REQUESSUCCESS, nil)
}

func (optionsc *ControlerOPTIONS) Do(ctx context.Context, req interface{}) (interface{}, httpsvr.APIErr) {
	logger.Info(ctx, logger.TAIHETagRequestIn, "Do")
	/*生成scheduler的信息并同步等待*/
	scheduleMsg, err := optionsc.genScheduleMsg(req)
	if  err !=  nil {
		return struct{}, common.GenErr(common.REQUESTFAILED, err)
	}
	res, err := scheduler.Options(scheduleMsg)
	if err !=  nil {
		return struct{}, common.GenErr(common.REQUESTFAILED, err)
	}
	return res, common.GenErr(common.REQUESSUCCESS, nil)
}

func (patchc *ControlerPATCH) Do(ctx context.Context, req interface{}) (interface{}, httpsvr.APIErr) {
	logger.Info(ctx, logger.TAIHETagRequestIn, "Do")
	/*生成scheduler的信息并同步等待*/
	scheduleMsg, err := patchc.genScheduleMsg(req)
	if  err !=  nil {
		return struct{}, common.GenErr(common.REQUESTFAILED, err)
	}
	res, err := scheduler.Patch(scheduleMsg)
	if err !=  nil {
		return struct{}, common.GenErr(common.REQUESTFAILED, err)
	}
	return res, common.GenErr(common.REQUESSUCCESS, nil)
}

func (deletec *ControlerDELETE) Do(ctx context.Context, req interface{}) (interface{}, httpsvr.APIErr) {
	logger.Info(ctx, logger.TAIHETagRequestIn, "Do")
	/*生成scheduler的信息并同步等待*/
	scheduleMsg, err := deletec.genScheduleMsg(req)
	if  err !=  nil {
		return struct{}, common.GenErr(common.REQUESTFAILED, err)
	}
	res, err := scheduler.Delete(scheduleMsg)
	if err !=  nil {
		return struct{}, common.GenErr(common.REQUESTFAILED, err)
	}
	return res, common.GenErr(common.REQUESSUCCESS, nil)
}

