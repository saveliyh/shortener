package logic

import "testing"

func Test_create_short_link(t *testing.T) {

	name := "create_short_link test"

	short_link := create_short_link()
	if len(short_link) != LINK_LENGTH {
		t.Errorf("%s: short_link length is not equal to %d", name, LINK_LENGTH)
	}
}

func Test_unshorten_link_same_long_link(t *testing.T) {
	name := "unshorten_link test for same input"
	storage := TestStorage{}
	short_link := "short_link1"
	long_link1, err := Unshorten_link(short_link, storage)
	if err != nil {
		t.Errorf("%s: error is not nil", name)
	}
	long_link2, err := Unshorten_link(short_link, storage)
	if err != nil {
		t.Errorf("%s: error is not nil", name)
	}
	if long_link1 != long_link2 {
		t.Errorf("%s: long_link1 is not equal to long_link2", name)
	}

}
func Test_unshorten_link_different_long_link(t *testing.T) {
	name := "unshorten_link test for different input"
	storage := TestStorage{}
	short_link1 := "short_link1"
	long_link1, err := Unshorten_link(short_link1, storage)
	if err != nil {
		t.Errorf("%s: error is not nil", name)
	}
	short_link2 := "short_link2"
	long_link2, err := Unshorten_link(short_link2, storage)
	if err != nil {
		t.Errorf("%s: error is not nil", name)
	}
	if long_link1 == long_link2 {
		t.Errorf("%s: long_link1 is equal to long_link2", name)
	}
}

func Test_unshorten_link_wrong_input(t *testing.T) {
	name := "unshorten_link test for wrong input"
	storage := TestStorage{}
	short_link := "short_link"
	_, err := Unshorten_link(short_link, storage)
	if err == nil {
		t.Errorf("%s: Had to return \"not found error\"", name)
	}
}
