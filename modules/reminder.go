package modules

import (
	"fmt"
	"strings"

	"github.com/olivia-ai/olivia/user"

	"github.com/olivia-ai/olivia/language"

	"github.com/olivia-ai/olivia/language/date"
)

var (
	reminderSetterTag = "reminder setter"
	reminderGetterTag = "reminder getter"
)

func init() {
	RegisterModule(Module{
		Tag: reminderSetterTag,
		Patterns: []string{
			"Remind me to call mom",
			"Note that I have an exam",
			"Remind me that I have to  tomorrow at 9pm",
		},
		Responses: []string{
			"Noted! I will remind you: “%s” for the %s",
		},
		Replacer: ReminderSetterReplacer,
	})

	RegisterModule(Module{
		Tag: reminderGetterTag,
		Patterns: []string{
			"What did I ask for you to remember",
			"Could you list my reminders",
		},
		Responses: []string{
			"You asked me to remember those things:\n%s",
		},
		Replacer: ReminderGetterReplacer,
	})
}

// ReminderSetterReplacer replaces the pattern contained inside the response by the date of the reminder
// and its reason.
// See modules/modules.go#Module.Replacer() for more details.
func ReminderSetterReplacer(entry, response, token string) (string, string) {
	// Search the time and
	sentence, date := date.SearchTime(entry)
	reason := language.SearchReason(sentence)

	// Format the date
	formattedDate := date.Format("01/02/2006 03:04")

	// Add the reminder inside the user's information
	user.ChangeUserInformation(token, func(information user.Information) user.Information {
		information.Reminders = append(information.Reminders, user.Reminder{
			Reason: reason,
			Date:   formattedDate,
		})

		return information
	})

	return reminderSetterTag, fmt.Sprintf(response, reason, formattedDate)
}

// ReminderSetterReplacer gets the reminders in the user's information and replaces the pattern in the
// response patterns by the current reminders
// See modules/modules.go#Module.Replacer() for more details.
func ReminderGetterReplacer(_, response, token string) (string, string) {
	reminders := user.GetUserInformation(token).Reminders
	var formattedReminders []string

	// Iterate through the reminders and parse them
	for _, reminder := range reminders {
		formattedReminder := fmt.Sprintf("- “%s” for the %s", reminder.Reason, reminder.Date)
		formattedReminders = append(formattedReminders, formattedReminder)
	}

	// If no reminder has been found
	if len(formattedReminders) == 0 {
		return reminderGetterTag, "You have no reminders saved."
	}

	return reminderGetterTag, fmt.Sprintf(response, strings.Join(formattedReminders, " "))
}
