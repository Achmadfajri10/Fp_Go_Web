# Fp_Go_Web

##How To Run:
1. **Clone the repository:**
    ```bash
    git clone https://github.com/Achmadfajri10/Fp_Go_Web
    cd Fp_Go_Web
    ```

2. **Install dependencies:**
    Make sure you have Go installed. If not, download and install it from [here](https://golang.org/dl/), and don't forget to run 
    ```bash
    go mod download
    ```

3. **Install Air: (Recommended)**
    Air is a live reload tool for Go applications.
    ```bash
    go install github.com/air-verse/air@latest
    ```

4. **Run the application with Air:**
    ```bash
    air
    ```

    **If you didn't install Air:**
    ```bash
    go run .
    ```

5. **Access the application:**
    Open your browser and navigate to `http://localhost:8080` (or the port specified in your application).

6. **Stop the application:**
    Press `Ctrl+C` in the terminal where Air is running.