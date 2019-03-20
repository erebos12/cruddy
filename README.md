# CRUDDY - CRUD REST Webservice for MongoDB

This webservice provides classical CRUD ops interface to write, read, update and delete JSON objects (here User-Objects) from/to MongoDB.

## Test Cruddy

- Clone it and Execute tests:

```
~/cruddy $ ./execute_tests.sh
Removing cruddy_mongodb_1 ... done
Removing network cruddy_default
mongodb uses an image, skipping
Building goapp_with_test
Step 1/15 : FROM golang:latest
 ---> d0e7a411e3da
...
...
[CRUDDY] - 2018/07/18 11:47:21 - INFO - getGollection - Successfully retrieved db name : 'test' and collection: 'user'
[GIN] 2018/07/18 - 11:47:21 | 200 |     401.582µs |                 | GET      /users?attribute=password&value=123
...
OK: 8 passed
--- PASS: TestGet (0.01s)
PASS
ok  	test/integration_test	0.049s
~/cruddy $
```

## Use Swagger UI

1. Start Cruddy locally:
```
~/cruddy $ docker-compose -f docker-compose-start.yml up
```

2. Open your browser with url 'http://localhost:8080/swagger/index.html'. You should see the Cruddy Swagger UI

![Alt text](cruddy_swagger_ui.png?raw=true "Screenshot Cruddy Swagger UI")
