/*
Copyright © 2022 niuzhiqiang <niuzhiqiang90@foxmail.com>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/niuzhiqiang90/yapi-user-operator/config"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete yapi user",
		Long: `Delete subcommand, user is required.
`,
	}

	cmd.AddCommand(NewDeleteUserCommand())
	return cmd
}

func NewDeleteUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Delete yapi user by userName.",
		Long: `Delete yapi user by userName.

For example:
yapi-user-manager delete user -u xxx@xxx.xxx
yapi-user-manager delete user --userName xxx@xxx.xxx`,
		Run: func(cmd *cobra.Command, args []string) {
			if userName == "" {
				fmt.Println("userName is required")
				fmt.Fprintln(cmd.OutOrStdout(), cmd.UsageString())
				return
			}
			deleteUser()
		},
	}

	cmd.Flags().StringVarP(&userName, "userName", "u", "", "userName (required)")
	return cmd
}

func deleteUser() {
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

	opts := options.FindOneAndDelete().
		SetProjection(bson.D{{"email", userName}})
	var deletedDocument bson.M
	err = collection.FindOneAndDelete(
		context.TODO(),
		bson.D{{"email", userName}},
		opts,
	).Decode(&deletedDocument)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments && strings.Contains(err.Error(), "mongo: no documents in result") {
			fmt.Printf("Account %s does not exists. \n", userName)
			return
		}

		log.Fatal(err)
	}
	fmt.Printf("Account %s deleted.\n", userName)
}
