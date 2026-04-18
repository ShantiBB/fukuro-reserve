# Gateway HTTP smoke in GoLand

File: `scripts/gateway_smoke.http`

Variable setup (important):
- HTTP Client variables in this project come from:
  - `http-client.env.json` (public defaults)
  - `http-client.private.env.json` (private secrets)
- Create private file from example:
  - `cp http-client.private.env.json.example http-client.private.env.json`
- Set valid admin credentials in `http-client.private.env.json` for `local`.

Variable priority in JetBrains HTTP Client:
- `@var` in `.http` has highest priority.
- `client.global.set("var", "...")` overrides environment values.
- Environment (`http-client.env.json` / private env) is used when variable is not defined in-file.

Why previous `unsubstituted variable` happened:
- `adminEmail/adminPassword` were not provided by environment.
- `adminAccess/ownerId` are runtime variables created by previous requests, so running a middle request alone can fail if setup requests were not executed.

What it does:
- Runs a full end-to-end smoke flow for gateway endpoints.
- Includes status assertions after each request.
- Reuses IDs/tokens between requests (owner, managed user, hotel, room, booking).

How to run:
1. Start all 4 services and DB.
2. Select environment `local` in GoLand HTTP Client.
3. Open `scripts/gateway_smoke.http` in GoLand.
4. Run from the first request (`Run All Requests in File`).

Notes:
- The file generates unique test emails and hotel titles per run.
- It cleans up created entities at the end.
- It does not delete the admin account used for login.
