# Mocker

Simple and lightweight tool to create and share mocked services for development or test environment.

## Usage
Mocker could be used through cli command or docker container.

### CLI
#### Installation
```bash
$ go get github.com/maykonlf/mocker/cmd/mocker
```

#### Flags
* `-addr`: mocker server API address (default ":8081")
* `-f`: mock service spec file path (default "./mocker.yaml")

#### Start mock
Go to a path containing the `mocker.yaml` spec file and run the following command on your terminal:
```bash
$ mocker
```

Then you can make requests to your mocked service and it will respond as configured on your `mocker.yaml` file.

### Docker
```bash
$ docker run -d -v .:/mocker -p 8081:8081 maykonlf/mocker
```

## Mock Specs - mocker.yaml
* `api`: describes mocked API specs
  * `routes`: list of matching api routes
  * `methods`: list of matching HTTP methods (must be in lowercase)
  * `response`: describes the response mocker should return when the received requests matches the specs above.
    * `status`: HTTP status code
    * `headers`: key-value response headers
    * `body`: response body string
    * `time`: time to wait before respond the request - useful to simulate API processing time (ex: `200ms`, `1s`)

### Example
```yaml
# mocker.yaml

api:
  - routes:
    - /v1/books
    - /v2/books
    methods:
    - get
    response:
      status: 200
      headers:
        Content-Type: application/json
      body: >
        [{
            "id: "7d619e9b-9fd2-4abb-b9b4-cfa5811e97a5",
            "title": "Clean Code: A Handbook of Agile Software Craftsmanship"
        }, {
            "id: "46c1bc9f-2ae0-4d73-bf9d-531fa1898537",
            "title": "Clean Architecture: A Craftsman's Guide to Software Structure and Design"
        }, {
            "id: "c8daf68a-5a28-4438-91c1-9f5835ad6439",
            "title": "Refactoring: Improving the Design of Existing Code"
        }]
      time: 200ms
```
