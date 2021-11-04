package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var logger log.Logger
var config = &Config{}

func main() {
	configPath := flag.String("config", "sd-config.yaml", "config file path")
	outputPath := flag.String("output", "sd.yaml", "output file path")
	flag.Parse()
	viper.SetConfigName("sd-config.yaml")
	viper.SetConfigFile(*configPath)
	viper.ReadInConfig()
	viper.Unmarshal(config)
	fmt.Println("debg", config.Filters)
	logger = log.NewLogfmtLogger(os.Stderr)
	Credential := common.NewCredential(config.Ak, config.Sk)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	fmt.Println("debug1", config.Region)
	client, err := cvm.NewClient(Credential, config.Region, cpf)
	if err != nil {
		panic(err)
	}
	step := func() error {
		instances, err := Getinstances(client)
		if err != nil {
			return err
		}
		f, err := os.Create(*outputPath)
		if err != nil {
			return err
		}
		return instances.Write(f)
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	err = step()
	if err != nil {
		level.Error(logger).Log("msg", "write sd file failed", "error", err)
	}
	for {
		select {
		case <-ch:
			level.Info(logger).Log("msg", "received a signal")
			os.Exit(0)
		case <-time.After(config.Interval):
			err := step()
			if err != nil {
				level.Error(logger).Log("msg", "write sd file failed", "error", err)
			}
		}
	}
}
