# golang-web

[`go`](https://docs.cloudfoundry.org/buildpacks/go/index.html) on CF.

There are [great examples in the `go-buildpack` repository.](https://github.com/cloudfoundry/go-buildpack/tree/c99109e8cbce1284c9a8da71273009722059e87e/fixtures)

Try goofing with the `manifest.yml` memory size. I got it as low as `6MB`, but whatever you wanna do.

local:

```sh
go run main.go
# http://localhost:8080
```

CF:

```sh
cf push
```
