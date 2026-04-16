# cube

> A lightweight Go CLI for automating Ubuntu VPS provisioning and secure initial server setup

## 📦 Installation

### Install via Script

```sh
curl -sSL https://raw.githubusercontent.com/parisikosto/cube/main/install.sh | bash
```

## Development

### Available commands

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

Run timezone commands

```sh
# The command to request for the time in America
cube timezone EST
# or
cube timezone America/New_York
```

```sh
# The command to request for the time in Tokyo
cube timezone Asia/Tokyo
```

```sh
# The command specifies the date format as YYYY/MM/DD,
# and the output should be the current date in that timezone in the prescribed format
cube timezone EST --date 2006/01/02
```
