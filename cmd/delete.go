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
		Short: "Delete yapi user by email.",
		Long: `Delete yapi user by email.

For example:
yapi-user-manager delete user -e xxx@xxx.xxx
yapi-user-manager delete user --email xxx@xxx.xxx`,
		Run: func(cmd *cobra.Command, args []string) {
			if email == "" {
				fmt.Println("email is required")
				fmt.Fprintln(cmd.OutOrStdout(), cmd.UsageString())
				return
			}
			deleteUser()
		},
	}

	cmd.Flags().StringVarP(&email, "email", "e", "", "email (required)")
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
	dbName := config.GetDBName()
	collection := client.Database(dbName).Collection("user")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.FindOneAndDelete().
		SetProjection(bson.D{{"email", email}})
	var deletedDocument bson.M
	err = collection.FindOneAndDelete(
		context.TODO(),
		bson.D{{"email", email}},
		opts,
	).Decode(&deletedDocument)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments && strings.Contains(err.Error(), "mongo: no documents in result") {
			fmt.Printf("Account %s does not exists. \n", email)
			return
		}

		log.Fatal(err)
	}
	fmt.Printf("Account %s deleted.\n", email)
}
