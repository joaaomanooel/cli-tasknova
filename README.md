# TaskNova CLI

TaskNova is a command-line interface (CLI) task management tool built in Go using the Cobra framework. It helps you manage tasks and notes efficiently through terminal commands.

## Features

- Create tasks with title, description, and priority
- List all tasks with formatted output
- Delete tasks by ID
- Persistent storage using JSON format
- Simple and intuitive command-line interface

## Installation

```bash
go install github.com/joaaomanooel/cli-tasknova@latest
```

## Usage

### Add a Task

```bash
tasknova add --title "Task Title" --description "Task Description" --priority "high"
```

Options:
- `-t, --title`: Task title (required)
- `-d, --description`: Task description (optional)
- `-p, --priority`: Task priority (optional, default: "low")

### List Tasks

```bash
tasknova list
```

This command displays all tasks with their details including:
- ID
- Title
- Description
- Priority
- Creation date
- Last update date

### Delete a Task

```bash
tasknova delete --id 1234
```

Options:
- `-i, --id`: Task ID (required)

## Development

### Prerequisites

- Go 1.16 or higher
- Git
- golangci-lint (for code linting)

### Code Quality

Run the linter:
```bash
make lint
```

### Setup

1. Clone the repository
```bash
git clone https://github.com/joaaomanooel/cli-tasknova.git
```

2. Navigate to the project directory
```bash
cd cli-tasknova
```

3. Install dependencies
```bash
go mod download
```

4. Build the project
```bash
go build
```

### Testing

Run the test suite:
```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
