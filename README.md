# isit
> The domain availability command-line uitlity.

## Build Source

```shell
# git clone this repo
cd isit
go build isit
```

## Command-line Options

The `available` command-line flag simply checks if the given domain(s) are available.

```shell
$ isit available google.com g000000000000000gle.com
false		google.com
true		g000000000000000gle.com
```

The `registered` command-line flag simply checks if the given domain(s) have been registered.

```shell
$ isit registered google.com g000000000000000gle.com
true		google.com
false		g000000000000000gle.com
```

The `resolvable` command-line flag simply checks if the given domain(s) are resolvable (to an IP address).

```shell
$ isit resolvable google.com g000000000000000gle.com
true		google.com
false		g000000000000000gle.com
```

## Help Menu

```shell
NAME:
   isit - domain availability command-line utility

USAGE:
   isit [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     available, a   check if the given domain(s) are available
     registered, r  check if the given domain(s) are registered
     resolvable, R  check if the given domain(s) are resolvable
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
