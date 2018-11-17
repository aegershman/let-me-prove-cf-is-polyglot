from flask import Flask
import os

app = Flask(__name__)

port = int(os.getenv("PORT", 8080))

@app.route('/')
def root():
    return 'I am instance ' + str(os.getenv("CF_INSTANCE_INDEX", 0))

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=port)
