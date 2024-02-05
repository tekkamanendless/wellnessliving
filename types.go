package wellnessliving

import (
	"fmt"
)

// BaseResponse is the base of all responses.
// The fields here will be present in every response.
type BaseResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// ErrorResponse is an error response.
type ErrorResponse struct {
	BaseResponse

	Errors  []Error `json:"a_error"`
	Class   string  `json:"class"`
	Code    *int    `json:"code"`
	Message string  `json:"message"`
}

type Error struct {
	HTMLMessage string  `json:"html_message"`
	ID          *int    `json:"id"`
	Field       *string `json:"s_field"`
	Message     string  `json:"s_message"`
	SID         string  `json:"sid"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%+v", *r)
}

// EventListResponse is the response from "/Wl/Event/EventList.json".
type EventListResponse struct {
	BaseResponse

	EnrollmentBlockList StringMap `json:"a_enrollment_block_list"`
	EventList           []Event   `json:"a_event_list"`
}

type Event struct {
	ClassTab []string   `json:"a_class_tab"`
	Logo     Logo       `json:"a_logo"`
	Schedule []Schedule `json:"a_schedule"`
	//TODO:"a_search_tag": [],
	CanCancel         bool      `json:"can_cancel"`
	EarlybirdEndDate  *Date     `json:"dl_early"`
	EndDate           Date      `json:"dl_end"`
	StartDate         Date      `json:"dl_start"`
	SessionDTU        DateTime  `json:"dtu_session"`
	SessionAll        int       `json:"i_session_all"`
	SessionFuture     int       `json:"i_session_future"`
	SessionPast       int       `json:"i_session_past"`
	IsAgeRestrict     bool      `json:"is_age_restrict"`
	IsAvailable       bool      `json:"is_available"`
	IsBlock           bool      `json:"is_block"`
	IsBookable        bool      `json:"is_bookable"`
	IsBooked          bool      `json:"is_booked"`
	IsClosed          bool      `json:"is_closed"`
	IsFull            bool      `json:"is_full"`
	IsOnline          bool      `json:"is_online"`
	IsOnlinePrivate   bool      `json:"is_online_private"`
	IsOpen            bool      `json:"is_open"`
	IsPromotionOnly   bool      `json:"is_promotion_only"`
	IsProrate         bool      `json:"is_prorate"`
	IsVirtual         bool      `json:"is_virtual"`
	ClassID           string    `json:"k_class"`
	ClassPeriodID     string    `json:"k_class_period"`
	EnrollmentBlockID string    `json:"k_enrollment_block"`
	LocationID        string    `json:"k_location"`
	PriceTotal        Currency  `json:"m_price_total"`
	PriceTotalEarly   *Currency `json:"m_price_total_early"`
	AgeRestrictText   string    `json:"text_age_restrict"`
	Title             string    `json:"text_title"`
	URLBook           string    `json:"url_book"`
	XMLDescription    string    `json:"xml_description"`
}

type Logo struct {
	Business string `json:"k_business"`
	Class    string `json:"k_class"`
	Image    Image  `json:"a_image"`
	IsOwn    bool   `json:"is_own"`

	Height       int    `json:"i_height"`
	HeightSource int    `json:"i_height_src"` // Not present if "a_image" is set.
	Rotate       int    `json:"i_rotate"`     // Not present if "a_image" is set.
	Width        int    `json:"i_width"`
	WidthSource  int    `json:"i_width_src"`   // Not present if "a_image" is set.
	IDTypeSource int    `json:"id_type_src"`   // Not present if "a_image" is set.
	IsResize     bool   `json:"is-resize"`     // Not present if "a_image" is set.
	URLView      string `json:"url-view"`      // Not present if "a_image" is set.
	URLThumbnail string `json:"url-thumbnail"` // Not present if "a_image" is set.
	IsOld        bool   `json:"is_old"`
	URL          string `json:"s_url"`
}

type Image struct {
	Height       int    `json:"i_height"`
	HeightSource int    `json:"i_height_src"`
	Rotate       int    `json:"i_rotate"`
	Width        int    `json:"i_width"`
	WidthSource  int    `json:"i_width_src"`
	IDTypeSource int    `json:"id_type_src"`
	IsResize     bool   `json:"is-resize"`
	URLView      string `json:"url-view"`
	URLThumbnail string `json:"url-thumbnail"`
}

type Schedule struct {
	Day           map[string]int `json:"a_day"`
	StaffMember   []StaffMember  `json:"a_staff_member"`
	EndDate       Date           `json:"dl_end"`
	StartDate     Date           `json:"dl_start"`
	IsDay         bool           `json:"is_day"`
	ClassPeriodID string         `json:"k_class_period"`
	LocationID    string         `json:"k_location"`
	LocationText  string         `json:"text_location"`
	TimeText      string         `json:"text_time"`
}

type StaffMember struct {
	StaffMemberID int    `json:"k_staff_member"`
	BusinessRole  string `json:"text_business_role"`
	Mail          string `json:"text_mail"`
	NameFirst     string `json:"text_name_first"`
	NameFull      string `json:"text_name_full"`
	NameLast      string `json:"text_name_last"`
	UID           string `json:"uid"`
}
