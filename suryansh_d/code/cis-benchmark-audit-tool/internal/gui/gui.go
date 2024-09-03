// package gui

// import (
// 	"bytes"
// 	"cis-benchmark-audit-tool/internal/benchmarks"
// 	"io/ioutil"
// 	"log"

//         "fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/layout"
// 	"fyne.io/fyne/v2/widget"
// )
// func CreateGUI() {
// 	// Create a new Fyne application
// 	myApp := app.New()

// 	// Create a new window
// 	myWindow := myApp.NewWindow("CIS Benchmark Audit Tool")
// 	myWindow.Resize(fyne.NewSize(800, 600)) // Set window size to full screen
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

// 	// Create a dropdown with checkboxes for selecting benchmarks
// 	benchmarkCheckboxes := map[string]*widget.Check{
// 		"RunLinuxCheck1":  widget.NewCheck("CheckLinuxFirewall", nil),
// 		"RunLinuxCheck2":  widget.NewCheck("DisableCramfs", nil),
// 		"RunLinuxCheck3":  widget.NewCheck("DisableFreevxfs", nil),
// 		"RunLinuxCheck4":  widget.NewCheck("Disablejffs2", nil),
// 		"RunLinuxCheck5":  widget.NewCheck("DisableHfs", nil),
// 		"RunLinuxCheck6":  widget.NewCheck("DisableSquashfs", nil),
// 		"RunLinuxCheck7":  widget.NewCheck("DisableUdf", nil),
// 		"RunLinuxCheck8":  widget.NewCheck("EnsureTmpIsSeparatePartition", nil),
// 		"RunLinuxCheck9":  widget.NewCheck("EnsureNodevOnTmp ", nil),
// 		"RunLinuxCheck10": widget.NewCheck("EnsureNoexecOnTmp", nil),
// 		"RunLinuxCheck11": widget.NewCheck("EnsureNosuidOnTmp", nil),
// 		"RunLinuxCheck12": widget.NewCheck("EnsureSeparateVarPartition", nil),
// 		"RunLinuxCheck13": widget.NewCheck("EnsureNodevOnVar", nil),
// 		"RunLinuxCheck14": widget.NewCheck("EnsureNosuidOnVar", nil),
// 		"RunLinuxCheck15": widget.NewCheck("EnsureSeparateVarTmpPartition", nil),
// 		"RunLinuxCheck16": widget.NewCheck("EnsureNodevOnVarTmp", nil),
// 		"RunLinuxCheck17": widget.NewCheck("EnsureSeparateVarLogPartition", nil),
// 		"RunLinuxCheck18": widget.NewCheck("EnsureNoexecOnVarLog", nil),
// 		"RunLinuxCheck19": widget.NewCheck("EnsureNosuidOnVarLog", nil),
// 		"RunLinuxCheck20": widget.NewCheck("EnsureSeparateVarLogAuditPartition", nil),
// 		"RunLinuxCheck21": widget.NewCheck("EnsureNodevOnVarLog", nil),
// 		"RunLinuxCheck22": widget.NewCheck("EnsureNoexecOnVarLogAudit", nil),
// 		"RunLinuxCheck23": widget.NewCheck("EnsureNosuidOnVarLogAudit", nil),
// 		"RunLinuxCheck24": widget.NewCheck("EnsureNodevOnHome", nil),
// 		"RunLinuxCheck25": widget.NewCheck("EnsureNosuidOnHome", nil),
// 		"RunLinuxCheck26": widget.NewCheck("EnsureNodevOnDevShm", nil),
// 		"RunLinuxCheck27": widget.NewCheck("EnsureNoexecOnDevShm", nil),
// 		"RunLinuxCheck28": widget.NewCheck("EnsureNosuidOnDevShm", nil),
// 		"RunLinuxCheck29": widget.NewCheck("EnsureAutomountingDisabled", nil),
// 		"RunLinuxCheck30": widget.NewCheck("EnsureUSBStorageDisabled", nil),

// 		// Add more benchmarks by referencing actual function names from internal/benchmarks/linux.go
// 	}

// 	// Create a VBox to hold the checkboxes
// 	checkboxContainer := container.NewVBox()

// 	// Create the "Select All" checkbox
// 	selectAllCheckbox := widget.NewCheck("Select All", func(checked bool) {
// 		for _, check := range benchmarkCheckboxes {
// 			check.SetChecked(checked)
// 		}
// 	})

// 	// Add the "Select All" checkbox to the container
// 	checkboxContainer.Add(selectAllCheckbox)

// 	// Add individual benchmark checkboxes to the container
// 	for _, check := range benchmarkCheckboxes {
// 		checkboxContainer.Add(check)
// 	}

// 	// Create a scrollable container for the checkboxes
// 	scrollableCheckboxContainer := container.NewScroll(checkboxContainer)
// 	scrollableCheckboxContainer.SetMinSize(fyne.NewSize(250, 350)) // Set a minimum size for visibility

// 	// Create a button to trigger the dropdown
// 	benchmarkButton := widget.NewButton("Select Benchmarks", func() {
// 		// Display the scrollable checkboxes as a pop-up
// 		benchmarkMenu := widget.NewPopUp(scrollableCheckboxContainer, myWindow.Canvas())
// 		benchmarkMenu.ShowAtPosition(fyne.NewPos(myWindow.Canvas().Size().Width-230, 40)) // Adjust position as needed
// 	})

// 	// Start Audit button
// 	startButton := widget.NewButton("Start Audit", func() {
// 		go func() {
// 			results := ""

// 			if benchmarkCheckboxes["CheckLinuxFirewall"] != nil && benchmarkCheckboxes["CheckLinuxFirewall"].Checked {
// 				result, err := benchmarks.CheckLinuxFirewall()
// 				if err != nil {
// 					results += "Error running CheckLinuxFirewall: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["DisableCramfs"] != nil && benchmarkCheckboxes["DisableCramfs"].Checked {
// 				result, err := benchmarks.DisableCramfs()
// 				if err != nil {
// 					results += "Error running DisableCramfs: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["DisableFreevxfs"] != nil && benchmarkCheckboxes["DisableFreevxfs"].Checked {
// 				result, err := benchmarks.DisableFreevxfs()
// 				if err != nil {
// 					results += "Error running DisableFreevxfs: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["DisableJffs2"] != nil && benchmarkCheckboxes["DisableJffs2"].Checked {
// 				result, err := benchmarks.DisableJffs2()
// 				if err != nil {
// 					results += "Error running DisableJffs2: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["DisableHfs"] != nil && benchmarkCheckboxes["DisableHfs"].Checked {
// 				result, err := benchmarks.DisableHfs()
// 				if err != nil {
// 					results += "Error running DisableHfs: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["DisableHfsplus"] != nil && benchmarkCheckboxes["DisableHfsplus"].Checked {
// 				result, err := benchmarks.DisableHfsplus()
// 				if err != nil {
// 					results += "Error running DisableHfsplus: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["DisableSquashfs"] != nil && benchmarkCheckboxes["DisableSquashfs"].Checked {
// 				result, err := benchmarks.DisableSquashfs()
// 				if err != nil {
// 					results += "Error running DisableSquashfs: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["DisableUdf"] != nil && benchmarkCheckboxes["DisableUdf"].Checked {
// 				result, err := benchmarks.DisableUdf()
// 				if err != nil {
// 					results += "Error running DisableUdf: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureTmpIsSeparatePartition"] != nil && benchmarkCheckboxes["EnsureTmpIsSeparatePartition"].Checked {
// 				result, err := benchmarks.EnsureTmpIsSeparatePartition()
// 				if err != nil {
// 					results += "Error running EnsureTmpIsSeparatePartition: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNodevOnTmp"] != nil && benchmarkCheckboxes["EnsureNodevOnTmp"].Checked {
// 				result, err := benchmarks.EnsureNodevOnTmp()
// 				if err != nil {
// 					results += "Error running EnsureNodevOnTmp: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNoexecOnTmp"] != nil && benchmarkCheckboxes["EnsureNoexecOnTmp"].Checked {
// 				result, err := benchmarks.EnsureNoexecOnTmp()
// 				if err != nil {
// 					results += "Error running EnsureNoexecOnTmp: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNosuidOnTmp"] != nil && benchmarkCheckboxes["EnsureNosuidOnTmp"].Checked {
// 				result, err := benchmarks.EnsureNosuidOnTmp()
// 				if err != nil {
// 					results += "Error running EnsureNosuidOnTmp: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureSeparateVarPartition"] != nil && benchmarkCheckboxes["EnsureSeparateVarPartition"].Checked {
// 				result, err := benchmarks.EnsureSeparateVarPartition()
// 				if err != nil {
// 					results += "Error running EnsureSeparateVarPartition: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNodevOnVar"] != nil && benchmarkCheckboxes["EnsureNodevOnVar"].Checked {
// 				result, err := benchmarks.EnsureNodevOnVar()
// 				if err != nil {
// 					results += "Error running EnsureNodevOnVar: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNosuidOnVar"] != nil && benchmarkCheckboxes["EnsureNosuidOnVar"].Checked {
// 				result, err := benchmarks.EnsureNosuidOnVar()
// 				if err != nil {
// 					results += "Error running EnsureNosuidOnVar: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureSeparateVarTmpPartition"] != nil && benchmarkCheckboxes["EnsureSeparateVarTmpPartition"].Checked {
// 				result, err := benchmarks.EnsureSeparateVarTmpPartition()
// 				if err != nil {
// 					results += "Error running EnsureSeparateVarTmpPartition: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNodevOnVarTmp"] != nil && benchmarkCheckboxes["EnsureNodevOnVarTmp"].Checked {
// 				result, err := benchmarks.EnsureNodevOnVarTmp()
// 				if err != nil {
// 					results += "Error running EnsureNodevOnVarTmp: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNoexecOnVarTmp"] != nil && benchmarkCheckboxes["EnsureNoexecOnVarTmp"].Checked {
// 				result, err := benchmarks.EnsureNoexecOnVarTmp()
// 				if err != nil {
// 					results += "Error running EnsureNoexecOnVarTmp: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNosuidOnVarTmp"] != nil && benchmarkCheckboxes["EnsureNosuidOnVarTmp"].Checked {
// 				result, err := benchmarks.EnsureNosuidOnVarTmp()
// 				if err != nil {
// 					results += "Error running EnsureNosuidOnVarTmp: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureSeparateVarLogPartition"] != nil && benchmarkCheckboxes["EnsureSeparateVarLogPartition"].Checked {
// 				result, err := benchmarks.EnsureSeparateVarLogPartition()
// 				if err != nil {
// 					results += "Error running EnsureSeparateVarLogPartition: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNoexecOnVarLog"] != nil && benchmarkCheckboxes["EnsureNoexecOnVarLog"].Checked {
// 				result, err := benchmarks.EnsureNoexecOnVarLog()
// 				if err != nil {
// 					results += "Error running EnsureNoexecOnVarLog: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNosuidOnVarLog"] != nil && benchmarkCheckboxes["EnsureNosuidOnVarLog"].Checked {
// 				result, err := benchmarks.EnsureNosuidOnVarLog()
// 				if err != nil {
// 					results += "Error running EnsureNosuidOnVarLog: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureSeparateVarLogAuditPartition"] != nil && benchmarkCheckboxes["EnsureSeparateVarLogAuditPartition"].Checked {
// 				result, err := benchmarks.EnsureSeparateVarLogAuditPartition()
// 				if err != nil {
// 					results += "Error running EnsureSeparateVarLogAuditPartition: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNoexecOnVarLogAudit"] != nil && benchmarkCheckboxes["EnsureNoexecOnVarLogAudit"].Checked {
// 				result, err := benchmarks.EnsureNoexecOnVarLogAudit()
// 				if err != nil {
// 					results += "Error running EnsureNoexecOnVarLogAudit: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNosuidOnVarLogAudit"] != nil && benchmarkCheckboxes["EnsureNosuidOnVarLogAudit"].Checked {
// 				result, err := benchmarks.EnsureNosuidOnVarLogAudit()
// 				if err != nil {
// 					results += "Error running EnsureNosuidOnVarLogAudit: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureSeparateHomePartition"] != nil && benchmarkCheckboxes["EnsureSeparateHomePartition"].Checked {
// 				result, err := benchmarks.EnsureSeparateHomePartition()
// 				if err != nil {
// 					results += "Error running EnsureSeparateHomePartition: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNodevOnHome"] != nil && benchmarkCheckboxes["EnsureNodevOnHome"].Checked {
// 				result, err := benchmarks.EnsureNodevOnHome()
// 				if err != nil {
// 					results += "Error running EnsureNodevOnHome: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNosuidOnHome"] != nil && benchmarkCheckboxes["EnsureNosuidOnHome"].Checked {
// 				result, err := benchmarks.EnsureNosuidOnHome()
// 				if err != nil {
// 					results += "Error running EnsureNosuidOnHome: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNodevOnDevShm"] != nil && benchmarkCheckboxes["EnsureNodevOnDevShm"].Checked {
// 				result, err := benchmarks.EnsureNodevOnDevShm()
// 				if err != nil {
// 					results += "Error running EnsureNodevOnDevShm: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNoexecOnDevShm"] != nil && benchmarkCheckboxes["EnsureNoexecOnDevShm"].Checked {
// 				result, err := benchmarks.EnsureNoexecOnDevShm()
// 				if err != nil {
// 					results += "Error running EnsureNoexecOnDevShm: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureNosuidOnDevShm"] != nil && benchmarkCheckboxes["EnsureNosuidOnDevShm"].Checked {
// 				result, err := benchmarks.EnsureNosuidOnDevShm()
// 				if err != nil {
// 					results += "Error running EnsureNosuidOnDevShm: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureAutomountingDisabled"] != nil && benchmarkCheckboxes["EnsureAutomountingDisabled"].Checked {
// 				result, err := benchmarks.EnsureAutomountingDisabled()
// 				if err != nil {
// 					results += "Error running EnsureAutomountingDisabled: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}

// 			if benchmarkCheckboxes["EnsureUSBStorageDisabled"] != nil && benchmarkCheckboxes["EnsureUSBStorageDisabled"].Checked {
// 				result, err := benchmarks.EnsureUSBStorageDisabled()
// 				if err != nil {
// 					results += "Error running EnsureUSBStorageDisabled: " + err.Error() + "\n"
// 				} else {
// 					results += result + "\n"
// 				}
// 			}
//         // Update the resultArea widget with the results
//         fyne.InvokeOnMainThread(func() {
//             resultArea.SetText(results)
//             })
//           }()
// 	})

// 	// Create the main container
// 	content := container.NewVBox(
// 		widget.NewLabelWithStyle("CIS Benchmark Audit Tool (LINUX)", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
// 		logo,
// 		container.NewHBox(
// 			layout.NewSpacer(),
// 			benchmarkButton,
// 		),
// 		resultArea,
// 		startButton,
// 	)

// 	// Set the content of the window
// 	myWindow.SetContent(content)

// 	// Show the window
// 	myWindow.ShowAndRun()
// }

package gui

import (
	"bytes"
	"cis-benchmark-audit-tool/internal/benchmarks"
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
		// List of benchmarks
		"RunLinuxCheck1":  widget.NewCheck("CheckLinuxFirewall", nil),
		"RunLinuxCheck2":  widget.NewCheck("DisableCramfs", nil),
		"RunLinuxCheck3":  widget.NewCheck("DisableFreevxfs", nil),
		"RunLinuxCheck4":  widget.NewCheck("DisableJffs2", nil),
		"RunLinuxCheck5":  widget.NewCheck("DisableHfs", nil),
		"RunLinuxCheck6":  widget.NewCheck("DisableSquashfs", nil),
		"RunLinuxCheck7":  widget.NewCheck("DisableUdf", nil),
		"RunLinuxCheck8":  widget.NewCheck("EnsureTmpIsSeparatePartition", nil),
		"RunLinuxCheck9":  widget.NewCheck("EnsureNodevOnTmp", nil),
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
	scrollableCheckboxContainer.SetMinSize(fyne.NewSize(250, 350)) // Set a minimum size for visibility

	// Create a button to trigger the dropdown
	benchmarkButton := widget.NewButton("Select Benchmarks", func() {
		// Display the scrollable checkboxes as a pop-up
		benchmarkMenu := widget.NewPopUp(scrollableCheckboxContainer, myWindow.Canvas())
		benchmarkMenu.ShowAtPosition(fyne.NewPos(myWindow.Canvas().Size().Width-230, 40)) // Adjust position as needed
	})
	var benchmarkFunctions = map[string]func() (string, error){
		"RunLinuxCheck1":  benchmarks.CheckLinuxFirewall,
		"RunLinuxCheck2":  benchmarks.DisableCramfs,
		"RunLinuxCheck3":  benchmarks.DisableFreevxfs,
		"RunLinuxCheck4":  benchmarks.DisableJffs2,
		"RunLinuxCheck5":  benchmarks.DisableHfs,
		"RunLinuxCheck6":  benchmarks.DisableSquashfs,
		"RunLinuxCheck7":  benchmarks.DisableUdf,
		"RunLinuxCheck8":  benchmarks.EnsureTmpIsSeparatePartition,
		"RunLinuxCheck9":  benchmarks.EnsureNodevOnTmp,
		"RunLinuxCheck10": benchmarks.EnsureNoexecOnTmp,
		"RunLinuxCheck11": benchmarks.EnsureNosuidOnTmp,
		"RunLinuxCheck12": benchmarks.EnsureSeparateVarPartition,
		"RunLinuxCheck13": benchmarks.EnsureNodevOnVar,
		"RunLinuxCheck14": benchmarks.EnsureNosuidOnVar,
		"RunLinuxCheck15": benchmarks.EnsureSeparateVarTmpPartition,
		"RunLinuxCheck16": benchmarks.EnsureNodevOnVarTmp,
		"RunLinuxCheck17": benchmarks.EnsureSeparateVarLogPartition,
		"RunLinuxCheck18": benchmarks.EnsureNoexecOnVarLog,
		"RunLinuxCheck19": benchmarks.EnsureNosuidOnVarLog,
		"RunLinuxCheck20": benchmarks.EnsureSeparateVarLogAuditPartition,
		"RunLinuxCheck21": benchmarks.EnsureNodevOnVarLog,
		"RunLinuxCheck22": benchmarks.EnsureNoexecOnVarLogAudit,
		"RunLinuxCheck23": benchmarks.EnsureNosuidOnVarLogAudit,
		"RunLinuxCheck24": benchmarks.EnsureNodevOnHome,
		"RunLinuxCheck25": benchmarks.EnsureNosuidOnHome,
		"RunLinuxCheck26": benchmarks.EnsureNodevOnDevShm,
		"RunLinuxCheck27": benchmarks.EnsureNoexecOnDevShm,
		"RunLinuxCheck28": benchmarks.EnsureNosuidOnDevShm,
		"RunLinuxCheck29": benchmarks.EnsureAutomountingDisabled,
		"RunLinuxCheck30": benchmarks.EnsureUSBStorageDisabled,
	}
	// Start Audit button
	startButton := widget.NewButton("Start Audit", func() {
		go func() {
			results := ""

			// Check and run each benchmark function
			for name, check := range benchmarkCheckboxes {
				if check != nil && check.Checked {
					benchmarkFunc, exists := benchmarkFunctions[name]
					if exists {
						result, err := benchmarkFunc()
						if err != nil {
							results += "Error running " + name + ": " + err.Error() + "\n"
						} else {
							results += result + "\n"
						}
					} else {
						results += "No function for " + name + "\n"
					}
				}
			}

			// Update the resultArea widget with the results
			fyne.CurrentApp().Driver().CanvasForObject(resultArea).Refresh(resultArea)
			resultArea.SetText(results)
		}()
	})

	// Create the main container
	content := container.NewVBox(
		widget.NewLabelWithStyle("CIS Benchmark Audit Tool (LINUX)", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
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

