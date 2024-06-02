from urllib import request
from flask import Flask, jsonify
import numpy as np
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

from sklearn.linear_model import LinearRegression
from sklearn.model_selection import train_test_split
from sklearn.metrics import mean_absolute_percentage_error
from sklearn.preprocessing import StandardScaler

import tensorflow as tf
from tensorflow import keras
from tensorflow.keras.models import Sequential # type: ignore
from tensorflow.keras.layers import Dense, Dropout # type: ignore

import joblib

# Загрузка данных из файла
df = pd.read_excel('/home/scrumpovi4/ML/hackit/src/ml/processing.xlsx')

# Разделение данных на признаки и целевую переменную
selected_columns = ['date', 'airline', 'from', 'to', 'class', 'price']
df_selected = df[selected_columns]
X = df_selected.drop('price', axis=1)
y = df_selected['price']

# Нормализация данных
scaler = StandardScaler()
X_scaled = scaler.fit_transform(X)

# Разделение данных на тренировочные и тестовые наборы
X_train, X_test, y_train, y_test = train_test_split(X_scaled, y, test_size=0.2, random_state=42)

# Построение и обучение модели Keras с регуляризацией и тюнингом гиперпараметров
def build_model(input_shape):
    model = Sequential()
    model.add(Dense(units=128, activation='relu', input_shape=input_shape))
    model.add(Dropout(0.3))
    model.add(Dense(units=64, activation='relu'))
    model.add(Dropout(0.3))
    model.add(Dense(units=32, activation='relu'))
    model.add(Dropout(0.3))
    model.add(Dense(units=1))

    model.compile(optimizer='adam', loss='mean_squared_error')
    return model

model_keras = build_model((X_train.shape[1],))
history = model_keras.fit(X_train, y_train, epochs=200, batch_size=32, validation_split=0.2, verbose=1)

# Добавление прогнозируемых значений в новый столбец "forecast"
df['forecast'] = model_keras.predict(scaler.transform(X)).astype(int)

# Сохранение данных в файл "processing_by_model.xlsx"
df.to_excel('/home/scrumpovi4/ML/hackit/src/ml/processing_by_model.xlsx', index=False)

# Сохранение нормализатора и модели 
model_keras.save('/home/scrumpovi4/ML/hackit/src/ml/model.h5')
joblib.dump(scaler, '/home/scrumpovi4/ML/hackit/src/ml/scaler.pkl')