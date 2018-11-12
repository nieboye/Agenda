package cmd

import (
	"agenda"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit for further use",
	Long: `edit for further use and u need to input username, password.
	it will be better if email and phone is provider`,
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Error[edit]： %v\n", err)
			}
		}()

		fmt.Println("edit called")

		titleFlag, _ := cmd.Flags().GetString("title")
		userNameFlag, _ := cmd.Flags().GetString("username")
		addBoolFlag, _ := cmd.Flags().GetBool("invite")
		delBoolFlag, _ := cmd.Flags().GetBool("del")

		if addBoolFlag && delBoolFlag {
			panic(errors.New("The flag invite(i) and del(d) can't appear at the same time"))

		} else if addBoolFlag {

			err := agenda.AddParticipatorToMeeting(agenda.MeetingTitle(titleFlag), agenda.Username(userNameFlag))
			if err != nil {
				panic(err)
			}
		} else {

			err := agenda.RemoveParticipatorFromMeeting(agenda.MeetingTitle(titleFlag), agenda.Username(userNameFlag))
			if err != nil {
				panic(err)
			}
		}

		// agenda.SaveAll()

	},
}

func init() {
	RootCmd.AddCommand(editCmd)

	editCmd.Flags().StringP("title", "t", "", "edit info for title")
	editCmd.Flags().StringP("username", "u", "", "edit info for username")
	editCmd.Flags().BoolP("invite", "i", false, "edit info for add")
	editCmd.Flags().BoolP("del", "d", false, "edit info for del")

}