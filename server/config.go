package main

import (
	"flag"
	"github.com/DAVFoundation/captain/config"
)

func initConfig () {

	config.LoadString(&config.Redis.Host, "REDIS_HOST", "redis-host", "Redis Host")
	config.LoadInt(&config.Redis.Port, "REDIS_PORT", "redis-port", "Redis Port")

	config.LoadString(&config.Log.Threshold, "LOG_THRESHOLD", "log-threshold", "Log level threshold")

	config.LoadString(&config.Server.Host, "SERVER_HOST", "server-host", "Server Host")
	config.LoadInt(&config.Server.Port, "SERVER_PORT", "server-port", "Server Port")

	config.LoadString(&config.MissionControl.Host, "MISSION_CONTROL_HOST", "mission-control-host", "Mission Control Host")
	config.LoadInt(&config.MissionControl.Port, "MISSION_CONTROL_PORT", "mission-control-port", "Mission Control Port")

	flag.Parse()

}