from flask import Flask
import os

app = Flask(__name__)

@app.route('/')
def root():
    return 'I am instance ' + str(os.getenv('CF_INSTANCE_INDEX', 0))

if __name__ == '__main__':
    # Bind to PORT if defined, otherwise default to 8080.
    port = int(os.getenv('PORT', 8080))
    app.run(host='0.0.0.0', port=port)
