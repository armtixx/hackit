from enum import Enum

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

class Airlines(Enum):
    VISTARA = 1
    AIR_INDIA = 2
    INDIGO = 3
    GO_FIRST = 4
    AIRASIA = 5
    SPICEJET = 6
    STARAIR = 7
    TRUJET = 8

class Locations(Enum):
    DELHI = 11
    MUMBAI = 12
    KOLKATA = 13
    BANGALORE = 14
    HYDERABAD = 15
    CHENNAI = 16

class FlightClasses(Enum):
    ECONOMY = 0
    BUSINESS = 1
