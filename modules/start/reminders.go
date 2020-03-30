package start

import (
	"time"

	"github.com/olivia-ai/olivia/user"
)

func init() {
	RegisterModule(Module{
		Action: CheckReminders,
	})
}

// CheckReminders will check the dates of the user's reminder and if they are outdated, remove them
func CheckReminders(token string) {
	reminders := user.GetUserInformation(token).Reminders

	// Iterate through the reminders to check if they are outdated
	for i, reminder := range reminders {
		date, _ := time.Parse("01/02/2006", reminder.Date)

		// If the date of the reminder has already been passed
		if time.Now().After(date) {
			user.ChangeUserInformation(token, func(information user.Information) user.Information {
				reminders := information.Reminders

				// Removes the element from the reminders slice
				if len(reminders) == 1 {
					reminders = []user.Reminder{}
				} else {
					reminders[i] = reminders[len(reminders)-1]
					reminders = reminders[:len(reminders)-1]
				}

				// Set the updated slice
				information.Reminders = reminders

				return information
			})
		}
	}
}
