# todopher-web

## Overview
is a web-based todo management application with a Go backend and a dynamic
frontend powered by HTMX, CSS, and JavaScript. It allows users to create, edit,
and manage their to-do lists efficiently.

## Features
- Add, edit, and delete todos
- Search for todos
- Filter todos by status (e.g., completed, pending)
- Inline updates without page reloads using HTMX
- Responsive and clean UI with custom CSS styling
- Go backend using the Gin framework for fast and efficient request handling
- Uses MySQL for storing todos data

## Tech Stack
### Backend:
- [Go](https://golang.org/) (Gin framework)
- MySQL for data persistence

### Frontend:
- [HTMX](https://htmx.org/) for seamless interactivity
- CSS for styling
- JavaScript for additional client-side logic

## Installation 
### Prerequisites
- Go installed (1.18+ recommended)
- MySQL database

### Steps
1. Clone the repository: 
```bash 
    git clone https://github.com/nendix/todopher-web.git 
    cd todopher-web 
```
2. Run the application:
```bash
   docker compose up --build
```
3. Open your browser and navigate to `http://localhost:8080`.

## Contributing
Pull requests are welcome! For major changes, please open an issue first to
discuss proposed modifications.


## License
This project is licensed under the Apache License 2.0.

