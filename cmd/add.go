/*
Copyright © 2022 niuzhiqiang <niuzhiqiang90@foxmail.com>

*/
package cmd

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/niuzhiqiang90/yapi-user-operator/config"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userName string

func NewAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "",
		Long: `Add subcommand, user is required
`,
	}

	cmd.AddCommand(NewAddUserCommand())
	return cmd
}

func NewAddUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Add user to yapi",
		Long: `For example:
yapi-user-operator add user -u xxx@xxx.xxx
yapi-user-operator add user --userName xxx@xxx.xxx`,
		Run: func(cmd *cobra.Command, args []string) {
			if userName == "" {
				fmt.Println("userName is required")
				fmt.Fprintln(cmd.OutOrStdout(), cmd.UsageString())
				return
			}
			addUser()
		},
	}

	cmd.Flags().StringVarP(&userName, "userName", "u", "", "userName (required)")
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
	DBName := config.GetDBName()
	collection := client.Database(DBName).Collection("user")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rand.Seed(time.Now().UnixNano())
	res, err := collection.InsertOne(ctx, bson.D{
		{"_id", rand.Intn(100)},
		{"study", true},
		{"type", "site"},
		{"username", userName},
		{"password", "224179069e921d923a2059de27d60ab2cb58cc4f"},
		{"email", userName},
		{"passsalt", "w4byep62al"},
		{"role", "member"},
		{"add_time", time.Unix(time.Now().Unix(), 0)},
		{"up_time", time.Unix(time.Now().Unix(), 0)},
		{"__v", 0}})

	if err != nil {
		fmt.Println(err)
	}

	if res != nil {
		fmt.Println("Add user success")
		fmt.Println("Account:", userName)
		fmt.Println("Password: 1234qwer!@#$")
		fmt.Println("Please change your password after login")
	}
}
