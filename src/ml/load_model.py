import os
import joblib
import numpy as np
from tensorflow.keras.models import load_model # type: ignore

# Определение пути до директории скрипта
script_dir = os.path.dirname(__file__)
model_file_path = os.path.join(script_dir, 'model.h5')
scaler_file_path = os.path.join(script_dir, 'scaler.pkl')

# Загрузка модели и нормализатора
def load_model_and_scaler():
    model = load_model(model_file_path)
    scaler = joblib.load(scaler_file_path)
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
    return int(predicted_price[0][0])
