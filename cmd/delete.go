package cmd

import (
	"agenda"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete for further use",
	Long: `delete for further use and u need to input username, password.
	it will be better if email and phone is provider`,
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Error[delete]： %v\n", err)
			}
		}()

		fmt.Println("delete called")

		meetingBoolFlag, _ := cmd.Flags().GetBool("meeting")
		userBoolFlag, _ := cmd.Flags().GetBool("user")
		titleFlag, _ := cmd.Flags().GetString("title")
		if meetingBoolFlag && userBoolFlag {
			panic(errors.New("The flag meeting(m) and user(u) can't appear at the same time"))

		} else if userBoolFlag {

			// agenda.ClearAllMeeting()
			if err := agenda.CancelAccount(); err != nil {
				panic(err)
			}
			fmt.Print("CancelAccount successfully\n")
		} else {

			title := agenda.MeetingTitle(titleFlag)

			if err := agenda.CancelMeeting(title); err != nil {
				panic(err)
			}

		}
		// agenda.SaveAll()

	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolP("meeting", "m", false, "delete info for meeting")
	deleteCmd.Flags().BoolP("user", "u", false, "delete info for user")
	deleteCmd.Flags().StringP("title", "t", "", "delete info for email")

}
