
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"agenda"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register for further use",
	Long: `register for further use and u need to input username, password.
	it will be better if email and phone is provider`,
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Error[register]： %v\n", err)
			}
		}()

		fmt.Println("register called")

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")

		// fmt.Println("register called by " + username)
		// fmt.Println("register with info password: " + password)
		// fmt.Println("register with info email: " + email)
		// fmt.Println("register with info phone: " + phone)

		info := agenda.MakeUserInfo(agenda.Username(username), agenda.Auth(password), email, phone)

		if err := agenda.RegisterUser(info); err != nil {
			panic(err)
		} else {
			fmt.Print("register sucessfully!\n")
			// agenda.SaveAll()
		}
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("username", "u", "Anonymous", "register info for username")
	registerCmd.Flags().StringP("password", "p", "", "register info for password")
	registerCmd.Flags().StringP("email", "e", "", "register info for email")
	registerCmd.Flags().StringP("phone", "t", "", "register info for phone")

}
