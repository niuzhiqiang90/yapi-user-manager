/*
Copyright © 2022 niuzhiqiang <niuzhiqiang90@foxmail.com>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/niuzhiqiang90/yapi-user-operator/config"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewBlockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block",
		Short: "Block yapi user",
		Long: `Block subcommand, user is required.
`,
	}

	cmd.AddCommand(NewBlockUserCommand())
	return cmd
}

func NewBlockUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Block yapi user by userName.",
		Long: `Block yapi user by userName.

For example:
yapi-user-manager block user -u xxx@xxx.xxx
yapi-user-manager block user --userName xxx@xxx.xxx`,
		Run: func(cmd *cobra.Command, args []string) {
			if userName == "" {
				fmt.Println("userName is required")
				fmt.Fprintln(cmd.OutOrStdout(), cmd.UsageString())
				return
			}
			blockUser()
		},
	}

	cmd.Flags().StringVarP(&userName, "userName", "u", "", "userName (required)")
	return cmd
}

func blockUser() {
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

	getPasswordOpts := options.FindOne().SetSort(bson.D{{"password", 1}})
	var result bson.M
	err = collection.FindOne(
		context.TODO(),
		bson.D{{"email", userName}},
		getPasswordOpts,
	).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}

	if result["password"].(string)[0] == '!' {
		fmt.Printf("Account %s has been blocked. \n", userName)
		os.Exit(0)
	}
	newPassword := "!" + result["password"].(string)

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"email", userName}}
	update := bson.D{{"$set", bson.D{{"password", newPassword}}}}
	var updatedDocument bson.M
	err = collection.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
		opts,
	).Decode(&updatedDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}
	fmt.Printf("Account %s blocked.\n", userName)
}
