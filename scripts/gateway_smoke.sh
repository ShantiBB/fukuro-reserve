#!/usr/bin/env bash

set -euo pipefail

BASE_URL="${BASE_URL:-http://localhost:8080}"
DB_CONTAINER="${DB_CONTAINER:-postgres}"
PASS="${PASS:-Passw0rd!123}"

if ! command -v jq >/dev/null 2>&1; then
  echo "jq is required"
  exit 1
fi

if ! command -v python3 >/dev/null 2>&1; then
  echo "python3 is required"
  exit 1
fi

if ! command -v docker >/dev/null 2>&1; then
  echo "docker is required"
  exit 1
fi

TS="$(date +%s)"
OWNER_EMAIL="owner${TS}@example.com"
ADMIN_EMAIL="admin${TS}@example.com"
MANAGED_EMAIL="managed${TS}@example.com"

OWNER_ID=""
ADMIN_ID=""
MANAGED_ID=""
HOTEL_ID=""
HOTEL_SLUG=""
RENAMED_HOTEL_SLUG=""
ROOM_ID=""
BOOKING_ID=""

OWNER_ACCESS=""
OWNER_REFRESH=""
ADMIN_ACCESS=""

register() {
  local email="$1"
  curl -sS -X POST "${BASE_URL}/api/v1/auth/register" \
    -H 'Content-Type: application/json' \
    -d "{\"email\":\"${email}\",\"password\":\"${PASS}\"}"
}

login() {
  local email="$1"
  curl -sS -X POST "${BASE_URL}/api/v1/auth/login" \
    -H 'Content-Type: application/json' \
    -d "{\"email\":\"${email}\",\"password\":\"${PASS}\"}"
}

decode_sub() {
  python3 - "$1" <<'PY'
import base64
import json
import sys

payload = sys.argv[1].split(".")[1]
payload += "=" * (-len(payload) % 4)
print(json.loads(base64.urlsafe_b64decode(payload))["Sub"])
PY
}

expect_status() {
  local name="$1"
  local expected="$2"
  shift 2

  local body_file
  body_file="$(mktemp)"

  local code
  code="$(curl -sS -o "${body_file}" -w "%{http_code}" "$@")"

  printf "=== %s ===\nSTATUS: %s\n" "${name}" "${code}"
  cat "${body_file}"
  printf "\n\n"

  if [[ "${code}" != "${expected}" ]]; then
    echo "Expected HTTP ${expected}, got ${code} for ${name}" >&2
    rm -f "${body_file}"
    exit 1
  fi

  rm -f "${body_file}"
}

cleanup() {
  set +e

  if [[ -n "${BOOKING_ID}" ]]; then
    curl -sS -o /dev/null -X DELETE "${BASE_URL}/api/v1/bookings/${BOOKING_ID}" \
      -H "Authorization: Bearer ${OWNER_ACCESS}" || true
  fi

  if [[ -n "${ROOM_ID}" ]]; then
    curl -sS -o /dev/null -X DELETE "${BASE_URL}/api/v1/rooms/${ROOM_ID}" \
      -H "Authorization: Bearer ${OWNER_ACCESS}" || true
  fi

  if [[ -n "${RENAMED_HOTEL_SLUG}" ]]; then
    curl -sS -o /dev/null -X DELETE "${BASE_URL}/api/v1/hotels/jp/tokyo/${RENAMED_HOTEL_SLUG}" \
      -H "Authorization: Bearer ${OWNER_ACCESS}" || true
  elif [[ -n "${HOTEL_SLUG}" ]]; then
    curl -sS -o /dev/null -X DELETE "${BASE_URL}/api/v1/hotels/jp/tokyo/${HOTEL_SLUG}" \
      -H "Authorization: Bearer ${OWNER_ACCESS}" || true
  fi

  if [[ -n "${OWNER_ID}" ]]; then
    curl -sS -o /dev/null -X DELETE "${BASE_URL}/api/v1/auth/users/${OWNER_ID}" \
      -H "Authorization: Bearer ${OWNER_ACCESS}" || true
  fi

  if [[ -n "${MANAGED_ID}" && -n "${ADMIN_ACCESS}" ]]; then
    curl -sS -o /dev/null -X DELETE "${BASE_URL}/api/v1/auth/users/${MANAGED_ID}" \
      -H "Authorization: Bearer ${ADMIN_ACCESS}" || true
  fi

  if [[ -n "${ADMIN_ID}" && -n "${ADMIN_ACCESS}" ]]; then
    curl -sS -o /dev/null -X DELETE "${BASE_URL}/api/v1/auth/users/${ADMIN_ID}" \
      -H "Authorization: Bearer ${ADMIN_ACCESS}" || true
  fi
}

trap cleanup EXIT

expect_status "health" 200 "${BASE_URL}/health"

OWNER_REGISTER="$(register "${OWNER_EMAIL}")"
OWNER_ACCESS="$(printf '%s' "${OWNER_REGISTER}" | jq -r '.access')"
OWNER_REFRESH="$(printf '%s' "${OWNER_REGISTER}" | jq -r '.refresh')"
OWNER_ID="$(decode_sub "${OWNER_ACCESS}")"

ADMIN_REGISTER="$(register "${ADMIN_EMAIL}")"
ADMIN_BOOTSTRAP_ACCESS="$(printf '%s' "${ADMIN_REGISTER}" | jq -r '.access')"
ADMIN_ID="$(decode_sub "${ADMIN_BOOTSTRAP_ACCESS}")"

docker exec "${DB_CONTAINER}" psql -U postgres -d auth \
  -c "UPDATE users SET role='USER_ROLE_ADMIN' WHERE email='${ADMIN_EMAIL}';" >/dev/null

ADMIN_LOGIN="$(login "${ADMIN_EMAIL}")"
ADMIN_ACCESS="$(printf '%s' "${ADMIN_LOGIN}" | jq -r '.access')"

expect_status "auth login owner" 200 \
  -X POST "${BASE_URL}/api/v1/auth/login" \
  -H 'Content-Type: application/json' \
  -d "{\"email\":\"${OWNER_EMAIL}\",\"password\":\"${PASS}\"}"

expect_status "auth refresh owner" 200 \
  -X POST "${BASE_URL}/api/v1/auth/refresh" \
  -H 'Content-Type: application/json' \
  -d "{\"refresh_token\":\"${OWNER_REFRESH}\"}"

expect_status "users list as owner expect403" 403 \
  "${BASE_URL}/api/v1/auth/users?page=1&limit=10" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "users list as admin" 200 \
  "${BASE_URL}/api/v1/auth/users?page=1&limit=10" \
  -H "Authorization: Bearer ${ADMIN_ACCESS}"

expect_status "user get self owner" 200 \
  "${BASE_URL}/api/v1/auth/users/${OWNER_ID}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "user update self owner" 200 \
  -X PUT "${BASE_URL}/api/v1/auth/users/${OWNER_ID}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}" \
  -H 'Content-Type: application/json' \
  -d "{\"email\":\"${OWNER_EMAIL}\",\"username\":\"owner-${TS}\"}"

MANAGED_CREATE="$(curl -sS -X POST "${BASE_URL}/api/v1/auth/users" \
  -H "Authorization: Bearer ${ADMIN_ACCESS}" \
  -H 'Content-Type: application/json' \
  -d "{\"email\":\"${MANAGED_EMAIL}\",\"username\":\"managed-${TS}\",\"password\":\"${PASS}\"}")"
MANAGED_ID="$(printf '%s' "${MANAGED_CREATE}" | jq -r '.id')"
printf "=== user create as admin ===\nSTATUS: 201\n%s\n\n" "${MANAGED_CREATE}"
if [[ "${MANAGED_ID}" == "null" || -z "${MANAGED_ID}" ]]; then
  echo "Failed to create managed user" >&2
  exit 1
fi

expect_status "user activity as admin" 200 \
  -X PATCH "${BASE_URL}/api/v1/auth/users/${OWNER_ID}/activity" \
  -H "Authorization: Bearer ${ADMIN_ACCESS}" \
  -H 'Content-Type: application/json' \
  -d '{"is_active":true}'

expect_status "user role as admin" 200 \
  -X PATCH "${BASE_URL}/api/v1/auth/users/${OWNER_ID}/role" \
  -H "Authorization: Bearer ${ADMIN_ACCESS}" \
  -H 'Content-Type: application/json' \
  -d '{"role":"USER_ROLE_MODERATOR"}'

HOTEL_CREATE="$(curl -sS -X POST "${BASE_URL}/api/v1/hotels" \
  -H "Authorization: Bearer ${OWNER_ACCESS}" \
  -H 'Content-Type: application/json' \
  -d '{"country_code":"jp","city_slug":"tokyo","title":"Gateway Smoke Hotel","owner_id":1,"description":"gateway smoke hotel","address":"1 Test St, Tokyo","location":{"latitude":35.66,"longitude":139.70}}')"
HOTEL_ID="$(printf '%s' "${HOTEL_CREATE}" | jq -r '.id')"
HOTEL_SLUG="$(printf '%s' "${HOTEL_CREATE}" | jq -r '.hotel_slug')"
printf "=== hotel create ===\nSTATUS: 201\n%s\n\n" "${HOTEL_CREATE}"
if [[ "${HOTEL_ID}" == "null" || -z "${HOTEL_ID}" ]]; then
  echo "Failed to create hotel" >&2
  exit 1
fi

expect_status "hotel get" 200 \
  "${BASE_URL}/api/v1/hotels/jp/tokyo/${HOTEL_SLUG}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "hotel list" 200 \
  "${BASE_URL}/api/v1/hotels?country_code=jp&city_slug=tokyo&sort_by=title&page=1&limit=10" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "hotel update" 200 \
  -X PUT "${BASE_URL}/api/v1/hotels/jp/tokyo/${HOTEL_SLUG}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}" \
  -H 'Content-Type: application/json' \
  -d '{"description":"updated hotel","address":"2 Test St, Tokyo","location":{"latitude":35.67,"longitude":139.71}}'

expect_status "hotel update title" 200 \
  -X PATCH "${BASE_URL}/api/v1/hotels/jp/tokyo/${HOTEL_SLUG}/title" \
  -H "Authorization: Bearer ${OWNER_ACCESS}" \
  -H 'Content-Type: application/json' \
  -d '{"title":"Gateway Smoke Hotel Renamed"}'

RENAMED_HOTEL_SLUG="gateway-smoke-hotel-renamed"

expect_status "hotel get after title update" 200 \
  "${BASE_URL}/api/v1/hotels/jp/tokyo/${RENAMED_HOTEL_SLUG}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

ROOM_CREATE="$(curl -sS -X POST "${BASE_URL}/api/v1/rooms" \
  -H "Authorization: Bearer ${OWNER_ACCESS}" \
  -H 'Content-Type: application/json' \
  -d '{"country_code":"jp","city_slug":"tokyo","hotel_slug":"gateway-smoke-hotel-renamed","title":"Room 101","description":"nice room","room_number":"101","type":"ROOM_TYPE_SINGLE","price":"150.00","capacity":2,"area_sqm":18.5,"floor":1,"amenities":["wifi"],"images":["https://example.com/room-101.jpg"]}')"
ROOM_ID="$(printf '%s' "${ROOM_CREATE}" | jq -r '.id')"
printf "=== room create ===\nSTATUS: 201\n%s\n\n" "${ROOM_CREATE}"
if [[ "${ROOM_ID}" == "null" || -z "${ROOM_ID}" ]]; then
  echo "Failed to create room" >&2
  exit 1
fi

expect_status "rooms list" 200 \
  "${BASE_URL}/api/v1/rooms?country_code=jp&city_slug=tokyo&hotel_slug=gateway-smoke-hotel-renamed&page=1&limit=10" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "room get" 200 \
  "${BASE_URL}/api/v1/rooms/${ROOM_ID}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "room update" 200 \
  -X PUT "${BASE_URL}/api/v1/rooms/${ROOM_ID}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}" \
  -H 'Content-Type: application/json' \
  -d '{"title":"Room 101 Updated","room_number":"101","type":"ROOM_TYPE_DOUBLE","description":"updated room","price":"175.00","capacity":3,"area_sqm":20.0,"floor":1,"amenities":["wifi","tv"],"images":["https://example.com/room-101-updated.jpg"]}'

expect_status "room status update" 200 \
  -X PATCH "${BASE_URL}/api/v1/rooms/${ROOM_ID}/status" \
  -H "Authorization: Bearer ${OWNER_ACCESS}" \
  -H 'Content-Type: application/json' \
  -d '{"status":"ROOM_STATUS_MAINTENANCE"}'

expect_status "room get after status" 200 \
  "${BASE_URL}/api/v1/rooms/${ROOM_ID}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

BOOKING_CREATE="$(curl -sS -X POST "${BASE_URL}/api/v1/bookings" \
  -H "Authorization: Bearer ${OWNER_ACCESS}" \
  -H 'Content-Type: application/json' \
  -d "{\"user_id\":${OWNER_ID},\"hotel_id\":\"${HOTEL_ID}\",\"check_in\":\"2026-06-01T12:00:00Z\",\"check_out\":\"2026-06-03T12:00:00Z\",\"guest_name\":\"Gateway Booker\",\"guest_email\":\"${OWNER_EMAIL}\",\"guest_phone\":\"+79990000003\",\"currency\":\"USD\",\"expected_total_amount\":\"350.00\",\"rooms\":[{\"room_id\":\"${ROOM_ID}\",\"adults\":2,\"children\":1,\"price_per_night\":\"175.00\"}]}")"
BOOKING_ID="$(printf '%s' "${BOOKING_CREATE}" | jq -r '.id')"
printf "=== booking create ===\nSTATUS: 201\n%s\n\n" "${BOOKING_CREATE}"
if [[ "${BOOKING_ID}" == "null" || -z "${BOOKING_ID}" ]]; then
  echo "Failed to create booking" >&2
  exit 1
fi

expect_status "bookings list all" 200 \
  "${BASE_URL}/api/v1/bookings?page=1&limit=10" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "bookings list user" 200 \
  "${BASE_URL}/api/v1/bookings?userId=${OWNER_ID}&page=1&limit=10" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "bookings list user+hotel" 200 \
  "${BASE_URL}/api/v1/bookings?userId=${OWNER_ID}&hotelId=${HOTEL_ID}&page=1&limit=10" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "booking get" 200 \
  "${BASE_URL}/api/v1/bookings/${BOOKING_ID}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "booking confirm" 200 \
  -X PATCH "${BASE_URL}/api/v1/bookings/${BOOKING_ID}/confirm" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "booking cancel" 200 \
  -X PATCH "${BASE_URL}/api/v1/bookings/${BOOKING_ID}/cancel" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"

expect_status "booking delete" 204 \
  -X DELETE "${BASE_URL}/api/v1/bookings/${BOOKING_ID}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"
BOOKING_ID=""

expect_status "room delete" 204 \
  -X DELETE "${BASE_URL}/api/v1/rooms/${ROOM_ID}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"
ROOM_ID=""

expect_status "hotel delete" 204 \
  -X DELETE "${BASE_URL}/api/v1/hotels/jp/tokyo/${RENAMED_HOTEL_SLUG}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"
RENAMED_HOTEL_SLUG=""
HOTEL_SLUG=""
HOTEL_ID=""

expect_status "user delete self owner" 204 \
  -X DELETE "${BASE_URL}/api/v1/auth/users/${OWNER_ID}" \
  -H "Authorization: Bearer ${OWNER_ACCESS}"
OWNER_ID=""

expect_status "managed user delete as admin" 204 \
  -X DELETE "${BASE_URL}/api/v1/auth/users/${MANAGED_ID}" \
  -H "Authorization: Bearer ${ADMIN_ACCESS}"
MANAGED_ID=""

expect_status "admin delete self" 204 \
  -X DELETE "${BASE_URL}/api/v1/auth/users/${ADMIN_ID}" \
  -H "Authorization: Bearer ${ADMIN_ACCESS}"
ADMIN_ID=""

echo "Gateway smoke test passed"
