package wellnessliving

import "encoding/json"

type StringMap map[string]string

func (m *StringMap) UnmarshalJSON(contents []byte) error {
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
