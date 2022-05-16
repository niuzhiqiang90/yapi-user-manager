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

	"github.com/niuzhiqiang90/yapi-user-manager/config"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewResetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "Reset yapi user's password",
		Long: `Reset subcommand, email is required.
`,
	}

	cmd.AddCommand(NewResetUserPasswordCommand())
	return cmd
}

func NewResetUserPasswordCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "password",
		Short: "Reset yapi user's password by email.",
		Long: `Reset yapi user's password by email.

For example:
yapi-user-manager reset password -e xxx@xxx.xxx
yapi-user-manager reset password --email xxx@xxx.xxx`,
		Run: func(cmd *cobra.Command, args []string) {
			if email == "" {
				fmt.Println("email is required")
				fmt.Fprintln(cmd.OutOrStdout(), cmd.UsageString())
				return
			}
			if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
				fmt.Println("Email is invalid")
				fmt.Fprintln(cmd.OutOrStdout(), cmd.UsageString())
			}
			resetUserPassword()
		},
	}

	cmd.Flags().StringVarP(&email, "email", "e", "", "email (required)")

	return cmd
}

func resetUserPassword() {
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

	var password string = "224179069e921d923a2059de27d60ab2cb58cc4f"
	var passsalt string = "w4byep62al"

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"email", email}}
	update := bson.D{{"$set", bson.D{{"password", password}, {"passsalt", passsalt}}}}
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

	fmt.Printf("Account %s password has been reset.\n", email)
	fmt.Println("Password: 1234qwer!@#$")
	fmt.Println("Please change your password after login.")

}
