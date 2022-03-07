package main

import (
	"github.com/gefion-tech/bfb-user-microservice/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.NewRootCmd().Execute())
}
