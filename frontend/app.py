import grpc
from proto import auth_pb2, auth_pb2_grpc
from flask import Flask, render_template, request, redirect, url_for, session

app = Flask(__name__)
app.secret_key = "секретная_строка_для_сессии"

channel = grpc.insecure_channel("localhost:50051")
stub = auth_pb2_grpc.AuthServiceStub(channel)

# users = {}
# messages = []

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




# @app.route("/register", methods=["GET", "POST"])
# def register():
#     if request.method == "POST":
#         username = request.form["username"]
#         password = request.form["password"]

#         if username in users:
#             return "такой пользователь существует"
        
#         users[username] = password
#         session["username"] = username
#         return redirect(url_for("messenger"))

#     return render_template("register.html")

@app.route("/login", methods=["GET", "POST"])
def login():
    if request.method == "POST":
        username = request.form["username"]
        password = request.form["password"]

        if username in users and users[username] == password:
            session["username"] = username
            return redirect(url_for("messenger"))
        return "неверный логин или пароль"

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

    return render_template("messenger.html", username=session["username"], messages=messages)

@app.route('/clear', methods=['POST'])
def clear_messages():
    messages.clear()
    return redirect('/messenger')

if __name__ == "__main__":
    app.run(debug=True, port=5500)
