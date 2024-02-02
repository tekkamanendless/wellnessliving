package wellnessliving

import (
	"encoding/json"
	"strconv"
	"time"
)

// Date is a specific date.
type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(contents []byte) error {
	var v string
	err := json.Unmarshal(contents, &v)
	if err != nil {
		return err
	}

	location, err := time.LoadLocation("GMT")
	if err != nil {
		return err
	}

	d.Time, err = time.ParseInLocation("2006-01-02", v, location)
	if err != nil {
		return err
	}
	return nil
}

// DateTime is a specific date/time.
type DateTime struct {
	time.Time
}

func (d *DateTime) UnmarshalJSON(contents []byte) error {
	var v string
	err := json.Unmarshal(contents, &v)
	if err != nil {
		return err
	}

	location, err := time.LoadLocation("GMT")
	if err != nil {
		return err
	}

	d.Time, err = time.ParseInLocation("2006-01-02 15:04:05", v, location)
	if err != nil {
		return err
	}
	return nil
}

// Currency is an amount of money.
type Currency float64

func (d *Currency) UnmarshalJSON(contents []byte) error {
	var v string
	err := json.Unmarshal(contents, &v)
	if err != nil {
		return err
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return err
	}
	*d = Currency(f)
	return nil
}
