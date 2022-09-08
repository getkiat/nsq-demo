Simple Demo, How to use NSQ as a Message Queue
---

# Run NSQ
  - Start NSQD, NSQLOOKUPD, and NSQADMIN
  ```shell
  docker-compose up
  ```

# Run Publisher
  - Start Backend API

    ```shell
    cd publisher/
    ```
    and

    ```shell
    go run main.go
    ```
## Test message
  - send message like tVhis
    ```curl
      curl -X POST \
      http://localhost:3000/api/food/send \
      -H 'cache-control: no-cache' \
      -H 'content-type: application/json' \
      -H 'postman-token: d8e4a3cc-fcad-2d07-29f7-ad2f47ba3e66' \
      -d '{
      	"from": "Weee",
        	"content":{
        		"header": "Makan Makan",
        		"body": "Makan nasi"
        	}
        }'
    ```
# Run Consumer
  - Start Consumer API
     ```shell
    cd consumer/
    ```
    and 
    ```shell
    go run main.go
    ```

    you'll see the result like this

    ```shell
    2018/02/04 15:49:01 INF    1 [grabfood/ch] (127.0.0.1:4150) connecting to nsqd
    2018/02/04 15:49:06 Got a message: {"from":"Weee","content":{"header":"Makan Makan","body":"Makan nasi"}}
    ```
