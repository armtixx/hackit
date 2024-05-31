import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error, r2_score
from sklearn.preprocessing import StandardScaler, OneHotEncoder
from sklearn.compose import ColumnTransformer
from sklearn.pipeline import Pipeline
import requests
import json

# URL вашего сервера
url = 'http://your_backend_url'

# GET запрос для получения данных
response_get = requests.get(url)
data_get = response_get.json()  # Получение данных в формате JSON

# PUT запрос для отправки данных
data_to_send = {"key": "value"}  # Пример данных для отправки, замените на свои данные
response_put = requests.put(url, data=json.dumps(data_to_send))

# Проверка статуса запроса
if response_put.status_code == 200:
    print("PUT request was successful.")
else:
    print(f"Error: {response_put.status_code}")


# Загрузка данных из файла
file_path = '/home/scrumpovi4/Загрузки/Train.xlsx'
df = pd.read_excel(file_path)

# Предобработка данных
# Преобразование 'price' в числовой формат
df['price'] = df['price'].str.replace(',', '').astype(int)

# Преобразование 'time_taken' в минуты
def convert_to_minutes(time_str):
    parts = time_str.split(' ')
    hours = int(parts[0].replace('h', ''))
    minutes = int(parts[1].replace('m', '')) if len(parts) > 1 else 0
    return hours * 60 + minutes

df['time_taken'] = df['time_taken'].apply(convert_to_minutes)

# Заполнение пропусков (если есть)
df.ffill(inplace=True)

# Исключение нечисловых столбцов перед расчетом корреляции
numeric_df = df.select_dtypes(include=['number'])

# Визуализация корреляционной матрицы
correlation_matrix = numeric_df.corr()
plt.figure(figsize=(10, 8))
sns.heatmap(correlation_matrix, annot=True, cmap='coolwarm', fmt=".2f")
plt.title('Correlation Matrix')
plt.show()

# Определение категориальных и числовых признаков
categorical_features = ['airline', 'ch_code', 'from', 'stop', 'to']
numeric_features = ['num_code', 'time_taken']

# Создание пайплайна для предобработки данных
preprocessor = ColumnTransformer(
    transformers=[
        ('num', StandardScaler(), numeric_features),  # Масштабирование числовых признаков
        ('cat', OneHotEncoder(), categorical_features)  # Преобразование категориальных признаков
    ])

# Добавление предобработки данных в пайплайн обучения модели
pipeline = Pipeline(steps=[('preprocessor', preprocessor),
                           ('model', LinearRegression())])

# Разделение данных на обучающую и тестовую выборки
X_train, X_test, y_train, y_test = train_test_split(df.drop('price', axis=1), df['price'], test_size=0.2, random_state=42)

# Обучение модели
pipeline.fit(X_train, y_train)

# Предсказание на тестовой выборке
y_pred = pipeline.predict(X_test)

# Оценка модели
mse = mean_squared_error(y_test, y_pred)
r2 = r2_score(y_test, y_pred)

print(f'Mean Squared Error: {mse}')
print(f'R^2 Score: {r2}')
