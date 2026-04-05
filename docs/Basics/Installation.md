# Installation
This document will guide you through installing PufferStarter, so you could use it to its full potential.

### Dependencies
For PufferStarter to work property, you need:
- A PufferPanel instance hosted somewhere
- Know its IP Address

### Installation
To install PufferStarter, you need to go to [GitHub Releases](https://github.com/smoliicek/pufferstarter-cli/releases/), and download the following files for your OS:
- `pufferstarter-cli-windows-amd64.exe` for Windows
- `pufferstarter-cli-OS-arch` for Linux and other UNIX systems

After you download those files, continue with these instructions for you OS of choice:

#### For Windows:

- You are all set, just start the app in a terminal and you are good to go!

#### For Linux and other UNIX systems:

- Run this command `chmod +x ./pufferstarter-cli*` so you can execute it
- Move the file to your Path `sudo mv ./pufferstarter-cli* /usr/local/bin/pufferstarter_cli`
- Start the `pufferstarter-cli` executable!

### And what next?
PufferStarter relies on enviroment variables or `.env` files. See [Enviroment variables](../Configuration/env.md) for instructiuons.
