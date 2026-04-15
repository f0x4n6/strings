# ustrings
Carve Unicode and/or ASCII strings from files.

```console
go install go.foxforensics.dev/ustrings@latest
```

## Usage
```console
$ ustrings [nmtao] FILE
```

### Options
* `-n` Minimum string length (default `4`)
* `-m` Maximum string length (default `256`)
* `-t` Trim spaces from both ends
* `-a` Only ASCII strings
* `-o` Show file offset

## Acknowledgements
The carving algorithm is based on the original [Strings](https://github.com/robpike/strings) by [Rob Pike](https://github.com/robpike).

## License
Released under the [MIT License](LICENSE.md).