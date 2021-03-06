package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	log "github.com/cihub/seelog"

	"github.com/anchnet/custom-agent/cron"
	"github.com/anchnet/custom-agent/funcs"
	"github.com/anchnet/custom-agent/g"
	"github.com/anchnet/custom-agent/http"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	//init seelog
	g.InitSeeLog()

	g.InitRootDir()
	g.InitLocalIps()
	g.InitRpcClients()
	if g.Config().StartTime != "undefined" {
		log.Info("collecting will start at :", g.Config().StartTime)
		for {
			if g.Config().StartTime == time.Now().Format("15:04") {
				break
			}
			time.Sleep(60)
		}
	}
	funcs.BuildMappers()
	cron.Collect()

	go http.Start()

	select {}

}
