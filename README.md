### Coding Test

- Deposit to specify user wallet
- Withdraw from specify user wallet
- Transfer from one user to another user
- Get specify user balance
- Get specify user transaction history

### Preparing the environment

- Postgres
- Redis

### Installation steps

Before use, PostgreSQL and Docker need to be prepared. Of course, you can also build binary programs directly without using Docker. Thank you

## Step 1: Create Db And Import Script

There are an SQL scripts in the scripts directory. Please create a database named app and import these two scripts into the database

## Step 2: Check Config

Please check the configuration information of the config file in the config directory and modify the configuration according to your actual situation

    db:
        host: 127.0.0.1
        user: postgres
        password: hs2024!@#
        database: app
    redis:
        addr: 192.168.31.106:6379
        password: ""
        db: 0

## Step 3: Create an image

    docker build -t app .

## Step 4: Run the image

    docker run -it -d -p 9090:9090 app

The forwarded port is 9090, you can adjust it according to the actual situation, and the default remains unchanged

### Some things that have been done

- [x] Completed golangci-lint verification
- [x] Provides a json script for the Postman interface
- [x] Provides Dockerfile image building, although it is a very simple deployment


### Some work that has not been implemented but is under consideration, I need to consider some time costs

- [ ] When adding or subtracting balance, a distributed lock is used and no renewal is done
- [ ] There is a data table t_transaction that may have performance problems in the later stage of the project. You can adapt to the actual scenario.
- [ ] No cache is used in querying balances and historical transactions, so consistency issues need to be considered, such as the initial time cost.


## Support

For support, email huangsheng7638@gmail.com or commit issue! Developer: Huang Sheng