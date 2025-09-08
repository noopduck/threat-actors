# threat-actors

A Go-based API for querying and parsing MITRE ATT&CK threat actor (group) data. This project fetches group information from [attack.mitre.org](https://attack.mitre.org/groups/) and exposes endpoints for listing, searching, and retrieving details about threat groups.

## Features

- Fetches and parses MITRE ATT&CK threat group data
- Provides a REST API for:
	- Listing all threat groups
	- Searching threat groups by keyword
	- Retrieving detailed information for a specific group

## Project Structure

```
cmd/threatactors/         # Main entrypoint for the API server
internal/parser/          # HTML table parsing logic
internal/webclient/       # HTTP client for fetching MITRE data
go.mod, go.sum            # Go module files
```

## Development Setup

1. **Clone the repository:**
	 ```sh
	 git clone <repo-url>
	 cd threat-actors
	 ```

2. **Install dependencies:**
	 ```sh
	 go mod tidy
	 ```

3. **Run the API server:**
	 ```sh
	 go run ./cmd/threatactors/
	 ```

	 The server will start on `http://localhost:8080`.

## API Usage

### Endpoints

- `GET /`  
	Returns API usage information.

- `GET /mitreThreatGroups`  
	Returns a list of all MITRE threat groups.

- `GET /mitreThreatGroupDetails?group=G0000`  
	Returns details for a specific group (replace `G0000` with the desired group ID).

- `GET /mitreThreatGroupSearch?search=word`  
	Searches threat groups by keyword.

### Example Request

```sh
curl http://localhost:8080/mitreThreatGroups
```

## License

MIT License. See [`LICENSE`](LICENSE) for details.
