from flask import Flask, request
from load_model import make_prediction

app = Flask(__name__)

@app.route("/predict")
def predict():
    data = request.get_json()
    
    # Extracting required fields
    date = data['date']
    airline = data['airline']
    from_id = data['from']
    to_id = data['to']
    flight_class = data['class']

    input_data = [date, airline, from_id, to_id, flight_class]
    
    return {"price": make_prediction(*input_data)}

if __name__ == '__main__':
    app.run(port=8080, debug=True)
    
