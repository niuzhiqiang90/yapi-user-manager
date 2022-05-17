/*
Copyright © 2022 niuzhiqiang <niuzhiqiang90@foxmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "1.1.1"

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version information.",

		Run: func(cmd *cobra.Command, args []string) {
			RunVersion(cmd)
		},
	}

	cmd.Flags().StringP("output", "o", "", "Output format; available options are 'yaml', 'json' and 'short'")
	return cmd
}

func RunVersion(cmd *cobra.Command) error {
	const flag = "output"
	of, err := cmd.Flags().GetString(flag)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	switch of {
	case "short":
		fmt.Printf("%v\n", version)
	default:
		fmt.Printf("yapi-user-manager version: %v\n", version)
	}

	return nil
}
