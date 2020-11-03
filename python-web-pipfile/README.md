# python-web-pipfile

Same as `python-web`, except using a `Pipfile` and `pipenv`.

- [pypi/pipenv](https://pypi.org/project/pipenv/)
- [pipenv docs homepage](https://pipenv.pypa.io/en/latest/)

## local

```sh
# local setup (assuming you have pipenv available):
pipenv --three
pipenv install flask
pipenv install gunicorn

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
