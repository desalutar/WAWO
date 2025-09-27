from flask import Blueprint, render_template, request, redirect, url_for, session

messenger_bp = Blueprint("messenger", __name__)
messages = []
chats = {}
@messenger_bp.route("/messenger", methods=["GET", "POST"])
def messenger():
    if "username" not in session:
        return redirect(url_for("auth.login"))

    if request.method == "POST":
        msg = request.form["message"]
        messages.append(f"{session['username']}: {msg}")

    return render_template("chat.html", username=session["username"], messages=messages, chats=chats)

@messenger_bp.route("/clear", methods=["POST"])
def clear_messages():
    messages.clear()
    return redirect(url_for("messenger.messenger"))
