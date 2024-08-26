package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
  "git.sr.ht/~ferdinandyb/throttle/server"
)

func init() {
  rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
  Use:   "server",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
    server.Start()

  },
}
