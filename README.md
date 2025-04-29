# Go Data Structures - A Learning Project

## Introduction

Welcome! This repository contains implementations of common data structures written in the Go programming language.

This project serves primarily as a **learning exercise** for me (Justin Frost aka `gneissguise`) as I work to deepen my understanding of Go. I'm focusing on idiomatic Go practices, exploring concepts like pointers, interfaces, error handling, package structure, and particularly leveraging **Go Generics** (introduced in Go 1.18+).

Implementations aim to be:

-   Clear and understandable.
-   Reasonably efficient (though pedagogical clarity might sometimes take precedence over bleeding-edge optimization).
-   Developed using a **Test-Driven Development (TDD)** approach, with comprehensive unit tests located alongside the code.

## Implemented Data Structures Checklist

This checklist tracks the progress of implementing various data structures.

-   [x] Singly Linked List (`linkedlist`)
-   [ ] Stack
-   [ ] Queue
-   [ ] Binary Search Tree (BST)
-   [ ] Hash Map

## Running Tests

Each data structure resides in its own package (subdirectory). Tests are written using Go's built-in `testing` package.

To run all tests for the entire project from the root directory (`go-datastructures`):

```bash
go test ./...
```

To run tests for a specific data structure (e.g., the linked list):

```bash
cd linkedlist
go test -v
```

(The `-v` flag provides verbose output.)

## Disclaimer

As this is a learning project, the implementations might evolve and are provided "as is" without guarantees of production-readiness. Feedback and suggestions (if you happen to find this repo!) are welcome but understand the primary goal is educational.
