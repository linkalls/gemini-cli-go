# Gemini CLI (Go)

[![License](https://img.shields.io/github/license/linkalls/gemini-cli-go)](https://github.com/linkalls/gemini-cli-go/blob/main/LICENSE)

Gemini CLI is a terminal interface for Google's Gemini AI models written in Go.
It provides lightweight access to Gemini, giving you the most direct path from
your prompt to the model.

This is a Go rewrite of the original [Gemini CLI](https://github.com/google-gemini/gemini-cli).

## 🚀 Why Gemini CLI Go?

- **🎯 Free tier**: 1,000 requests/day with Gemini API key
- **🧠 Powerful Gemini models**: Access to latest Gemini models
- **⚡ Fast**: Written in Go for performance and low memory footprint
- **💻 Terminal-first**: Designed for developers who live in the command line
- **🛡️ Open source**: Apache 2.0 licensed

## 📦 Installation

### Pre-requisites before installation

- Go 1.21 or higher
- macOS, Linux, or Windows

### Build from source

```bash
git clone https://github.com/linkalls/gemini-cli-go
cd gemini-cli-go
make build
```

Or install directly with Go:

```bash
go install github.com/linkalls/gemini-cli-go@latest
```


## 🔐 Authentication

Currently, the Go version supports authentication via API key only.

### Gemini API Key

**✨ Best for:** Developers who need specific model control

**Benefits:**

- **Free tier**: 1,500 requests/day with Gemini 2.0
- **Model selection**: Choose specific Gemini models
- **Simple setup**: Just set an environment variable

```bash
# Get your key from https://aistudio.google.com/apikey
export GEMINI_API_KEY="YOUR_API_KEY"
gemini
```

You can also use `GOOGLE_API_KEY` instead of `GEMINI_API_KEY`.

## 🚀 Getting Started

### Basic Usage

#### Start interactive mode

```bash
gemini
```

#### Use specific model

```bash
gemini -m gemini-2.0-flash-exp
```

#### Non-interactive mode for scripts

Get a simple text response:

```bash
gemini -p "Explain the architecture of this codebase"
```

For structured output, use the `--output-format json` flag:

```bash
gemini -p "What is the meaning of life?" --output-format json
```

For real-time event streaming, use `--output-format stream-json`:

```bash
gemini -p "Tell me a story" --output-format stream-json
```

### Quick Examples

#### Interactive Chat

```bash
gemini
> What is the capital of France?
> Tell me more about its history
> /quit
```

#### Ask a Quick Question

```bash
gemini -p "What's the weather like in Tokyo?"
```

## 🔧 Building and Development

### Build from source

```bash
# Clone the repository
git clone https://github.com/linkalls/gemini-cli-go
cd gemini-cli-go

# Install dependencies
make install

# Build the binary
make build

# Run it
./gemini
```

### Available Make commands

```bash
make help      # Show available commands
make build     # Build the binary
make test      # Run tests
make format    # Format code
make clean     # Clean build artifacts
```

## 📚 Available Commands

In interactive mode, use these commands:

- `/help` - Show help message
- `/quit` or `/exit` - Exit the program
- `/clear` - Clear the screen

## 🚧 Current Limitations

This is a minimal Go rewrite focused on core functionality. The following features from the original TypeScript version are not yet implemented:

- OAuth login with Google Account
- Vertex AI support
- File system tools
- Shell command execution
- Web fetching and Google Search grounding
- MCP (Model Context Protocol) server integration
- Custom commands and skills
- Conversation checkpointing
- GEMINI.md context files
- VS Code integration

These features may be added in future versions.

## 📖 Resources

- **[Original Gemini CLI](https://github.com/google-gemini/gemini-cli)** - The TypeScript version
- **[Gemini API Documentation](https://ai.google.dev/gemini-api/docs)** - Official API docs
- **[GitHub Issues](https://github.com/linkalls/gemini-cli-go/issues)** - Report bugs or request features

## 🤝 Contributing

We welcome contributions! This project is fully open source (Apache 2.0), and we
encourage the community to:

- Report bugs and suggest features
- Improve documentation
- Submit code improvements
- Add new features

Please open an issue or pull request on [GitHub](https://github.com/linkalls/gemini-cli-go).

## 📄 Legal

- **License**: [Apache License 2.0](LICENSE)
- **Security**: [Security Policy](SECURITY.md)

---

<p align="center">
  Go rewrite by linkalls - Original TypeScript version by Google
</p>
