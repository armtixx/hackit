import numpy as np
from tensorflow.keras.models import load_model # type: ignore
import joblib

# Загрузка модели и нормализатора
def load_model_and_scaler():
    model = load_model('/home/scrumpovi4/ML/hackit/src/ml/model.h5')  # Загрузите всю модель
    scaler = joblib.load('/home/scrumpovi4/ML/hackit/src/ml/scaler.pkl')
    return model, scaler

model_keras, scaler = load_model_and_scaler()

# Определение маршрута для предсказания
def make_prediction(date, airline, from_id, to_id, flight_class):
    # Преобразование входных данных в нужный формат
    input_data = np.array([[date, airline, from_id, to_id, flight_class]])
    input_data_scaled = scaler.transform(input_data)
    
    # Предсказание цены
    predicted_price = model_keras.predict(input_data_scaled)
    
    # Возвращение предсказания
    return (int(predicted_price[0][0])/2)
