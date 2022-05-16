/*
Copyright © 2022 niuzhiqiang <niuzhiqiang90@foxmail.com>

*/
package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/niuzhiqiang90/yapi-user-operator/config"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userName string
	email    string
)

func NewAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add yapi user",
		Long:  `Add subcommand, user is required.`,
	}

	cmd.AddCommand(NewAddUserCommand())
	return cmd
}

func NewAddUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Add yapi user by user's email.",
		Long: `Add yapi user by user's email.

For example:
yapi-user-manager add user -u name -e xxx@xxx.xxx
yapi-user-manager add user --userName name --email xxx@xxx.xxx`,
		Run: func(cmd *cobra.Command, args []string) {
			if userName == "" {
				fmt.Println("userName is required")
				fmt.Fprintln(cmd.OutOrStdout(), cmd.UsageString())
				return
			}
			if email == "" {
				fmt.Println("email is required")
				fmt.Fprintln(cmd.OutOrStdout(), cmd.UsageString())
				return
			}
			if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
				fmt.Println("Email is invalid")
				fmt.Fprintln(cmd.OutOrStdout(), cmd.UsageString())
			}
			addUser()
		},
	}

	cmd.Flags().StringVarP(&userName, "username", "u", "", "username (required)")
	cmd.Flags().StringVarP(&email, "email", "e", "", "email (required)")

	return cmd
}

func addUser() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoUri := config.GetMongoUri()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		fmt.Println(err)
		return
	}
	dbName := config.GetDBName()
	collection := client.Database(dbName).Collection("user")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userId int
	for userId = 20; userId <= 30000; userId++ {
		res, err := collection.InsertOne(ctx, bson.D{
			{"_id", userId},
			{"study", true},
			{"type", "site"},
			{"username", userName},
			{"password", "224179069e921d923a2059de27d60ab2cb58cc4f"},
			{"email", email},
			{"passsalt", "w4byep62al"},
			{"role", "member"},
			{"add_time", time.Unix(time.Now().Unix(), 0)},
			{"up_time", time.Unix(time.Now().Unix(), 0)},
			{"__v", 0}})

		if err != nil && strings.Contains(err.Error(), "_id_ dup key") {
			continue
		}

		if err != nil && strings.Contains(err.Error(), "email_1 dup key") {
			fmt.Printf("Account %s already exists. \n", email)
			break
		}

		if res != nil {
			fmt.Println("Add user successfully.")
			fmt.Println("Username:", userName)
			fmt.Println("Account:", email)
			fmt.Println("Password: 1234qwer!@#$")
			fmt.Println("Please change your password after login.")
			break
		}
	}
}
