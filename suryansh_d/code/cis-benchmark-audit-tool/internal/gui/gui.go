package gui

import (
//	"bytes" // Added to convert []byte to io.Reader
	"cis-benchmark-audit-tool/internal/benchmarks"
	"cis-benchmark-audit-tool/internal/report"
//	"io/ioutil"
//	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
//	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// func CreateGUI() {
// 	// Create a new Fyne application
// 	myApp := app.New()

// 	// Create a new window
// 	myWindow := myApp.NewWindow("CIS Benchmark Audit Tool")
// 	myWindow.Resize(fyne.NewSize(100, 100)) // Set window size to full screen
// 	myWindow.CenterOnScreen()

// 	// Load and display a logo from the assets directory
// 	imageData, err := ioutil.ReadFile("assets/logo.png")
// 	if err != nil {
// 		log.Fatal("Failed to load logo image:", err)
// 	}
// 	logo := canvas.NewImageFromReader(bytes.NewReader(imageData), "logo.png")
// 	logo.FillMode = canvas.ImageFillOriginal

// 	// Create a text area for displaying audit results
// 	resultArea := widget.NewMultiLineEntry()
// 	resultArea.SetPlaceHolder("Audit results will be displayed here...")
// 	resultArea.Wrapping = fyne.TextWrapWord
// 	resultArea.Disable()

// 	// Start Audit button
// 	startButton := widget.NewButton("Start Audit", func() {
// 		go func() {
// 			results := benchmarks.RunLinuxChecks()
// 			resultArea.SetText(results)

// 			// Generate report after checks are done
// 			r := report.Report{}
// 			for _, result := range results {
// 				r.AddResult(string(result)) // Ensure the result is treated as a string
// 			}

// 			err := r.GenerateReport("audit_report.txt")
// 			if err != nil {
// 				resultArea.SetText("Error generating report: " + err.Error())
// 			} else {
// 				resultArea.SetText(resultArea.Text + "\nAudit report generated successfully.")
// 			}
// 		}()
// 	})

// 	// Create the main container
// 	content := container.NewVBox(
// 		widget.NewLabelWithStyle("CIS Benchmark Audit Tool", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
// 		logo,
// 		resultArea,
// 		startButton,
// 	)

// 	// Set the content of the window
// 	myWindow.SetContent(content)

// 	// Show the window
// 	myWindow.ShowAndRun()
// }

func CreateGUI() {
	// Create a new Fyne application
	myApp := app.New()

	// Create a new window
	myWindow := myApp.NewWindow("CIS BENCHMARK AUDIT TOOL")
	//myWindow.Resize(fyne.NewSize(80, 60)) // Set window size to 800x600
  	myWindow.CenterOnScreen()

	// Load and display a logo from the assets directory
//	imageData, err := ioutil.ReadFile("assets/logo.png")
//	if err != nil {
//		log.Fatal("Failed to load logo image:", err)
//	}
	//logo := canvas.NewImageFromReader(bytes.NewReader(imageData), "logo.png")
	//ogo.FillMode = canvas.ImageFillOriginal

	// Create a text area for displaying audit results
	resultArea := widget.NewMultiLineEntry()
	resultArea.SetPlaceHolder("Audit results will be displayed here...")
	resultArea.Wrapping = fyne.TextWrapWord
	resultArea.Disable()

	// Start Audit button
	startButton := widget.NewButton("Start Audit", func() {
		go func() {
			results := benchmarks.RunLinuxChecks()
			resultArea.SetText(results)

			// Generate report after checks are done
			r := report.Report{}
			for _, result := range results {
				r.AddResult(string(result)) // Ensure the result is treated as a string
			}

			err := r.GenerateReport("audit_report.txt")
			if err != nil {
				resultArea.SetText("Error generating report: " + err.Error())
			} else {
				resultArea.SetText(resultArea.Text + "\nAudit report generated successfully.")
			}
		}()
	})

	// Create the main container
	content := container.NewVBox(
		widget.NewLabelWithStyle("CIS Benchmark Audit Tool", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		//logo,
		resultArea,
		startButton,
	)

	// Set the content of the window
	myWindow.SetContent(content)

	// Show the window
	myWindow.ShowAndRun()
}
