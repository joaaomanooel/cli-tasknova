# TaskNova CLI

TaskNova is a powerful command-line task management tool built with Go, designed to help you manage your tasks efficiently from the terminal.

## Features

- Add tasks with title, description, and priority levels
- Tasks are automatically timestamped with creation and update times
- Persistent storage using JSON format
- Simple and intuitive command-line interface

## Installation

To install TaskNova, make sure you have Go installed on your system, then run:

```bash
go install github.com/yourusername/cli-tasknova@latest
```

Or clone the repository and build from source:

```bash
git clone https://github.com/yourusername/cli-tasknova.git
cd cli-tasknova
make build
```

## Usage

### Adding a Task

```bash
tasknova add --title "My Task" --description "Task description" --priority high
```

Options:
- `--title, -t`: Task title (required)
- `--description, -d`: Task description (optional)
- `--priority, -p`: Task priority (optional, default: "low")

## Development

### Prerequisites

- Go 1.x or higher
- Make

### Building

```bash
make build
```

This will create binaries for multiple platforms in the `bin` directory.

### Running Tests

```bash
make test
```

### Available Make Commands

- `make test`: Run tests with coverage
- `make build`: Build for all platforms
- `make build-macos`: Build for macOS
- `make build-linux`: Build for Linux
- `make build-windows`: Build for Windows
- `make test-watch`: Watch for changes and run tests
- `make test-coverage`: Show test coverage

## License

[MIT](LICENSE.md)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
