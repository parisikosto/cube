<div align="center">

# cube

**A lightweight Go CLI for automating Ubuntu VPS provisioning and secure initial server setup.**

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![Ubuntu 24.04 LTS](https://img.shields.io/badge/Ubuntu-24.04%20LTS-E95420?logo=ubuntu&logoColor=white)](https://ubuntu.com)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

</div>

---

A bare Ubuntu box, two commands later, it's a server: system updates, a
sudo user, SSH, the firewall, Docker — all the tedious setup you'd
otherwise re-type every time you spin up a droplet, gone.

## Install

```sh
curl -sSL https://raw.githubusercontent.com/parisikosto/cube/main/install.sh | bash
```

Installs the latest release to `/usr/local/bin/cube` and sets up bash
completion automatically. Run `cube --help` to get started.

## Quickstart

The two setup commands are numbered on purpose — run them one after the
other, on the same fresh box.

**1. As root:**

```sh
cube ubuntu-initial-setup
```

Runs a full system update, creates a new user with sudo privileges, and
configures the firewall.

**2. As the new user:**

```sh
cube ubuntu-standard-setup
```

Installs and configures Git, installs Docker, and adds the user to the
`docker` group.

## Commands

Every step above is also its own command, grouped exactly like
`cube --help` groups them.

#### Setup

| Command | What it does |
|---|---|
| `setup-firewall` | Allows OpenSSH through the firewall, enables UFW, shows status |
| `setup-git` | Sets `git config --global user.name`/`user.email` |
| `setup-github-ssh` | Generates an RSA-4096 SSH key for GitHub and configures agent auto-load |
| `create-user` | Prompts for a username, runs `adduser` + `usermod -aG sudo` |

#### System

| Command | What it does |
|---|---|
| `update-system` | `apt update && upgrade && dist-upgrade && autoremove` |
| `refresh-ssh-agent` | Re-adds your GitHub key when the ssh-agent dies mid-session |
| `timezone <tz>` | Prints the current date in a given timezone (`--date` to customize the format) |

#### Git

| Command | What it does |
|---|---|
| `install-git` | Installs Git via apt, verifies the version |
| `uninstall-git` | Removes Git and unused dependencies |
| `git-tips` | Prints a categorized reference of useful Git commands |

#### Docker

| Command | What it does |
|---|---|
| `install-docker` | Installs Docker CE from the official repo, verifies the service |
| `uninstall-docker` | Removes Docker CE and unused dependencies |
| `docker-add-user` | Adds the current user to the `docker` group |
| `docker-check-user` | Verifies the current user is in the `docker` group |
| `docker-prune-all` | Prunes all unused Docker resources (system/container/volume/network/image) |
| `docker-tips` | Prints a categorized reference of useful Docker commands |

#### Info

| Command | What it does |
|---|---|
| `linux-tips` | Prints a categorized reference of useful Linux/SSH commands |
| `version` | Prints the installed version |

Example:

```sh
cube timezone Europe/Athens
cube timezone America/New_York --date "2006-01-02 15:04:05"
```

## Why the `*-tips` commands exist

`git-tips`, `docker-tips`, and `linux-tips` print a categorized
reference of the commands you always forget — no browser tab required.
Small detail, but it's the one people actually remember.

## Built with

- [Go](https://go.dev) 1.26
- [Cobra](https://github.com/spf13/cobra) for the CLI
- [promptui](https://github.com/manifoldco/promptui) for interactive prompts

## Contributing

Want to build from source, cross-compile, or see the full command
reference? See [CONTRIBUTING.md](./CONTRIBUTING.md).

## License

[MIT](./LICENSE) © [Paris Kostopoulos](https://github.com/parisikosto)
