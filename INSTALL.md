# How to install Ecla

## Download a release

You can download the Ecla interpreter as a prebuilt standalone binary for your system from the latest release in https://github.com/Eclalang/Ecla/releases

Rename the executable file to `ecla` or `ecla.exe` depending on your os, and add the file path to your PATH.

Restart your terminal to apply the changes.

## Build from sources

Alternatively you can build the Ecla interpreter from the sources

### Install Requirements

Install Go 1.21.1 or later.

### Windows:

- Download the Go installer for Windows from the official Go website: https://go.dev/dl/

- Open the installation file and follow the instructions to install Go.

- After the installation, open the Windows command prompt (cmd) and type "go version". This should display the version of Go you just installed.

### Linux:
- Use the following command to install Go:

```bash
wget https://dl.google.com/go/go1.21.1.linux-amd64.tar.gz
sudo tar -xvf go1.21.1.linux-amd64.tar.gz
sudo mv go /usr/local
```

- Add the following line to your ~/.bashrc file for adding Go to the PATH environment variable:

```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

- Save and close the file. Then, reload your terminal and type "go version". This should display the version of Go you just installed.

### Build the Ecla Interpreter

- Clone the repository:

```bash
git clone https://github.com/Eclalang/Ecla.git
cd Ecla
```

- Execute the following command to create the executable file:

```bash
go build -o .
```

> If you want you can also export the executable file to your environment variables.

## You are ready to use Ecla!
