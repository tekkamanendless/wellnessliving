package wellnessliving

type ErrorResponse struct {
	Errors  []Error `json:"a_error"`
	Class   string  `json:"class"`
	Code    *int    `json:"code"`
	Message string  `json:"message"`
	Status  string  `json:"status"`
	Version string  `json:"version"`
}

type Error struct {
	HTMLMessage string  `json:"html_message"`
	ID          *int    `json:"id"`
	Field       *string `json:"s_field"`
	Message     string  `json:"s_message"`
	SID         string  `json:"sid"`
}
