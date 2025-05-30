package main

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/hello-xone/xone-chain/app"
	"github.com/hello-xone/xone-chain/cmd/xoned/cmd"
)

const (
	Version = "v0.2.0"
)

var (
	BuildTime = ""
	GitCommit = ""
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	version.Version = fmt.Sprintf("%s %s %s", Version, BuildTime, GitCommit)
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
