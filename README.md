# Password Generator

This package is aimed to be a utility for quickly generating passwords within a command-line environment

## Installation

Releases coming soon...

## Manual Installation

To manually install this binary, follow these steps.

1. Make sure you have Golang installed; you can get that [here](https://golang.org/dl/)
1. Make sure you have `dep` installed; you can get that [here](https://golang.github.io/dep/docs/introduction.html)
1. Clone this repository
1. Install vendor dependencies using `dep ensure`
1. Run tests (optional) using `go test`
1. Build the binary using `go build`. For help with building, you can run `go help build`

You can now run the compiled binary. The default binary name will be `password-generator` on *NIX platforms, and `password-generator.exe` on Windows. The binary will reside within the same folder as the repository, unless you override where it was built using the `-o` option.

Alternately, if you want to rely on your Go environment, you can install the binary at the GOBIN directory by running `go install`. Be sure your GOBIN directory is in your PATH.

### Install & Run Example

#### Using Go Build

```
git clone https://github.com/theoscion/password-generator.git
cd ./password-generator
dep ensure
go build
./password-generator
```

#### Using Go Build with Override

```
git clone https://github.com/theoscion/password-generator.git
cd ./password-generator
dep ensure
go build -o /usr/local/bin/pwgen
pwgen
```

#### Using Go Install

```
git clone https://github.com/theoscion/password-generator.git
cd ./password-generator
dep ensure
go install
password-generator
```

## Usage

Once installed, this application runs as a standard binary in a command line environment. Logs are output to `STDERR`; the generated password is output to `STDOUT`. If the password fails to generate, the application will exit with a `1` status code, and `STDERR` will show the reason, if verbose logging is on. 

### Flags
| Flag         | Purpose                                                                     | Optional | Default            |
|--------------|-----------------------------------------------------------------------------|:--------:|-------------------:|
| --length=### | Specifies the length of the password. Minimum allowed is 8; maximum is 4096 | Yes      | 32                 |
| --no-symbols | Specifies that the password should only contain letters and numbers         | Yes      | Symbols Included   |
| -v           | Specifies to output verbose logging                                         | Yes      | No verbose logging |

### Usage Examples

#### Generate a 16-character password, without symbols, and output to a file; then print the file contents that contains the generated password

```
password-generator --length=16 --no-symbols > ~/pw.txt
cat ~/pw.txt
```

####  Generate a default 32-character password, with symbols, and copy to clipboard on MacOS; then print the generated password from clipboard with a newline after output is printed

```
password-generator | pbcopy
pbpaste ; echo 
```

####  Generate a 512-character password, without symbols, with verbose logging and a newline after output is printed

```
password-generator --length=512 --no-symbols -v ; echo
```