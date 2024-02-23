package wellnessliving

import (
	"encoding/json"
	"fmt"
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
		return fmt.Errorf("date: could not unmarshal string: %w", err)
	}

	location, err := time.LoadLocation("GMT")
	if err != nil {
		return fmt.Errorf("date: could not load location: %w", err)
	}

	d.Time, err = time.ParseInLocation("2006-01-02", v, location)
	if err != nil {
		return fmt.Errorf("date: could not parse string: %w", err)
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
		return fmt.Errorf("datetime: could not unmarshal string: %w", err)
	}

	location, err := time.LoadLocation("GMT")
	if err != nil {
		return fmt.Errorf("datetime: could not load location: %w", err)
	}

	d.Time, err = time.ParseInLocation("2006-01-02 15:04:05", v, location)
	if err != nil {
		return fmt.Errorf("datetime: could not parse string: %w", err)
	}
	return nil
}

// Currency is an amount of money.
type Currency float64

func (d *Currency) UnmarshalJSON(contents []byte) error {
	var v string
	err := json.Unmarshal(contents, &v)
	if err != nil {
		return fmt.Errorf("currency: could not unmarshal string: %w", err)
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return fmt.Errorf("currency: could not parse string: %w", err)
	}
	*d = Currency(f)
	return nil
}

// Float is an amount of money.
type Float float64

func (d *Float) UnmarshalJSON(contents []byte) error {
	{
		var v float64
		err := json.Unmarshal(contents, &v)
		if err == nil {
			*d = Float(v)
			return nil
		}
	}
	{
		var v string
		err := json.Unmarshal(contents, &v)
		if err == nil {
			f, err := strconv.ParseFloat(v, 64)
			if err == nil {
				*d = Float(f)
				return nil
			}
		}
	}
	return fmt.Errorf("float: could not parse: %s", contents)
}

// Integer is an integer, which could be represented as an integer or a string string.
type Integer int

func (d *Integer) UnmarshalJSON(contents []byte) error {
	{
		var v int
		err := json.Unmarshal(contents, &v)
		if err == nil {
			*d = Integer(v)
			return nil
		}
	}
	{
		var v string
		err := json.Unmarshal(contents, &v)
		if err == nil {
			f, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return err
			}
			*d = Integer(f)
			return nil
		}
	}
	return fmt.Errorf("integer: could not parse: %s", contents)
}
