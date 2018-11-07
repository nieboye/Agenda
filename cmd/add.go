package cmd

import (
	"agenda"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// addCmd represents the create Meeting command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "create Meeting for further use",
	Long: `create Meeting for further use and u need to input username, password.
	it will be better if email and phone is provider`,
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Error[create Meeting]： %v\n", err)
			}
		}()

		fmt.Println("create Meeting called")

		startTimeFlag, _ := cmd.Flags().GetString("startTime")
		endTimeFlag, _ := cmd.Flags().GetString("endTime")
		participatorsFlag, _ := cmd.Flags().GetStringSlice("participator")
		titleFlag, _ := cmd.Flags().GetString("title")
		// fmt.Printf("start: %v end: %v\n", startTimeFlag, endTimeFlag)
		startTime, _ := time.Parse("2006-01-02 15:04:05", startTimeFlag)
		endTime, _ := time.Parse("2006-01-02 15:04:05", endTimeFlag)

		fmt.Printf("start: %v end: %v\n", startTime, endTime)

		participators := make([]agenda.Username, 0, len(participatorsFlag))
		participators = append(participators, agenda.LoginedUser().Name)
		for _, participator := range participatorsFlag {
			participators = append(participators, agenda.Username(participator))
		}
		meetingInfo := agenda.MakeMeetingInfo(agenda.MeetingTitle(titleFlag),
			agenda.LoginedUser().Name, participators,
			startTime, endTime)

		if _, err := agenda.SponsorMeeting(meetingInfo); err != nil {
			panic(err)
		}
		// agenda.SaveAll()
		fmt.Print("sucessfully create meeting\n")

	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().StringP("startTime", "s", "", "create Meeting info for startTime")
	addCmd.Flags().StringP("endTime", "e", "", "create Meeting info for endTime")
	addCmd.Flags().StringP("participator", "p", "", "create Meeting info for participator")
	addCmd.Flags().StringP("title", "t", "", "create Meeting info for title")

}