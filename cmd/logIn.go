package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"agenda"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login for further use",
	Long: `login for further use and u need to input username, password.
	it will be better if email and phone is provider`,
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Error[login]： %v\n", err)
			}
		}()

		fmt.Println("login called")

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		fmt.Println("login called by " + username)
		fmt.Println("login with info password: " + password)

		if err := agenda.LogIn(agenda.Username(username), agenda.Auth(password)); err != nil {
			panic(err)
		} else {
			fmt.Print("login sucessfully!\n")
			// agenda.SaveAll()
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "Anonymous", "login info for username")
	loginCmd.Flags().StringP("password", "p", "", "login info for password")

}