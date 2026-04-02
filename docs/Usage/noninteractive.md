# Non-Interactive mode

Puffer-starter features an easy to understand CLI.

!> Please see the docs about [configuring PufferStarter](../Configuration/configuration-file.md) before continuing

### Available options
```
Usage:
  pufferstarter-cli [flags]

Flags:
      --env string         Path to the .env file
  -g, --getInfo            Gets info about a server (requires --id)
  -h, --help               help for pufferstarter-cli
      --id string          Set the Server ID (length: 8)
  -l, --listAll            Lists all servers and IDs
  -s, --setStatus string   Set status: on, off, restart, kill (requires --id)


```

### Explaining time!
- No options
  - Displays the help dialog, as seen above.
- `--id` option
  - Sets the ServerID. This option cannot be the only option set while starting PufferStarter.
- `-l, --listAll` option
  - Lists all server names and IDs
- `-g, --getInfo` option
  - Gets information about the specified server
  - Must be run with `--id` set
- `-s, --setStatus` option
  - The only valid arguments for this option is `off`, `on` and `kill`
  - This option will set the server status to the specified status
    - off - Stop the server
    - on - Start the server
    - kill - Kill the server
  - Must be run with `--id` set
