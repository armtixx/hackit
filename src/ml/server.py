from flask import Flask, request
from model import make_prediction
from config import HOST, PORT

app = Flask(__name__)

needed_args = [ "date", "airline", "from_id", "to_id", "flight_class" ]

@app.route("/predict")
def predict():
    args = request.args

    if len(args) != 5 or not all(arg in args for arg in needed_args):
        return {"ok": False}, 400

    return {"ok": True, "price": make_prediction(**args)}, 200

if __name__ == "__main__":
    app.run(host=HOST, port=PORT, debug=True)
