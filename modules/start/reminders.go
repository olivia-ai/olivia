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

	for i, reminder := range reminders {
		date, _ := time.Parse("01-02-2006", reminder.Date)

		// If the date of the reminder has already been passed
		if time.Now().After(date) {
			// Change the user information to remove the current reminder
			user.ChangeUserInformation(token, func(information user.Information) user.Information {
				information.Reminders = append(information.Reminders[:i], information.Reminders[i+1:]...)

				return information
			})
		}
	}
}
