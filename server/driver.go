package TPserver

import (
	"fmt"
	"time"

	fptr10 "github.com/EuginKostomarov/ftpr10"
)

func GotConnectionForGetDeviceState() {
	for {

		if fptr.IsOpened() {

		}
		time.Sleep(50 * time.Millisecond)
	}

}

func InitDriver() {
	var err error
	fptr, err = fptr10.NewSafe()
	if err != nil {
		PrintErrToForm(`Не удаётся найти установленную программу "Тест Драйвера АТОЛ" на ПК.`, err)
	}
	if err := fptr.Open(); err != nil {
		fmt.Println(err)
	}
}

func GotConnectionForGetKassaState() {

	CheckStates()

}

func ReadPerviusataFromKKT() {

	CheckStates()

}

func CheckStates() {

	fptr.SetParam(fptr10.LIBFPTR_PARAM_DATA_TYPE, fptr10.LIBFPTR_DT_SHIFT_STATE)
	fptr.QueryData()

	state := fptr.GetParamInt(fptr10.LIBFPTR_PARAM_SHIFT_STATE)
	if state == 1 {
		Shift.state = true
	} else {
		Shift.state = false
	}

	Shift.updated = time.Now()

}

func GetALastSessionStartFromDriver() error {

	fptr, err := fptr10.NewSafe()
	fmt.Println(fptr.Version())
	if err != nil {
		HeaderWorkSpace.info.Text = fmt.Sprint("Не удаётся подключится к драйверу. Ошибка: ", err)
		HeaderWorkSpace.info.Refresh()
		time.Sleep(5 * time.Second)
		HeaderWorkSpace.info.Text = ""
		HeaderWorkSpace.info.Refresh()
		return err
	}
	if !fptr.IsOpened() {
		fptr.Open()
		if fptr.IsOpened() {
			fmt.Println("Sucsess")
		}
		// 	HeaderWorkSpace.info.Text = fmt.Sprint("Не удаётся подключится к ККТ. Ошибка: ", err)
		// 	HeaderWorkSpace.info.Refresh()
		// 	time.Sleep(5 * time.Second)
		// 	HeaderWorkSpace.info.Text = ""
		// 	HeaderWorkSpace.info.Refresh()
		// 	return err
		// }
		fptr.Close()
	}
	return nil

}
