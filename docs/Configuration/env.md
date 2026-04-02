# Enviroment variables
PufferStarter uses enviroment variables for configuration.

!> First, get your [OAuth2 credentials](../Configuration/getcreds.md), then come back to configure PufferStarter.

## Order of loading

1) Eviroment variables set by the shell - PufferStarter will never overwrite the variables you already set
2) An `.env` file set by the `--env` flag - If you set a custom location for the enviroment file, that file will be used.
3) Enviroment file in the current directory - If you didn't specify an `.env` file location, and have no shell variables set, the contents of the `.env` file in your working directory will be loaded
4) User config directory - if neither of mentioned above is found, PufferStarter will look in your users config directory, respectivly:
   * Linux: ~/.config/pufferstarter/.env
   * macOS: ~/Library/Application Support/pufferstarter/.env
   * Windows: %AppData%\pufferstarter\.env

## Setup

Copy the `.env.example` to your destined location as `.env` and fill out the variables.

### Setting shell variables

If you want to use shell variables for PufferStarter, you can do so using the following:
```
export CLIENT_ID="some_value"
export CLIENT_SECRET="some_value"
export SERVER_IP="some_value"
```
