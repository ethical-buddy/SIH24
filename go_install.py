import platform
import os
import subprocess

platform_name = platform.system()

if platform_name == "Linux":
    os.system("sudo rm -rf /usr/local/go")
    os.system("sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz")
    os.system("echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile")
    os.system("source ~/.profile")
    os.system("go version")

    print("Go installed successfully on Linux.")

elif platform_name == "Windows":
    go_installer_url = "https://golang.org/dl/go1.23.0.windows-amd64.msi"
    os.system(f"curl -O {go_installer_url}")
    os.system("msiexec /i go1.23.0.windows-amd64.msi /quiet /norestart")
    go_path = r"C:\Go\bin"
    os.system(f"setx PATH \"%PATH%;{go_path}\"")
    os.system("go version")

    print("Go installed successfully on Windows.")

elif platform_name == "Darwin":  # "Darwin" is the platform name for macOS
    os.system("sudo rm -rf /usr/local/go")
    os.system("curl -O https://golang.org/dl/go1.23.0.darwin-amd64.tar.gz")
    os.system("sudo tar -C /usr/local -xzf go1.23.0.darwin-amd64.tar.gz")
    os.system("echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc")
    os.system("source ~/.zshrc")
    os.system("go version")

    print("Go installed successfully on macOS.")

else:
    print("Unsupported platform.")
