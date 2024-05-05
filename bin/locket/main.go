package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	greetingBanner = `
██╗      ██████╗  ██████╗██╗  ██╗███████╗████████╗
██║     ██╔═══██╗██╔════╝██║ ██╔╝██╔════╝╚══██╔══╝
██║     ██║   ██║██║     █████╔╝ █████╗     ██║
██║     ██║   ██║██║     ██╔═██╗ ██╔══╝     ██║
███████╗╚██████╔╝╚██████╗██║  ██╗███████╗   ██║
╚══════╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝╚══════╝   ╚═╝
	`
)

var (
	mode   string
	addr   string
	port   int
	data   string
	driver string
	dsn    string
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP()
}

func initConfig() {
	viper.AutomaticEnv()
	var err error
	instanceProfile, err = profile.GetProfile()
	if err != nil {
		fmt.Printf("failed to get profile, error: %+v\n", err)
		return
	}

	fmt.Printf(`---
Server profile
version: %s
data: %s
dsn: %s
addr: %s
port: %d
mode: %s
driver: %s
frontend: %t
---
`, instanceProfile.Version, instanceProfile.Data, instanceProfile.DSN, instanceProfile.Addr, instanceProfile.Port, instanceProfile.Mode, instanceProfile.Driver, instanceProfile.Frontend)
}

func printGreetings() {
	print(greetingBanner)
	if len(instanceProfile.Addr) == 0 {
		fmt.Printf("Version %s has been started on port %d\n", instanceProfile.Version, instanceProfile.Port)
	} else {
		fmt.Printf("Version %s has been started on address '%s' and port %d\n", instanceProfile.Version, instanceProfile.Addr, instanceProfile.Port)
	}
	fmt.Printf(`---
See more in:
👉Website: %s
👉GitHub: %s
---
`, "https://usememos.com", "https://github.com/usememos/memos")
}

func main() {
	err := os.Execute()
	if err != nil {
		panic(err)
	}
}
