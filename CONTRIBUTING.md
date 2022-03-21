# Contributing

## Setup your machine

`cleancontacts` is written in [Go](https://golang.org/).

Prerequisites:

- [Task](https://taskfile.dev/#/installation)
- [Go 1.18+](https://golang.org/doc/install)

Clone `cleancontacts`:

```sh
git clone git@github.com:staticdev/cleancontacts.git
```

Install the dependencies with:

```sh
cd cleancontacts
task setup
```

A good way of making sure everything is all right is running the test suite:

```sh
task test
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

```sh
task build
```

When you are satisfied with the changes, we suggest you run:

```sh
task ci
```

Before you commit the changes, we also suggest you run:

```sh
task fmt
```

## Create a commit

Commit messages should be well formatted, and to make that "standardized", we
are using Conventional Commits.

You can follow the documentation on
[their website](https://www.conventionalcommits.org).

## Submit a pull request

Push your branch to your `cleancontacts` fork and open a pull request against the main branch.
