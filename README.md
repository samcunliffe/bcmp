# rfmtbcmp - Reformat Bandcamp Files

[![Build and test](https://github.com/samcunliffe/rfmtbcmp/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/samcunliffe/rfmtbcmp/actions/workflows/build-and-test.yml) [![codecov](https://codecov.io/gh/samcunliffe/rfmtbcmp/graph/badge.svg?token=NESNLRXF4V)](https://codecov.io/gh/samcunliffe/rfmtbcmp)

This is a Go command-line utility for extracting and reformatting music files purchased and downloaded from Bandcamp.

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
