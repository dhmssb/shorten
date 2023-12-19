# URL Shortener Service

This project is a URL shortening service developed in Go.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)

## Overview

This just a shorten Url and can modify with slug

## Prerequisites

- Go (version X.X.X)
- Postman (for API testing)

## Installation

Guide users on how to install and set up the project:
1. Clone the repository: `git clone https://github.com/your_username/your_repository.git`
2. Navigate to the project directory: `cd your_repository`
3. Run the Go application: `go run main.go`

## Usage

how to use the application:
1. Start the Go application.
2. Use Postman to interact with the API endpoints.

## Endpoints

List and describe the available endpoints:
- `POST /shorten`: Create a shortened URL.
- `PUT /slug/:slug`: Modify the slug of a shortened URL.
- `GET /:slug`: Redirect to the original URL associated with the slug.
- `GET /urls`: Retrieve a list of created short URLs.


