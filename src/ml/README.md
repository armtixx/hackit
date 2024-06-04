# FlyEasy
----

```
┬╔╦╗┌─┐┌─┐┌┬┐
│ ║ ├┤ ├─┤│││
o ╩ └─┘┴ ┴┴ ┴
```

```
╔═╗┬ ┬ ┬╔═╗┌─┐┌─┐┬ ┬
╠╣ │ └┬┘║╣ ├─┤└─┐└┬┘
╚  ┴─┘┴ ╚═╝┴ ┴└─┘ ┴
```

# Navigation
----

<!--toc:start-->
- [Пошаговая инструкция по запуску](#пошаговая-инструкция-по-запуску)
  - [Создание виртуального окружения и установка зависимостей](#создание-виртуального-окружения-и-установка-зависимостей)
  - [Запуск](#запуск)
    - [Запуск обучения модели](#запуск-обучения-модели)
    - [Запуск сервера](#запуск-сервера)
    - [Отправка запросов к серверу](#отправка-запросов-к-серверу)
- [Типы передаваемых данных](#типы-передаваемых-данных)
    - [Когда (Смотрится по месяцам)](#когда-смотрится-по-месяцам)
    - [АвиаКомпании](#авиакомпании)
    - [Откуда](#откуда)
    - [Куда](#куда)
    - [Выбор Класса](#выбор-класса)
<!--toc:end-->


# Пошаговая инструкция по запуску
----
## Создание виртуального окружения и установка зависимостей
```sh
python -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

## Запуск
----
- Входные тестовые данные подаются в файле `model_training/assets/train.xlsx`

### Запуск обучения модели
----
```sh
cd model_training
python processing_data_frame.py
python train_model.py
```

- Обученная модель сохраняется в папке `model_training/model_out/`
- Для использования новой модели переместите файлы из папки `model_training/model_out/` в папку `model/`

### Запуск сервера
----
```sh
python server.py
```

### Отправка запросов к серверу
----
```sh
python send_request.py
```
- Подставьте необходимые тестовые данные в тестовый запрос из файла `send_request.py`
Тестовые данные для тестового запроса из файла `send_request.py` были взяты на основе первой строки входных данных из файла `model_training/out/processing_by_model.xlsx`
Убедитьесь в том, что вы вводите корректные данные на основе ниже предоставленных данных!


# Типы передаваемых данных
----

##### Когда (Смотрится по месяцам)
----
```python
class Months(Enum):
    JANUARY = 1
    FEBRUARY = 2
    MARCH = 3
    APRIL = 4
    MAY = 5
    JUNE = 6
    JULY = 7
    AUGUST = 8
    SEPTEMBER = 9
    OCTOBER = 10
    NOVEMBER = 11
    DECEMBER = 12
```

##### АвиаКомпании
----
```python
class Airlines(Enum):
    VISTARA = 1
    AIR_INDIA = 2
    INDIGO = 3
    GO_FIRST = 4
    AIRASIA = 5
    SPICEJET = 6
    STARAIR = 7
    TRUJET = 8
}
```

##### Откуда
----
```python
class Locations(Enum):
    DELHI = 11
    MUMBAI = 12
    KOLKATA = 13
    BANGALORE = 14
    HYDERABAD = 15
    CHENNAI = 16
```


##### Куда
----
```python
class Locations(Enum):
    DELHI = 11
    MUMBAI = 12
    KOLKATA = 13
    BANGALORE = 14
    HYDERABAD = 15
    CHENNAI = 16
}
```

##### Выбор Класса
----
```python
class FlightClasses(Enum):
    ECONOMY = 0
    BUSINESS = 1
```
