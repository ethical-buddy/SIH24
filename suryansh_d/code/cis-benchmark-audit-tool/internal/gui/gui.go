package gui

import (
    "bytes"
    "cis-benchmark-audit-tool/internal/benchmarks"
    "html/template"
    "io/ioutil"
    "log"
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
//        "RunLinuxCheck21": benchmarks.EnsureNodevOnVarLog,
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
            // Create a slice to hold benchmark results
            var benchmarksResults []struct {
                Name    string
                Result  string
                IsError bool
            }

            // Check and run each benchmark function
            for name, check := range benchmarkCheckboxes {
                if check != nil && check.Checked {
                    benchmarkFunc, exists := benchmarkFunctions[name]
                    if exists {
                        result, err := benchmarkFunc()
                        if err != nil {
                            benchmarksResults = append(benchmarksResults, struct {
                                Name    string
                                Result  string
                                IsError bool
                            }{Name: name, Result: "Error running " + name + ": " + err.Error(), IsError: true})
                        } else {
                            benchmarksResults = append(benchmarksResults, struct {
                                Name    string
                                Result  string
                                IsError bool
                            }{Name: name, Result: result, IsError: false})
                        }
                    } else {
                        benchmarksResults = append(benchmarksResults, struct {
                            Name    string
                            Result  string
                            IsError bool
                        }{Name: name, Result: "No function for " + name, IsError: true})
                    }
                }
            }

            // Generate the HTML report
            htmlFileName := "audit_report.html"
            tmpl := template.Must(template.New("report").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>CIS Benchmark Audit Report</title>
    <style>
        body { font-family: Arial, sans-serif; }
        h1 { text-align: center; }
        .benchmark { margin: 20px; padding: 10px; border: 1px solid #ddd; border-radius: 5px; }
        .error { color: red; }
        .pass { color: green; }
    </style>
</head>
<body>
    <h1>CIS Benchmark Audit Report</h1>
    {{range .}}
    <div class="benchmark {{if .IsError}}error{{else}}pass{{end}}">
        <h2>{{.Name}}</h2>
        <p>{{.Result}}</p>
    </div>
    {{end}}
</body>
</html>
            `))
            var htmlContent bytes.Buffer
            if err := tmpl.Execute(&htmlContent, benchmarksResults); err != nil {
                log.Println("Failed to generate HTML report:", err)
                return
            }

            if err := ioutil.WriteFile(htmlFileName, htmlContent.Bytes(), 0644); err != nil {
                log.Println("Failed to write HTML report to file:", err)
                return
            }

            // Open the HTML report in the default web browser
            cmd := exec.Command("xdg-open", htmlFileName)
            if runtime.GOOS == "windows" {
                cmd = exec.Command("cmd", "/c", "start", htmlFileName)
            }
            if err := cmd.Run(); err != nil {
                log.Println("Failed to open HTML report:", err)
            }

            // Update the result area in the GUI
            fyne.CurrentApp().SendNotification(&fyne.Notification{
                Title:   "Audit Completed",
                Content: "The audit has completed and the report is available.",
            })
        }()
    })

    // Layout components
    content := container.NewVBox(
        logo,
        benchmarkButton,
        startButton,
        resultArea,
    )
    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}

