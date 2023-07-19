WIP

# webmachine

`webmachine` is a filesystem-based HTTP router that aims to support all of your favourite programming languages allowing you to easily create apps using multiple programming languages.

## Programming Language Support

- [x] Go
- [x] Ruby
- [x] Python
- [x] JavaScript
- [ ] TypeScript

## Install

<!-- ### Via Go Toolchain -->
```bash
go install github.com/mikerybka/webmachine@latest
```

#### A note on dependencies

For any programming language you plan to you use, make sure it's binary (i.e. `ruby`, `python3`) is available in $PATH.

| Language | Executable Name |
| --- | --- |
| Go | go |
| Ruby | ruby |
| Python | python3 |
| JavaScript | node |

## Usage

```
webmachine serve [arguments]
```

Arguments:
- `email` (optional): The email to share with Let's Encrypt
- `dir` (optional): The root directory to serve. Defaults to /etc/webmachine
- `port` (optional): The port to listen on. Defaults to both 443 and 80. If a port is not provided, HTTPS is served on 443 and HTTP served on 80. If the port provided is 443, HTTPS is served on port 443. If any other port is provided, HTTP is served on that port.

### Routing Logic

The filesystem is organized such that there should be one file for each "route".
A route is defined as HTTP Method + Hostname + URL Path.
Dynamic paths are supported with path variables and the ":"-prefix naming.

If a request routes to a file, only GET requests are allowed on that path and the file is served as-is.

If the request routes to a folder, the child directory with the name matching the request's HTTP method is read.
If that folder does not exist, 404.

If it does, the directory is searched for a main.[go|rb|py|js] file (in that order).

#### Example Directory Structure

```
example.org/python/GET/main.py
example.org/ruby/GET/main.rb
example.org/node/GET/main.js
dynamicsite.com/GET/main.go
dynamicsite.com/favicon.ico
dynamicsite.com/items/GET/main.go
dynamicsite.com/items/POST/main.go
dynamicsite.com/items/new/GET/main.go
dynamicsite.com/items/:itemID/GET/main.go
dynamicsite.com/items/:itemID/PUT/main.go
dynamicsite.com/items/:itemID/DELETE/main.go
staticsite.com/index.html
staticsite.com/products
staticsite.com/services
staticsite.com/about
staticsite.com/contact
staticsite.com/faq
staticsite.com/privacy
staticsite.com/terms
```

### Handler Logic

The handler interface is designed to be simple.
Standard input is filled with the request body.
Standard output is expected to contain the response body.

The exit code indicates the response status.
- If the exit code is 0: `200 OK`
- If the exit code is 1XX, 2XX, 3XX, 4XX or 5XX, the response status is set to that code.

Cookies and headers can be set through via Standard Error.

URL and query params as well as request headers are encoded in the command with arguments like `--key=value`.
