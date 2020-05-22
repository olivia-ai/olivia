package start

import (
	"fmt"
	"github.com/olivia-ai/olivia/util"
	"strings"
	"time"

	"github.com/olivia-ai/olivia/user"
)

func init() {
	RegisterModule(Module{
		Action: CheckReminders,
	})
}

// CheckReminders will check the dates of the user's reminder and if they are outdated, remove them
func CheckReminders(token, locale string) {
	reminders := user.GetUserInformation(token).Reminders
	var messages []string

	// Iterate through the reminders to check if they are outdated
	for i, reminder := range reminders {
		date, _ := time.Parse("01/02/2006 03:04", reminder.Date)

		now := time.Now()
		// If the date is today
		if date.Year() == now.Year() && date.Day() == now.Day() && date.Month() == now.Month() {
			messages = append(messages, fmt.Sprintf("“%s”", reminder.Reason))

			// Removes the current reminder
			RemoveUserReminder(token, i)
		}
	}

	// Send the startup message
	if len(messages) != 0 {
		// If the message is already filled in return.
		if GetMessage() != "" {
			return
		}

		// Set the message with the current reminders
		listRemindersMessage := util.GetMessage(locale, "list reminders")
		if listRemindersMessage == "" {
			return
		}

		message := fmt.Sprintf(
			listRemindersMessage,
			user.GetUserInformation(token).Name,
			strings.Join(messages, ", "),
		)
		SetMessage(message)
	}
}

// RemoveUserReminder removes the reminder at a specific index in the user's information
func RemoveUserReminder(token string, index int) {
	user.ChangeUserInformation(token, func(information user.Information) user.Information {
		reminders := information.Reminders

		// Removes the element from the reminders slice
		if len(reminders) == 1 {
			reminders = []user.Reminder{}
		} else {
			reminders[index] = reminders[len(reminders)-1]
			reminders = reminders[:len(reminders)-1]
		}

		// Set the updated slice
		information.Reminders = reminders

		return information
	})
}
