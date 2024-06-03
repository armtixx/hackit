query_to_db() {
  mariadb -u root -h 127.0.0.1 -P 3307 -pAa12344321 < "$1" 
}

main() {
  docker-compose up -d
  query_to_db "create_db.sql"
  query_to_db "tables.sql"
  query_to_db "data.sql"
}
