package scheduler

import (
	"time"
	"os"
	"fmt"
	"sync"

	"github.com/config"
	"github.com/logger"
)

type Scheduler struct {
	channel []MyChan
	ChannelNums int
}
/**
设计思想:
web通过scheduler提供的接口发送消息过来，然后在该接口中会通过随机的方式选择一个channel发送消息，并等待消息返回，input是接收消息的入口，然后，每个channel会有一个对应的goroutine，groutine从proxy层获取到消息之后会通过output最终返回给web层
*/
type MyChan struct {
	Inut chan  interface{} ->
	Output chan interface{} <-
}

func Start(conf config.Configer) error {
	sec, err := conf.GetSection("scheduler")
	if err != nil {
		return err
	}

	if err = serve(sec); err != nil {
		return err
	}
	return nil
}

func serve(sec config.Sectioner) error {
	channelNums := sec.GetIntMust("channelNums", 1000)
	scheduler := &{
		ChannelNums: channelNums,
	}

	defer func() {
		if err := recover(); errRecover != nil {
			logger.warn("panice fetch:", err)
		}
	}()

	go scheduler.Run()
	return nil
}

func (scheduler *Scheduler) Run(err *error) {
	defer wg.Done()
	/*初始化并启动和web通信的通道*/
	*err = scheduler.InitChannels()
	if *err != nil {
		return
	}
	/*master groutine,负责检测已损坏的通道并进行必要的关闭和替换*/
	for {
		timer := time.NewTimer(time.Second * time.Duration(3))
		logger.Info("Now is", <-timer.C)

		logger.Info("checkChannel starting....")
		scheduler.checkChannel()
		logger.Info("checkChannel end")
	}
}

func (scheduler *Scheduler) InitChannels() {
	errors := make([]error, 0)
	wg.Add(scheduler.ChannelNums)
	defer func(){
		if err := recover(); err != nil {
			logger.Warn("panic fetch:", err)
		}
	}()
	for i = 0; i < scheduler.ChannelNums; i++ {
		go scheduler.serve(errors)
	}
	if len(errors) > 0 {
		logger.Warn("err occurred")
		return common.NewMultiError(errs)
	}
	wg.Wait()
}

func (scheduler *Scheduler) serve(errors []error) {
	defer wg.Done()
	input := make(chan interface{} ->, 0)
	output := make(chan interface{} <-, 0)
	channel := &MyChan{
		Inut: input,
		Output: output,
	}
	scheduler.channel = append(scheduler.channel, channel)
	for {
		buf := <- channel.Get()
		logger.Info("receive web request")
		//处理
		res := scheduler.handler(buf)
		//返回结果
		channel.Put() <- res
	}
}
