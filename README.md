    ```markdown
    # Weather CLI

    A simple command-line interface (CLI) application to fetch and display weather information using the OpenWeatherMap API.

    ## Features

    - Fetches current weather data for a given city
    - Displays temperature, pressure, humidity, and weather description

    ## Getting Started

    ### Prerequisites

    - Go (version 1.16 or later)
    - An API key from [OpenWeatherMap](https://openweathermap.org/)

    ### Installation

    1. Clone the repository:
        ```sh
        git clone https://github.com/AustinKingOry/weather-cli.git
        cd weather-cli
        ```

    2. Initialize the Go module:
        ```sh
        go mod init weather-cli
        ```

    3. Replace `"YOUR_API_KEY_HERE"` in `main.go` with your actual API key from OpenWeatherMap.

    ### Running the Application

    1. Build and run the application:
        ```sh
        go run main.go "CityName"
        ```

        Replace `"CityName"` with the name of the city you want to fetch weather information for.

    ### Example Usage

    ```sh
    go run main.go "Nairobi"
    ```

    Example output:
    ```
    Weather in Nairobi:
    Temperature: 23.45°C
    Pressure: 1012 hPa
    Humidity: 80%
    Description: scattered clouds
    ```

    ## Contributing

    Contributions are welcome! Please fork the repository and create a pull request with your changes.

    ## License

    This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

    ## Acknowledgments

    - Thanks to the OpenWeatherMap community for their excellent API and documentation.
    ```