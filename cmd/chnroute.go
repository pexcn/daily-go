package cmd

import (
	"daily/chnroute"
	"daily/config"

	"github.com/spf13/cobra"
)

var chnrouteCmd = &cobra.Command{
	Use:   "chnroute",
	Short: "China routing table generator",
	Long:  `China routing table generator.`,
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		chnrouteFlag.GlobalFlag = *globalFlag
		// workaround for https://github.com/spf13/cobra/issues/1752
		if chnrouteFlag.Ipv6 {
			chnrouteFlag.Ipv4 = false
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		chnroute.Run(cmd, chnrouteFlag)
	},
}

var chnrouteFlag = &config.ChnrouteFlag{}

func init() {
	chnrouteCmd.Flags().SortFlags = false
	rootCmd.AddCommand(chnrouteCmd)

	chnrouteCmd.Flags().StringSliceVarP(&chnrouteFlag.Url, "url", "u", []string{}, "Input your URLs")
	chnrouteCmd.Flags().StringSliceVarP(&chnrouteFlag.File, "file", "f", []string{}, "Input your files")
	chnrouteCmd.Flags().StringVarP(&chnrouteFlag.Output, "output", "o", "", "Output file")
	chnrouteCmd.Flags().BoolVarP(&chnrouteFlag.Ipv4, "ipv4", "4", true, "Parse as IPv4 address")
	chnrouteCmd.Flags().BoolVarP(&chnrouteFlag.Ipv6, "ipv6", "6", false, "Parse as IPv6 address")

	chnrouteCmd.MarkFlagsOneRequired("url", "file")
	//chnrouteCmd.MarkFlagsMutuallyExclusive("url", "file")
	chnrouteCmd.MarkFlagsMutuallyExclusive("ipv4", "ipv6")
}
