package web

import (
	"time"
	"fmt"
	"os"

	"github.com/config"
	"github.com/httpsvr"
)

func Start(conf config.Configer) error {
	sec, err := conf.GetSection("web")
	if err != nil {
		return err
	}

	if err = httpServe(sec); err != nil {
		return err
	}

	return nil
}

func httpServe(sec config.Sectioner) error {
	addr := sec.GetStringMust("addr", "0.0.0.0:8090")
	readTimeout := sec.GetIntMust("readTimeout", 3600 * 1000000)
	writeTimeot := sec.GetIntMust("writeTimeot", 3600 * 1000000)
	svr := httpsvr.New(
		addr, 
		httpsvr.SetServerReadTimeout(time.Millisecond * time.Duration(readTimeout)),
		httpsvr.SetServerWriteTimeout(time.Millisecond * time.Duration(writeTimeot)),
		httpsvr.SetHandleDefaultTimeout(2000),
	)
	/*设置路由*/
	svr.AddRoute(
		"GET",
		"/web/controler",
		&web.ControlerGET{},
	)
	svr.AddRoute(
		"POST",
		"/web/controler",
		&web.ControlerPOST{},
	)
	svr.AddRoute(
		"PATCH",
		"/web/controler",
		&web.ControlerPATCH{},
	)
	svr.AddRoute(
		"OPTIONS",
		"/web/controler",
		&web.ControlerOPTIONS{},
	)
	svr.AddRoute(
		"DELETE",
		"/web/controler",
		&web.ControlerDELETE{},
	)
	go svr.Serve()

	return nil
}