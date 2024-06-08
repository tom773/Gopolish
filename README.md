# Gopolish

Polish up those ugly Go test outputs with Gopolish. Simple tool that adds some color and much needed formatting to the output of your Go tests.

## Installation

### Using Go
```
> go install github.com/tom773/gopolish@latest
```
Or build from the source:

```
> git clone git@github.com:tom773/Gopolish.git
> cd Gopolish
> go build -o <path-to-your-project>/gopolish cmd/main.go
```
## Usage

```
> go test ./... -v | gopolish <or ./gopolish if you built from source>
```
That's it. Ideally you'd pipe the output in with -v, as this makes it look cooler, but it's not necessary.
