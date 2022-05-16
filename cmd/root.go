/*
Copyright © 2022 niuzhiqiang <niuzhiqiang90@foxmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "yapi-user-manager ",
		Short: "yapi-user-manager is a command line tool for yapi user management.",
		Long: `yapi-user-manager is a command line tool for yapi user management.
It makes your operation simple, convenient and fast.

For more information about yyapi-user-manager, please visit https://github.com/niuzhiqiang90/yapi-user-manager
For more information about yapi, please visit https://github.com/YMFE/yapi
`,
	}

	rootCmd.AddCommand(NewAddCommand())
	rootCmd.AddCommand(NewBlockCommand())
	rootCmd.AddCommand(NewUnBlockCommand())
	rootCmd.AddCommand(NewDeleteCommand())
	rootCmd.AddCommand(NewResetCommand())

	return rootCmd
}
