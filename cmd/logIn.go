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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	loginCmd.Flags().StringP("username", "u", "Anonymous", "login info for username")
	loginCmd.Flags().StringP("password", "p", "", "login info for password")

}