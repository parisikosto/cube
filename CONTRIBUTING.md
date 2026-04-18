# Contributing

## Development

Install packages

```sh
go install
```

Build app

```sh
go build
```

Run help to list commands

```sh
cube --help
```

## Available commands

Run timezone commands

```sh
# Get the current date in a given timezone
cube timezone EST
# or
cube timezone America/New_York
```

```sh
# Customize the date format
cube timezone EST --date 2006/01/02
```
