package cmd

import (
	"egen.io/cx/pkg/aws"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var RegionsCmd = &cobra.Command{
	Use:   "regions",
	Short: "View and Set regions for the AWS CLI",
	Run: func(cmd *cobra.Command, args []string) {
		code := viper.GetString("code")
		if code == "" {
			log.Fatal("ERROR: code is missing")
		}
		regions := aws.ListRegions(code)
		var selectedRegion string

		prompt := &survey.Select{
			Message: "change aws region:",
			Options: regions,
		}

		err := survey.AskOne(prompt, &selectedRegion, survey.WithIcons(func(icons *survey.IconSet) {
			// you can set any icons
			icons.SelectFocus.Text = "❯"
		}))

		if err != nil {
			log.Fatal("ERROR: selecting options")
		}

		aws.UpdateRegion(selectedRegion)
	},
}

func init() {
	RegionsCmd.Flags().String("code", "", "Region Code: us, eu, ap")
	//_ = RegionsCmd.MarkFlagRequired("code")
	_ = viper.BindPFlag("code", RegionsCmd.Flags().Lookup("code"))
}
