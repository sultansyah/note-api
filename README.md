## Setup Instructions

### **1. Copy `.env.example` to `.env`:**

   To set up the environment variables, create a `.env` file by copying the provided `.env.example` file. Run the following command:

   ```bash
   cp .env.example .env
   ```

### **2. Run Database Migration for MySQL:**

   To run the database migration for MySQL, use the following command:

   ```bash
   migrate -database "mysql://username:password@tcp(localhost:3306)/database_name" -path database/migrations up
   ```