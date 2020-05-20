# liner

Source Code: https://github.com/aerth/liner

License: MIT

### Usage:

`liner <n> <filename>`

skips n-1 lines and prints single line to stdout

if n > number of lines, exit with fatal error.

note: line numbers begin at 1

the `-x` flag outputs hex, useful for non-printable characters

the `-q` flag outputs quoted and escaped output, also useful for non-printable characters

the `-n` flag prints the new line in the hex or quoted output, but has no effect with normal usage

### Example

`liner -n -q 13 main.go`

outputs: 
```
"\tprintln(`Usage of liner:\n"
```

### Free

The MIT License (MIT)

Copyright (c) 2020 aerth <aerth@riseup.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
