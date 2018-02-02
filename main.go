package main

import (
	"github.com/2at2/telegacli/cmd"
	"github.com/mono83/xray/std/xcobra"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{}

func main() {
	mainCmd.AddCommand(cmd.SendMessageCmd)
	xcobra.Start(mainCmd)
}
