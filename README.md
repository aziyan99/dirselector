# dirselector

`dirselector` is a simple cli app that show directory selector and print the selected
directory into the console.

## Installation

Download the latest release from the release page or clone this repository and build the project using `Make` command. The executeable will be available inside the `build` directory.

To build the dirselector without using `Make`:
```
go build -mod=readonly -ldflags "-s -w" -o .\build\dirselector.exe .\cmd\dirselector\dirselector.go
```

## Usage

```
.\dirselector.exe
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

## License
[MIT](https://github.com/aziyan99/dirselector/blob/main/LICENSE)
