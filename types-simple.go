package wellnessliving

import "encoding/json"

type StringToStringMap map[string]string

func (m *StringToStringMap) UnmarshalJSON(contents []byte) error {
	if string(contents) == "[]" {
		*m = map[string]string{}
	} else {
		v := map[string]string{}
		err := json.Unmarshal(contents, &v)
		if err != nil {
			return err
		}
		*m = v
	}
	return nil
}

type StringToAnyMap map[string]interface{}

func (m *StringToAnyMap) UnmarshalJSON(contents []byte) error {
	if string(contents) == "[]" {
		*m = map[string]interface{}{}
	} else {
		v := map[string]interface{}{}
		err := json.Unmarshal(contents, &v)
		if err != nil {
			return err
		}
		*m = v
	}
	return nil
}
