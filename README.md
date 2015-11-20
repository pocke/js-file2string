js-file2string
===============

Installation
------------

```sh
go get github.com/pocke/js-file2string
```

Usage
------

```
Usage of js-file2string:
  -f, --filename-only[=false]: trim directory
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
