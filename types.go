package wellnessliving

import (
	"fmt"
)

// BaseResponse is the base of all responses.
// The fields here will be present in every response.
type BaseResponse struct {
	Status  string `json:"status"`
	Version string `json:"s_version"`
}

// ErrorResponse is an error response.
type ErrorResponse struct {
	BaseResponse

	Errors  []Error  `json:"a_error"`
	Class   string   `json:"class"`
	Code    *Integer `json:"code"`
	Message string   `json:"message"`
}

type Error struct {
	HTMLMessage string   `json:"html_message"`
	ID          *Integer `json:"id"`
	Field       *string  `json:"s_field"`
	Message     string   `json:"s_message"`
	SID         string   `json:"sid"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%+v", *r)
}

// NotepadResponse is the response from "/Core/Passport/Login/Enter/Notepad.json".
type NotepadResponse struct {
	BaseResponse

	RegionID *string `json:"id_region"`
	Hash     string  `json:"s_hash"`
	Notepad  string  `json:"s_notepad"`
}

// EnterResponse is the response from "/Core/Passport/Login/Enter/Enter.json".
type EnterResponse struct {
	URLRedirect string `json:"url_redirect"`
}

// EventListResponse is the response from "/Wl/Event/EventList.json".
type EventListResponse struct {
	BaseResponse

	EnrollmentBlockList StringToStringMap `json:"a_enrollment_block_list"`
	EventList           []Event           `json:"a_event_list"`
}

type SearchTag struct {
	SearchTagID string `json:"k_search_tag"`
	Title       string `json:"text_title"`
}

type Event struct {
	ClassTab          []string        `json:"a_class_tab"`
	Logo              Logo            `json:"a_logo"`
	Schedule          []EventSchedule `json:"a_schedule"`
	SearchTags        []SearchTag     `json:"a_search_tag"`
	CanCancel         Bool            `json:"can_cancel"`
	EarlybirdEndDate  *Date           `json:"dl_early"`
	EndDate           Date            `json:"dl_end"`
	StartDate         Date            `json:"dl_start"`
	SessionDTU        *DateTime       `json:"dtu_session"` // Date of the closest session of the event.
	SessionAll        Integer         `json:"i_session_all"`
	SessionFuture     Integer         `json:"i_session_future"`
	SessionPast       Integer         `json:"i_session_past"`
	IsAgeRestrict     Bool            `json:"is_age_restrict"`
	IsAvailable       Bool            `json:"is_available"`
	IsBlock           Bool            `json:"is_block"`
	IsBookable        Bool            `json:"is_bookable"`
	IsBooked          Bool            `json:"is_booked"`
	IsClosed          Bool            `json:"is_closed"`
	IsFull            Bool            `json:"is_full"`
	IsOnline          Bool            `json:"is_online"`
	IsOnlinePrivate   Bool            `json:"is_online_private"`
	IsOpen            Bool            `json:"is_open"`
	IsPromotionOnly   Bool            `json:"is_promotion_only"`
	IsProrate         Bool            `json:"is_prorate"`
	IsVirtual         Bool            `json:"is_virtual"`
	ClassID           string          `json:"k_class"`
	ClassPeriodID     string          `json:"k_class_period"`
	EnrollmentBlockID string          `json:"k_enrollment_block"`
	LocationID        string          `json:"k_location"`
	PriceTotal        Currency        `json:"m_price_total"`
	PriceTotalEarly   *Currency       `json:"m_price_total_early"`
	AgeRestrictText   string          `json:"text_age_restrict"`
	Title             string          `json:"text_title"`
	URLBook           string          `json:"url_book"`
	XMLDescription    string          `json:"xml_description"`
}

type Logo struct {
	Business string `json:"k_business"`
	Class    string `json:"k_class"`
	Image    Image  `json:"a_image"`
	IsOwn    Bool   `json:"is_own"`

	Height       Integer `json:"i_height"`
	HeightSource Integer `json:"i_height_src"` // Not present if "a_image" is set.
	Rotate       Integer `json:"i_rotate"`     // Not present if "a_image" is set.
	Width        Integer `json:"i_width"`
	WidthSource  Integer `json:"i_width_src"`   // Not present if "a_image" is set.
	IDTypeSource Integer `json:"id_type_src"`   // Not present if "a_image" is set.
	IsResize     Bool    `json:"is-resize"`     // Not present if "a_image" is set.
	URLView      string  `json:"url-view"`      // Not present if "a_image" is set.
	URLThumbnail string  `json:"url-thumbnail"` // Not present if "a_image" is set.
	IsOld        Bool    `json:"is_old"`
	URL          string  `json:"s_url"`
}

type Image struct {
	Height       Integer `json:"i_height"`
	HeightSource Integer `json:"i_height_src"`
	Rotate       Integer `json:"i_rotate"`
	Width        Integer `json:"i_width"`
	WidthSource  Integer `json:"i_width_src"`
	IDTypeSource Integer `json:"id_type_src"`
	IsResize     Bool    `json:"is-resize"`
	URLView      string  `json:"url-view"`
	URLThumbnail string  `json:"url-thumbnail"`
}

type EventSchedule struct {
	Day           map[string]Integer `json:"a_day"`
	StaffMember   []StaffMember      `json:"a_staff_member"`
	EndDate       Date               `json:"dl_end"`
	StartDate     Date               `json:"dl_start"`
	IsDay         Bool               `json:"is_day"`
	ClassPeriodID string             `json:"k_class_period"`
	LocationID    string             `json:"k_location"`
	LocationText  string             `json:"text_location"`
	TimeText      string             `json:"text_time"`
}

type StaffMember struct {
	StaffMemberID Integer `json:"k_staff_member"`
	BusinessRole  string  `json:"text_business_role"`
	Mail          string  `json:"text_mail"`
	NameFirst     string  `json:"text_name_first"`
	NameFull      string  `json:"text_name_full"`
	NameLast      string  `json:"text_name_last"`
	UID           string  `json:"uid"`
}

// ClassResponse is the response from "/Wl/Classes/ClassView/Element.json".
type ClassResponse struct {
	BaseResponse

	ClassList map[string]Class `json:"a_class_list"`
}

type Class struct {
	ClassTab []string `json:"a_class_tab"`
	// TODO: "a_config": null,
	Schedule   []ClassSchedule `json:"a_schedule"`
	SearchTags []SearchTag     `json:"a_search_tag"`
	// TODO: "a_visits_required": [],
	HasOwnImage            Bool   `json:"has_own_image"`
	HTMLDescription        string `json:"html_description"`
	HTMLSpecialInstruction string `json:"html_special_instruction"`
	IsAgePublic            Bool   `json:"is_age_public"` // "0"
	// TODO: "i_age_from": null,
	// TODO: "i_age_to": null,
	IsBookable              Bool      `json:"is_bookable"`
	IsEvent                 Bool      `json:"is_event"`
	IsOnlinePrivate         Bool      `json:"is_online_private"`
	IsPromotionClient       Bool      `json:"is_promotion_client"`
	IsPromotionOnly         Bool      `json:"is_promotion_only"`
	IsPromotionStaff        Bool      `json:"is_promotion_staff"`
	IsSingleBuy             Bool      `json:"is_single_buy"`
	IsVirtual               Bool      `json:"is_virtual"`
	ClassID                 string    `json:"k_class"`
	Price                   *Currency `json:"m_price"`
	ShowSpecialInstructions Bool      `json:"show_special_instructions"` // "1"
	Title                   string    `json:"text_title"`
	XMLDescription          string    `json:"xml_description"`
	XMLSpecialInstruction   string    `json:"xml_special_instruction"`
	URLImage                string    `json:"url_image"`
}

type ClassSchedule struct {
	Repeat struct {
		RepeatAmount   Integer `json:"i_repeat"`  // "2" (for every 2)
		RepeatInterval Integer `json:"id_repeat"` // 7 (for weeks)
	} `json:"a_repeat"`
	StaffIDs          []Integer `json:"a_staff_key"`
	EndDate           Date      `json:"dl_end"`
	StartDate         Date      `json:"dl_start"`
	DayOfWeek         Integer   `json:"i_day"` // 1 is Monday; 7 is Sunday.
	DurationInMinutes Integer   `json:"i_duration"`
	IsCancel          Bool      `json:"is_cancel"`
	ClassID           string    `json:"k_class"`
	ClassPeriodID     string    `json:"k_class_period"`
	LocationID        string    `json:"k_location"`
	Price             Currency  `json:"m_price"`
	TextTimeRange     string    `json:"text_time_range"` // 7:00pm - 9:00pm
	TextTimeStart     string    `json:"text_time_start"` // 7:00pm
}

// ScheduleClassListResponse is the response from "/Wl/Schedule/ClassList/ClassList.json".
type ScheduleClassListResponse struct {
	BaseResponse

	Calendar            StringToAnyMap         `json:"a_calendar"`
	Sessions            []ScheduleClassSession `json:"a_session"`
	IsTimezoneDifferent Bool                   `json:"is_timezone_different"`
	IsVirtualService    Bool                   `json:"is_virtual_service"`
}

type ScheduleClassSession struct {
	StartTime         DateTime `json:"dt_date"`  // This is in UTC.
	TimeString        string   `json:"dt_time"`  // "19:15:00"
	LocalStartTime    DateTime `json:"dtl_date"` // "2024-02-23 19:15:00"
	DayOfWeek         Integer  `json:"i_day"`
	DurationInMinutes Integer  `json:"i_duration"`
	IsCancel          Bool     `json:"is_cancel"` // "0"
	ClassID           string   `json:"k_class"`
	ClassPeriodID     string   `json:"k_class_period"`
	LocationID        string   `json:"k_location"`
	Title             string   `json:"s_title"`
	Timezone          string   `json:"text_timezone"` // "EDT"
	URLBook           string   `json:"url_book"`
	Staff             []string `json:"a_staff"`
	// TODO: "a_virtual_location": []
	HideApplication Bool     `json:"hide_application"`
	IsVirtual       Bool     `json:"is_virtual"`
	ClassTab        []string `json:"a_class_tab"`
}

// ScheduleClassViewResponse is the response from "/Wl/Schedule/ClassView/ClassView.json".
type ScheduleClassViewResponse struct {
	BaseResponse

	// TODO: "a_asset"
	// TODO: "a_class"
	// TODO: "a_location"
	SessionResult []struct {
		// TODO: "a_asset"
		Class struct {
			// TODO: "a_class_tab"
			Image struct {
				Height  Integer `json:"i_height"`
				Width   Integer `json:"i_width"`
				IsEmpty Bool    `json:"is_empty"`
				IsOwn   Bool    `json:"is_own"`
				URL     string  `json:"https://d12lnanyhdwsnh.cloudfront.net/1/kXI.png"`
			} `json:"a_image"`
			// TODO: "a_search_tag"
			// TODO: "a_tag"
			CanBook           Bool     `json:"can_book"`
			GlobalDate        DateTime `json:"dt_date_global"`
			LocalDate         DateTime `json:"dt_date_local"`
			HTMLDenyReason    string   `json:"html_deny_reason"`
			HTMLDescription   string   `json:"html_description"`
			HTMLSpecial       string   `json:"html_special"`
			AgeFrom           *Integer `json:"i_age_from"`
			AgeTo             *Integer `json:"i_age_to"`
			Book              Integer  `json:"i_book"`
			BookActive        Integer  `json:"i_book_active"`
			Capacity          Integer  `json:"i_capacity"`
			Duration          Integer  `json:"i_duration"` // In minutes.
			Wait              Integer  `json:"i_wait"`
			WaitLimit         *Integer `json:"i_wait_limit"`
			WaitSpot          Integer  `json:"i_wait_spot"`
			DenyReasonID      Integer  `json:"id_deny_reason"`
			IsAgePublic       Bool     `json:"is_age_public"`
			IsBook            Bool     `json:"is_book"`
			IsCancel          Bool     `json:"is_cancel"`
			IsEvent           Bool     `json:"is_event"`
			IsPromotionOnly   Bool     `json:"is_promotion_only"`
			VirtualProviderID *Integer `json:"id_virtual_provider"`
			IsVirtual         Bool     `json:"is_virtual"`
			IsWait            Bool     `json:"is_wait"`
			IsWaitList        Bool     `json:"is_wait_list"`
			IsWaitListEnabled Bool     `json:"is_wait_list_enabled"`
			ClassID           string   `json:"k_class"`
			Price             Currency `json:"m_price"`
			HidePrice         Bool     `json:"hide_price"`
			DurationString    string   `json:"s_duration"`
			Title             string   `json:"s_title"`
			RoomText          string   `json:"text_room"`
			TimezoneString    string   `json:"text_timezone"`
			VirtualJoinURL    string   `json:"url_virtual_join"`
		} `json:"a_class"`
		Location struct {
			Latitude   Float  `json:"f_latitude"`
			Longitude  Float  `json:"f_longitude"`
			Rate       Float  `json:"f_rate"`
			LocationID string `json:"k_location"`
			Address    string `json:"s_address"`
			Map        string `json:"s_map"`
			Phone      string `json:"s_phone"`
			Title      string `json:"s_title"`
		} `json:"a_location"`
		Staff []struct {
			StaffID      Integer `json:"k_staff"`
			Family       string  `json:"s_family"`
			Name         string  `json:"s_name"`
			FullName     string  `json:"s_name_full"`
			Position     string  `json:"s_position"`
			UID          string  `json:"uid"`
			XMLBiography string  `json:"xml_biography"`
		} `json:"a_staff"`
		// TODO: "a_visits_required"
		Date          DateTime `json:"dt_date"`
		ClassPeriodID string   `json:"k_class_period"`
	} `json:"a_session_result"`
	// TODO: "a_staff"
	// TODO: "a_visits_required"
}

type TabResponse struct {
	BaseResponse

	Tabs []Tab `json:"a_tab"`
}

type Tab struct {
	IDClassTabObject  Integer  `json:"id_class_tab_object"`
	IDClassTabSystem  Integer  `json:"id_class_tab_system"`
	ClassTabID        *Integer `json:"k_class_tab"`
	ResourceTypeID    *Integer `json:"k_resource_type"`
	ServiceCategoryID *Integer `json:"k_service_category"`
	Title             string   `json:"s_title"`
	ID                Integer  `json:"k_id"`
	Order             Integer  `json:"i_order"`
	URLOrigin         string   `json:"url_origin"`
}

type LocationListResponse struct {
	BaseResponse

	LocationMap map[string]Location `json:"a_location"`
}

type Location struct {
	Latitude       Float   `json:"f_latitude"`
	Longitude      Float   `json:"f_longitude"`
	Order          Integer `json:"i_order"`
	BusinessID     string  `json:"k_business"`
	CountryID      string  `json:"k_country"`
	LocationID     string  `json:"k_location"`
	TimezoneID     string  `json:"k_timezone"`
	RegionID       string  `json:"k_region"`
	URLLogo        string  `json:"url_logo"`
	Shift          Integer `json:"i_shift"`
	Title          string  `json:"s_title"`
	FullAddress    string  `json:"text_address"`
	AddressStreet  string  `json:"text_address_individual"`
	AddressCity    string  `json:"text_city"`
	AddressCountry string  `json:"text_country"`
	AddressPostal  string  `json:"text_postal"` // Zip code in the US.
	AddressRegion  string  `json:"text_region"` // State in the US.
}

type LocationResponse struct {
	BaseResponse

	// TODO: "a_age": [],
	// TODO: "a_amenities": [],
	// TODO: "a_level": [],
	// TODO: "a_logo": {
	/*
	    "is_empty": false,
	    "k_business": "6470",
	    "k_location": "6627",
	    "a_image": {
	      "i_height": 100,
	      "i_height_src": 465,
	      "i_rotate": 0,
	      "i_width": 220,
	      "i_width_src": 1022,
	      "id_type_src": 3,
	      "is-resize": true,
	      "url-view": (string),
	      "url-thumbnail": (string).
	    },
	    "i_height": 100,
	    "i_width": 220,
	    "s_url": (string)
	  },
	*/
	Slides []struct {
		Height     Integer `json:"i_height"`
		Width      Integer `json:"i_width"`
		URLPreview string  `json:"url_preview"`
		URLSlide   string  `json:"url_slide"`
	} `json:"a_slide"`
	// TODO: "a_work": {
	/*
	    "1": [
	      {
	        "s_end": "18:00:00",
	        "s_start": "09:00:00"
	      }
	    ],
	    "2": [
	      {
	        "s_end": "18:00:00",
	        "s_start": "09:00:00"
	      }
	    ],
	    "3": [
	      {
	        "s_end": "18:00:00",
	        "s_start": "09:00:00"
	      }
	    ],
	    "4": [
	      {
	        "s_end": "18:00:00",
	        "s_start": "09:00:00"
	      }
	    ],
	    "5": [
	      {
	        "s_end": "18:00:00",
	        "s_start": "09:00:00"
	      }
	    ]
	  },
	*/
	Latitude               Float   `json:"f_latitude"`
	Longitude              Float   `json:"f_longitude"`
	HTMLDescriptionFull    string  `json:"html_description_full"`
	HTMLDescriptionPreview string  `json:"html_description_preview"`
	IndustryID             Integer `json:"id_industry"`
	IsPhone                Bool    `json:"is_phone"`
	IsTopChoice            Bool    `json:"is_top_choice"`
	BusinessID             string  `json:"k_business"`
	BusinessTypeID         string  `json:"k_business_type"`
	TimezoneID             string  `json:"k_timezone"`
	Address                string  `json:"s_address"`
	Map                    string  `json:"s_map"`
	PhoneNumber            string  `json:"s_phone"`
	Timezone               string  `json:"s_timezone"` // PHP timezone identifier.
	Title                  string  `json:"s_title"`
	AddressStreet          string  `json:"text_address_individual"`
	Alias                  string  `json:"text_alias"`
	BusinessType           string  `json:"text_business_type"`
	AddressCity            string  `json:"text_city"`
	AddressCountry         string  `json:"text_country"`
	Industry               string  `json:"text_industry"`
	EmailAddress           string  `json:"text_mail"`
	AddressPostal          string  `json:"text_postal"`      // Zip code in the US.
	AddressRegion          string  `json:"text_region"`      // State in the US.
	AddressRegionCode      string  `json:"text_region_code"` // State abbreviaion in the US.
	URLFacebook            string  `json:"url_facebook"`
	URLInstagram           string  `json:"url_instagram"`
	URLLinkedIn            string  `json:"url_linkedin"`
	URLMap                 string  `json:"url_map"`
	URLMicrosite           string  `json:"url_microsite"`
	URLSite                string  `json:"url_site"`
	URLTwitter             string  `json:"url_twitter"`
	URLWeb                 string  `json:"url_web"`
	URLYouTube             string  `json:"url_youtube"`
}

type AttendanceListResponse struct {
	ListActive      []*AttendanceListPerson `json:"a_list_active"`
	ListConfirm     []*AttendanceListPerson `json:"a_list_confirm"`
	ListWait        []*AttendanceListPerson `json:"a_list_wait"`
	Capacity        Integer                 `json:"i_capacity"`
	ClientCount     Integer                 `json:"i_client"`
	WaitListLimit   Integer                 `json:"wait_list_limit"`
	IsWaitListLimit Bool                    `json:"is_wait_list_limit"`
	LocationID      string                  `json:"k_location"`
}

type AttendanceListPerson struct {
	Photo struct {
		Login   string  `json:"s_login"`
		Height  Integer `json:"i_height"`
		Width   Integer `json:"i_width"`
		URL     string  `json:"s_url"`
		IsEmpty Bool    `json:"is_empty"`
	} `json:"a_photo"`
	// TODO: "a_wait_confirm": [],
	BookedDate           DateTime `json:"dt_book"`     // In UTC.
	Date                 DateTime `json:"dt_date"`     // In UTC.
	ExpireDate           *Date    `json:"dt_expire"`   // Can be "".
	RegisterDate         DateTime `json:"dt_register"` // Can be "0000-00-00 00:00:00".
	HTMLAge              string   `json:"html_age"`
	HTMLBookedBy         string   `json:"html_book_by"`
	HTMLGenderClass      string   `json:"html_gender_class"`
	HTMLMember           string   `json:"html_member"`
	HTMLTooltipBookedBy  string   `json:"html_tooltip_book_by"`
	Remaining            *Integer `json:"i_left"`
	Total                Integer  `json:"i_total"`
	GenderID             *Integer `json:"id_gender"`
	ProgramID            Integer  `json:"id_program"`
	IDVisit              Integer  `json:"id_visit"` // TODO: Find a better name for this.
	IsAttend             Bool     `json:"is_attend"`
	IsDeposit            Bool     `json:"is_deposit"`
	IsEarly              Bool     `json:"is_early"`
	IsFree               Bool     `json:"is_free"`
	IsHidden             Bool     `json:"is_hidden"`
	PassProspectID       Integer  `json:"id_pass_prospect"`
	IsPenalty            Bool     `json:"is_penalty"`
	IsPending            Bool     `json:"is_pending"`
	IsPromotion          Bool     `json:"is_promotion"`
	IsPromotionChange    *Bool    `json:"is_promotion_change"`
	IsRestrict           Bool     `json:"is_restrict"`
	IsTruancy            Bool     `json:"is_truancy"`
	IsUnpaid             Bool     `json:"is_unpaid"`
	IsVisit              Bool     `json:"is_visit"`
	IsWait               Bool     `json:"is_wait"`
	IsWaitConfirm        Bool     `json:"is_wait_confirm"`
	IsWaitPriority       Bool     `json:"is_wait_priority"`
	LocationID           string   `json:"k_location"`
	LoginPromotionID     *string  `json:"k_login_promotion"`
	VisitID              string   `json:"k_visit"`
	Expire               string   `json:"s_expire"`
	FirstName            string   `json:"s_firstname"` // Deprecated: use "text_firstname" instead.
	LastName             string   `json:"s_lastname"`  // Deprecated: use "text_lastname" instead.
	Login                string   `json:"s_login"`
	EmailAddress         string   `json:"s_mail"`
	Note                 string   `json:"s_note"`
	Phone                string   `json:"s_phone"`
	Promotion            string   `json:"s_promotion"`
	ModeSID              string   `json:"sid_mode"` // For example: "web-backend"
	TextAge              *Integer `json:"text_age"`
	TextExpire           string   `json:"text_expire"`
	TextFirstName        string   `json:"text_firstname"`
	TextIconClass        string   `json:"text_icon_class"`
	TextLastName         string   `json:"text_lastname"`
	TextMember           *string  `json:"text_member"`
	TextPromition        string   `json:"text_promotion"`
	TextVisitStatusClass string   `json:"text_visit_status_class"`
	TextVisitStatusIcon  string   `json:"text_visit_status_icon"`
	UID                  string   `json:"uid"`
	UIDBook              string   `json:"uid_book"`
	URLCancel            string   `json:"url-cancel"`
	URLCancelAdmin       string   `json:"url-cancel-admin"`
	URLLoginView         string   `json:"url-login-view"`
	URLMail              string   `json:"url-mail"`
	URLProfile           string   `json:"url-profile"`
	I                    Integer  `json:"i"` // Deprecated; use "i_order" instead.
	Order                Integer  `json:"i_order"`
	// TODO: "a_resource": [],
	CanProfile Bool `json:"can_profile"`
	// TODO: "a_wearable": [],
	Icon struct {
		ColorBackground string `json:"s_color_background"`
		ColorForeground string `json:"s_color_foreground"`
		Letter          string `json:"s_letter"`
		Shape           string `json:"s_shape"`
		Title           string `json:"s_title"`
		ShapeSID        string `json:"sid_shape"`
	} `json:"icon"`
}
