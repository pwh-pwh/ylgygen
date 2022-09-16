package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/pwh-pwh/ylgygen"
	"os"
	"strconv"
)

func main() {
	os.Setenv("FYNE_FONT", "C:\\Windows\\Fonts\\msyh.ttc")
	a := app.New()
	w := a.NewWindow("羊了个羊")
	titleL := widget.NewLabel("token")
	titleE := widget.NewEntry()
	ctL := widget.NewLabel("次数")
	ctE := widget.NewEntry()
	tiL := widget.NewLabel("通关时间")
	tiE := widget.NewEntry()
	sLa := widget.NewLabel("状态")
	bt := widget.NewButton("刷", func() {
		num, err := strconv.Atoi(ctE.Text)
		if err != nil {
			sLa.Text = "次数填写错误，请填入数字"
		}
		time, err := strconv.Atoi(tiE.Text)
		if err != nil {
			sLa.Text = "时间填写错误，请填入数字"
		}
		sLa.SetText("开始刷")
		ylgygen.BrushScore(num, titleE.Text, time)
		sLa.SetText("刷完")
	})
	cte := container.New(layout.NewFormLayout(), titleL, titleE,
		ctL, ctE,
		tiL, tiE,
		bt, sLa,
	)
	w.SetContent(cte)
	w.ShowAndRun()
}
