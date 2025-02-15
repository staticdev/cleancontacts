# Clean Contacts

[![Tests](https://github.com/staticdev/cleancontacts/workflows/Tests/badge.svg)](https://github.com/staticdev/cleancontacts/actions?workflow=Tests) [![Codecov](https://codecov.io/gh/staticdev/cleancontacts/branch/main/graph/badge.svg)](https://codecov.io/gh/staticdev/cleancontacts) [![Go Reference](https://pkg.go.dev/badge/github.com/staticdev/cleancontacts.svg)](https://pkg.go.dev/github.com/staticdev/cleancontacts) [![Go Report Card](https://goreportcard.com/badge/github.com/staticdev/cleancontacts)](https://goreportcard.com/report/github.com/staticdev/cleancontacts)

Do not want to share all your contact info to mobile apps? This software is for you!

Export your contacts in VCard format and run the program. BANG! You have a new VCard file with cleaned contacts with just their names and phone numbers.

## Installation

Download the relevant binary on Assets of latest [release](https://github.com/staticdev/cleancontacts/releases).

Alternatively, if you have golang 1.23+ installed:

```sh
go install github.com/staticdev/cleancontacts
```

## Usage

Go to your contacts provider and export your contacts in VCard format. The you will have a contacts file with extension `.vcf`.

Run on command prompt:

```sh
cleancontacts ContactsFilename.vcf
# or full path
cleancontacts /Downloads/ContactsFilename.vcf
```

It generates a new cleaned contact file with the name ContactsFilename_cleaned.vcf.

Next step would be stop syncing your contacts on your phone, delete all contacts and import the cleaned contacts. This will prevent all apps from getting extra information other than contact names and phone numbers.

## License

Distributed under the terms of the [MIT license](LICENSE.md), Clean Contacts is free and open source software.
