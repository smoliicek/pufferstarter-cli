# Building instructions for Windows
This document will guide you through the process of building PufferStarter on your Windows system.

### Prerequisites
This guide requires you to have Administrator privileges on your user account (for installing necessary software), and also know how to navigate in the command line. Even though all the commands in this file ****are safe to run****, it's not recommended to copy and run code and/or commands from a random website you find on the internet. With this out of the way, we can continue.

### Building dependencies

- [git](https://git-scm.com/)
- [go](https://go.dev/doc/install)
- GNU make

!> Make sure you have all of these programs installed, and in your PATH variable on your system before continuing.

## Building process

- Clone PufferStarters repository by running `git clone https://github.com/smoliicek/pufferstarter-cli`
- Change your working directory to the cloned repo `cd pufferstarter-cli`
- Build the app `make windows/amd64`
- Your .exe file is now present at this location `.\dist\pufferstartel_cli.exe`

### And what next?
PufferStarter relies on enviroment variables or `.env` files. See [Enviroment variables](../Configuration/env.md) for instructiuons.
