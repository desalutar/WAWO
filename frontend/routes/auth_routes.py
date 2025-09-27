from flask import Blueprint, render_template, request, redirect, url_for, session
from proto import auth_pb2
from grpc_client import stub

auth_bp = Blueprint("auth", __name__, url_prefix="/auth")

@auth_bp.route("/register", methods=["GET", "POST"])
def register():
    if request.method == "POST":
        username = request.form["username"]
        password = request.form["password"]
        try:
            stub.Register(auth_pb2.RegisterRequest(username=username, password=password))
            session["username"] = username
            return redirect(url_for("messenger.messenger"))
        except Exception as e:
            return f"Ошибка регистрации: {str(e)}"
    return render_template("register.html")

@auth_bp.route("/login", methods=["GET", "POST"])
def login():
    if request.method == "POST":
        username = request.form["username"]
        password = request.form["password"]
        try:
            response = stub.Login(auth_pb2.LoginRequest(username=username, password=password))
            if response.success:
                session["username"] = username
                return redirect(url_for("messenger.messenger"))
            else:
                return "Неверный логин или пароль"
        except Exception as e:
            return f"Ошибка входа: {str(e)}"
    return render_template("login.html")

@auth_bp.route("/logout")
def logout():
    session.pop("username", None)
    return redirect(url_for("auth.login"))
