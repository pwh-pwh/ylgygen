package main

import (
	"fmt"
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
	titleL := widget.NewLabel("uid或者token")
	titleE := widget.NewEntry()
	ctL := widget.NewLabel("次数")
	ctE := widget.NewEntry()
	tiL := widget.NewLabel("通关时间")
	tiE := widget.NewEntry()
	sLa := widget.NewLabel("状态")

	msLab := widget.NewLabel("模式选择")
	flag := true
	cgg := widget.NewSelect([]string{"t模式", "uid模式"}, func(s string) {
		if s == "t模式" {
			flag = false
		} else {
			flag = true
		}
	})

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
		if flag {
			//uid模式
			fmt.Println("uid model")
			ylgygen.BrushScore2(titleE.Text, num, time)
		} else {
			fmt.Println("token model")
			ylgygen.BrushScore(num, titleE.Text, time)
		}

		sLa.SetText("刷完")
	})
	cte := container.New(layout.NewFormLayout(), titleL, titleE,
		ctL, ctE,
		tiL, tiE,
		bt, sLa,
		msLab, cgg,
	)
	w.SetContent(cte)
	w.ShowAndRun()
}
