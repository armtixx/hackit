export FLYEASY_HOST="127.0.0.1"
export FLYEASY_ML_PORT="8082"
export FLYEASY_FRONTEND_PORT="8080"

_base="$(dirname "$0")"
_base="$(realpath "$_base")"

main() {
  echo -n "[I]: Running backend: "
  cd "./src/backend/" && bash "start.sh" &
  echo "OK"

  echo -n "[I]: Running ML server: "
  cd "$_base" && bash "./src/ml/start.sh" &
  echo "OK"

  echo -n "[I]: Running frontend: "
  bash "./src/frontend/start.sh" &
  echo "OK"
  wait "$!"
}

main
