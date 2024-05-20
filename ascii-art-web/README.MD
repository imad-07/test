# ASCII Art Web Application

## Description
This web application generates ASCII art based on user input. Users can enter a text string and select a banner style to generate corresponding ASCII art, which is displayed on the result page. The application is built using Go and serves HTML templates for the user interface.

## Authors
- Imad Tounsadi
- Amine Habchi

## Usage

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/ascii-art-webapp.git
   cd ascii-art-webapp
Install Go if you haven't already:
Download and install Go

Run the application:

bash
Copy code
go run main.go
Access the application:
Open your web browser and navigate to http://localhost:8080.

Using the Application
On the homepage, enter the text you want to convert to ASCII art.
Select a banner style from the options provided.
Submit the form to generate and display the ASCII art.
Implementation Details
Algorithm
Input Processing:

The user inputs text and selects a banner style.
The input text is cleaned by removing carriage return characters (\r).
Text Validation:

Each character in the input text is checked to ensure it is within the valid ASCII range (32 to 128).
Banner File Reading:

The banner file corresponding to the selected style is read.
The file content is processed to remove carriage return characters (\r).
Character Mapping:

The banner file content is parsed to map each ASCII character to its corresponding ASCII art representation.
Each character's representation consists of 8 lines of ASCII art.
ASCII Art Generation:

The input text is split into words.
For each character in each word, the corresponding ASCII art lines are concatenated to form the final ASCII art output.
Functions
indexHandler: Handles the requests to the homepage, serves the HTML template.
asciiArtHandler: Handles the form submission, processes input text and banner style, and generates ASCII art.
generate: Main function that coordinates the ASCII art generation process.
output: Constructs the final ASCII art output from the parsed characters.
split: Splits the input text into individual words for processing.
getLettres: Parses the banner file content into a mapping of ASCII characters to their ASCII art representations.

didikas lkhouya othman wnas hay l9ods wahed wahed
