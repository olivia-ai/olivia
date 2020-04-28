package locales

import "testing"

func TestGetNameByTag(t *testing.T) {
	name := "english"
	excepted := "en"
	tag := GetTagByName(name)

	if tag != excepted {
		t.Errorf("GetNameByTag() failed, excepted %s got %s.", excepted, tag)
	}
}

func TestGetTagByName(t *testing.T) {
	tag := "en"
	excepted := "english"
	name := GetNameByTag(tag)

	if name != excepted {
		t.Errorf("GetTagByName() failed, excepted %s got %s.", excepted, name)
	}
}

func TestExists(t *testing.T) {
	tag := "en"
	excepted := true
	exists := Exists(tag)

	if exists != excepted {
		t.Errorf("Exists() failed, excepted %t got %t.", excepted, exists)
	}
}
