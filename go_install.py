import platform
import subprocess
import sys
import os

def run_command(command, shell=True):
    result = subprocess.run(command, shell=shell, text=True, capture_output=True)
    print(result.stdout)
    if result.returncode != 0:
        print(f"Error occurred while running command: {result.stderr}", file=sys.stderr)
        sys.exit(result.returncode)

platform_name = platform.system()

if platform_name == "Linux":
    go_tarball_url = "https://go.dev/dl/go1.23.0.linux-amd64.tar.gz"
    go_tarball = "go1.23.0.linux-amd64.tar.gz"    
    run_command("sudo rm -rf /usr/local/go")    
    run_command(f"wget {go_tarball_url} -O {go_tarball}")    
    run_command(f"sudo tar -C /usr/local -xf {go_tarball}")    
    run_command("echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile", shell=True)
    run_command("echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc", shell=True)
    run_command("echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc", shell=True)    
    run_command("source ~/.profile", shell=True)
    run_command("source ~/.bashrc", shell=True)
    run_command("source ~/.zshrc", shell=True)    
    run_command("go version")

    print("Go installed successfully on Linux.")

elif platform_name == "Windows":
    go_installer_url = "https://go.dev/dl/go1.23.0.windows-amd64.msi"
    go_installer = "go1.23.0.windows-amd64.msi"
    
    run_command(f"wget {go_installer_url} -O {go_installer}")    
    run_command(f"msiexec /i {go_installer} /quiet /norestart")    
    go_path = r"C:\Go\bin"
    run_command(f"setx PATH \"%PATH%;{go_path}\"", shell=True)    
    run_command("go version")

    print("Go installed successfully on Windows.")

elif platform_name == "Darwin":  # "Darwin" is the platform name for macOS
    go_tarball_url = "https://golang.org/dl/go1.23.0.darwin-amd64.tar.gz"
    go_tarball = "go1.23.0.darwin-amd64.tar.gz"    
    run_command("sudo rm -rf /usr/local/go")    
    run_command(f"wget {go_tarball_url} -O {go_tarball}")
    
    run_command(f"sudo tar -C /usr/local -xf {go_tarball}")    
    run_command("echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc", shell=True)
    run_command("source ~/.zshrc", shell=True)
    
    run_command("go version")
    print("Go installed successfully on macOS.")

else:
    print("Unsupported platform.")
    sys.exit(1)

# Change to the directory containing this script
script_dir = os.path.dirname(os.path.abspath(__file__))
os.chdir(script_dir)

# Execute the main.go file if it exists
if os.path.isfile("main.go"):
    run_command("go run main.go")
else:
    print("main.go file not found.")
