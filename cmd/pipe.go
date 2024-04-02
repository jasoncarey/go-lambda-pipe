package cmd

import "github.com/spf13/cobra"

var pipeCmd = &cobra.Command{
	Use:   "pipe [lambda function name] [input file]",
	Short: "Process a JSON file and invoke a Lambda function",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		processJSONFile(args[1], "temp.json")
		invokeLambdaFunction(args[0], "temp.json")
	},
}

func init() {
	rootCmd.AddCommand(pipeCmd)
}
