# A-CLI-Todo-App-in-Go
A simple CLI application for managing a todo list using GoLang
# CLI Todo App in Go

## Project Overview

This project is a simple command-line interface (CLI) application for managing a todo list, created as a learning exercise to familiarize myself with the Go programming language. The application demonstrates basic CRUD (Create, Read, Update, Delete) operations on a list of tasks, with data stored in a JSON file.

## Language

This project is written in **Go** (Golang), a statically typed, compiled programming language designed for simplicity and efficiency. For more information about Go, you can visit the [Go Programming Language](https://golang.org/) website.

## Purpose

The primary goal of this project was to gain hands-on experience with Go by building a practical application. This todo app serves as a straightforward example of how to use Go for file handling, command-line arguments, and basic data manipulation.

## Features

- **Add a New Todo**: Add new tasks to your todo list.
- **Edit an Existing Todo**: Update the title of a task by specifying its index.
- **Delete a Todo**: Remove a task from the list by its index.
- **Toggle Task Completion**: Mark a task as completed or not completed.
- **List All Todos**: Display all tasks in a formatted table.

## How It Works

### 1. Data Storage

Tasks are saved in a JSON file, which allows for easy persistence and retrieval of the todo list between application runs.

### 2. Command-Line Interface

The application uses command-line flags to execute various operations. Users can interact with the app via terminal commands to manage their tasks.

### 3. Main Components

- **`main.go`**: Contains the core logic for managing todos, including the `Todo` struct and methods for adding, editing, deleting, toggling, and listing tasks.
- **`storage.go`**: Manages saving and loading data to/from a JSON file using the `Storage` struct.
- **`cmdflags.go`**: Handles parsing command-line flags and executing corresponding actions using the `CmdFlags` struct.

## Installation

 **Clone the Repository**:
   ```bash
   git clone https://github.com/amiraa205/A-CLI-Todo-App-in-Go.git
   cd A-CLI-Todo-App-in-Go

