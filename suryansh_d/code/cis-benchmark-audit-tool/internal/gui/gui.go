package gui

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func CreateGUI() {
    myApp := app.New()
    myWindow := myApp.NewWindow("CIS Benchmark Audit Tool")

    startButton := widget.NewButton("Start Audit", func() {
        // Trigger audit functions here
    })

    myWindow.SetContent(container.NewVBox(
        widget.NewLabel("Welcome to the CIS Benchmark Audit Tool!"),
        startButton,
    ))

    myWindow.ShowAndRun()
}

