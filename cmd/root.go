/*
Copyright © 2022 niuzhiqiang <niuzhiqiang90@foxmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "yapi-user-operator",
		Short: "yapi-user-operator is a command line tool for yapi.",
		Long: `yapi-user-operator is a command line tool for yapi accout management.
It makes your operation simple, convenient and fast.

For more information about yapi-user-operator, please visit https://github.com/niuzhiqiang90/yapi-user-operator
For more information about yapi, please visit https://github.com/YMFE/yapi
`,
	}

	rootCmd.AddCommand(NewAddCommand())
	rootCmd.AddCommand(NewBlockCommand())
	rootCmd.AddCommand(NewUnBlockCommand())
	rootCmd.AddCommand(NewDeleteCommand())

	return rootCmd
}
