package utils

import (
  "fmt"
  "crypto/sha256"
  "encoding/hex"
  "time"
  "strings"
  "strconv"
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

  return h * 60 + m
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
