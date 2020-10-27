# pass-checker

Get the entropy of your password, the maximum time needed to break it, and if the SHA1 sum appears in HaveIBeenPwnd API.

## Usage

### Installation and compilation

Go is needed.

```
$ git clone https://github.com/devops-works/pass-checker.git
$ cd pass-checker/
$ go build .
```

### Run

```
./pass-checker [options]

options:
    -g
        guesses per second
        default: 1000000000
```

### Demo

[![asciicast](https://asciinema.org/a/fJ8ZUt69EsbEcl13cLgiLeYGD.svg)](https://asciinema.org/a/fJ8ZUt69EsbEcl13cLgiLeYGD)

## Licence

[MIT](https://choosealicense.com/licenses/mit/)