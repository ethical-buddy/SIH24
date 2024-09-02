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
                "RunLinuxCheck1": widget.NewCheck("Linux Check 1", nil),
                "RunLinuxCheck2": widget.NewCheck("Linux Check 2", nil),
                "RunLinuxCheck3": widget.NewCheck("Linux Check 3", nil),
                // Add more benchmarks by referencing actual function names from internal/benchmarks/linux.go
        }

        // Create a VBox to hold the checkboxes
        checkboxContainer := container.NewVBox()
        for _, check := range benchmarkCheckboxes {
                checkboxContainer.Add(check)
        }

        // Create a button to trigger the dropdown
        benchmarkButton := widget.NewButton("Select Benchmarks", func() {
                // Display the checkboxes as a pop-up
                benchmarkMenu := widget.NewPopUp(checkboxContainer, myWindow.Canvas())
                benchmarkMenu.ShowAtPosition(fyne.NewPos(myWindow.Canvas().Size().Width-200, 50)) // Adjust position as needed
        })

        // // Start Audit button
        // startButton := widget.NewButton("Start Audit", func() {
        //      go func() {
        //              results := ""
        //              // Run selected benchmarks
        //              if benchmarkCheckboxes["RunLinuxCheck1"].Checked {
        //                      results += benchmarks.DisableFreevxfs() // Replace with actual function
        //              }
        //              if benchmarkCheckboxes["RunLinuxCheck2"].Checked {
        //                      results += benchmarks.DisableCramfs() // Replace with actual function
        //              }
        //              if benchmarkCheckboxes["RunLinuxCheck3"].Checked {
        //                      results += benchmarks.CheckLinuxFirewall() // Replace with actual function
        //              }

        //              resultArea.SetText(results)

        //              // Generate report after checks are done
        //              r := report.Report{}
        //              for _, result := range results {
        //                      r.AddResult(string(result)) // Ensure the result is treated as a string
        //              }

        //              err := r.GenerateReport("audit_report.txt")
        //              if err != nil {
        //                      resultArea.SetText("Error generating report: " + err.Error())
        //              } else {
        //                      resultArea.SetText(resultArea.Text + "\nAudit report generated successfully.")
        //              }
        //      }()
        // })
// ... (other imports and setup)

startButton := widget.NewButton("Start Audit", func() {
        go func() {
                var results string

                // Handle DisableFreevxfs
                if benchmarkCheckboxes["DisableFreevxfs"].Checked {
                        result, err := benchmarks.DisableFreevxfs()
                        if err != nil {
                                results += "Error running DisableFreevxfs: " + err.Error() + "\n"
                        } else {
                                results += result + "\n"
                        }
                }

                // Handle DisableCramfs
                if benchmarkCheckboxes["DisableCramfs"].Checked {
                        result, err := benchmarks.DisableCramfs()
                        if err != nil {
                                results += "Error running DisableCramfs: " + err.Error() + "\n"
                        } else {
                                results += result + "\n"
                        }
                }

                // Handle CheckLinuxFirewall
                if benchmarkCheckboxes["CheckLinuxFirewall"].Checked {
                        result, err := benchmarks.CheckLinuxFirewall()
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

// ... (rest of the code)

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
//sorry


