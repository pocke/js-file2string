js-file2string
===============

[![Build Status](https://travis-ci.org/pocke/js-file2string.svg)](https://travis-ci.org/pocke/js-file2string)

Installation
------------

```sh
go get github.com/pocke/js-file2string
```

Usage
------

```sh
$ js-file2string --help
Usage of js-file2string:
  -f, --filename-only[=false]: trim directory
  -r, --replace[=false]: replace as javascript identifier
  -t, --typing[=false]: output .d.ts for TypeScript
```

### In command line

```sh
$ js-file2string file1 file2 > files.js
```

### In JavaScript

```javascript
var files = require('files');
var file1Value = files['file1'];
```

### Options

- `-f`, `--filename-only`
- `-r`, `--replace`
- `-t`, `--typing`


Development
-----------

### Test

```sh
$ go test -v
$ node tset.js
```

License
-------

These codes are licensed under CC0.

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](http://creativecommons.org/publicdomain/zero/1.0/deed.en)
