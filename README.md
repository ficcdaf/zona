# Zona

Zona is a small tool for building a static blog website. It allows users to write pages and blog posts in Markdown and automatically build them into a static, lightweight blog.

> **Warning:** Implementing v1 functionality is still WIP. There will be binaries available to download from the Releases tab when Zona is ready to be used. Until then, you are welcome to download the source code, but I can't promise anything will work!

## Table of Contents

- [Features](#v1-features)
- [Installation](#installation)
- [Roadmap](#roadmap)

## v1 Features

- Write your pages in Markdown.
- Build a lightweight website with zero JavaScript.
- Simple CLI build interface.
- HTML layout optimized for screen readers and Neocities.

## Getting Started

### Dependencies

- `go 1.23.2`

```Bash
# On Arch Linux
sudo pacman -S go

# On Ubuntu/Debian
sudo apt install go
```

### Installation

First, download the repository and open it:

```Bash
git clone https://github.com/ficcdaf/zona.git && cd zona
```

On Linux:

```Bash
# run the provided build script
./build.sh
```

On other platforms:

```Bash
go build -o bin/zona cmd/zona
```

The resulting binary can be found at `bin/zona`.

## Roadmap

- [ ] Zona configuration file to define build options.
- [ ] Image optimization & dithering options.
- [ ] AUR package after first release
- [ ] Automatic RSS/Atom feed generation.

## Inspirations

- [Zoner](https://git.sr.ht/~ryantrawick/zoner)
- Zonelets
