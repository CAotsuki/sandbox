#!/bin/zsh
source .env
Host=localhost
Port=${API_PORT}

function waitStdin() {
  echo ""
  read 'KEY?press enter key: '
}

# echo "home page."
# curl -i ${Host}:${Port}/; waitStdin

echo "GET"
curl -i ${Host}:${Port}/todos/; waitStdin

echo "POST"
curl -i -X POST -H "Content-Type: application/json" -d '{"title":"test", "content":"テストです。"}' ${Host}:${Port}/todos/
curl -i ${Host}:${Port}/todos/; waitStdin

echo "PUT"
curl -i -X PUT -H "Content-Type: application/json" -d '{"title":"test", "content":"updateテスト"}' ${Host}:${Port}/todos/3
curl -i ${Host}:${Port}/todos/; waitStdin

echo "DELETE"
curl -i -X DELETE ${Host}:${Port}/todos/3
curl -i ${Host}:${Port}/todos/; waitStdin
