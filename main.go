package main

import (
	"fmt"
	"os"
	"smoliicek/pufferstarter/pkg/auth"
	"smoliicek/pufferstarter/pkg/operator"

	"github.com/spf13/cobra"

	"github.com/joho/godotenv"
)

func main() {
	var serverID string
	var status string

	var rootCmd = &cobra.Command{
		Use:   "pufferstarter-cli",
		Short: "A CLI tool for server management",
		Run: func(cmd *cobra.Command, args []string) {

			if cmd.Flags().Changed("listAll") {
				fmt.Println("Listing all servers...")
				// getAllServers(0)
				return
			}

			if cmd.Flags().Changed("getInfo") {
				if serverID == "" {
					fmt.Println("You need to specify --id to run this function.")
					os.Exit(1)
				}
				fmt.Printf("Getting info for server %s\n", serverID)
				// getServerInfo(ip, token, 0, serverID)
				return
			}

			if cmd.Flags().Changed("setStatus") {
				if serverID == "" {
					fmt.Println("You need to specify --id to run this function.")
					os.Exit(1)
				}

				switch status {
				case "on":
					fmt.Println("Starting server", serverID)
					// setServerStatus(ip, token, 0, serverID, "on")
				case "off":
					fmt.Println("Stopping server", serverID)
					operator.ChangeServerStatus("serverIP", "authToken", serverID, "off")
				case "kill":
					fmt.Println("Killing server", serverID)
					// setServerStatus(ip, token, 0, serverID, "kill")
				default:
					fmt.Println("Invalid status. Use on, off, or kill.")
				}
				return
			}

			if serverID != "" {
				fmt.Println("You can't run --id by itself.")
			}
			cmd.Help()
		},
	}

	rootCmd.Flags().StringVar(&serverID, "id", "", "Set the Server ID (length: 8)")
	rootCmd.Flags().BoolP("listAll", "l", false, "Lists all servers and IDs")
	rootCmd.Flags().BoolP("getInfo", "g", false, "Gets info about a server (requires --id)")
	rootCmd.Flags().StringVarP(&status, "setStatus", "s", "", "Set status: on, off, kill (requires --id)")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func getToken() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	serverIP := os.Getenv("SERVER_IP")

	token, err := auth.GetAuthToken(clientID, clientSecret, serverIP)
	if err != nil {
		panic(err)
	}

	fmt.Println("Auth Token:", token)
}
