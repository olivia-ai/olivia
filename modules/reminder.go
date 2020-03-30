package modules

import (
	"fmt"

	"github.com/olivia-ai/olivia/user"

	"github.com/olivia-ai/olivia/language"

	"github.com/olivia-ai/olivia/language/date"
)

var reminderSetterTag = "reminder setter"

func init() {
	RegisterModule(Module{
		Tag: reminderSetterTag,
		Patterns: []string{
			"Remind me to call mom",
			"Note that I have an exam",
		},
		Responses: []string{
			"Noted! I will remind you: “%s” for the %s",
		},
		Replacer: ReminderSetterReplacer,
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
	formattedDate := date.Format("01/02/2006")

	// Add the reminder inside the user's information
	user.ChangeUserInformation(token, func(information user.Information) user.Information {
		information.Reminders = append(information.Reminders, user.Reminder{
			Reason: reason,
			Date:   formattedDate,
		})

		return information
	})

	return areaTag, fmt.Sprintf(response, reason, formattedDate)
}
