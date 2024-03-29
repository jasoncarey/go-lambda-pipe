package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type JSONData struct {
	Body interface{} `json:"body"`
}

var processjsonCmd = &cobra.Command{
	Use:   "processjson [input file] [output file]",
	Short: "Stringify the body of a JSON file and write to new file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		processJSONFile(args[0], args[1])
	},
}

func processJSONFile(inputFile, outputFile string) {
	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(fileContent, &jsonData); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if bodyValue, exists := jsonData["body"]; exists && bodyValue != nil {
		jsonData["body"] = jsonStringify(bodyValue)
	}

	modifiedContent, err := json.MarshalIndent(jsonData, "", " ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	if err := os.WriteFile(outputFile, modifiedContent, 0644); err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}
}

func jsonStringify(value interface{}) string {
	bytes, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("Failed to stringify JSON: %v", err)
	}
	return string(bytes)
}

func init() {
	rootCmd.AddCommand(processjsonCmd)
}
