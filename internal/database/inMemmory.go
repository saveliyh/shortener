package database

import "errors"

func (d InMemory) Connect_db() (Storage, error) {
	d.longToShort = make(map[string]string)
	d.shortToLong = make(map[string]string)
	return d, nil
}

func (d InMemory) Check_long_link(long_link string) bool {
	_, ok := d.longToShort[long_link]
	return ok
}

func (d InMemory) Check_short_link(short_link string) bool {
	_, ok := d.shortToLong[short_link]
	return ok
}

func (d InMemory) Get_short_link(long_link string) (string, error) {
	short_link, ok := d.longToShort[long_link]
	if !ok {
		return "", errors.New("Long link not exist in database")
	}
	return short_link, nil
}

func (d InMemory) Get_long_link(short_link string) (string, error) {
	long_link, ok := d.shortToLong[short_link]
	if !ok {
		return "", errors.New("Short link not exist in database")
	}
	return long_link, nil
}

func (d InMemory) Store_in_db(short_link string, long_link string) error {
	if _, ok := d.shortToLong[short_link]; ok {
		return errors.New("Short link already exist in database")
	}
	if _, ok := d.longToShort[long_link]; ok {
		return errors.New("Long link already exist in database")
	}
	d.shortToLong[short_link] = long_link
	d.longToShort[long_link] = short_link
	return nil
}
