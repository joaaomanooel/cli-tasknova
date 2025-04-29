% TASKNOVA(1) TaskNova CLI
% João Manoel
% January 2024

# NAME
tasknova - A CLI task manager

# SYNOPSIS
**tasknova** [*command*] [*flags*]

# DESCRIPTION
TaskNova is a command-line task manager that helps you organize tasks with priorities and descriptions.

# COMMANDS
**add**
: Add a new task with title, description, and priority.

**list**
: List all tasks in a formatted table.

**update**
: Update an existing task by ID.

**delete**
: Delete a task by ID.

**completion**
: Generate shell completion scripts.

# OPTIONS
**--help**
: Show help message and exit.

**--version**
: Show program's version number and exit.

# EXAMPLES
**tasknova add --title "Learn Go" --description "Start with basics" --priority high**
: Add a new high-priority task.

**tasknova list**
: List all tasks.

**tasknova update --id 1 --title "New Title"**
: Update task title.

**tasknova delete --id 1**
: Delete a task.

# FILES
~/.tasknova/tasks.json
: Tasks storage file

# SEE ALSO
Full documentation at: <https://github.com/joaaomanooel/cli-tasknova>