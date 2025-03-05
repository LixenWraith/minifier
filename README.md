# Minifier

A command-line tool that recursively minifies files in a directory structure. Powered by [tdewolff/minify](https://github.com/tdewolff/minify).

## Supported File Types

- HTML (.html)
- CSS (.css)
- JavaScript (.js)
- JSON (.json)
- SVG (.svg)
- XML (.xml)

## Installation

Pre-compiled binaries are available for:
- Linux AMD64: `bin/minifier_linux_amd64`
- FreeBSD AMD64: `bin/minifier_freebsd_amd64`

Download the appropriate binary for your system and make it executable:
```bash
chmod +x minifier_*_amd64
```

## Building from Source

Requires Go 1.21 or later.

```bash
go get -u github.com/tdewolff/minify/v2
go build -o minifier main.go
```

For FreeBSD builds, use the provided script:
```bash
./bsd_make.sh
```

## Usage

```bash
minifier <source_directory> <target_directory>
```

### Example

```bash
./minifier ./my_website ./my_website_minified
```

This will:
1. Create `my_website_minified` if it doesn't exist
2. Recursively copy the directory structure from `my_website`
3. Minify supported files during the copy process
4. Leave unsupported file types unchanged

## License

BSD-3