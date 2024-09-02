# Automated CIS Benchmark Auditing Tool For Linux And Windows 

## Overview

Maintaining a robust cybersecurity posture is crucial for organizations across various industries. Compliance with industry standards and guidelines, such as those provided by the Center for Internet Security (CIS), is essential for ensuring the security and integrity of IT infrastructure. Manually auditing and ensuring adherence to these benchmarks is time-consuming, error-prone, and resource-intensive. This project aims to address these challenges by providing automated auditing scripts tailored to CIS benchmarks for different operating systems.

## Description

This software solution provides automated auditing capabilities for CIS benchmarks across various operating systems. The tool lists control guidelines as per CIS benchmarks for the following operating systems:

- **Windows**
  - Windows 11 (Enterprise version)
  - Windows 11 (Standalone version)
- **Linux**
  - Red Hat Enterprise Linux (8 and 9)
  - Ubuntu Desktop (20.04 LTS, 22.04 LTS)
  - Ubuntu Server (12.04 LTS and 14.04 LTS)

## Features

- **Fast**: Utilizes Go concurrency to achieve faster and more efficient audits.
- **Automated Auditing**: Efficiently automate the process of auditing system configurations against CIS benchmarks.
- **User-Friendly GUI**: Provides a graphical user interface for easy interaction and report generation.
- **Customizable**: Adaptable to organizational needs and scalable for diverse IT environments.
- **Accurate Reporting**: Reliable and precise in identifying deviations from best practices.
- **Easy Updates**: Facilitates straightforward updates to accommodate changes in CIS benchmarks over time.

## Installation

To set up the Automated CIS Benchmark Auditing Tool, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/ethical-buddy/SIH24.git
   ```
2. Navigate to the project directory:
   ```sh
   cd automated-cis-benchmark-auditing-tool
   ```
3. Install dependencies:
   ```sh
   python go_install.py
   ```

## What This Script Does

This Python script automates the setup and execution of an auditing tool for CIS benchmarks. Here’s a step-by-step overview of its functionality:

1. **Operating System Detection and Go Installation**:
   - Detects your operating system.
   - Installs Go (Golang) if it is not already present on your system.

2. **Run the Main Application**:
   - Executes the `main.go` file, which launches the graphical user interface (GUI) for the auditing tool.

3. **Run Audit**:
   - Within the GUI, initiate the auditing process. The tool will scan your system and compare configurations against the selected CIS benchmarks.

4. **Generate Report**:
   - Upon completion of the audit, the tool will generate a detailed report of the findings, which you can review and save.

## Usage

This script performs system tests to ensure compliance with CIS Benchmarks. It does not modify any system files.

### GUI Options

**Filters:**

- `help`: Show this help message and exit.
- `level {1,2}`: Run tests for the specified CIS benchmark level only.
- `include INCLUDES [INCLUDES ...]`: Specify a space-delimited list of tests to include in the audit.
- `exclude EXCLUDES [EXCLUDES ...]`: Specify a space-delimited list of tests to exclude from the audit.
- `l {DEBUG,INFO,WARNING,CRITICAL}`, `log-level {DEBUG,INFO,WARNING,CRITICAL}`: Set the logging output level.
- `debug`: Enable debug output. This is equivalent to setting the log level to DEBUG.
- `nice`: Lower the CPU priority for test execution. This is the default behavior.
- `no-nice`: Do not lower CPU priority. This may speed up the tests but increase server load. Overrides `--nice`.
- `no-colour`, `no-color`: Disable colored output for STDOUT. Note that output redirected to a file or pipe will not be colored regardless.
- `system-type {server,workstation}`: Specify which test levels to use.
  - `server`: Use "server" levels for tests. This is the default.
  - `workstation`: Use "workstation" levels for tests.
- `outformat {csv,json,psv,text,tsv}`: Define the format for output results.
  - `text`: Output results as plain text (default).
  - `json`: Output results as JSON.
  - `csv`: Output results as comma-separated values.
  - `psv`: Output results as pipe-separated values.
  - `tsv`: Output results as tab-separated values.
- `version`: Print the script version and exit.
- `config CONFIG`: Specify the location of the configuration file to load.

**Example Usage:**





## License

This project is licensed under the [MIT License](LICENSE).

---

Thank you for using the Automated CIS Benchmark Auditing Tool!
