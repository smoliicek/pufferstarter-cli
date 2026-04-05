# PufferStarter-cli

PufferStarter-cli is a command-line interface tool designed to manage servers on [PufferPanel](https://www.pufferpanel.com/) (v3.0+). It allows you to list servers, view detailed server information (including resource usage and JVM stats), and control server power states (start, stop, restart, kill) directly from your terminal.

> [!NOTE]
> For usage with older PufferPanel versions, use the older C++ based version - <https://github.com/smoliicek/pufferstarter-cli/tree/OLD-PS-cpp-version>

- **List All Servers**: Quickly see all servers on your panel, their IDs, status, and ports.
- **Server Information**: Get detailed stats for a specific server, including:
  - CPU and Memory usage.
  - JVM Heap and Metaspace statistics (for supported server types).
  - Node information and IP/Port details.
- **Power Management**: Control your servers with simple commands (`on`, `off`, `restart`, `kill`).

> [!WARNING]
> This is a really quick guide, if you don't know what you're doing, read the [docs](https://docs.smoliicek.cz).

## Installation

### From Binaries

Download the latest binary for your operating system from the [Releases](https://github.com/smoliicek/pufferstarter-cli/releases) page.

#### Windows

Just download `pufferstarter-cli.exe` and run it in your preferred terminal (PowerShell or CMD).

#### Linux and other UNIX systems

1. Download the `tar.gz` archive and extract it.
2. Make the binary executable:

   ```bash
   chmod +x ./pufferstarter-cli
   ```

3. (Optional) Move it to your PATH:

   ```bash
   sudo mv ./pufferstarter-cli /usr/local/bin/
   ```

### From Source

If you have Go installed, you can build it yourself:

```bash
git clone https://github.com/smoliicek/pufferstarter-cli.git
cd pufferstarter-cli
go build -o pufferstarter-cli main.go
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

## Badges

![Made with Go](https://forthebadge.com/badges/made-with-go.svg) ![Built with <3](https://forthebadge.com/badges/built-with-love.svg)
