# Clean Contacts

[![Go Reference](https://pkg.go.dev/badge/github.com/staticdev/cleancontacts.svg)](https://pkg.go.dev/github.com/staticdev/cleancontacts) [![Go Report Card](https://goreportcard.com/badge/github.com/staticdev/cleancontacts)](https://goreportcard.com/report/github.com/staticdev/cleancontacts)

Do not want to share all your contact info to mobile apps? This software is for you!

Export your contacts in VCard format and run the program. BANG! You have a new VCard file with cleaned contacts with just their names and telephones.

## Installation

If you have golang 1.17+ installed:

```sh
go install github.com/staticdev/cleancontacts
```

## Usage

Run on command prompt:

```sh
cleancontacts contacts.vcf
# or full path
cleancontacts /Downloads/contacts.vcf
```

## License

Distributed under the terms of the [MIT license](LICENSE.md), Clean Contacts is free and open source software.
