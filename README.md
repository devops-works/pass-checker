# pass-checker

Get the entropy of your password, the maximum time needed to break it, and if the SHA1 sum appears in HaveIBeenPwned API.

## Usage

### Installation and compilation

If you have Golang, you can clone the repo and build it locally :

```
$ git clone https://github.com/devops-works/pass-checker.git
$ cd pass-checker/
$ go build .
```

Else, you can get the binary from the releases.

### Run

```
$ ./pass-checker [options]

options:
    -g
        guesses per second
        default: 1000000000
```

### Demo

[![asciicast](https://asciinema.org/a/fJ8ZUt69EsbEcl13cLgiLeYGD.svg)](https://asciinema.org/a/fJ8ZUt69EsbEcl13cLgiLeYGD)

## Licence

[MIT](https://choosealicense.com/licenses/mit/)