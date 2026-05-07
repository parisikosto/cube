# Contributing

## Development

Clone the repository

```sh
git clone git@github.com:parisikosto/cube.git
cd cube
```

Install dependencies

```sh
go mod tidy
```

Build the app

```sh
go build -o cube main.go
```

Run help to list commands

```sh
./cube --help
```

---

## Build with ldflags

Inject version metadata at build time using `-ldflags`:

```sh
go build \
  -ldflags "-X github.com/parisikosto/cube/build.Version=0.1.0 \
            -X 'github.com/parisikosto/cube/build.User=$(id -u -n)' \
            -X 'github.com/parisikosto/cube/build.Time=$(date)'" \
  -o cube main.go
```

Variables are defined in `build/build.go`:

```go
package build

var Version = "development"
var User string
var Time string
var GitCommit string
var TargetOS string
```

### Verify injected values with go tool nm

Build the binary and inspect it:

```sh
go build -o cube main.go
go tool nm ./cube | grep build
```

Expected output — confirms the variables are embedded:

```
1001000 T _go:buildid
1445000 D _go:buildinfo
1467970 B github.com/parisikosto/cube/build.GitCommit
1467980 B github.com/parisikosto/cube/build.TargetOS
1467960 B github.com/parisikosto/cube/build.Time
...
```

---

## Cross-compile

Build for different target architectures:

```sh
# Ubuntu / Linux (amd64)
env GOOS=linux GOARCH=amd64 go build -o cube-linux-amd64 main.go

# macOS (amd64)
env GOOS=darwin GOARCH=amd64 go build -o cube-darwin-amd64 main.go

# Raspberry Pi (arm)
env GOOS=linux GOARCH=arm GOARM=5 go build -o cube-linux-arm main.go
```

---

## Available commands

```sh
cube --help
```

| Group  | Command               | Description                                             |
| ------ | --------------------- | ------------------------------------------------------- |
| Setup  | ubuntu-initial-setup  | Initial VPS setup as root [1] (Ubuntu 24.04.4 LTS)      |
| Setup  | ubuntu-standard-setup | Standard VPS setup as new user [2] (Ubuntu 24.04.4 LTS) |
| System | create-user           | Create a new system user and grant sudo privileges      |
| System | setup-firewall        | Configure and enable the UFW firewall                   |
| System | update-system         | Update and upgrade all system packages                  |
| Git    | install-git           | Install Git version control system                      |
| Git    | setup-git             | Configure global Git user name and email                |
| Git    | setup-github-ssh      | Generate an SSH key pair for GitHub access              |
| Git    | refresh-ssh-agent     | Add your GitHub SSH key to the running ssh-agent        |
| Git    | uninstall-git         | Uninstall Git and remove unused dependencies            |
| Docker | install-docker        | Install Docker CE from the official Docker repository   |
| Docker | docker-add-user       | Add the current user to the docker group                |
| Docker | docker-check-user     | Verify the current user is in the docker group          |
| Docker | docker-prune-all      | Prune all unused Docker resources                       |
| Docker | uninstall-docker      | Uninstall Docker CE and remove unused dependencies      |
| Tips   | linux-tips            | Display useful Linux system and SSH command tips        |
| Tips   | git-tips              | Display useful Git command tips                         |
| Tips   | docker-tips           | Display useful Docker command tips                      |
| Info   | version               | Print version information                               |
| Info   | timezone              | Get the current date in a given timezone                |
