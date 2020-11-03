# python-web-pipfile

Same as `python-web`, except using a `Pipfile` and [`pipenv`](https://pypi.org/project/pipenv/).

## local

```sh
# local setup (assuming you have pipenv available):
pipenv --three

# run locally
gunicorn main:app

# local cleanup:
pipenv uninstall --all
```

a few helpful commands to have on hand:

```sh
pipenv --venv
pipenv --where
pipenv graph
pipenv lock
```
