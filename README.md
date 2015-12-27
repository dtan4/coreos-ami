# coreos-ami
[![Build Status](https://travis-ci.org/dtan4/coreos-ami.svg?branch=master)](https://travis-ci.org/dtan4/coreos-ami)

__coreos-ami__ is a CLI to get [the current official CoreOS AMI IDs](https://coreos.com/os/docs/latest/booting-on-ec2.html).

## Usage

```bash
$ coreos-ami -c <channel> [-r <region>] [-t <type>]
```

### Options

- `-c <channel>` (Required)
  - CoreOS channel (alpha, beta, stable)
- `-r <region>` (Optional)
  - AWS region
- `-t <type>` (Optional)
  - AMI type (PV, HVM)

## Install

To install, use `go get`:

```bash
$ go get -d github.com/dtan4/coreos-ami
```

## Contribution

1. Fork ([https://github.com/dtan4/coreos-ami/fork](https://github.com/dtan4/coreos-ami/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[dtan4](https://github.com/dtan4)
