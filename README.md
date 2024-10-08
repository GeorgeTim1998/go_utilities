# go_utilities
# GoLang Project Readme

## Project Overview
This project demonstrates various tasks aimed at developing Go-based utilities, tools, and design patterns implementation. The project spans across solving real-world problems, implementing design patterns, and building utilities that mimic standard UNIX tools, HTTP servers, and other essential software tasks. The goal is to showcase expertise in Go programming, concurrency, data management, and system interaction using Go's robust ecosystem.

## Table of Contents
1. [Design Patterns](#design-patterns)
2. [Development Tasks](#development-tasks)
    - [NTP Time Program](#ntp-time-program)
    - [String Unpacking](#string-unpacking)
    - [Sorting Utility](#sorting-utility)
    - [Anagram Finder](#anagram-finder)
    - [Grep Utility](#grep-utility)
    - [Cut Utility](#cut-utility)
    - [OR Channel](#or-channel)
    - [Shell Utility](#shell-utility)
    - [Wget Utility](#wget-utility)
    - [Telnet Utility](#telnet-utility)
3. [HTTP Server for Calendar Events](#http-server-for-calendar-events)
4. [Installation](#installation)
5. [Usage](#usage)
6. [Testing](#testing)
7. [Stack](#stack)

---

## Design Patterns

Implemented design patterns, each with code examples and documentation:
- **Facade**: Simplifies the interaction with complex subsystems.
- **Builder**: Constructs complex objects step by step.
- **Visitor**: Defines new operations without changing the class structure.
- **Command**: Encapsulates requests as objects to allow for queuing, logging, and undoing.
- **Chain of Responsibility**: Passes requests along a chain of handlers.
- **Factory Method**: Defines an interface for creating objects in a superclass but lets subclasses alter the type of created objects.
- **Strategy**: Enables selection of algorithms at runtime.
- **State**: Allows an object to alter its behavior when its internal state changes.

Each pattern includes a description of:
- Applicability
- Real-world use cases
- Pros and cons

---

## Development Tasks

### NTP Time Program
This program prints the exact time using the NTP library from `github.com/beevik/ntp`. It is designed as a Go module and handles errors by printing them to STDERR and returning a non-zero exit code.

### String Unpacking
A Go function is implemented that unpacks a string with repeating characters. It supports escape sequences and returns an error for invalid strings. Unit tests ensure correctness.

Examples:
- `"a4bc2d5e" => "aaaabccddddde"`
- `"abcd" => "abcd"`
- `"45" => ""` (invalid input)

### Sorting Utility
A Go utility mimicking the Unix `sort` command. It supports sorting by column, numeric values, reverse order, and removes duplicates. Additional features include month sorting, ignoring trailing spaces, and human-readable sorting.

### Anagram Finder
A function that finds all sets of anagrams from a dictionary of Russian words in UTF-8. Results are returned as a map, and words are sorted and deduplicated.

### Grep Utility
A Go utility that mimics the behavior of Unix `grep`, supporting flags like:
- `-A` for printing lines after matches
- `-B` for printing lines before matches
- `-C` for context
- `-c` for counting lines
- `-i` for case insensitivity
- `-v` for inverting match results

### Cut Utility
A Go utility mimicking Unix `cut`, allowing extraction of fields/columns based on a delimiter (default is TAB). It supports custom delimiters and only processes lines with separators.

### OR Channel
This function combines one or more done channels into a single channel. When any channel closes, the combined channel also closes. This solution is ideal for dynamic channel management at runtime.

### Shell Utility
A custom Unix shell implemented in Go, supporting basic commands like:
- `cd` to change directories
- `pwd` to print the current directory
- `echo` to output text
- `kill` to terminate a process
- `ps` to list running processes

It also supports piping commands (`cmd1 | cmd2 | ... | cmdN`).

### Wget Utility
Implements a simple version of `wget`, downloading websites in their entirety.

### Telnet Utility
A simple `telnet` client implemented in Go. It supports connecting to TCP servers, writing to sockets from STDIN, and handling timeout options.

---

## HTTP Server for Calendar Events

A fully functional HTTP server for managing calendar events. The server exposes the following API methods:
- **POST** `/create_event`: Creates a new event.
- **POST** `/update_event`: Updates an existing event.
- **POST** `/delete_event`: Deletes an event.
- **GET** `/events_for_day`: Retrieves events for a specific day.
- **GET** `/events_for_week`: Retrieves events for a week.
- **GET** `/events_for_month`: Retrieves events for a month.

All methods return JSON responses in the form of `{"result": "..."}` or `{"error": "..."}` for success and error cases respectively.

Middleware for logging requests is implemented, and the server is configurable to run on custom ports.

---

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/go-project.git
   cd go-project
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```
