# mdtoc
Toy utility to create Table of contents based on provided MD file.
Can work with literals in header, but don't expect it to work with something
too complicated.

# Install
```bash
go install github.com/insomnes/mdtoc@latest
```

This command will download the package, install its dependencies, compile it, and place the binary in your `$GOPATH/bin` directory.

# Usage
## With defaults
```bash
mdtoc README.md
...

- [mdtoc](#mdtoc)
- [Install](#install)
- [Usage](#usage)
  - [With defaults](#with-defaults)
  - [With depth provided](#with-depth-provided)
```
## With depth provided
```bash
mdtoc -d 1 README.md
...

- [mdtoc](#mdtoc)
- [Install](#install)
- [Usage](#usage)
```
