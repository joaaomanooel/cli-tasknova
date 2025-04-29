# TaskNova CLI

A powerful command-line task manager built with Go.

## Features

- Create tasks with title, description, and priority
- List all tasks with formatted output
- Update tasks by ID
- Delete tasks by ID
- Persistent storage using JSON format
- Simple and intuitive command-line interface

## Installation

```bash
go install github.com/joaaomanooel/cli-tasknova@latest
```

## Usage

### Add Task

```bash
tasknova add --title "Task Title" --description "Task Description" --priority "high"
```

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

### Update Task

```bash
tasknova update --id <task_id> --title "New Title" --description "New Description" --priority "low"
```

### Delete Task

```bash
tasknova delete --id <task_id>
```

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
