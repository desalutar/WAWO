from flask import Flask, redirect, url_for
from routes.auth_routes import auth_bp
from routes.messenger_routes import messenger_bp

app = Flask(__name__)
app.secret_key = "your_secret_key_here"

app.register_blueprint(auth_bp)
app.register_blueprint(messenger_bp)

@app.route("/")
def index():
    return redirect(url_for("auth.register"))

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5500, debug=True)
