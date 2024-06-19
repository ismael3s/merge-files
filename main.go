package main

import (
	"github.com/spf13/cobra"
)

func executeCommand() *cobra.Command {
	var outputFilenameFromCli string
	command := &cobra.Command{
		Use:   "execute {source_dir} ",
		Short: "Merge all files inside of a folder into a single file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := NewMerger(args[0], outputFilenameFromCli).Execute()
			if err != nil {
				panic(err)
			}
		},
	}
	command.Flags().StringVarP(&outputFilenameFromCli, "output", "o", "output", "output filename")
	return command
}

func main() {
	rootCommand := &cobra.Command{}
	rootCommand.AddCommand(executeCommand())
	rootCommand.Execute()
}
