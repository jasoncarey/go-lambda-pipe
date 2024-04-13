# Go Lambda Pipe

Go Lambda Pipe is a CLI tool for invoking AWS Lambda Functions.

## Setup
1. Download and install Go
2. Build go-lambda-pipe
```bash
go build
```
3. Install go-lambda-pipe
```bash
go install
```
4. Run the CLI tool
```bash
go-lambda-pipe [command]
```
5. Invoking Lambda functions requires AWS CLIv2 to be installed and configured in your environment

## Available Commands
`invokeLambda` - invokes a lambda function with a JSON file event
`processJson` - stringifies the body key of a JSON and writes to a new file
`pipe` - invokes the `processJson` and `invokeLambda` command in succession

