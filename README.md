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

3. **Setup Database**
   ```sql
    CREATE DATABASE `go_products`;
    USE `go_products`;
    
    CREATE TABLE `categories` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

    CREATE TABLE `products` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `category_id` int NOT NULL,
    `stock` int NOT NULL,
    `description` text,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
   ```

4. **Install Air: (Recommended)**
    Air is a live reload tool for Go applications.
    ```bash
    go install github.com/air-verse/air@latest
    ```

5. **Run the application with Air:**
    ```bash
    air
    ```

    **If you didn't install Air:**
    ```bash
    go run .
    ```

6. **Access the application:**
    Open your browser and navigate to `http://localhost:8080` (or the port specified in your application).

7. **Stop the application:**
    Press `Ctrl+C` in the terminal where Air is running.