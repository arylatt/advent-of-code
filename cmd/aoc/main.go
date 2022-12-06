package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "aoc [--chdir DIR]",
	PersistentPreRunE: persistentPreRunERoot,
}

func init() {
	rootCmd.PersistentFlags().String("chdir", "", "Switch to a different working directory before executing the given subcommand")
	rootCmd.MarkFlagDirname("chdir")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func persistentPreRunERoot(cmd *cobra.Command, args []string) error {
	if chdir := cmd.Flag("chdir").Value.String(); chdir != "" {
		return os.Chdir(chdir)
	}

	return nil
}
