# nodejs

Nodejs on cf. You might need to change params in the `engine` block of `package.json` to match whichever appropriate nodejs buildpacks you have available. Also you should probably change the scaling & memory config in `manifest.yml`.

- [Nodejs buildpack tips](https://docs.cloudfoundry.org/buildpacks/node/node-tips.html)
- [Nodejs service bindings](https://docs.cloudfoundry.org/buildpacks/node/node-service-bindings.html)

local:

```sh
npm install
npm start
# http://localhost:3000
```

pushing to cloud foundry:

```sh
cf push
```
