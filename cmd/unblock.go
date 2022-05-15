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

func NewUnBlockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unblock",
		Short: "Unblock yapi user",
		Long: `Unblock subcommand, user is required.
`,
	}

	cmd.AddCommand(NewUnBlockUserCommand())
	return cmd
}

func NewUnBlockUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "UnBlock user to yapi",
		Long: `For example:
yapi-user-manager unblock user -u xxx@xxx.xxx
yapi-user-manager unblock user --userName xxx@xxx.xxx`,
		Run: func(cmd *cobra.Command, args []string) {
			if userName == "" {
				fmt.Println("userName is required")
				fmt.Fprintln(cmd.OutOrStdout(), cmd.UsageString())
				return
			}
			unBlockUser()
		},
	}

	cmd.Flags().StringVarP(&userName, "userName", "u", "", "userName (required)")
	return cmd
}

func unBlockUser() {
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

	findPasswordOpts := options.FindOne().SetSort(bson.D{{"password", 1}})
	var result bson.M
	err = collection.FindOne(
		context.TODO(),
		bson.D{{"email", userName}},
		findPasswordOpts,
	).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}

	if result["password"].(string)[0] != '!' {
		fmt.Printf("Account %s is not blocked.\n", userName)
		os.Exit(0)
	}
	newPassword := result["password"].(string)[1:]

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
	fmt.Printf("Account %s is unlocked.\n", userName)

}
