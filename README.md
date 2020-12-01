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
* Card Number: 5527532727647849
* Pin: 2222
* CVV: 222
* Exp Date: 8/24/2028
 