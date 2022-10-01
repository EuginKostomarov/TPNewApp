package TPclient

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PrintErrToForm(our, err any) {
	HeaderWorkSpace.info.Text = fmt.Sprint("Ошибка: ", err)
	HeaderWorkSpace.info.Refresh()
	if fmt.Sprint(our) != "" {
		HeaderWorkSpace.Helper.Text = fmt.Sprint("Подсказка: ", fmt.Sprint(our))
		HeaderWorkSpace.info.Refresh()
	}
	time.Sleep(5 * time.Second)
	HeaderWorkSpace.info.Text = ""
	HeaderWorkSpace.info.Refresh()
}

func CreateWorkspaceWindow() {

	WorkSpaceApp = App.NewWindow("Workspace")
	WorkSpaceApp.Resize(fyne.NewSize(700, 700))

	InitSVGStatus()

	AddHeaderToWorkSpace()
	WorkSpaceApp.SetContent(
		container.NewVBox( // Вертикальный контейнер всего приложения
			container.NewHBox( // 1ая строка общая панель.
				container.NewHBox( // Верхняя горизонтальная панель состояния программы (Левая часть)
					HeaderWorkSpace.info,
					HeaderWorkSpace.Timer,
				),
				container.NewHBox( // Верхняя горизонтальная панель состояния программы (Правая часть)
					HeaderWorkSpace.DiskImage,
					HeaderWorkSpace.PrinterImage),
				container.NewHBox( // 2ая сверху горизонтальная панель состояния авторизации
					HeaderWorkSpace.LoginAccordion,
					HeaderWorkSpace.ShiftAccordion,
				),
			)))

	WorkSpaceApp.Show()
	// ReadPerviusataFromKKT()

}

func InitSVGStatus() {

	disk := canvas.NewImageFromFile("diskoff.svg")
	disk.FillMode = canvas.ImageFillOriginal
	disk.Resize(fyne.NewSize(50, 20))
	HeaderWorkSpace.DiskImage = disk

	printer := canvas.NewImageFromFile("printeroff.svg")
	printer.FillMode = canvas.ImageFillOriginal
	printer.Resize(fyne.NewSize(50, 20))
	HeaderWorkSpace.PrinterImage = printer
}

func PrintErr(s string) {

}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Время: 03:04:05")
	clock.SetText(formatted)
}

func CreateALoginForm() {
	HeaderWorkSpace.info = widget.NewLabel("")
	HeaderWorkSpace.Timer = MakeATimer()

	LoginItem := widget.NewAccordionItem(fmt.Sprint("Пользователь: ", Authorized.UserData.FullName),
		container.NewVBox(widget.NewLabel(fmt.Sprint("ФИО: ", Authorized.UserData.Username)),
			widget.NewLabel(fmt.Sprint("Почта: ", Authorized.UserData.Email)),
			widget.NewLabel(fmt.Sprint("Роль: ", Authorized.UserData.Role)),
		))

	HeaderWorkSpace.LoginAccordion = widget.NewAccordion(LoginItem)
}

func CreateAShiftForm() {

	open := widget.NewButton("Открыть сессию", func() {
		if ActualStateOfSession {
			OpenAWindowWithQuestionAboutOpeningSession()
		} else {
			HeaderWorkSpace.info.Text = "Смена в данный момент открыта"
			HeaderWorkSpace.info.Refresh()
			time.Sleep(5 * time.Second)
			HeaderWorkSpace.info.Text = ""
			HeaderWorkSpace.info.Refresh()
		}
	})
	open.Importance = widget.HighImportance

	close := widget.NewButton("Закрыть сессию", func() {
		if ActualStateOfSession {
			OpenAWindowWithQuestionAboutSession()
		} else {
			HeaderWorkSpace.info.Text = "Смена в данный момент закрыта"
			HeaderWorkSpace.info.Refresh()
			time.Sleep(5 * time.Second)
			HeaderWorkSpace.info.Text = ""
			HeaderWorkSpace.info.Refresh()
		}
	})
	ShiftItem := widget.NewAccordionItem(fmt.Sprint("Смена: ", Shift.state),
		container.NewVBox(
			container.NewHBox(open, close),
			container.NewHBox(widget.NewLabel(fmt.Sprint("Обновлено: ", Shift.updated.Format("03:04:05")))),
		))

	HeaderWorkSpace.ShiftAccordion = widget.NewAccordion(ShiftItem)
}

func CreateADeviceForm() {

}

func AddHeaderToWorkSpace() {

	CreateALoginForm()

	CreateAShiftForm()

	CreateADeviceForm()

	// HeaderWorkSpace.back.Importance = widget.HighImportance
	// HeaderWorkSpace.back.Resize(fyne.NewSize(150, 30))
	// HeaderWorkSpace.back.Move(fyne.NewPos(250, 30))

}

func OpenAWindowWithQuestionAboutOpeningSession() {
	w := App.NewWindow("Открытие смену")
	w.Resize(fyne.NewSize(400, 100))
	label := widget.NewLabel(("Вы точно хотите открыть смену?"))
	yes := widget.NewButton("Да", func() {
		ActualStateOfSession = !ActualStateOfSession
		// HeaderWorkSpace.start.Text = time.Now().Format("Смена открыта: 03:04:05")
		// HeaderWorkSpace.start.Refresh()
	})
	no := widget.NewButton("Нет", func() { w.Close() })
	no.Importance = widget.HighImportance
	w.SetContent(container.NewVBox(label, container.NewHBox(yes, no)))
}

func OpenAWindowWithQuestionAboutSession() {
	w := App.NewWindow("Закрытие смену")
	w.Resize(fyne.NewSize(400, 100))
	label := widget.NewLabel(("Вы точно хотите закрыть смену?"))
	yes := widget.NewButton("Да", func() { ActualStateOfSession = !ActualStateOfSession })
	no := widget.NewButton("Нет", func() { w.Close() })
	no.Importance = widget.HighImportance
	w.SetContent(container.NewVBox(label, container.NewHBox(yes, no)))
}

func MakeATimer() *widget.Label {
	clock := widget.NewLabel("")
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	return clock
}
