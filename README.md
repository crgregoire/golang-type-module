# Satya

## The source of _truth_

## What is Satya

Satya is our module for our data models. It is the source of truth between services, such as ***Vijñāna*** and ***Buddha***. It is written in Golang and handles MySQL functionality through *GORM*, and provides data to requests routed through ***Buddha***.

## Developer Setup

### Before you start

You will need [Git](https://git-scm.com/), [MySQL 5.7 (or later)](https://dev.mysql.com/downloads/), and [Golang 1.12 or later](https://golang.org/dl/). It is also recommended that you use [GoLand](https://www.jetbrains.com/go/) or [Visual Studio Code](https://code.visualstudio.com) for development.

Once you have those tools installed you should clone the repo [here](https://github.com/tespo/satya.git).

### Setting up Satya

To begin, open up the IDE of your choice (preferably Visual Studio Code, as GoLand requires a module import in the settings when you clone from GitHub) and use the terminal to pull the code using:

``` git
git pull
```

This ensures you have the latest code available.

After pulling run `make init`.  This will setup a couple different tools we've configured for this repository.

To begin, get the files `docker.env`, `local.env`, and `docker-compose.yml` from one of your teammates. When you've received those, place the `docker.env & local.env` files into a folder named `config` in the root directory of your project.

For the database, which will store every model Satya creates, you must have an instance of MySQL running. Name the database `tespo_docker` (per the `docker.env`). You can do this by `docker-compose up -d mysqldb`. This requires the correct `docker-compose.yml` file in your project's root directory which you'll get from a teammate.

If you followed the directions to the letter, you should now direct your terminal instance to your working directory if you haven't already using:

``` bash
cd path/to/your/working/directory's/root
```

When you're in this directory you can use the command:

``` golang
go build
```

This builds Satya so you can run commands to seed the database, test, and nuke the database.
The commands are:

``` bash
./satya -type=nuke
```

This drops all tables from the database and removes all foreign key constraints. This command is to be run first, as it gives you a fresh database to work with.

``` bash
./satya -type=build
```

This runs both the migration and seeder functions, essentially taking the form of the next two commands. However if you wish you can also use:

``` bash
./satya -type=migrate
```

This runs the migrate functions, and:

``` bash
./satya -type=seed
```

Which runs the seeder function, seeding the database.

After building, seeding, nuking, et cetera, you may want to run tests. Luckily Satya comes with a testing suite right out of the box. You can run these tests by using the command:

``` golang
go test ./tests
```

All tests must pass before you can merge `your-branch` into `develop`, then `develop` into `staging`, then `staging` into `master`.
