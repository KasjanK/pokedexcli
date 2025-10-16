# Pokedex REPL (Go)
A command-line Pokedex that lets you explore locations, catch Pokémon, and inspect your collection. Data is fetched from the public PokéAPI and cached locally for fast, offline-friendly lookups.

## Features
- Explore location areas and move between them
- Catch Pokémon and store their details locally
- Inspect caught Pokémon (name, height, weight, stats, types)
- Helpful REPL with built-in commands and help text
- Simple in-memory cache for API responses

## Commands
- `help` — Show available commands
- `map` — List the next page of location areas
- `mapb` — Go back to the previous page of location areas
- `explore <area>` — List Pokémon that can appear in the specified area
- `catch <pokemon>` — Try to catch a Pokémon by name
- `inspect <pokemon>` — Show details for a caught Pokémon
- `exit` — Quit the program

Example:

- `inspect pidgey` → “you have not caught that pokemon”
- `catch pidgey` → attempts capture and stores details on success
- `inspect pidgey` → prints full details (stats, height, weight, types)

## Tech Overview
- Language: Go
- API: https://pokeapi.co/
- Caching: internal/pokecache package (time-based cache)
- API Client: internal/pokeapi package (typed responses, pagination helpers)
- REPL: command_* files and repl.go orchestrate user input

## Getting Started
### 1. Prerequisites:
- Go 1.20+ (or your project’s specified version)

### 2. Install:
```
git clone <your-repo-url>
cd <repo>
go mod download
```

### 3. Run:
```
go run .
```

You’ll see a prompt like:

```
Pokedex >

Type `help` to list commands.
```

## Notes
- The inspect command uses locally stored Pokémon data captured via catch; it does not re-fetch from the API.
- Caching reduces redundant API calls. Cache TTL is managed in internal/pokecache.

## Testing
```
go test ./...
```

## Roadmap Ideas
- Persist caught Pokémon to disk
- Smarter catch mechanics (probabilities, items)
- Autocomplete for commands and Pokémon names
- Better error messages and input validation# pokedexcli
