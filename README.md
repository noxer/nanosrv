# nanosrv

A tiny static file server written in Go. It serves a directory over HTTP, or over HTTPS with HTTP/2 when given a certificate and key, and logs each request to stdout.

## Install

Requires Go 1.27 or newer.

```sh
go install github.com/noxer/nanosrv@latest
```

## Usage

Serve the current directory on port 8888:

```sh
nanosrv
```

Serve a specific directory over HTTPS:

```sh
nanosrv -d ./public -p :8443 -c cert.pem -k key.pem
```

| Flag | Default | Description |
|------|---------|-------------|
| `-d` | current directory | Directory to serve |
| `-p` | `:8888` | Address and port to listen on |
| `-c` | | TLS certificate file |
| `-k` | | TLS private key file |

HTTPS is enabled only when both `-c` and `-k` are set.

## License

GPL-2.0. See [LICENSE](LICENSE).
