
## Setup

### Docker
1. Clone the repository:
    ```sh
    git clone https://github.com/voltento/go-blog-project
    cd go-blog-project
    ```
   
2. Build and run service in Docker
  ```sh
    docker build -t go-blog-app .
    docker run -p 8080:8080 go-blog-app
  ```

### Prerequisites
- Go installed on your machine (version 1.22 or higher).

### Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/voltento/go-blog-project
    cd go-blog-project
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    GIN_MODE=release go run cmd/blog/main.go --port=8080 --migration=resourses/blog_data.json
    ```

## API Endpoints and `curl` Examples

### Retrieve a specific post
- **Endpoint:** `GET /v1/posts/{id}`
- **Curl Command:**
    ```sh
    curl -X GET http://localhost:8080/v1/posts/1
    ```
- **Response:**
    ```json
    {
        "id": 1,
        "title": "Title 1",
        "content": "Content of the post",
        "author": "Author 1"
    }
    ```

### Create a new post
- **Endpoint:** `POST /v1/posts`
- **Curl Command:**
    ```sh
    curl -X POST http://localhost:8080/v1/posts -H "Content-Type: application/json" -d '{"title":"New Post","content":"New Content","author":"New Author"}'
    ```
- **Response:**
    ```json
    {
        "id": 6,
        "title": "New Post",
        "content": "New Content",
        "author": "New Author"
    }
    ```

### Update an existing post
- **Endpoint:** `PUT /v1/posts/{id}`
- **Curl Command:**
    ```sh
    curl -X PUT http://localhost:8080/v1/posts/1 -H "Content-Type: application/json" -d '{"title":"Updated Post","content":"Updated Content","author":"Updated Author"}'
    ```
- **Response:**
    ```json
    {
        "id": 1,
        "title": "Updated Post",
        "content": "Updated Content",
        "author": "Updated Author"
    }
    ```

### Delete a post
- **Endpoint:** `DELETE /v1/posts/{id}`
- **Curl Command:**
    ```sh
    curl -X DELETE http://localhost:8080/v1/posts/1
    ```
- **Response:** `204 No Content`


### Retrieve all posts
The Posts method returns all available posts. Pagination is not implemented to maintain the simplicity of the Storage. The task specifies using an in-memory data store, and for the scope of this technical test, pagination is omitted to focus on core CRUD operations and ensure straightforward data handling.

Implementing an optimal storage solution with pagination would require additional considerations for space and synchronization efficiency. It's better to reuse ready solutions like SQL databases rather than implementing from scratch.
- **Endpoint:** `GET /v1/posts`
- **Curl Command:**
    ```sh
    curl -X GET http://localhost:8080/v1/posts
    ```
- **Response:**
    ```json
    {
      "posts": [
        {
            "id": 1,
            "title": "Title 1",
            "content": "Content of the post",
            "author": "Author 1"
        },
        {
            "id": 2,
            "title": "Title 2",
            "content": "Content of the post",
            "author": "Author 2"
        }
     ]
    }
    ```

## Running Tests
To run the tests, use the following command:
```sh
go test ./tests
