# ATM BANK API

***

TODO

* Add the middleware for different groups
* Add uuid id for models

***

# HOW TO RUN LOCALLY

* Have docker and docker-compose installed on your machine
* git clone & cd
* create .env file from .env.example 
* docker-compose up
* DB is running on localhost:3306 and Server is running on localhost:8080

.env example 
```js
MYSQL_DATABASE=bank
MYSQL_USER=bank-root
MYSQL_PASSWORD=atmSecret
MYSQL_ROOT_PASSWORD=atmSecret
SECRET=avufyts4srtfuygbu7r76r5fiygiu
```

If you want to populate DB you can use the routes or go to
*localhost:8080/init/data* to insert 1 account with 1 card and 1 transaction

***

Credentials for that /init/data account would then be
* Card Number: 5527532727647849 5527 5327 2764 7849
* Pin: 2222
* CVV: 222 	
* Exp Date: 8/24/2028
 
 OR
 
 * Card Number: 5507770244348111 5507 7702 4434 8111
 * Pin: 2222
 * CVV: 222
 * Exp Date: 8/24/2028
 
 After first time it's recommended to change the data in controllers/Init.go file
 The automated generation of data is yet to be done.
 
 ***
 
 To DROP tables use /init/clean but then you would have to restart docker-compose
 as tables are to be created(if there are none) only within the server set up time
 
 ***
 
 To run tests use *go test ./...*
   