# Heimdall

`heimdall` is the gatekeeper of the graphlog platform.

# Pre-requisites

- Go v1.13
- Google Wire

# Setup

```bash
    go mod vendor
    go run cmd/main.go
```

_The application should now be served on http://localhost:9000_

# Tech stack

- Google Wire: For Dependency Injection
- FastHTTP: For our http router - built to be faster than native `net/http` but lacks a lot of features many router libraries like Gorrila/mux. Used because the intention is to be a fast routing mechanism to validate requests and send messages.
- Go v1.13: significant improvements over v1.12 and Go modules!
