# Building instructions for Windows
This document will guide you through the process of building PufferStarter on your Windows system.

### Prerequisites
This guide requires you to have Administrator privileges on your user account (for installing necessary software), and also know how to navigate in the command line. Even though all the commands in this file ****are safe to run****, it's not recommended to copy and run code and/or commands from a random website you find on the internet. With this out of the way, we can continue.

### Building dependencies

- [git](https://git-scm.com/)
- [go](https://go.dev/doc/install)

!> Make sure you have all of these programs installed, and in your PATH variable on your system before continuing.

## Building process

1) Automatic way

- Clone PufferStarters repository by running `git clone https://github.com/smoliicek/pufferstarter-cli`
- Change your working directory to the cloned repo `cd pufferstarter-cli`
- Run the automatic build file `.\build.ps1`
- Your compiled .exe is now present at `.\pufferstarter_cli.exe`

2) Manual way

- Clone PufferStarters repository by running `git clone https://github.com/smoliicek/pufferstarter-cli`
- Change your working directory to the cloned repo `cd pufferstarter-cli`
- Create a new folder for your build files `mkdir build`, and change your working directory into it `cd build`
- Generate CMake build files by running `go build ../main.go -o pufferstarter_cli.exe`
- Your .exe file is now present at this location `.\pufferstartel_cli.exe`

### And what next?
PufferStarter relies on enviroment variables or `.env` files. See [Enviroment variables](../Configuration/env.md) for instructiuons.
