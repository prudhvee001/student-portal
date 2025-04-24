from flask import Flask, jsonify

app = Flask(__name__)

@app.route("/user/<int:user_id>", methods=["GET"])
def get_user(user_id):
    user_data = {"user_id": user_id, "name": "John Doe"}
    return jsonify(user_data)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5001)
