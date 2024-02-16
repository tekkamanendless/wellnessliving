package wellnessliving

type ADateWeekSID int

const (
	ADateWeekSIDFriday    ADateWeekSID = 5
	ADateWeekSIDMonday    ADateWeekSID = 1
	ADateWeekSIDSaturday  ADateWeekSID = 6
	ADateWeekSIDSunday    ADateWeekSID = 7
	ADateWeekSIDThursday  ADateWeekSID = 4
	ADateWeekSIDTuesday   ADateWeekSID = 2
	ADateWeekSIDWednesday ADateWeekSID = 3
)

type ADurationSID int

const (
	ADurationSIDDay    ADurationSID = 4
	ADurationSIDHour   ADurationSID = 3
	ADurationSIDMinute ADurationSID = 2
	ADurationSIDMonth  ADurationSID = 5
	ADurationSIDSecond ADurationSID = 1
	ADurationSIDWeek   ADurationSID = 7
	ADurationSIDWeek4  ADurationSID = 8
	ADurationSIDYear   ADurationSID = 6
)

type AFlagSID int

const (
	AFlagSIDAll AFlagSID = 1
	AFlagSIDOff AFlagSID = 2
	AFlagSIDOn  AFlagSID = 3
)

type AGenderSID int

const (
	AGenderSIDFemale    AGenderSID = 2
	AGenderSIDMale      AGenderSID = 1
	AGenderSIDUndefined AGenderSID = 3
)

type CurrencySID int

const (
	CurrencySIDAED CurrencySID = 11
	CurrencySIDAUD CurrencySID = 6
	CurrencySIDCAD CurrencySID = 4
	CurrencySIDEGP CurrencySID = 8
	CurrencySIDEUR CurrencySID = 13
	CurrencySIDGBP CurrencySID = 3
	CurrencySIDKYD CurrencySID = 5
	CurrencySIDNZD CurrencySID = 10
	CurrencySIDPHP CurrencySID = 12
	CurrencySIDUSD CurrencySID = 1
	CurrencySIDZAR CurrencySID = 7
)

type ProjectSID int

const (
	ProjectSIDWellnessLiving ProjectSID = 4
)

type RegionSID int

const (
	RegionSIDAPSoutheast2 RegionSID = 2
	RegionSIDUSEast1      RegionSID = 1
)

type SaleSID int

const (
	SaleSIDAppointment        SaleSID = 8  // Single appointment reservation.
	SaleSIDAppointmentDeposit SaleSID = 11 // Single appointment deposit reservation.
	SaleSIDAppointmentTip     SaleSID = 12 // Tips for the appointment.
	SaleSIDClassPeriod        SaleSID = 6  // Single class visit.
	SaleSIDCoupon             SaleSID = 7  // Gift card.
	SaleSIDEnrollment         SaleSID = 3  // Enrollments. Classes where flag event is <tt>true</tt>.
	SaleSIDPackage            SaleSID = 5  // Promotions with program {@link WlProgramSid::PACKAGE}.
	SaleSIDProduct            SaleSID = 4  // Products: water, t-shirts, etc.
	SaleSIDPromotionClass     SaleSID = 1  // Promotions with program category {@link WlProgramCategorySid::CLASS} and {@link WlProgramCategorySid::VISIT}.
	SaleSIDPromotionResource  SaleSID = 9  // Promotions with program category {@link WlProgramCategorySid::RESOURCE}.
	SaleSIDPromotionService   SaleSID = 2  // Promotions with program category {@link WlProgramCategorySid::SERVICE}.
	SaleSIDQuickBuy           SaleSID = 10 // Products: water, t-shirts, etc. That available for quick buy.
)

type ServiceSID int

const (
	ServiceSIDAppointment ServiceSID = 1
	ServiceSIDClass       ServiceSID = 2
	ServiceSIDEnrollment  ServiceSID = 3
	ServiceSIDResource    ServiceSID = 5
	ServiceSIDVisit       ServiceSID = 4
)

type YesNoSID int

const (
	YesNoSIDNo  YesNoSID = 2
	YesNoSIDYes YesNoSID = 1
)
