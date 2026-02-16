package database

import (
	"fmt"
)

func (d *InMemory) Connect_db() (Storage, error) {
	d.longToShort = make(map[string]string)
	d.shortToLong = make(map[string]string)
	return d, nil
}

func (d *InMemory) Check_long_link(long_link string) bool {
	d.RLock()
	defer d.RUnlock()
	_, ok := d.longToShort[long_link]
	return ok
}

func (d *InMemory) Check_short_link(short_link string) bool {
	d.RLock()
	defer d.RUnlock()
	_, ok := d.shortToLong[short_link]
	return ok
}

func (d *InMemory) Get_short_link(long_link string) (string, error) {
	d.RLock()
	defer d.RUnlock()
	short_link, ok := d.longToShort[long_link]
	if !ok {
		return "", fmt.Errorf("Long link %s not exist in database", long_link)
	}
	return short_link, nil
}

func (d *InMemory) Get_long_link(short_link string) (string, error) {
	d.RLock()
	defer d.RUnlock()
	long_link, ok := d.shortToLong[short_link]
	if !ok {
		return "", fmt.Errorf("Short link %s not exist in database", short_link)
	}
	return long_link, nil
}

func (d *InMemory) Store_in_db(short_link string, long_link string) error {
	d.Lock()
	defer d.Unlock()
	if _, ok := d.shortToLong[short_link]; ok {
		return fmt.Errorf("Short link %s already exist in database", short_link)
	}
	if _, ok := d.longToShort[long_link]; ok {
		return fmt.Errorf("Long link %s already exist in database", long_link)
	}
	d.shortToLong[short_link] = long_link
	d.longToShort[long_link] = short_link
	return nil
}
