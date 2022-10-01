package TPserver

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	fptr10 "github.com/EuginKostomarov/ftpr10"
)

var ErrChan chan (string)

var fptr *fptr10.IFptr

var Connection, ActualStateOfSession bool

var InformationAbotLogined Info

var Authorized Logined

var AuthForm authForm

var App fyne.App

var HeaderWorkSpace WorkSpaceHeader

var Shift ShiftInfo

var WorkSpaceApp fyne.Window

var Kassa Device

type ShiftInfo struct {
	state   bool
	updated time.Time
}

type authForm struct {
	form *widget.Form
	info *widget.Label
	w    fyne.Window
}

type Info struct {
	Userinfo   Userinfo   `toml:"userinfo"`
	Driverinfo Driverinfo `toml:"driverinfo"`
}

type Userinfo struct {
	Login      string    `toml:"login"`
	Password   string    `toml:"password"`
	Auth       bool      `toml:"auth"`
	Authorized time.Time `toml:"authorized"`
	Session    bool      `toml:"session"`
	Activated  time.Time `toml:"activated"`
}

type Driverinfo struct {
	Path string `toml:"path"`
	Com  string `toml:"com"`
	Time string `toml:"time"`
}

type Logined struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"token_type"`
	UserData    struct {
		ID       int    `json:"id"`
		FullName string `json:"fullName"`
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		Email    string `json:"email"`
		Role     string `json:"role"`
		Ability  []struct {
			Subject string `json:"subject"`
			Action  string `json:"action"`
		} `json:"ability"`
		Extras struct {
			ECommerceCartItemsCount int `json:"eCommerceCartItemsCount"`
		} `json:"extras"`
		Contact struct {
			ID           int         `json:"id"`
			FullName     string      `json:"full_name"`
			DeletedAt    interface{} `json:"deleted_at"`
			CreatedAt    time.Time   `json:"created_at"`
			UpdatedAt    time.Time   `json:"updated_at"`
			Phone        string      `json:"phone"`
			Email        string      `json:"email"`
			Gender       string      `json:"gender"`
			DateBirth    interface{} `json:"date_birth"`
			PassNum      interface{} `json:"pass_num"`
			PassDate     interface{} `json:"pass_date"`
			PassIssued   interface{} `json:"pass_issued"`
			PassDivision interface{} `json:"pass_division"`
			Inn          int64       `json:"inn"`
			Snils        interface{} `json:"snils"`
			Address      interface{} `json:"address"`
			Type         string      `json:"type"`
		} `json:"contact"`
	} `json:"userData"`
}

type WorkSpaceHeader struct {
	Timer          *widget.Label
	info           *widget.Label
	Helper         *widget.Label
	DiskImage      *canvas.Image
	PrinterImage   *canvas.Image
	LoginAccordion *widget.Accordion
	ShiftAccordion *widget.Accordion
	DriverInfo     *widget.Accordion
	KassaInfo      *widget.Accordion
}

type Device struct {
	logicalNumber, model, mode, submode, shiftNumber uint
	isFiscalDevice, isFiscalFN,
	isFNPresent, isInvalidFN,
	isCashDrawerOpened, isPaperPresent,
	isPaperNearEnd, isCoverOpened,
	isPrinterConnectionLost, isPrinterError,
	isCutError, isPrinterOverheat, isDeviceBlocked bool
	serialNumber, modelName, firmwareVersion string
	updated                                  time.Time
}
