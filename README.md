# Fp_Go_Web

Link Demo FP:
[Youtube](https://www.youtube.com/live/IvIPQFeuOug?si=3cPXopVnB9peyNhB&t=75)

##How To Run:
1. **Clone the Repository:**
    ```bash
        git clone https://github.com/Achmadfajri10/Fp_Go_Web
        cd Fp_Go_Web
    ```

2. **Install Dependencies:**
    Make sure you have Go installed. If not, download and install it from [here](https://golang.org/dl/), and don't forget to run
    ```bash
        go mod download
    ```

3. **Make Sure MYSQL is turned on**

4. **Init Setup**
    this step is used to generate a .env file for the user when it didnt detect a .env file, when it's done you can use `Ctrl + C` to stop the program 
    ```bash
        go run .
    ```

5. **Install Air: (Optional)**
    Air is a live reload tool for Go applications. You can skip this step if you just want to run the program, but it's recommended to use while updating the code as it will help you reload the go program when a change is detected.
    ```bash
        go install github.com/air-verse/air@latest

        # Wait until the program is installed
        air
    ```

6. **Access the application:**
    Open your browser and navigate to `http://localhost:8080` (or the port specified in your application).

7. **Stop the application:**
    Press `Ctrl+C` in the terminal where Air/Go is running.