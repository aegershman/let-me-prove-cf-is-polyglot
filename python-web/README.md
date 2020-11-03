# python-web

- [`python` buildpack docs](https://docs.cloudfoundry.org/buildpacks/python/index.html)
- [`python` buildpack source code](https://github.com/cloudfoundry/python-buildpack)
- [`python` buildpack fixtures as examples](https://github.com/cloudfoundry/python-buildpack/tree/master/fixtures)
- [heroku `python` buildpack tutorials](https://devcenter.heroku.com/articles/python-gunicorn#adding-gunicorn-to-your-application)

Note that in the `Procfile` we can goof around with some flags for `gunicorn`:

```yml
web: gunicorn main:app --timeout 10 --preload --max-requests 1200
```

There's no reason for them; you don't need them. They're just examples picked up from [this really great `heroku` tutorial](https://devcenter.heroku.com/articles/python-gunicorn#adding-gunicorn-to-your-application).

## local work

local dev setup, assuming you're using `python3` as per `brew` or something to that effect:

```sh
# local setup:
python3 -m venv getting-started
source getting-started/bin/activate
pip freeze > requirements.txt

# run locally; a few examples:
gunicorn main:app
gunicorn main:app --timeout 10 --access-logfile -

# local cleanup:
pip uninstall -r requirements.txt -y
deactivate
rm -rf getting-started/ __pycache__/
```

## vendoring

You should also be able to `vendor` dependencies and store them along your source code or in a packaged archive output. In this example we're still assuming we're in our virtual env:

```sh
mkdir -p vendor

# vendors pip *.tar.gz into vendor/
pip download -r requirements.txt --no-binary=:none: -d vendor

# git add vendor/; git commit -m "vendor content"; git push; etc., etc.
```

>cf push uploads your vendored dependencies. The buildpack installs them directly from the vendor/ directory.

See the [`python buildpack` docs on vendoring and private dependencies](https://docs.cloudfoundry.org/buildpacks/python/index.html#vendoring) for more details.

## cf push

whelp... not an _enormous_ amount to say here.

CF:

```sh
cf push
```

By the way, if you have any questions on any of this, if you're searching around, you should definitely add the keyword `"heroku"` to your queries. They have fantastic articles, tutorials, and direct examples out there for `python` support with `buildpacks` and cloud-native `python` in general.
