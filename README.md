# coreos-ami
[![Build Status](https://travis-ci.org/dtan4/coreos-ami.svg?branch=master)](https://travis-ci.org/dtan4/coreos-ami)

__coreos-ami__ is a CLI to print [CoreOS Container Linux AMI IDs](https://coreos.com/os/docs/latest/booting-on-ec2.html).

## Usage

```bash
$ coreos-ami [alpha|beta|stable] [-v]

# Example
$ coreos-ami beta | head -n5
ap-northeast-1 hvm ami-91f3bbf6
ap-northeast-1 pv  ami-45f3bb22
ap-northeast-2 hvm ami-4c16c622
ap-northeast-2 pv  ami-1614c478
ap-south-1     hvm ami-4a334325

# Print release information
$ coreos-ami beta -v | head -n10
Version:      1325.2.0
Release date: 2017-03-01 18:48:26 +0000

AMIs:
ap-northeast-1 hvm ami-77742010
ap-northeast-1 pv  ami-98481cff
ap-northeast-2 hvm ami-f12dfd9f
ap-northeast-2 pv  ami-3f2efe51
ap-south-1     hvm ami-df0b7bb0
ap-south-1     pv  ami-1a097975
```

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
