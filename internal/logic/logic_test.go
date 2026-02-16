package logic

import "testing"

func Test_create_short_link(t *testing.T) {

	name := "create_short_link test"

	short_link := create_short_link()
	if len(short_link) != LINK_LENGTH {
		t.Errorf("%s: short_link length is not equal to %d", name, LINK_LENGTH)
	}
}

func Test_unshorten_link_same_input(t *testing.T) {
	name := "unshorten_link test for same input"
	storage := TestStorage{}
	short_link := "short_link1"
	long_link1, err := Unshorten_link(short_link, storage)
	if err != nil {
		t.Errorf("%s: error is not nil. %s", name, err)
	}
	long_link2, err := Unshorten_link(short_link, storage)
	if err != nil {
		t.Errorf("%s: error is not nil. %s", name, err)
	}
	if long_link1 != long_link2 {
		t.Errorf("%s: got different long_link when retrieved", name)
	}

}
func Test_unshorten_link_different_input(t *testing.T) {
	name := "unshorten_link test for different input"
	storage := TestStorage{}
	short_link1 := "short_link1"
	long_link1, err := Unshorten_link(short_link1, storage)
	if err != nil {
		t.Errorf("%s: error is not nil. %s", name, err)
	}
	short_link2 := "short_link2"
	long_link2, err := Unshorten_link(short_link2, storage)
	if err != nil {
		t.Errorf("%s: error is not nil. %s", name, err)
	}
	if long_link1 == long_link2 {
		t.Errorf("%s: got same long_link for different short_links", name)
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

func Test_get_short_link_add_and_retrieve(t *testing.T) {
	name := "get_short_link test for add and retrieve"
	storage := TestStorage{}
	long_link := "long_link3"
	short_link1, err := Get_short_link(long_link, storage)
	storage.test_store_in_db(short_link1)
	if err != nil {
		t.Errorf("%s: error is not nil. %s", name, err)
	}
	short_link2, err := Get_short_link(long_link, storage)
	if err != nil {
		t.Errorf("%s: error is not nil. %s", name, err)
	}
	println(short_link1, short_link2)
	if short_link1 != short_link2 {
		t.Errorf("%s: got new short_link when retrieved", name)
	}
}

func Test_get_short_link_correct_link_length(t *testing.T) {
	name := "get_short_link test for correct link length"
	storage := TestStorage{}
	long_link := "long_link"
	short_link, err := Get_short_link(long_link, storage)
	if err != nil {
		t.Errorf("%s: error is not nil. %s", name, err)
	}
	if len(short_link) != LINK_LENGTH {
		t.Errorf("%s: short_link length is not equal to %d", name, LINK_LENGTH)
	}
}
