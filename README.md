# Gopolish

Gopolish is a lightweight and fast tool for formatting and prettifying Go test results.

## Installation

This is currently a work in progress. You can build the binary yourself by following the instructions below.

```
> git clone git@github.com:tom773/Gopolish.git
> cd Gopolish
> go build -o <path-to-your-project>/gopolish cmd/main.go
```
## Usage

```
> go test ./... -v | gopolish
```
That's it, no need to worry about the flags or anything else (-v makes it look cooler). Just pipe the output of your tests to Gopolish and it will do the rest.
