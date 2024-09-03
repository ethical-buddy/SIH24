package gui

import (
    "bytes"
    "cis-benchmark-audit-tool/internal/benchmarks"
    "html/template"
    "io/ioutil"
    "log"
//    "os"
    "os/exec"
    "runtime"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func CreateGUI() {
    // Create a new Fyne application
    myApp := app.New()

    // Create a new window
    myWindow := myApp.NewWindow("CIS Benchmark Audit Tool")
    myWindow.Resize(fyne.NewSize(800, 600)) // Set window size
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
        // Benchmark functions
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

    // Create a mapping from function identifiers to descriptive names
    functionNameMapping := map[string]string{
        "RunLinuxCheck1":  "Check Linux Firewall",
        "RunLinuxCheck2":  "Disable Cramfs",
        "RunLinuxCheck3":  "Disable Freevxfs",
        "RunLinuxCheck4":  "Disable Jffs2",
        "RunLinuxCheck5":  "Disable Hfs",
        "RunLinuxCheck6":  "Disable Squashfs",
        "RunLinuxCheck7":  "Disable Udf",
        "RunLinuxCheck8":  "Ensure /tmp is a Separate Partition",
        "RunLinuxCheck9":  "Ensure nodev option on /tmp",
        "RunLinuxCheck10": "Ensure noexec option on /tmp",
        "RunLinuxCheck11": "Ensure nosuid option on /tmp",
        "RunLinuxCheck12": "Ensure /var is a Separate Partition",
        "RunLinuxCheck13": "Ensure nodev option on /var",
        "RunLinuxCheck14": "Ensure nosuid option on /var",
        "RunLinuxCheck15": "Ensure /var/tmp is a Separate Partition",
        "RunLinuxCheck16": "Ensure nodev option on /var/tmp",
        "RunLinuxCheck17": "Ensure /var/log is a Separate Partition",
        "RunLinuxCheck18": "Ensure noexec option on /var/log",
        "RunLinuxCheck19": "Ensure nosuid option on /var/log",
        "RunLinuxCheck20": "Ensure /var/log/audit is a Separate Partition",
        "RunLinuxCheck22": "Ensure noexec option on /var/log/audit",
        "RunLinuxCheck23": "Ensure nosuid option on /var/log/audit",
        "RunLinuxCheck24": "Ensure nodev option on /home",
        "RunLinuxCheck25": "Ensure nosuid option on /home",
        "RunLinuxCheck26": "Ensure nodev option on /dev/shm",
        "RunLinuxCheck27": "Ensure noexec option on /dev/shm",
        "RunLinuxCheck28": "Ensure nosuid option on /dev/shm",
        "RunLinuxCheck29": "Ensure automounting is Disabled",
        "RunLinuxCheck30": "Ensure USB Storage is Disabled",
    }

    // Create a button to start the audit
    startButton := widget.NewButton("Start Audit", func() {
        go func() {
            var results string
            for id, check := range benchmarkCheckboxes {
                if check.Checked {
                    if benchmarkFunc, exists := benchmarkFunctions[id]; exists {
                        result, err := benchmarkFunc()
                        if err != nil {
                            results += functionNameMapping[id] + " failed: " + err.Error() + "\n"
                        } else {
                            results += functionNameMapping[id] + " passed: " + result + "\n"
                        }
                    }
                }
            }

            // Generate the HTML report
            htmlFileName := "audit_report.html"
            tmpl := template.Must(template.New("report").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CIS Benchmark Audit Report</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
        }
        .container {
            max-width: 800px;
            margin: auto;
        }
        .report-header {
            text-align: center;
        }
        .report-content {
            font-size: 16px;
            white-space: pre-wrap; /* Preserve line breaks and spaces */
        }
        .text-center {
            text-align: center;
            margin-top: 20px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="report-header">
            <h1>CIS Benchmark Audit Report</h1>
        </div>
        <div class="report-content">
            {{ .Results }}
        </div>
        <div class="text-center">
            <a href="audit_report.pdf" class="btn btn-primary" download>Download PDF</a>
        </div>
    </div>
</body>
</html>
`))
            var buf bytes.Buffer
            data := struct {
                Results string
            }{
                Results: results,
            }
            err = tmpl.Execute(&buf, data)
            if err != nil {
                log.Println("Failed to generate HTML report:", err)
                return
            }

            err = ioutil.WriteFile(htmlFileName, buf.Bytes(), 0644)
            if err != nil {
                log.Println("Failed to save HTML report:", err)
                return
            }

            resultArea.SetText(results)

            // Add functionality to convert HTML to PDF
            if runtime.GOOS == "linux" {
                err = exec.Command("wkhtmltopdf", htmlFileName, "audit_report.pdf").Run()
                if err != nil {
                    log.Println("Failed to convert HTML to PDF:", err)
                }
            } else {
                log.Println("PDF conversion is only supported on Linux.")
            }
        }()
    })

    // Create a button to open the file
    openFileButton := widget.NewButton("Open Report", func() {
        if err := exec.Command("xdg-open", "audit_report.html").Run(); err != nil {
            log.Println("Failed to open HTML report:", err)
        }
    })

    // Create a container for the buttons
    buttonsContainer := container.NewHBox(startButton, openFileButton)

    // Create the main layout
    mainContent := container.NewVBox(
        logo,
        benchmarkButton,
        buttonsContainer,
        resultArea,
    )

    myWindow.SetContent(mainContent)
    myWindow.ShowAndRun()
}

