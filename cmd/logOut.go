package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"agenda"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout ",
	Long: `logout for further use and u need to input username, password.
	it will be better if email and phone is provider`,
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Error[logout]： %v\n", err)
			}
		}()

		fmt.Println("logout called")

		// fmt.Println("logout called by " + username)
		// fmt.Println("logout with info password: " + password)

		if err := agenda.LogOut(agenda.LoginedUser().Name); err != nil {
			panic(err)
		} else {
			fmt.Print("logout sucessfully!\n")
			// agenda.SaveAll()
		}
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)
}
