# python-web-pipfile

Similar to `python-web` except using `pipenv` with `Pipfile`.

- [pypi/pipenv](https://pypi.org/project/pipenv/)
- [pipenv docs homepage](https://pipenv.pypa.io/en/latest/)

## local

```sh
# local setup and teardown (assuming you have pipenv installed):
pipenv install
pipenv shell
gunicorn main:app
exit

# alternatively, after install, could have just ran as one-off:
pipenv run gunicorn main:app
```

a few helpful commands to have on hand:

```sh
pipenv --three
pipenv --venv
pipenv --where
pipenv graph
pipenv lock
pipenv update --outdated
pipenv install flask gunicorn
pipenv install -r ../python-web/requirements.txt
```

## cf push

```sh
cf push
```
