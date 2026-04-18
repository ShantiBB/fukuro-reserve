# Gateway HTTP smoke in GoLand

Files:
- `scripts/gateway_smoke/00_setup.http`
- `scripts/gateway_smoke/10_auth.http`
- `scripts/gateway_smoke/20_hotels.http`
- `scripts/gateway_smoke/30_rooms.http`
- `scripts/gateway_smoke/40_bookings.http`
- `scripts/gateway_smoke/90_cleanup.http`

Variable setup (important):
- HTTP Client variables in this project come from:
  - `scripts/http-client.env.json` (public defaults)
  - `scripts/http-client.private.env.json` (private secrets)
- Create private file from example:
  - `cp scripts/http-client.private.env.json.example scripts/http-client.private.env.json`
- Set valid admin credentials in `scripts/http-client.private.env.json` for `local`.

Variable priority in JetBrains HTTP Client:
- `@var` in `.http` has highest priority.
- `client.global.set("var", "...")` overrides environment values.
- Environment (`scripts/http-client.env.json` / private env) is used when variable is not defined in-file.

Why previous `unsubstituted variable` happened:
- `adminEmail/adminPassword` were not provided by environment.
- `adminAccess/ownerId` are runtime variables created by previous requests, so running a middle request alone can fail if setup requests were not executed.

What it does:
- Runs a full end-to-end smoke flow for gateway endpoints split by domain.
- Includes status assertions after each request.
- Reuses IDs/tokens between requests (owner, managed user, hotel, room, booking).

How to run:
1. Start all 4 services and DB.
2. Select environment `local` in GoLand HTTP Client.
3. Open files in `scripts/gateway_smoke/` in GoLand.
4. Run files in order: `00_setup -> 10_auth -> 20_hotels -> 30_rooms -> 40_bookings -> 90_cleanup`.

CLI run with `ijhttp`:
```bash
ijhttp \
  scripts/gateway_smoke/00_setup.http \
  scripts/gateway_smoke/10_auth.http \
  scripts/gateway_smoke/20_hotels.http \
  scripts/gateway_smoke/30_rooms.http \
  scripts/gateway_smoke/40_bookings.http \
  scripts/gateway_smoke/90_cleanup.http \
  -e local \
  -v scripts/http-client.env.json \
  -p scripts/http-client.private.env.json \
  -L BASIC \
  -r /tmp/ijhttp-report-env
```

Notes:
- The file generates unique test emails and hotel titles per run.
- It cleans up created entities at the end.
- It does not delete the admin account used for login.
