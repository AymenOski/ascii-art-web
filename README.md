# Ascii-Art Web

## Description
The **Ascii-Art Web** project is a web-based GUI for creating ASCII art banners using different styles like **shadow**, **standard**, and **thinkertoy**. The application allows users to input text, choose a banner style, and view the resulting ASCII art. The backend is powered by a Go HTTP server, while the frontend uses HTML forms to facilitate user interaction.

## Authors
- [Aymen Azizi , Github.com/AymenOski]

## Project Structure
```
├── BuildAndCleanup.sh
├── Dockerfile
├── README.md
├── functions
    ├── GenerateHandler.go
    ├── GeneratingTheAsciiArt.go
    ├── MainPage.go
    ├── Mapping.go
    ├── Printing.go
    ├── RenderErrPage.go
    ├── StyleHandler.go
    └── Templates.go
├── go.mod
├── main.go
├── shadow.txt
├── standard.txt
├── static
    ├── background.jpg
    ├── index-back-ground.jpg
    ├── page-not-found.css
    ├── styles-result.css
    └── styles.css
├── templates
    ├── index.html
    ├── page-not-found.html
    └── result.html
└── thinkertoy.txt
```
## Usage

To run this project locally, follow these steps:

1. **Install Go**: Ensure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).

2. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/ascii-art-web.git
   cd ascii-art-web
3. **Run the server**: 
   ```go
   go run .
3. **Open** your browser and go to: 
   ```go
   http://localhost:8020
## HTTP Endpoints

### 1. `GET /`
- **Purpose:** Serves the main HTML page.
- **Response:** HTML with input for text, banner selection (radio buttons), and submit button.
- **Status Codes:**
  - `200 OK` on success.
  - `404 Not Found` if the template is missing.

### 2. `POST /ascii-art`
- **Purpose:** Processes user input and generates ASCII art.
- **Input:** Text from the form and the selected banner.
- **Response:** Displays ASCII art result.
- **Status Codes:**
  - `200 OK` on success.
  - `400 Bad Request` for invalid input.
  - `404 Not Found` if banner is not found.
  - `500 Internal Server Error` for unhandled errors.

## Implementation Details:
1. **Main Page Handler (`GET /`):**
   - Serves an HTML page using Go templates.
   - Contains a form with:
     - Text input
     - Banner selection (radio buttons)
     - Submit button

2. **ASCII Art Handler (`POST /ascii-art`):**
   - Receives form data (text and banner type).
   - Validates input.
   - Reads the appropriate banner file (`shadow`, `standard`, `thinkertoy`).
   - Generates ASCII art using the algorithm from the original `ascii-art` project.
   - Displays the result in the template.

3. **Error Handling:**
   - Uses proper HTTP status codes.
   - Displays error messages in the web page.

4. **Templates:**
   - `index.html`: Main page with form.
   - `result.html`: Displays ASCII art results.

## ⚙️ Export Feature (ascii-art-web-export)

The project supports exporting the ASCII result directly as a **.txt** file.  
This is handled via the same HTML form by selecting the "Download"
- **Export Method**: Form-based
- **MIME Headers Used**:
  - `Content-Type`: `text/plain`
  - `Content-Length`: Size of the file
  - `Content-Disposition`: `attachment; filename="ascii-art.txt"`

> Exported files include read & write permissions for the user.
---