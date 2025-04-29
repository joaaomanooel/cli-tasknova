# TaskNova User Guide

## Introduction

TaskNova is a powerful CLI for task and note management, developed in Go. This guide provides a detailed overview of all available features.

## Main Commands

### add
Adds a new task to the system.

**Flags:**
- `-t, --title` (required): Task title
- `-d, --description`: Detailed description
- `-p, --priority`: Priority (low, normal, high)

### list
Lists all existing tasks.

### update
Updates an existing task.

**Flags:**
- `-i, --id` (required): Task ID
- `-t, --title`: New title
- `-d, --description`: New description
- `-p, --priority`: New priority

### delete
Removes a task from the system.

**Flags:**
- `-i, --id` (required): Task ID

## Usage Examples

### Basic Flow
```bash
# Add a task
tasknova add -t "Team Meeting" -d "Project discussion" -p high

# List tasks
tasknova list

# Update task
tasknova update -i 1 -t "Team Meeting - Urgent"

# Delete task
tasknova delete -i 1
```