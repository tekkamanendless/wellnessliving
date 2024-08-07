package wellnessliving

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Bool is a boolean, which could be represented as a bool, an integer, or a string.
type Bool bool

func (d *Bool) UnmarshalJSON(contents []byte) error {
	{
		var v bool
		err := json.Unmarshal(contents, &v)
		if err == nil {
			*d = Bool(v)
			return nil
		}
	}
	{
		var v int
		err := json.Unmarshal(contents, &v)
		if err == nil {
			*d = Bool(v != 0)
			return nil
		}
	}
	{
		var v string
		err := json.Unmarshal(contents, &v)
		if err == nil {
			if v == "" {
				return nil
			}
			f, err := strconv.ParseBool(v)
			if err != nil {
				return err
			}
			*d = Bool(f)
			return nil
		}
	}
	return fmt.Errorf("bool: could not parse: %q", contents)
}

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

	if v == "" {
		return nil
	}
	if v == "0000-00-00" {
		return nil
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

	if v == "" {
		return nil
	}
	if v == "0000-00-00 00:00:00" {
		return nil
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

	if v == "" {
		return nil
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
			if v == "" {
				return nil
			}
			f, err := strconv.ParseFloat(v, 64)
			if err == nil {
				*d = Float(f)
				return nil
			}
		}
	}
	return fmt.Errorf("float: could not parse: %q", contents)
}

// Integer is an integer, which could be represented as an integer or a string.
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
			if v == "" {
				return nil
			}
			f, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return err
			}
			*d = Integer(f)
			return nil
		}
	}
	return fmt.Errorf("integer: could not parse: %q", contents)
}
