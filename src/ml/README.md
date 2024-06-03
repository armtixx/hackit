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
- [Изменение данных](#изменение-данных)
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
pip install requirements.txt
```

## Запуск
----
1. Входные тестовые данные подаются в файле "model_training/assets/train.xlsx"
2. Запустите файл "model_training/processing_data_frame.py"
3. Запустите файл "model_training/model.py"
4. Запустите файл "server.py"
5. Запустите файл "request.py"

# Изменение данных
----
- Подставьте необходимые тестовые данные в тестовый запрос из файла "request.py"
Тестовые данные для тестового запроса из файла "request.py" были взяты на основе первой строки входных данных из файла "./model_training/out/processing_by_model.xlsx"
Убедитьесь в том, что вы вводите корректные данные на основе ниже предоставленных данных!

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
