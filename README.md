## How to run?
First, clone the repository:
```sh
git clone git@github.com:parsaakbari1209/crud-echo.git
```
Next, change the current directory to the repository:
```sh
cd crud-echo
```
Next, install the dependencies:
```sh
go get ./...
```
Finally, run the app on port `8000`:
```sh
go run .
```
## Endpoints:
```sh
GET    /users/:email
POST   /users
PUT    /users/:email
DELETE /users/:email
```
### Get User
This endpoint retrieves a user given the email.  
Send a `GET` request to `/users/:email`:
```sh
curl -X GET 'http://127.0.0.1:8000/users/bob@gmail.com'
```
Response:
```sh
{
    "name": "Bob",
    "email": "bob@gmail.com",
    "password": "secure"
}
```
### Create User
This endpoint creates a user given the information.
Send a `POST` request to `/users`:
```sh
curl -X POST 'http://127.0.0.1:8000/users' -H "Content-Type: application/json" -d '{"name": "Bob", "email": "bob@gmail.com", "password": "secure"}'
```
Response:  
```sh
{
    "name": "Bob",
    "email": "bob@gmail.com",
    "password": "secure"
}
```
### Update User
This endpoint updates the provided fields for the given user email.  
Send a `PUT` request to `/users/:email`:
```sh
curl -X PUT 'http://127.0.0.1:8000/users/bob@gmail.com' -H "Content-Type: application/json" -d '{"password": "secure_v2"}'
```
Response:
```sh
{
    "name": "Bob",
    "email": "bob@gmail.com",
    "password": "secure_v2"
}
```
### Delete User
This endpoint deletes the user given the email.  
Send a `DELETE` request to `/users/:email`:
```sh
curl -X DELETE 'http://127.0.0.1:8000/users/bob@gmail.com'
```
Response:
```sh
{
    "name": "Bob",
    "email": "bob@gmail.com",
    "password": "secure_v2"
}
```
### Errors
All of the endpoints return an error in json format with a proper http status code, if something goes wrong:
```sh
{
  "message": "user not found"
}
```
