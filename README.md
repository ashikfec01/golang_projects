# Authentication RESTful API with Go and CQRS

Welcome to the Authentication RESTful API repository! This project demonstrates how to build a simple authentication API using the Go programming language and the Command Query Responsibility Segregation (CQRS) pattern.

authentication-api/
├── cmd/
│   ├── main.go      (Entry point for the authentication service)
├── handlers/
│   ├── register.go
│   ├── login.go
│   └── common.go
├── models/
│   ├── commands.go
│   └── events.go
└── main.go               (Common setup for the entire application)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This repository provides a basic example of implementing user registration and login functionality within an authentication API. It follows the CQRS pattern to separate read and write operations, enhancing scalability and maintainability.

## Features

- User registration with unique usernames
- User login with JWT-based authentication
- Project organized following best practices and modularity

## Prerequisites

Before you begin, ensure you have the following:

- Go programming language (installed and configured)
- Basic understanding of RESTful APIs and Go

## Getting Started

1. Clone this repository:

   ```bash
   git clone https://github.com/your-username/authentication-restful-api.git
