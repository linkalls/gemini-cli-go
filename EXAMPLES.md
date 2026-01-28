# Gemini CLI Go - Examples

## Building the Project

```bash
$ make build
go build -o gemini .
```

## Getting Help

```bash
$ ./gemini --help
Gemini CLI is a terminal interface for Google's Gemini AI models.
It provides code understanding, generation, automation, and integration capabilities.

Usage:
  gemini [flags]

Flags:
  -h, --help                   help for gemini
  -m, --model string           Specify the Gemini model to use
      --output-format string   Output format: text, json, or stream-json (default "text")
  -p, --prompt string          Non-interactive mode: provide a prompt
  -v, --version                version for gemini
```

## Version Information

```bash
$ ./gemini --version
gemini version 0.1.0
```

## Non-Interactive Mode

### Simple text output

```bash
$ export GEMINI_API_KEY="your-api-key-here"
$ ./gemini -p "What is the capital of France?"
Paris
```

### JSON output

```bash
$ ./gemini -p "What is 2+2?" --output-format json
{"response": "4"}
```

### Stream JSON output

```bash
$ ./gemini -p "Tell me a short joke" --output-format stream-json
{"type": "response", "text": "Why don't scientists trust atoms?\n\nBecause they make up everything!"}
```

## Interactive Mode

```bash
$ ./gemini
Gemini CLI - Interactive Mode
Using model: gemini-2.0-flash-exp
Type your message and press Enter. Type /quit or /exit to quit.

> Hello! How are you?

I'm doing well, thank you for asking! As a large language model, I don't have feelings 
in the way humans do, but I'm functioning properly and ready to help you with any 
questions or tasks you might have.

How can I assist you today?

> /help

Available commands:
  /help   - Show this help message
  /quit   - Exit the program
  /exit   - Exit the program
  /clear  - Clear the screen

> /quit
Goodbye!
```

## Using Different Models

```bash
$ ./gemini -m gemini-1.5-flash -p "Explain quantum computing in one sentence"
Quantum computing uses quantum-mechanical phenomena like superposition and entanglement 
to perform computations that are intractable for classical computers.
```

## Running Tests

```bash
$ make test
go test -v ./...
?   	github.com/linkalls/gemini-cli-go	[no test files]
?   	github.com/linkalls/gemini-cli-go/cmd	[no test files]
?   	github.com/linkalls/gemini-cli-go/pkg/client	[no test files]
=== RUN   TestLoadConfig
=== RUN   TestLoadConfig/with_GEMINI_API_KEY
=== RUN   TestLoadConfig/with_GOOGLE_API_KEY
=== RUN   TestLoadConfig/without_API_key
--- PASS: TestLoadConfig (0.00s)
    --- PASS: TestLoadConfig/with_GEMINI_API_KEY (0.00s)
    --- PASS: TestLoadConfig/with_GOOGLE_API_KEY (0.00s)
    --- PASS: TestLoadConfig/without_API_key (0.00s)
=== RUN   TestDefaultModel
--- PASS: TestDefaultModel (0.00s)
PASS
ok  	github.com/linkalls/gemini-cli-go/pkg/config	0.002s
?   	github.com/linkalls/gemini-cli-go/pkg/interactive	[no test files]
```

## Installation

```bash
$ ./install.sh
Installing Gemini CLI (Go version)...
Building binary...
Installing to /usr/local/bin (requires sudo)...
[sudo] password for user:

✅ Gemini CLI installed successfully!

To get started:
  1. Set your API key: export GEMINI_API_KEY=your-api-key
  2. Run: gemini

For help: gemini --help
```
