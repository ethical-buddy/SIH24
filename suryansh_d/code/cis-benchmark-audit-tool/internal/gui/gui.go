package gui

import (
	"bytes"
	"cis-benchmark-audit-tool/internal/benchmarks"
	"cis-benchmark-audit-tool/internal/report"
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateGUI() {
	// Create a new Fyne application
	myApp := app.New()

	// Create a new window
	myWindow := myApp.NewWindow("CIS Benchmark Audit Tool")
	myWindow.Resize(fyne.NewSize(800, 600)) // Set window size to full screen
	myWindow.CenterOnScreen()

	// Load and display a logo from the assets directory
	imageData, err := ioutil.ReadFile("assets/logo.png")
	if err != nil {
		log.Fatal("Failed to load logo image:", err)
	}
	logo := canvas.NewImageFromReader(bytes.NewReader(imageData), "logo.png")
	logo.FillMode = canvas.ImageFillOriginal

	// Create a text area for displaying audit results
	resultArea := widget.NewMultiLineEntry()
	resultArea.SetPlaceHolder("Audit results will be displayed here...")
	resultArea.Wrapping = fyne.TextWrapWord
	resultArea.Disable()

	// Create a dropdown with checkboxes for selecting benchmarks
	benchmarkCheckboxes := map[string]*widget.Check{
		"RunLinuxCheck1":  widget.NewCheck("CheckLinuxFirewall", nil),
		"RunLinuxCheck2":  widget.NewCheck("DisableCramfs", nil),
		"RunLinuxCheck3":  widget.NewCheck("DisableFreevxfs", nil),
		"RunLinuxCheck4":  widget.NewCheck("Disablejffs2", nil),
		"RunLinuxCheck5":  widget.NewCheck("DisableHfs", nil),
		"RunLinuxCheck6":  widget.NewCheck("DisableSquashfs", nil),
		"RunLinuxCheck7":  widget.NewCheck("DisableUdf", nil),
		"RunLinuxCheck8":  widget.NewCheck("EnsureTmpIsSeparatePartition", nil),
		"RunLinuxCheck9":  widget.NewCheck("EnsureNodevOnTmp ", nil),
		"RunLinuxCheck10": widget.NewCheck("EnsureNoexecOnTmp", nil),
		"RunLinuxCheck11": widget.NewCheck("EnsureNosuidOnTmp", nil),
		"RunLinuxCheck12": widget.NewCheck("EnsureSeparateVarPartition", nil),
		"RunLinuxCheck13": widget.NewCheck("EnsureNodevOnVar", nil),
		"RunLinuxCheck14": widget.NewCheck("EnsureNosuidOnVar", nil),
		"RunLinuxCheck15": widget.NewCheck("EnsureSeparateVarTmpPartition", nil),
		"RunLinuxCheck16": widget.NewCheck("EnsureNodevOnVarTmp", nil),
		"RunLinuxCheck17": widget.NewCheck("EnsureSeparateVarLogPartition", nil),
		"RunLinuxCheck18": widget.NewCheck("EnsureNoexecOnVarLog", nil),
		"RunLinuxCheck19": widget.NewCheck("EnsureNosuidOnVarLog", nil),
		"RunLinuxCheck20": widget.NewCheck("EnsureSeparateVarLogAuditPartition", nil),
		"RunLinuxCheck21": widget.NewCheck("EnsureNodevOnVarLog", nil),
		"RunLinuxCheck22": widget.NewCheck("EnsureNoexecOnVarLogAudit", nil),
		"RunLinuxCheck23": widget.NewCheck("EnsureNosuidOnVarLogAudit", nil),
		"RunLinuxCheck24": widget.NewCheck("EnsureNodevOnHome", nil),
		"RunLinuxCheck25": widget.NewCheck("EnsureNosuidOnHome", nil),
		"RunLinuxCheck26": widget.NewCheck("EnsureNodevOnDevShm", nil),
		"RunLinuxCheck27": widget.NewCheck("EnsureNoexecOnDevShm", nil),
		"RunLinuxCheck28": widget.NewCheck("EnsureNosuidOnDevShm", nil),
		"RunLinuxCheck29": widget.NewCheck("EnsureAutomountingDisabled", nil),
		"RunLinuxCheck30": widget.NewCheck("EnsureUSBStorageDisabled", nil),

		// Add more benchmarks by referencing actual function names from internal/benchmarks/linux.go
	}

	// Create a VBox to hold the checkboxes
	checkboxContainer := container.NewVBox()

	// Create the "Select All" checkbox
	selectAllCheckbox := widget.NewCheck("Select All", func(checked bool) {
		for _, check := range benchmarkCheckboxes {
			check.SetChecked(checked)
		}
	})

	// Add the "Select All" checkbox to the container
	checkboxContainer.Add(selectAllCheckbox)

	// Add individual benchmark checkboxes to the container
	for _, check := range benchmarkCheckboxes {
		checkboxContainer.Add(check)
	}

	// Create a scrollable container for the checkboxes
	scrollableCheckboxContainer := container.NewScroll(checkboxContainer)
	scrollableCheckboxContainer.SetMinSize(fyne.NewSize(300, 350)) // Set a minimum size for visibility

	// Create a button to trigger the dropdown
	benchmarkButton := widget.NewButton("Select Benchmarks", func() {
		// Display the scrollable checkboxes as a pop-up
		benchmarkMenu := widget.NewPopUp(scrollableCheckboxContainer, myWindow.Canvas())
		benchmarkMenu.ShowAtPosition(fyne.NewPos(myWindow.Canvas().Size().Width-220, 50)) // Adjust position as needed
	})

	// Start Audit button
	startButton := widget.NewButton("Start Audit", func() {
		go func() {
			results := ""
			// Run selected benchmarks
			if benchmarkCheckboxes["RunLinuxCheck1"].Checked {
				result, err := benchmarks.DisableFreevxfs() // Replace with actual function
				if err != nil {
					results += "Error running DisableFreevxfs: " + err.Error() + "\n"
				} else {
					results += result + "\n"
				}
			}
			if benchmarkCheckboxes["RunLinuxCheck2"].Checked {
				result, err := benchmarks.DisableCramfs() // Replace with actual function
				if err != nil {
					results += "Error running DisableCramfs: " + err.Error() + "\n"
				} else {
					results += result + "\n"
				}
			}
			if benchmarkCheckboxes["RunLinuxCheck3"].Checked {
				result, err := benchmarks.CheckLinuxFirewall() // Replace with actual function
				if err != nil {
					results += "Error running CheckLinuxFirewall: " + err.Error() + "\n"
				} else {
					results += result + "\n"
				}
			}

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
		logo,
		container.NewHBox(
			layout.NewSpacer(),
			benchmarkButton,
		),
		resultArea,
		startButton,
	)

	// Set the content of the window
	myWindow.SetContent(content)

	// Show the window
	myWindow.ShowAndRun()
}

