
<p align="center">
  <img src=".github/img/header.png" alt="Banner" width="100%">
</p>

<p align="center">
  <!-- Add shields from https://shields.io/ -->
  <a href="https://github.com/sponsors/KyleTryon">
    <img alt="GitHub Sponsors" src="https://img.shields.io/github/sponsors/KyleTryon">
  </a>
  <a href="https://github.com/TechSquidTV/uhs-cli/actions/workflows/golangci-lint.yml">
  <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/TechSquidTV/uhs-cli/golangci-lint.yml">
  </a>
  <a href="https://discord.gg/CTC9DVvYZz">
    <img alt="Discord" src="https://img.shields.io/discord/415249366840901643?style=plastic&logo=discord">
  </a>
</p>

<p align="center">
UltimateHomeServer CLI
</p>

<p align="center">
An interactive CLI to assist in configuring services for the <a href="https://github.com/TechSquidTV/UltimateHomeServer">UHS stack</a>.
</p>


**Beta:** This CLI is currently in active development and is subject to change.


# Getting Started

## Installation


### Binary

Utilize the install script to download the latest release for your platform.

```bash
wget https://raw.githubusercontent.com/TechSquidTV/uhs-cli/main/install.sh
```

```bash
chmod +x install.sh
```

```bash
./install.sh
```

### Go

```bash
go install github.com/TechSquidTV/uhs-cli@latest
```

## Usage

```bash
Usage:
  uhs [command]

Available Commands:
  configure   Configure your UHS instance
  default     Get the default configuration for UHS
  help        Help about any command

Flags:
  -h, --help     help for uhs-cli
  -t, --toggle   Help message for toggle
```

Generate configurations for individual services:

```bash
uhs-cli configure <service> -o values.yaml
```

Generate a default configuration for all services:

```bash
uhs-cli default -o values.yaml
```