# webmachine

Webmachine is a [polygot](https://en.wikipedia.org/wiki/Polyglot_(computing)) web application framework written in Go.

## Dependencies

For any programming language you plan to you use, make sure it's binary (i.e. `ruby`, `python3`) is available in $PATH.

| File Name | Executable Name |
| --- | --- |
| main.go | go |
| main.rb | ruby |
| main.py | python3 |
| main.js | node |

## Install

<!-- ### Via Go Toolchain -->

```bash
go install github.com/mikerybka/webmachine@latest
```

## Usage

```
webmachine <command> [arguments]
```

### Commands

#### serve

Arguments:
- `email` (optional): The email to share with Let's Encrypt
- `dir` (optional): The root directory to serve. Defaults to /etc/webmachine
- `port` (optional): The port to listen on. Defaults to both 443 and 80. If a port is not provided, HTTPS is served on 443 and HTTP served on 80. If the port provided is 443, HTTPS is served on port 443. If any other port is provided, HTTP is served on that port.

### Directory structure

#### Example

Here's an example directory setup.

```
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

#### Rules

The filesystem is organized such that there should be one file for each "route".
A route is generally defined as HTTP Method + URL Path.
Dynamic paths are supported with path variables and the ":"-prefix naming.

If a request routes to a file, only GET requests are allowed on that path and the file is served as-is.

If the request routes to a folder instead, the child directory with the name matching the request's HTTP method is read.
If that folder does not exist, 404.

If it does, the directory is searched for a main.* file matching one of our supported environments.

### Language Support

- Go
- Node.js (planned)
- Python (planned)
- Ruby (planned)

