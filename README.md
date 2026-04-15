# ustrings
Carve Unicode and/or ASCII strings from files.

```console
go install go.foxforensics.dev/ustrings@latest
```

## Usage
```console
$ ustrings [nmao] FILE
```

### Options
* `-n` Minimum string length (default `4`)
* `-m` Maximum string length (default `255`)
* `-a` Only ASCII strings
* `-o` Show file offset

## License
Released under the [MIT License](LICENSE.md).