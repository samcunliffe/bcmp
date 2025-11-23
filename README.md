# bcmptidy - Tidy Bandcamp Files

[![Build and test](https://github.com/samcunliffe/bcmptidy/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/samcunliffe/bcmptidy/actions/workflows/build-and-test.yml) [![codecov](https://codecov.io/gh/samcunliffe/bcmptidy/graph/badge.svg?token=NESNLRXF4V)](https://codecov.io/gh/samcunliffe/bcmptidy)

This is a Go command-line utility for extracting, renaming, and organising music files purchased and downloaded from [Bandcamp](https://bandcamp.com).

<p align="center">
  <img src="assets/bandcamp-logo-gopher.svg" alt="Gopher wearing a bandcamp tshirt" />
</p>

It's a little weekend project for [@samcunliffe](https://github.com/samcunliffe) to learn Go, it's very opinionated in the way the files are formatted. If you like your music files named and/or arranged in some other way it's probably not supported. (Sorry!)

I used GitHub Copilot for autocompletion and for PR review. But I do not use LLMs for generation of any substantial parts of the code. That would defeat the purpose.

## Install

You should be able to run

```sh
go install github.com/samcunliffe/bcmptidy@latest
```

then you'll have `bcmptidy` in your path.

## Usage

Run `bcmptidy` without any arguyments, or with `-h,--help` to get usage information:

```sh
bcmptidy --help
```

To run over a file and organise to the default location (`$HOME/Music`) run:

```sh
bcmptidy '/path/to/Downloads/Artist - Album.zip'
```

To extract to some other location, use `-p,--music-path`:

```sh
bcmptidy --music-path /somewhere/else/Music '/path/to/Downloads/Artist - Album.zip'
```

## Support

Should work with Linux and MacOS.
I don't support Windows in my free time projects. No idea if this will work in Windows. Probably not.
I don't have precompiled binaries, but maybe I will in the future.

## Specification

Music from Bandcamp is downloaded either as individual music files or as a zip archive of the whole album.

### Zip archive

Of the format:

```
Album Artist - Album Name.zip
```

Containing files of the format:

```
Album Artist - Album Name - 01 First Track Name.flac
Album Artist - Album Name - 02 Second Track Name.flac
...
```

### Destination

Files will be extracted and organised into a directory structure:

```
$HOME/Music/Album Artist/Album Name/01 First Track Name.flac
```

Or:

```
.
├── Album Artist
│   └── Album Name
│       ├── 01 First Track.flac
│       ├── 02 Second Track.flac
│       └── ...
└── Another Artist
```

## References

As this was a learning exercise, here are the things I used:

- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)
- [Go By Example](https://gobyexample.com/)
- [Official Go Docs](https://go.dev/doc)

## Reuse

This tool is released under [MIT-0](./LICENSE.md). Hopefully the code is useful to you, and no need to attribute it to me (you can if you like).

The Gopher logo is taken from [keygx/Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector), was originally designed by Renee French and released under [CC-by-3.0](https://creativecommons.org/licenses/by/3.0/). The Bandcamp logo is adapted from [Wikimedia Commons](https://commons.wikimedia.org/wiki/File:Bandcamp-logotype-aqua.svg).
