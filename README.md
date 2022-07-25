# golang-backend

My first experience with golang.

The project uses a CRUD app which lets you
## Operations for Users
- Create User - `POST /users`
```
curl localhost:8080/users -i -X POST -H 'Content-Type: application/json' -d '{
    "email": "test@example.com",
    "password": "12345",
    "name": "john doe",
    "age": 18
}'
```
- Get User `GET /users`
```
curl localhost:8080/users/{user_email}
```
- Update user - `PUT /users/{user_email}`
```
curl 'localhost:8080/users/test@example.com' -i -X PUT -H 'Content-Type: application/json' -d '{
    "email": "test@example.com",
    "password": "123456",
    "name": "john doe",
    "age": 18
}'
```
- Delete User - `DELETE /users/{user_email}`
```
curl -X DELETE localhost:8080/users/{user_email}
```
---

## Operations for Posts

- Create Post - `POST /posts`
```
curl localhost:8080/posts -i -X POST -H 'Content-Type: application/json' -d '{
    "userEmail": "test@example.com",
    "text": "say something"
}'
```
- Get Posts by a User `GET /posts/{user_email}`
```
curl localhost:8080/posts/{user_email}
```
- Delete Post - `DELETE /posts/{post_id}`
```
curl -X DELETE localhost:8080/posts/{user_email}
```
---










