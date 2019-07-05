# golang-web

[Golang on PCF. You might need to change the `manifest.yml` memory size. I got it as low as `6MB`, but whatever you wanna do.](https://docs.cloudfoundry.org/buildpacks/go/index.html)

local:

```sh
go run main.go
# http://localhost:8080
```

pushing to cloud foundry:

```sh
cf push
```
