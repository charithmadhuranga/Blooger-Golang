#!/usr/bin/env bash
set -euo pipefail

BASE_URL="http://127.0.0.1:3000"
COOKIE_JAR="/tmp/blogger_cookies.txt"
USERNAME="user_$RANDOM"
PASSWORD="pass1234"

# Ensure server
if ! lsof -iTCP:3000 -sTCP:LISTEN >/dev/null 2>&1; then
  nohup go run ./cmd/server > /tmp/blogger_server.log 2>&1 &
  sleep 1
fi

echo "[1/7] Register user"
REG_CODE=$(curl -s -o /dev/null -w "%{http_code}" -c "$COOKIE_JAR" -X POST -d "username=$USERNAME" -d "password=$PASSWORD" "$BASE_URL/register")
echo "register: $REG_CODE"

echo "[2/7] Login"
LOGIN_CODE=$(curl -s -o /dev/null -w "%{http_code}" -c "$COOKIE_JAR" -b "$COOKIE_JAR" -X POST -d "username=$USERNAME" -d "password=$PASSWORD" "$BASE_URL/login")
echo "login: $LOGIN_CODE"

TOKEN=$(awk '($6=="token"){print $7}' "$COOKIE_JAR" | tail -n1)
echo "token_len: ${#TOKEN}"

echo "[3/7] Access protected root"
ROOT_CODE=$(curl -s -o /dev/null -w "%{http_code}" -b "$COOKIE_JAR" "$BASE_URL/")
echo "root: $ROOT_CODE"

echo "[4/7] Create post"
CREATE_CODE=$(curl -s -o /dev/null -w "%{http_code}" -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{"title":"Hello","content":"World"}' "$BASE_URL/api/posts")
echo "create: $CREATE_CODE"

echo "[5/7] List posts"
LIST_JSON=$(curl -s -H "Authorization: Bearer $TOKEN" "$BASE_URL/api/posts")
echo "list_len: ${#LIST_JSON}"
FIRST_ID=$(echo "$LIST_JSON" | sed -n 's/.*"id":\([0-9][0-9]*\).*/\1/p' | head -n1)
echo "id: $FIRST_ID"

echo "[6/7] Update post"
UPDATE_CODE=$(curl -s -o /dev/null -w "%{http_code}" -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -X PUT -d '{"title":"Updated","content":"Post"}' "$BASE_URL/api/posts/$FIRST_ID")
echo "update: $UPDATE_CODE"

echo "[7/7] Delete post"
DELETE_CODE=$(curl -s -o /dev/null -w "%{http_code}" -H "Authorization: Bearer $TOKEN" -X DELETE "$BASE_URL/api/posts/$FIRST_ID")
echo "delete: $DELETE_CODE"

# Success criteria summary
if [[ "$ROOT_CODE" == "200" && "$CREATE_CODE" == "201" && "$UPDATE_CODE" == "200" && "$DELETE_CODE" == "204" ]]; then
  echo "SMOKE: OK"
  exit 0
else
  echo "SMOKE: FAILED"
  exit 1
fi
