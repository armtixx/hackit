import os
import joblib
import numpy as np
from tensorflow.keras.models import load_model

base_path = os.path.dirname(os.path.abspath(__file__))

# Загрузка модели и нормализатора
def load_model_and_scaler():
    model = load_model(f"{base_path}/model.h5")
    scaler = joblib.load(f"{base_path}/scaler.pkl")
    return model, scaler

model_keras, scaler = load_model_and_scaler()

def make_prediction(date, airline, from_id, to_id, flight_class):
    # Преобразование входных данных в нужный формат
    input_data = np.array([[date, airline, from_id, to_id, flight_class]])
    input_data_scaled = scaler.transform(input_data)
    
    # Предсказание цены
    predicted_price = model_keras.predict(input_data_scaled)
    
    # Возвращение предсказания
    return int(predicted_price[0][0])
