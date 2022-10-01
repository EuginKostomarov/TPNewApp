package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func InitilizeApp() {
	a := app.New()
	App = a
}

func CreateNewAuthWindow() {

	w := App.NewWindow("Авторизация")
	w.Resize(fyne.NewSize(400, 100))
	MakeAnAuthForm()
	AuthForm.w = w
	w.SetContent(AuthForm.form)
	w.ShowAndRun()

}

func Authorization(l string, p string) (Logined, error) {

	url := fmt.Sprint("https://ticket-place.ru/api/auth/login?email=", l, "&password=", p)
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return Logined{}, errors.New("Не удалось скомпоновать строку запроса")
	}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return Logined{}, errors.New("Не удалось сделать запрос")
	}
	err = ReturnHTTPERR(res)
	if err != nil {
		return Logined{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Logined{}, errors.New("Не удалось прочитать тело ответа")
	}
	var ans Logined
	// fmt.Println(string(body))
	err = json.Unmarshal(body, &ans)
	if err != nil {
		fmt.Println(err)
		return Logined{}, errors.New("Не удалось разобрать тело ответа")
	}
	// fmt.Println(ans.UserData.Contact.Inn)
	return ans, nil
}

func MakeAnAuthForm() {
	// label empty
	label := widget.NewLabel((""))
	log := WhatsATOMLLogin()
	ActualizeCondition()
	loginw := widget.NewEntry()
	if log != "" {
		loginw.Text = log
	}
	passw := widget.NewPasswordEntry()
	loginw.SetPlaceHolder("Email")
	// passw.SetPlaceHolder("")
	login := widget.NewFormItem("Пользователь", loginw)
	password := widget.NewFormItem("Пароль", passw)
	form := widget.NewForm(
		login,
		password,
	)
	form.SubmitText = "Войти"
	form.OnSubmit = func() {
		// fmt.Println(loginw.Text)
		if len(loginw.Text) < 10 {
			label.Text = "Заполните логин"
			label.Refresh()

		} else if len(passw.Text) < 1 {
			label.Text = "Заполните пароль"
			label.Refresh()
		} else {

			u, err := Authorization(loginw.Text, passw.Text)
			if err != nil {
				AuthForm.info.Text = fmt.Sprint(err)
				AuthForm.info.Refresh()
				return
			}
			Authorized = u
			if err != nil {
				fmt.Println(err)
				var r Userinfo
				label.Text = fmt.Sprint(err)
				r.Login = loginw.Text
				r.Password = passw.Text
				r.Auth = false
				SetLoginDataInToml(r)
				ActualizeCondition()
			}
			AuthForm.w.Close()
			CreateWorkspaceWindow()

		}
	}
	AuthForm.form = form
	AuthForm.info = label

}
