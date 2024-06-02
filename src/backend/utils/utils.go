package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GenerateUserToken(email, passwordHash, salt string) string {
	seq := []byte(fmt.Sprintf("%v%v%v%v%v",
		salt, email, salt, passwordHash, salt,
	))
	sha := sha256.New()
	sha.Write([]byte(seq))
	hash := sha.Sum(nil)
	return hex.EncodeToString(hash)
}

func ParseTime(s string) int {
	tmp := strings.Replace(strings.Replace(s, "h", "", 1), "m", "", 1)

	hm := strings.Split(tmp, " ")
	h, errm := strconv.Atoi(hm[0])
	if errm != nil {
		h = 0
	}

	m, errm := strconv.Atoi(hm[1])
	if errm != nil {
		m = 0
	}

	return h*60 + m
}

func GetArriveTime(d, t string, timet string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	dateTimeStr := fmt.Sprintf("%s %s", d, t)

	dateTime, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		return time.Time{}, &time.ParseError{}
	}

	unixTime := dateTime.Unix()
	unixTime += int64(ParseTime(timet) * 60)
	return time.Unix(unixTime, 0).UTC(), nil
}

func CheckEmail(email string) bool {
	emRgx := `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emRgx, email)
	if match {
		return true
	} else {
		return false
	}
}

func CheckDate(email string) bool {
	dateRegex := `^(\d{4})[:-](\d{2})[:-](\d{2})$`
	dateStr := "2022-05-23"
	if match, _ := regexp.MatchString(dateRegex, dateStr); match {
		fmt.Println("Дата в формате 'yyyy-mm-dd' валидна:", dateStr)
	} else {
		fmt.Println("Дата в формате 'yyyy-mm-dd' невалидна:", dateStr)
	}

	// Проверка времени в формате "12h 30m"
	timeRegex := `^(\d{1,2})h\s(\d{1,2})m$`
	timeStr := "12h 30m"
	if match, _ := regexp.MatchString(timeRegex, timeStr); match {
		fmt.Println("Время в формате '12h 30m' валидно:", timeStr)
		return true
	} else {
		fmt.Println("Время в формате '12h 30m' невалидно:", timeStr)
		return false
	}
}
