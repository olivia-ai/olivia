package user

import "testing"

func TestUserInformation(t *testing.T) {
	SetUserInformation("1", Information{
		Name: "Hugo",
	})

	if GetUserInformation("1").Name != "Hugo" {
		t.Errorf("SetUserInformation and/or GetUserInformation failed.")
	}

	ChangeUserInformation("1", func(information Information) Information {
		information.Name = "Steve"
		return information
	})

	if GetUserInformation("1").Name != "Steve" {
		t.Errorf("ChangeUserInformation failed.")
	}
}
