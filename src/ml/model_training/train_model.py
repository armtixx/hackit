import os
import joblib
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import Dense, Dropout
from tensorflow.keras.callbacks import EarlyStopping

base_path = os.path.dirname(os.path.abspath(__file__))
out_path = f"{base_path}/out"
model_path = f"{base_path}/model_out"

# Загрузка данных из файла
df = pd.read_excel(f"{out_path}/processing.xlsx")

# Разделение данных на признаки и целевую переменную
selected_columns = ["date", "airline", "from", "to", "class", "price"]
df_selected = df[selected_columns]
X = df_selected.drop("price", axis=1)
y = df_selected["price"]

# Нормализация данных
scaler = StandardScaler()
X_scaled = scaler.fit_transform(X)

# Разделение данных на тренировочные и тестовые наборы
X_train, X_test, y_train, y_test = train_test_split(X_scaled, y, test_size=0.2, random_state=42)

# Построение и обучение модели Keras с регуляризацией и тюнингом гиперпараметров
def build_model(input_shape):
    model = Sequential()
    model.add(Dense(units=128, activation="relu", input_shape=input_shape))
    model.add(Dropout(0.3))
    model.add(Dense(units=64, activation="relu"))
    model.add(Dropout(0.3))
    model.add(Dense(units=32, activation="relu"))
    model.add(Dropout(0.3))
    model.add(Dense(units=1))

    model.compile(optimizer="adam", loss="mean_squared_error")
    return model

# Создание коллбэка EarlyStopping
early_stopping = EarlyStopping(monitor='val_loss', patience=10, restore_best_weights=True)

model_keras = build_model((X_train.shape[1],))
history = model_keras.fit(
    X_train, y_train, 
    epochs=150, batch_size=32, 
    validation_split=0.2, verbose=1, 
    callbacks=[early_stopping]
)

# Оценка модели на тестовых данных
test_loss = model_keras.evaluate(X_test, y_test)
print(f"Test Loss: {test_loss}")

# Добавление прогнозируемых значений в новый столбец "forecast"
df["forecast"] = model_keras.predict(scaler.transform(X)).astype(int)

# Сохранение данных в файл "processing_by_model.xlsx"
df.to_excel(f"{out_path}/processing_by_model.xlsx", index=False)

# Сохранение нормализатора и модели 
model_keras.save(f"{model_path}/model.h5")
joblib.dump(scaler, f"{model_path}/scaler.pkl")
