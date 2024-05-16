/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"github.com/fdaygon/rift/pkg/spotify"
	"github.com/fdaygon/rift/pkg/terminal"
	"github.com/spf13/cobra"
)

// topCmd represents the top command
var (
	toptype string
	limit   string
	length  string
	topCmd  = &cobra.Command{
		Use:   "top",
		Short: "Return your top tracks or artists",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			terminal.RefreshSource()
			topChoice := spotify.TopParams{
				Type:      toptype,
				Limit:     limit,
				TimeRange: length,
			}

			spotify.GetTopItems(topChoice)
		},
	}
)

func init() {
	rootCmd.AddCommand(topCmd)

	topCmd.Flags().StringVarP(&toptype, "type", "t", "", "Tracks or Artists. Only Valid Options")
	topCmd.Flags().StringVarP(&limit, "limit", "l", "", "Number of items to return (1-50)")
	topCmd.Flags().StringVarP(&length, "range", "r", "", "long_term (~1 year of data), medium_term (approximately last 6 months), short_term (approximately last 4 weeks)")
	topCmd.MarkFlagRequired("limit")
	topCmd.MarkFlagRequired("type")
	topCmd.MarkFlagRequired("range")

}
