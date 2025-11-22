# rfmtbcmp - Reformat Bandcamp Files

[![Build and test](https://github.com/samcunliffe/rfmtbcmp/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/samcunliffe/rfmtbcmp/actions/workflows/build-and-test.yml) [![codecov](https://codecov.io/gh/samcunliffe/rfmtbcmp/graph/badge.svg?token=NESNLRXF4V)](https://codecov.io/gh/samcunliffe/rfmtbcmp)

This is a Go command-line utility for extracting and reformatting music files purchased and downloaded from [Bandcamp](https://bandcamp.com).

<p align="center">
  <img src="assets/bandcamp-logo-gopher.svg" alt="Gopher wearing a bandcamp tshirt" />
</p>

It's mostly a little utility project for @samcunliffe to learn Go, it's very opinionated in the way the files are formatted. If you like your music files named in some other way it's probably not supported. (Sorry!)

I used GitHub Copilot for autocompletion and for PR review. But I do not use LLMs for generation of any substantial parts of the code. That would defeat the purpose.

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

## Reuse

This tool is released under [MIT-0](./LICENSE.md). Hopefully the code is useful to you, and no need to attribute it to me (you can if you like).

The Gopher logo is taken from [keygx/Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector), was originally designed by Renee French and released under [CC-by-3.0](https://creativecommons.org/licenses/by/3.0/). The Bandcamp logo is adapted from [Wikimedia Commons](https://commons.wikimedia.org/wiki/File:Bandcamp-logotype-aqua.svg).
