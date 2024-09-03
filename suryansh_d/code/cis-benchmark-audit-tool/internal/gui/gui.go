package gui

import (
        "bytes"
        "cis-benchmark-audit-tool/internal/benchmarks"
        "html/template"
        "io/ioutil"
        "log"
        "os"
        "os/exec"
        "runtime"

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

                        // Save results to HTML file
                        htmlFileName := "audit_report.html"
                        tmpl, _ := template.New("report").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CIS Benchmark Audit Report</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
    <style>
        body {
            font-family: 'Arial', sans-serif;
            background-color: #121212;
            color: #e0e0e0;
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 900px;
            margin: auto;
            padding: 20px;
        }
        h1 {
            color: #bb86fc;
            text-align: center;
            margin-bottom: 20px;
        }
        .benchmark {
            background-color: #1e1e1e;
            border-radius: 8px;
            padding: 15px;
            margin-bottom: 20px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
        }
        .benchmark-title {
            color: #03dac6;
            font-size: 1.4em;
            margin-bottom: 10px;
        }
        .benchmark-output {
            color: #e0e0e0;
            font-size: 1.1em;
        }
        .error {
            color: #cf6679;
            font-size: 1.4em;
            margin-bottom: 10px;
        }
        .error-message {
            color: #ffebee;
            font-size: 1.1em;
        }
        .btn-primary {
            background-color: #bb86fc;
            border-color: #bb86fc;
        }
        .btn-primary:hover {
            background-color: #a070f5;
            border-color: #a070f5;
        }
        .pdf-download {
            text-align: center;
            margin-top: 20px;
        }
        #downloadPdf {
            background-color: #03dac6;
            border: none;
            color: #121212;
            padding: 10px 20px;
            font-size: 1.1em;
            border-radius: 5px;
            cursor: pointer;
            transition: background-color 0.3s;
        }
        #downloadPdf:hover {
            background-color: #018786;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>CIS Benchmark Audit Report</h1>
        <div id="reportContent">
            <!-- Benchmark results will be inserted here -->
            <!-- Example of dynamic content -->
            <div class="benchmark">
                <h3 class="benchmark-title">CheckLinuxFirewall</h3>
                <p class="benchmark-output">Passed: Firewall is active and configured correctly.</p>
            </div>
            <!-- Add more benchmarks as needed -->
        </div>
        <div class="pdf-download">
            <button id="downloadPdf">Download as PDF</button>
        </div>
    </div>

    <script src="https://cdnjs.cloudflare.com/ajax/libs/jspdf/2.4.0/jspdf.umd.min.js"></script>
    <script>
        document.getElementById('downloadPdf').addEventListener('click', () => {
            const { jsPDF } = window.jspdf;
            const doc = new jsPDF();

            doc.html(document.querySelector('#reportContent'), {
                callback: function (doc) {
                    doc.save('audit_report.pdf');
                },
                x: 10,
                y: 10
            });
        });

        // Function to dynamically add benchmark results
        function addBenchmarkResult(title, output, isError = false) {
            const container = document.getElementById('reportContent');
            const benchmarkDiv = document.createElement('div');
            benchmarkDiv.className = 'benchmark';
            
            const titleElem = document.createElement('h3');
            titleElem.className = 'benchmark-title';
            titleElem.textContent = title;
            
            const outputElem = document.createElement('p');
            outputElem.className = isError ? 'error' : 'benchmark-output';
            outputElem.textContent = output;
            
            benchmarkDiv.appendChild(titleElem);
            benchmarkDiv.appendChild(outputElem);
            
            container.appendChild(benchmarkDiv);
        }

        // Example usage of the addBenchmarkResult function
        addBenchmarkResult('CheckLinuxFirewall', 'Passed: Firewall is active and configured correctly.');
        addBenchmarkResult('EnsureNoexecOnTmp', 'Failed: Noexec option is not set on /tmp partition.', true);
    </script>
</body>
</html>
`)
                        file, _ := os.Create(htmlFileName)
                        defer file.Close()
                        tmpl.Execute(file, results)

                        // Update the resultArea widget with the results
                        fyne.CurrentApp().Driver().CanvasForObject(resultArea).Refresh(resultArea)
                        resultArea.SetText(results)
                }()
        })

        // Create a button to open the HTML report
        openReportButton := widget.NewButton("Open Report", func() {
                var cmd *exec.Cmd
                switch runtime.GOOS {
                case "linux":
                        cmd = exec.Command("xdg-open", "audit_report.html")
                case "windows":
                        cmd = exec.Command("cmd", "/c", "start", "audit_report.html")
                case "darwin":
                        cmd = exec.Command("open", "audit_report.html")
                default:
                        log.Printf("Unsupported OS: %s", runtime.GOOS)
                        return
                }
                cmd.Start()
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
                openReportButton, // Add the button to open the HTML report
        )

        // Set the content of the window
        myWindow.SetContent(content)

        // Show the window
        myWindow.ShowAndRun()
}



