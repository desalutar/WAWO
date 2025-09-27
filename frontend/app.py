import grpc
from proto import auth_pb2, auth_pb2_grpc
from flask import Flask, render_template, request, redirect, url_for, session

app = Flask(__name__)

channel = grpc.insecure_channel("auth:50051")
stub = auth_pb2_grpc.AuthServiceStub(channel)
app.secret_key = "любая_уникальная_строка_секретная"

@app.route("/")
def index():
    return redirect(url_for("register"))

@app.route("/register", methods=["GET", "POST"])
def register():
    if request.method == "POST":
        username = request.form["username"]
        password = request.form["password"]
        try:
            stub.Register(auth_pb2.RegisterRequest(username=username, password=password))
            session["username"] = username
            return redirect(url_for("messenger"))
        except grpc.RpcError as e:
            return f"Ошибка Логина: {e.details()}"
    
    return render_template("register.html")

@app.route("/login", methods=["GET", "POST"])
def login():
    if request.method == "POST":
        username = request.form["username"]
        password = request.form["password"]
        try: 
            response = stub.Login(auth_pb2.LoginRequest(username=username, password=password))
            if response.token: 
                session["username"] = username
                session["token"] = response.token
                return redirect(url_for("messenger"))
            else:
                return "Неверный логин или пароль"
        except grpc.RpcError as e:
            return f"Ошибка входа: {e.details()}"

    return render_template("login.html")
 
@app.route("/logout")
def logout():
    session.pop("username", None)
    return redirect(url_for("login"))

@app.route("/messenger", methods=["GET", "POST"])
def messenger():
    if "username" not in session:
        return redirect(url_for("login"))

    if request.method == "POST":
        msg = request.form["message"]
        messages.append(f"{session['username']}: {msg}")

    return render_template(
        "messenger.html",
        username=session["username"],
        messages=messages,
        chats=chats,
        chat_user=None
    )

@app.route('/clear', methods=['POST'])
def clear_messages():
    messages.clear()
    return redirect('/messenger')

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5500, debug=True)
