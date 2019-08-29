package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cx",
	Short: "A Collective CLI",
	Long: `A Go CLI built during the meetup focusing on: 
    	   Developing CLIs using Go.`,
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(RegionsCmd)
}

func initConfig() {

}
