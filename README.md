# isit
> The domain availability command-line uitlity.

The command-line application helps with domain availability discovery which can be helpful for a variety of different things. It provides three methods of lookup which all work concurrently.

## Build from Source

Until I've figured out a better option:

```shell
git clone https://github.com/picatz/isit.git
cd isit
go build isit.go
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

```
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
