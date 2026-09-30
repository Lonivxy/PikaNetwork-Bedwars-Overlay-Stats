# PikaStats Overlay

A standalone Windows BedWars stats overlay for **PikaNetwork** designed for **Lunar Client 1.8.x**.

It reads Minecraft/Lunar log output to detect the current TAB roster, fetches public PikaNetwork BedWars statistics, and displays them in a compact always-on-top overlay.

> **Unofficial project.** Not affiliated with, endorsed by, or maintained by PikaNetwork, Lunar Client, Mojang, or Microsoft.

## Features

- Detects BedWars players from the Minecraft/Lunar TAB-completion roster
- Shows:
  - A separate minigames-only Pika rank column (`Champion`, `Titan`, `Elite`, or `VIP`); other gamemode ranks are ignored
  - Level
  - FKDR
  - Kills
  - Final Kills
  - Beds Destroyed
  - Wins
- Shows the global leaderboard place beside Kills, Final Kills, Beds, and Wins whenever the API returns one, including places above #100
- Makes large places compact when fuzzy leaderboard data is enabled (`#1001` becomes `#1k+`)
- Clickable column sorting with DESC/ASC toggling, including Champion > Titan > Elite > VIP rank order
- A bounded, scrollable player list with draggable right/bottom edges for resizing
- Stat-based colors
- `[!]` threat highlighting
- `[ALT]` suspicious-stat-mismatch heuristic
- Small Minecraft player heads next to usernames when available
- Global `X` key to show/hide the overlay
- `-stats <player>` chat command support
- Background API fetching and caching
- Remembers overlay window position
- Debug logging for troubleshooting
- No Forge mod, DLL injection, account login, or Minecraft credentials required

## Requirements

- Windows 64-bit
- Lunar Client / Minecraft 1.8.x
- PikaNetwork BedWars
- Internet connection for PikaNetwork stats and player-head lookups

## Installation

1. Download `PikaStatsOverlay-3.11.0.exe`.
2. Run the EXE.
3. Start Lunar Client and join PikaNetwork BedWars.
4. In the BedWars waiting queue:
   - Focus the Minecraft chat input using your normal chat key
   - Type one space
   - Press `TAB`
   - Press `Esc`
5. The detected players and their stats should populate in the overlay.

Press **X** to show or hide the overlay.

## In-app settings

Click **Settings** in the overlay header to open the in-app settings screen. Changes save immediately.

- Show or hide: Level, FKDR, Kills, Final Kills, Beds, and Wins.
- Choose leaderboard interval: Lifetime, Monthly, or Weekly.
- Choose leaderboard mode: All Modes, Solo, Doubles, Triples, or Quads.
- Change the global show/hide shortcut by clicking **Set custom shortcut** and pressing the key you want.
- Toggle **Fuzzy LB data**. It is enabled by default; places through `#1000` stay exact and larger places use the `#12k+` form.

## Leaderboard places and settings

The overlay shows a player's global place immediately after **Kills**, **Final Kills**, **Beds**, and **Wins**, for example `9,999 #2`. The numeric place is retained even when it is outside the Top 100. With fuzzy data enabled, `#1001` is displayed as `#1k+`, while `#1` through `#1000` remain exact.

- `#1` is deep purple.
- `#2–#3` are bold dark red.
- `#4–#10` are red.
- `#11–#25` are orange-red.
- `#26–#50` are orange.
- `#51–#100` are gold.
- Places after `#100` stay visible; places outside the danger ladder use the neutral leaderboard color.

The default leaderboard selection is **Lifetime / All Modes**. It uses Pika's `total` interval, displayed in the overlay as **Lifetime**.

After running and closing the program once, edit this file while the overlay is closed to change the selection:

```text
%APPDATA%\PikaStatsOverlay\config.json
```

```json
{
  "leaderboardInterval": "total",
  "leaderboardMode": "ALL_MODES"
}
```

Supported values are:

- `leaderboardInterval`: `total` (Lifetime), `monthly`, or `weekly`
- `leaderboardMode`: `ALL_MODES`, `SOLO`, `DOUBLES`, `TRIPLES`, or `QUADS`

Only **BedWars** leaderboard placements are used. The overlay deliberately does not scan SkyPvP, Practice, Prison, Survival, or other non-minigame modes.

## Stat Colors

The colors are meant to make strong players easy to spot quickly.

| Stat | Gray | Green | Orange | Red | Purple |
|---|---:|---:|---:|---:|---:|
| **Level** | `< 20` | `20+` | `30+` | `50+` | `70+` |
| **FKDR** | `< 3` | `3+` | `5+` | `8+` | `15+` / INF |
| **Final Kills** | `< 1,000` | `1,000+` | `3,000+` | `5,000+` | `15,000+` |
| **Beds** | `< 500` | `500+` | `1,500+` | `3,000+` | `7,500+` |
| **Wins** | `< 500`* | `500+` | `1,000+` | `3,000+` | `5,000+` |

\*The UI keeps the lower wins range neutral/gray rather than inventing another tier.

## `[!]` Threat Highlight

A player is highlighted as a threat when **any** of these conditions is true:

- At least **one stat is purple**
- At least **3 stats are red or higher**
- At least **1 stat is red or higher** and at least **3 total stats are orange or higher**

Examples that will highlight:

- 1 purple + anything else
- 2 purple stats
- 3 red stats
- 1 red + 2 orange
- 2 red + 1 orange
- 1 green + 2 orange + 1 red

The `[!]` marker is only a quick stats-based warning. It does not mean the player is cheating.

## `[ALT]` Detection

`[ALT]` is a **heuristic**, not proof that an account is an alt.

The detector mainly looks for a mismatch such as:

- Very high FKDR while Finals / Beds / Wins are still gray or green
- High FKDR on an account with relatively low accumulated stats
- Very high efficiency over a small number of games
- Low level combined with unusually strong BedWars performance

For example, an account with roughly **8–10+ FKDR** but only a few hundred wins and otherwise low-volume stats may be flagged `[ALT]`.

## Automatic Ranking

The table automatically ranks players strongest → weakest until a header is clicked. Click a header once for descending order, then click it again for ascending order. Clickable columns are Player, Rank, Level, FKDR, Kills, Finals, Beds, and Wins. Scrolling is available when more than 12 players are shown.

The ranking considers:

- Level tier
- FKDR tier
- Final Kills tier
- Beds tier
- Wins tier
- `[!]` threat status
- `[ALT]` status

The chosen sort is for the current overlay session; a new TAB roster starts at the top of the list.

## `-stats` Chat Command

While Minecraft is focused, send:

```text
-stats PlayerName
```

The overlay looks up the player and sends a normal Minecraft chat response from your client, for example:

```text
PlayerName stats ; FKDR: 5.42, FINAL KILLS: 3,421, BEDS: 1,304, WINS: 824
```

Because this is a standalone application rather than an injected Minecraft mod, it uses normal Windows input to briefly open Minecraft chat, insert the finished response, and press Enter.

It does **not** need your Minecraft password or session token.

## Player Heads

The overlay can display a small Minecraft avatar next to a player's username.

Heads are downloaded in the background and cached in memory. If a username cannot be resolved to a skin/avatar, the stats row still works normally.

## How It Works

The overlay does **not** inject into Minecraft.

It:

1. Watches known Lunar Client / Minecraft `latest.log` locations.
2. Detects PikaNetwork / BedWars activity.
3. Uses the TAB-completion player roster to identify players.
4. Requests public player statistics from PikaNetwork's stats API.
5. Optionally requests a small avatar image for each username.
6. Draws the results in a native Windows overlay.

Public endpoints used by the application include:

```text
https://stats.pika-network.net/api/profile/{username}
https://stats.pika-network.net/api/profile/{username}/leaderboard?type=bedwars&interval=total&mode=ALL_MODES
https://mc-heads.net/avatar/{username}
```

## Privacy

The program does **not** contain or request:

- Minecraft passwords
- Lunar account passwords
- Microsoft account passwords
- Authentication cookies
- Session tokens
- API keys
- Payment information

The application does read local Minecraft/Lunar log files in order to detect the current username, server activity, and TAB roster.

### Debug log warning

For troubleshooting, the overlay creates:

```text
%LOCALAPPDATA%\PikaStatsOverlay\debug.log
```

The debug log can contain information such as:

- Your Windows user path
- Your Minecraft username
- Minecraft/Lunar log locations
- Player usernames detected in lobbies
- Stats/API diagnostic messages

**Do not commit or upload `debug.log` publicly unless you have checked/redacted it first.**

The debug log itself is **not included** in the normal source package or release ZIP.

The overlay also stores its window position in a per-user `PikaStatsOverlay/config.json`. This does not contain account credentials.

## Recommended `.gitignore`

If publishing the project on GitHub, use:

```gitignore
# Local logs / configuration
*.log
debug.log
config.json

# Build output
*.exe

# Backups / temporary files
*.bak
*.tmp

# Editor / OS files
.vscode/
.idea/
.DS_Store
Thumbs.db
```

If you want to distribute the EXE through GitHub Releases, remove `*.exe` from `.gitignore` or upload the executable directly to a Release instead of committing binaries to the source tree.

## Building From Source

The project is written in Go.

Example Windows GUI build:

```powershell
$env:GOOS="windows"
$env:GOARCH="amd64"
$env:CGO_ENABLED="0"
go build -ldflags="-H=windowsgui -s -w" -o PikaStatsOverlay-3.6.2.exe .
```

No third-party DLLs are required by the current native build.

## Repository Files

Recommended public repository:

```text
PikaStatsOverlay/
├── main.go
├── go.mod
├── README.md
├── LICENSE
└── .gitignore
```

Avoid publishing local logs, generated configuration, temporary backup files, or crash dumps.

## Troubleshooting

If player detection fails:

1. Confirm Lunar Client is running Minecraft 1.8.x.
2. Join a PikaNetwork BedWars waiting lobby.
3. Open chat with your normal chat key, type one space, press `TAB`, then `Esc`.
4. Check:

```text
%LOCALAPPDATA%\PikaStatsOverlay\debug.log
```

If sharing that file for support, remember that it may contain your Windows username/path and Minecraft username.

## Disclaimer

This project only presents publicly available game statistics and locally observed Minecraft roster information. Server rules can change, so users are responsible for checking whether overlays or external tools are permitted by the server they play on.

`[ALT]` and `[!]` are statistical heuristics/labels only and should not be treated as proof of cheating, account ownership, or rule violations.

## License

Choose a license before publishing if you want other people to have clear permission to use, modify, or redistribute the source.

For a simple open-source project, the **MIT License** is a common option.
