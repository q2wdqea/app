### Coding Test

- Deposit to specify user wallet
- Withdraw from specify user wallet
- Transfer from one user to another user
- Get specify user balance
- Get specify user transaction history

### Installation steps

Before use, PostgreSQL and Docker need to be prepared. Of course, you can also build binary programs directly without using Docker. Thank you

## Step 1: Create Db And Import Script

There are two SQL scripts in the scripts directory. Please create a database named app and import these two scripts into the database

## Step 2: Check Config

Please check the configuration information of the config file in the config directory and modify the configuration according to your actual situation

    db:
      host: 127.0.0.1
      user: postgres
      password: hs2024!@#
      database: app

## Step 3: Create an image

    docker build -t app .

## Step 4: Run the image

    docker run -it -d -p 9090:9090 app

The forwarded port is 9090, you can adjust it according to the actual situation, and the default remains unchanged