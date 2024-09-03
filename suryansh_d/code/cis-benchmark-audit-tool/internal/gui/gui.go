package gui

import (
	"bytes"
	"cis-benchmark-audit-tool/internal/benchmarks"
	"fmt"
	"io/ioutil"
	"log"
//	"os"
	"time"

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

	// Define the benchmark functions and additional information
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
	benchmarkDescriptions := map[string]string{
		"RunLinuxCheck1":  "Checks if the Linux firewall is active and correctly configured.",
		"RunLinuxCheck2":  "Disables the cramfs filesystem to enhance security.",
		"RunLinuxCheck3":  "Disables the freevxfs filesystem to prevent its use.",
		"RunLinuxCheck4":  "Disables the jffs2 filesystem to avoid potential vulnerabilities.",
		"RunLinuxCheck5":  "Disables the hfs filesystem to reduce attack surfaces.",
		"RunLinuxCheck6":  "Disables the squashfs filesystem to secure the system.",
		"RunLinuxCheck7":  "Disables the udf filesystem for enhanced security.",
		"RunLinuxCheck8":  "Ensures that /tmp is mounted on a separate partition.",
		"RunLinuxCheck9":  "Verifies that the nodev option is set on /tmp for security.",
		"RunLinuxCheck10": "Ensures the noexec option is set on /tmp to prevent execution of binaries.",
		"RunLinuxCheck11": "Verifies that the nosuid option is set on /tmp for security.",
		"RunLinuxCheck12": "Ensures that /var is mounted on a separate partition.",
		"RunLinuxCheck13": "Verifies that the nodev option is set on /var for security.",
		"RunLinuxCheck14": "Ensures the nosuid option is set on /var to enhance security.",
		"RunLinuxCheck15": "Ensures that /var/tmp is mounted on a separate partition.",
		"RunLinuxCheck16": "Verifies that the nodev option is set on /var/tmp for security.",
		"RunLinuxCheck17": "Ensures that /var/log is mounted on a separate partition.",
		"RunLinuxCheck18": "Verifies the noexec option is set on /var/log to prevent execution.",
		"RunLinuxCheck19": "Ensures the nosuid option is set on /var/log for enhanced security.",
		"RunLinuxCheck20": "Ensures that /var/log/audit is on a separate partition.",
		"RunLinuxCheck21": "Verifies that the nodev option is set on /var/log.",
		"RunLinuxCheck22": "Ensures the noexec option is set on /var/log/audit.",
		"RunLinuxCheck23": "Verifies the nosuid option is set on /var/log/audit.",
		"RunLinuxCheck24": "Ensures that the nodev option is set on /home.",
		"RunLinuxCheck25": "Verifies that the nosuid option is set on /home.",
		"RunLinuxCheck26": "Ensures the nodev option is set on /dev/shm.",
		"RunLinuxCheck27": "Verifies the noexec option is set on /dev/shm.",
		"RunLinuxCheck28": "Ensures the nosuid option is set on /dev/shm.",
		"RunLinuxCheck29": "Ensures automounting is disabled to prevent unauthorized access.",
		"RunLinuxCheck30": "Verifies that USB storage is disabled for security purposes.",
	}

	runButton := widget.NewButton("Run Selected Benchmarks", func() {
	results := ""
	htmlResults := ""

	// Run selected benchmarks and collect results
	for key, check := range benchmarkCheckboxes {
		if check.Checked {
			// Append the description of the benchmark
			results += fmt.Sprintf("%s:\n", benchmarkDescriptions[key])
			htmlResults += fmt.Sprintf("<h3>%s</h3><p>%s</p>", key, benchmarkDescriptions[key])

			// Run the benchmark function and collect the result
			if benchmarkFunc, ok := benchmarkFunctions[key]; ok {
				result, err := benchmarkFunc()
				if err != nil {
					// Append only the error message, without the word "Error"
					results += fmt.Sprintf("%v\n\n", err)
					htmlResults += fmt.Sprintf("<p style='color: red;'>%v</p>", err)
				} else {
					results += fmt.Sprintf("Result: %s\n\n", result)
					htmlResults += fmt.Sprintf("<p>Result: %s</p>", result)
				}
			}
		}
	}

	// Display the results in the text area
	resultArea.SetText(results)

	// Save the HTML report
	saveHTMLReport(htmlResults)
	// Button to open the HTML report in the default web browser
	openReportButton := widget.NewButton("Open Report", func() {
		openHTMLReport("CIS_Benchmark_Audit_Report.html")
	})

	// Add the open report button to the main container
	content := container.NewVBox(
		widget.NewLabelWithStyle("CIS Benchmark Audit Tool (LINUX)", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logo,
		container.NewHBox(
			layout.NewSpacer(),
			benchmarkButton,
		),
		resultArea,
		startButton,
		openReportButton, // Added the button to open the HTML report
	)
})


	// Create a button to clear the results
	clearButton := widget.NewButton("Clear Results", func() {
		resultArea.SetText("")
	})

	// Arrange all the components in a grid layout
	content := container.New(layout.NewVBoxLayout(),
		logo,
		benchmarkButton,
		resultArea,
		container.NewHBox(runButton, clearButton),
	)

	// Set the window content
	myWindow.SetContent(content)

	// Show and run the window
	myWindow.ShowAndRun()
}

// Function to save HTML report to a file
func saveHTMLReport(reportContent string) {
	// Add HTML boilerplate
	htmlContent := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>CIS Benchmark Audit Report</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 20px;
			}
			h1, h3 {
				color: #2c3e50;
			}
			p {
				font-size: 14px;
				color: #34495e;
			}
		</style>
	</head>
	<body>
		<h1>CIS Benchmark Audit Report</h1>
		<p>Generated on: %s</p>
		%s
	</body>
	</html>`, time.Now().Format("January 2, 2006, 15:04:05"), reportContent)

	// Write the HTML content to a file
	err := ioutil.WriteFile("CIS_Benchmark_Audit_Report.html", []byte(htmlContent), 0644)
	if err != nil {
		log.Fatalf("Failed to save HTML report: %v", err)
	}
	fmt.Println("HTML report saved successfully.")
}
// openHTMLReport opens the HTML report in the default web browser
func openHTMLReport(filename string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", filename)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", filename)
	case "darwin":
		cmd = exec.Command("open", filename)
	default:
		log.Fatalf("Unsupported platform")
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to open report: %v", err)
	}
}

