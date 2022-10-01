package TPserver

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	fptr10 "github.com/EuginKostomarov/ftpr10"
)

func StartParseData() {
	InitElements()
	for {

		fptr, err := fptr10.NewSafe()
		// fmt.Println(fptr.Version())
		if err != nil {
			fmt.Println(err)
		}
		if err := fptr.Open(); err != nil {
			Connection = false
			fmt.Println(err)
		} else {
			Connection = true
		}
		InitilizeTOML()
		InitileStateOfKassa()
		InitilezeDeviceState()
		if err := fptr.Close(); err != nil {
			Connection = true
			fmt.Println(err)
		} else {
			Connection = true
		}
		time.Sleep(1 * time.Minute)

	}
}

func InitilezeDeviceState() {
	GotConnectionForGetDeviceState()
}

func InitileStateOfKassa() {
	GotConnectionForGetKassaState()
}

func WhatsATOMLLogin() string {
	InitilizeTOML()
	var r Info
	_, err := toml.DecodeFile(".toml", &r)
	if err != nil {
		fmt.Println(err)
	}
	return r.Userinfo.Login

}

func ReturnStateString(state uint) string {

	if state == fptr10.LIBFPTR_SS_CLOSED {
		return "Смена закрыта"
	} else if state == fptr10.LIBFPTR_SS_OPENED {
		return "Cмена открыта"
	} else {
		return "Cмена истекла, откройте новую"
	}

}

func SetLoginDataInToml(u Userinfo) {
	InitilizeTOML()
	var r Info
	_, err := toml.DecodeFile(".toml", &r)
	if err != nil {
		fmt.Println(err)
	}
	u.Activated = r.Userinfo.Activated
	u.Session = r.Userinfo.Session
	f, err := os.OpenFile(".toml", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	r.Userinfo = u
	if err := toml.NewEncoder(f).Encode(r); err != nil {
		// failed to encode
		fmt.Println("slomalos")
		log.Fatal(err)
	}
	f.Close()

}

func InitilizeTOML() {
	if _, err := os.Stat(".toml"); err != nil {
		os.Create(".toml")
	}
	var i Info
	_, err := toml.DecodeFile(".toml", &i)
	if err != nil {
		panic(err)
	}
	InformationAbotLogined = i
}

func ActualizeCondition() {
	var i Info
	_, err := toml.DecodeFile(".toml", &i)
	if err != nil {
		panic(err)
	}
	InformationAbotLogined = i
}
