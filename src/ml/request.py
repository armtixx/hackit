import requests
from data_types import *
from config import HOST, PORT

date = Months.JANUARY
airline = Airlines.INDIGO
from_id = Locations.MUMBAI
to_id = Locations.DELHI
flight_class = FlightClasses.BUSINESS

data = {
    "date": date,
    "airline": airline,
	"from_id": from_id,
	"to_id": to_id,
	"flight_class": flight_class,
}
data = {k: v.value for k,v in data.items()}
print(data)

url = f"http://{HOST}:{PORT}/predict"

try:
    req = requests.get(url, params=data)
    if req.status_code == 200:
        price = req.json().get("price")
        print(f"Цена: {price}")
    else:
        print("Сервер вернул ошибку, проверьте данные")
except requests.exceptions.ConnectionError:
    print("Не удалось отправить запрос")
except Exception as e:
    print("Не удалось отправить запрос: ", e)
