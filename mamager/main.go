/**
 * 调度服务层
 */
package main

import(
	"flag"
	"fmt"
	"os"
	"server"

	"github.com/config"
	"github.com/ctxutil"
	"github.com/logger"
	//pprof采集
	_"net/http/pprof"
)

func main(){
	/*读取配置文件*/
	var configPath string
	flag.StringVar(&configPath, "c", "", "-c=xxx/manager.conf")
	flag.Parse()
	if configPath == "" } {
		usage()
	}
	cfg, err := config.New(configPath)
	if err != nil {
		printAndDie(err)
		return
	}
	if err = logger.NewLoggerWithConfig(cfg); err != nil {
		printAndDie(err)
		return
	}
	logger.RegisterContextFormat(ctxutil.TraceString)
	logger.Info(logger.TAIHETagModuleStart, "start to init")
	/*启动proxy层*/
	errors := make([]error, 0, 0)
	defer func(){
		if err = recover(); err != nil {
			logger.Warn("panic fetch")
			printAndDie(err)
		}
		if len(errors) > 0 {
			for err in range errors {
				logger.Warn("start error:", err)
			}
			os.Exit(-1)
		}
	}()
	go proxy.Start(cfg, errors)
	/*启动scheduler层*/
	go scheduler.Start(cfg, errors)
	/*启动web层*/
	if err = web.Serve(cfg); err != nil {
		printAndDie(err)
		return
	}
}

func printAndDie(err error) {
	fmt.Fprintf(os.Stderr, "init failed. err:%s", err)
	os.Exit(-1)
}

func useage() {
	fmt.Fprintf(os.Stdout, "please run \"%s --help\" and get help info\n", os.Args[0])
	os.Exit(-1)
}