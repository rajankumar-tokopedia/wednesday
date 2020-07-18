# wednesday
How to setup.
1. `mkdir tokopdia && cd tokopedia` inside your workspace folder

2. `git clone https://github.com/rajankumar-tokopedia/wednesday` 

3. `cd wednesday`

4. `dep insure`

5. `go run server.go`

Health Check 
```js
http://localhost:5176/healthcheck
```
Response 
```json
// 20200718173108
// http://localhost:5176/healthcheck

{
  "header": {
    "process_time": "0.019261",
    "server_time": "1595073667",
    "status_code": 200,
    "message": ""
  },
  "data": {
    "ping": "pong"
  }
}
```
1. Get All Cabs NearBy
```shell script
curl --location --request POST 'localhost:5176/api/v1/cabs/near' \
--header 'Content-Type: application/json' \
--data-raw '{
	"lon":"3.44566",
	"lat":"1.44566"
}'
```
Sample Response 
```json
{
    "header": {
        "process_time": "0.071909",
        "server_time": "1595073088",
        "status_code": 200,
        "message": ""
    },
    "data": [
        {
            "id": 1,
            "number": "UP-14-DS-6246",
            "distance": 100,
            "unit": "m"
        },
        {
            "id": 2,
            "number": "UP-14-DS-6245",
            "distance": 0.5,
            "unit": "km"
        },
        {
            "id": 3,
            "number": "UP-14-DS-6243",
            "distance": 400,
            "unit": "m"
        }
    ]
}
```

2. Book a Cab
```shell script
curl --location --request POST 'localhost:5176/api/v1/cab/book' \
--header 'Content-Type: application/json' \
--data-raw '{
	"from":"A",
	"to":"B",
	"user_id":1
}'
```

Response 
```json
{
    "header": {
        "process_time": "0.25294",
        "server_time": "1595073829",
        "status_code": 200,
        "message": ""
    },
    "data": {
        "id": 3,
        "user_id": 1,
        "cab_id": 1,
        "Time": "2020-07-18T01:33:49.037316+05:30",
        "from": "K",
        "to": "J",
        "fare": 50,
        "currency": "INR",
        "state": "BOOKED"
    }
}
```

3. User Previous Bookings
```shell script
curl --location --request POST 'localhost:5176/api/v1/user/bookings' \
--header 'Content-Type: application/json' \
--data-raw '{
	"user_id":1,
	"limit":3
}'
```
Response
```json
{
    "header": {
        "process_time": "0.300683",
        "server_time": "1595072923",
        "status_code": 200,
        "message": ""
    },
    "data": [
        {
            "id": 1,
            "user_id": 1,
            "cab_id": 1,
            "Time": "2020-07-18T05:18:43.047361+05:30",
            "from": "A",
            "to": "B",
            "fare": 10,
            "currency": "INR",
            "state": "COMPLETED"
        },
        {
            "id": 2,
            "user_id": 1,
            "cab_id": 2,
            "Time": "2020-07-18T02:18:43.047361+05:30",
            "from": "K",
            "to": "A",
            "fare": 30,
            "currency": "INR",
            "state": "COMPLETED"
        },
        {
            "id": 3,
            "user_id": 1,
            "cab_id": 1,
            "Time": "2020-07-18T01:18:43.047361+05:30",
            "from": "K",
            "to": "J",
            "fare": 50,
            "currency": "INR",
            "state": "FAILED"
        }
    ]
}
```
Author : Rajan Kumar

If you have any query.
```
Mob : +91-95401-52552
Email : rajankumar549@gmail.com
```