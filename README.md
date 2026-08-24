# AetherLang

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-00ADD8?style=for-the-badge)
![CI](https://img.shields.io/badge/CI-GitHub%20Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white)
![CodeQL](https://img.shields.io/badge/CodeQL-Security-00ADD8?style=for-the-badge)
![Version](https://img.shields.io/badge/Version-v2.1.0-00ADD8?style=for-the-badge)

> Custom programming language with lexer, parser, AST, bytecode compiler, and VM - written in Go

`programming-language` `compiler` `virtual-machine` `bytecode` `lexer` `parser` `golang` `ast`

---

## What is it?

**AetherLang** is A complete programming language ecosystem: lexer, parser, AST, bytecode compiler, and virtual machine. Write .aether files and run them on your own runtime.

## Why should you care?

- **Fast** - Compiled Go binary, zero overhead
- **Secure** - CodeQL analysis + Dependabot
- **Offline-first** - Works without internet
- **Lightweight** - Single binary deployment
- **Developer-friendly** - Clean CLI with docs

---

## Features

- Custom lexer with token types
- Recursive descent parser
- AST generation
- Bytecode compiler
- Stack-based VM
- Built-in functions (print, len, type)
- Variables, loops, conditionals
- Functions with closures
- String manipulation
- Error handling with line numbers

---

## Quick Start

### Prerequisites
- Go 1.21 or higher

### Install from source
```bash
git clone https://github.com/AetherCodeHQ/aetherlang.git
cd aetherlang
go build -o aetherlang .
```

### Run
```bash
./aetherlang --help
```

---

## Usage

./aether run hello.aether  OR  ./aether repl  OR  ./aether compile hello.aether -o hello.ab

---

## CLI Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--path` | Target directory | `.` |
| `--format` | Output format (json, text) | `text` |
| `--output` | Output filename | `stdout` |
| `--verbose` | Enable verbose output | `false` |

---

## Development

```bash
git clone https://github.com/AetherCodeHQ/aetherlang.git
cd aetherlang
go build -o aetherlang .
go test ./...
golangci-lint run
```

---

## Contributing

Contributions welcome! See [CONTRIBUTING.md](CONTRIBUTING.md).

## Security

Report to: aethercode.core@gmail.com | See [SECURITY.md](SECURITY.md)

## License

MIT License - see [LICENSE](LICENSE)

---

<p align="center">
  Built with love by <a href="https://github.com/AetherCodeHQ">AetherCode</a> | <a href="https://github.com/AetherCode-Core">AetherCode Core</a>
</p>


---

## What's New in v2.0.0

- Professional documentation with badges
- CI/CD pipeline with GitHub Actions
- Security analysis with CodeQL
- Dependency management with Dependabot
- Code quality with GolangCI-Lint
- Issue and PR templates
- Contributing guidelines

## Categories

| Category | Description |
|----------|-------------|
| Other | This project is part of the AetherCode ecosystem |

## Related Projects

Part of [AetherCode Core](https://github.com/AetherCode-Core) ecosystem.

---
## Changelog
### v2.0.0-2208
- Auto updated
- Tag: `v2.0.0-2208`

## What's New in v2.0.0

- **Real Implementation**: All code rewritten from stubs to working tools
- **CLI Interface**: Full command-line interface with subcommands
- **Error Handling**: Comprehensive error handling throughout
- **Documentation**: Updated with usage examples

See [CHANGELOG.md](CHANGELOG.md) for full details.


## What's new in v2.1.0

- Clean CLI with subcommands
- Robust error handling
- Fast, standard-library-only implementation
