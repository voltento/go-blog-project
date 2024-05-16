
## Setup

### Docker
1. Build and run
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
    GIN_MODE=release go run cmd/blog/main.go --port=8080 --migration=<file_to_migration>
    ```

## API Endpoints and `curl` Examples

### Retrieve a specific post
- **Endpoint:** `GET /posts/{id}`
- **Curl Command:**
    ```sh
    curl -X GET http://localhost:8080/posts/1
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
- **Endpoint:** `POST /posts`
- **Curl Command:**
    ```sh
    curl -X POST http://localhost:8080/posts -H "Content-Type: application/json" -d '{"title":"New Post","content":"New Content","author":"New Author"}'
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
- **Endpoint:** `PUT /posts/{id}`
- **Curl Command:**
    ```sh
    curl -X PUT http://localhost:8080/posts/1 -H "Content-Type: application/json" -d '{"title":"Updated Post","content":"Updated Content","author":"Updated Author"}'
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
- **Endpoint:** `DELETE /posts/{id}`
- **Curl Command:**
    ```sh
    curl -X DELETE http://localhost:8080/posts/1
    ```
- **Response:** `204 No Content`


### Retrieve all posts
The Posts method returns all available posts. Pagination is not implemented to maintain the simplicity of the Storage. The task specifies using an in-memory data store, and for the scope of this technical test, pagination is omitted to focus on core CRUD operations and ensure straightforward data handling.

Implementing an optimal storage solution with pagination would require additional considerations for space and synchronization efficiency. It's better to reuse ready solutions like SQL databases rather than implementing from scratch.
- **Endpoint:** `GET /posts`
- **Curl Command:**
    ```sh
    curl -X GET http://localhost:8080/posts
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
        },
        ...
     ]
    }
    ```

## Running Tests
To run the tests, use the following command:
```sh
go test ./tests
