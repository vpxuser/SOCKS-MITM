package main

import (
	"fmt"
	yaklog "github.com/yaklang/yaklang/common/log"
	"socks2https/setting"
	"socks2https/socks"
)

const PROGRAM_NAME = "\n███████╗ ██████╗  ██████╗██╗  ██╗███████╗██████╗ ██╗  ██╗████████╗████████╗██████╗ ███████╗\n██╔════╝██╔═══██╗██╔════╝██║ ██╔╝██╔════╝╚════██╗██║  ██║╚══██╔══╝╚══██╔══╝██╔══██╗██╔════╝\n███████╗██║   ██║██║     █████╔╝ ███████╗ █████╔╝███████║   ██║      ██║   ██████╔╝███████╗\n╚════██║██║   ██║██║     ██╔═██╗ ╚════██║██╔═══╝ ██╔══██║   ██║      ██║   ██╔═══╝ ╚════██║\n███████║╚██████╔╝╚██████╗██║  ██╗███████║███████╗██║  ██║   ██║      ██║   ██║     ███████║\n╚══════╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝     ╚══════╝\n                                                                                           \n"

func init() {
	fmt.Println(PROGRAM_NAME)
	yaklog.SetLevel(setting.Config.Log.Level)
}

func main() {
	mitmSocks := socks.NewMITMSocks()
	mitmSocks.Host = setting.Config.Socks.Host
	mitmSocks.Threads = setting.Config.Socks.Threads
	mitmSocks.ClientTimeout = setting.Config.Socks.Timeout.Client
	mitmSocks.TargetTimeout = setting.Config.Socks.Timeout.Target
	mitmSocks.Run()
}
