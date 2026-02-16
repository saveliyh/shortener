package database

import "testing"

func Test_inMemory_Connect_db(t *testing.T) {
	name := "inMemory test for Connect_db"
	d, err := (&InMemory{}).Connect_db()
	if err != nil {
		t.Errorf("%s: error is not nil. %s", name, err)
	}
	if inMemmory, ok := d.(*InMemory); ok {
		if inMemmory.longToShort == nil {
			t.Errorf("%s: longToShort is nil", name)
		}
		if inMemmory.shortToLong == nil {
			t.Errorf("%s: shortToLong is nil", name)
		}
	}
}

func Test_inMemory_Check_long_link_exist_input(t *testing.T) {
	name := "inMemory test for Check_long_link exist input"
	d := InMemory{
		longToShort: map[string]string{
			"long_link1": "short_link1",
		},
		shortToLong: map[string]string{
			"short_link1": "long_link1",
		},
	}
	if !d.Check_long_link("long_link1") {
		t.Errorf("%s: long_link1 not found in storage", name)
	}
}

func Test_inMemory_Check_long_link_not_exist_input(t *testing.T) {
	name := "inMemory test for Check_long_link not exist input"
	d := InMemory{
		longToShort: map[string]string{},
		shortToLong: map[string]string{},
	}
	if d.Check_long_link("long_link1") {
		t.Errorf("%s: not existed long_link found in storage", name)
	}
}

func Test_inMemory_Check_short_link_exist_input(t *testing.T) {
	name := "inMemory test for Check_short_link exist input"
	d := InMemory{
		longToShort: map[string]string{
			"long_link1": "short_link1",
		},
		shortToLong: map[string]string{
			"short_link1": "long_link1",
		},
	}
	if !d.Check_short_link("short_link1") {
		t.Errorf("%s: short_link1 not found in storage", name)
	}
}

func Test_inMemory_Check_short_link_not_exist_input(t *testing.T) {
	name := "inMemory test for Check_short_link not exist input"
	d := InMemory{
		longToShort: map[string]string{},
		shortToLong: map[string]string{},
	}
	if d.Check_short_link("short_link1") {
		t.Errorf("%s: not existed short_link found in storage", name)
	}
}

func Test_inMemory_Get_short_link_correct_input(t *testing.T) {
	name := "inMemory test for Get_short_link correct input"
	d := InMemory{
		longToShort: map[string]string{
			"long_link1": "short_link1",
		},
		shortToLong: map[string]string{
			"short_link1": "long_link1",
		},
	}
	short_link, err := d.Get_short_link("long_link1")
	if err != nil {
		t.Errorf("%s: error is not nil. %s", name, err)
	}
	if short_link != "short_link1" {
		t.Errorf("%s: returned short_link is not equal to short_link in storage", name)
	}
}

func Test_inMemory_Get_short_link_wrong_input(t *testing.T) {
	name := "inMemory test for Get_short_link wrong input"
	d := InMemory{
		longToShort: map[string]string{},
		shortToLong: map[string]string{},
	}
	_, err := d.Get_short_link("long_link1")
	if err == nil {
		t.Errorf("%s: error is nil. Must be \"not found\"", name)
	}
}

func Test_inMemory_Get_long_link_correct_input(t *testing.T) {
	name := "inMemory test for Get_long_link correct input"
	d := InMemory{
		longToShort: map[string]string{
			"long_link1": "short_link1",
		},
		shortToLong: map[string]string{
			"short_link1": "long_link1",
		},
	}
	long_link, err := d.Get_long_link("short_link1")
	if err != nil {
		t.Errorf("%s: error is not nil. %s", name, err)
	}
	if long_link != "long_link1" {
		t.Errorf("%s: returned long_link is not equal to long_link in storage", name)
	}
}

func Test_inMemory_Get_long_link_wrong_input(t *testing.T) {
	name := "inMemory test for Get_long_link wrong input"
	d := InMemory{
		longToShort: map[string]string{},
		shortToLong: map[string]string{},
	}
	_, err := d.Get_long_link("short_link1")
	if err == nil {
		t.Errorf("%s: error is nil. Must be \"not found\"", name)
	}
}

func Test_inMemory_Store_in_db_correct_input(t *testing.T) {
	name := "inMemory test for Store_in_db correct input"
	d := InMemory{
		longToShort: map[string]string{},
		shortToLong: map[string]string{},
	}
	err := d.Store_in_db("short_link1", "long_link1")
	if err != nil {
		t.Errorf("%s: error is not nil. %s", name, err)
	}
	if _, ok := d.longToShort["long_link1"]; !ok {
		t.Errorf("%s: long_link1 not added to storage", name)
	}
	if _, ok := d.shortToLong["short_link1"]; !ok {
		t.Errorf("%s: short_link1 not added to storage", name)
	}
}

func Test_inMemory_Store_in_db_add_exist_short_link(t *testing.T) {
	name := "inMemory test for Store_in_db add exist short link"
	d := InMemory{
		longToShort: map[string]string{
			"long_link1": "short_link1",
		},
		shortToLong: map[string]string{
			"short_link1": "long_link1",
		},
	}
	err := d.Store_in_db("short_link1", "long_link2")
	if err == nil {
		t.Errorf("%s: error is nil. Must be \"short link already exist\"", name)
	}
}

func Test_inMemory_Store_in_db_add_exist_long_link(t *testing.T) {
	name := "inMemory test for Store_in_db add exist long link"
	d := InMemory{
		longToShort: map[string]string{
			"long_link1": "short_link1",
		},
		shortToLong: map[string]string{
			"short_link1": "long_link1",
		},
	}
	err := d.Store_in_db("short_link2", "long_link1")
	if err == nil {
		t.Errorf("%s: error is nil. Must be \"long link already exist\"", name)
	}
}
