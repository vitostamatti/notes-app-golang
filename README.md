
<h1 align="center">Notes API Golang</h1>


<p align="center"> 
    A basic CRUD API of Notes with users JWT authentication.
    <br> 
</p>


## Getting Started <a name = "getting_started"></a>

First you need to clone this repository.

```bash
git clone https://github.com/vitostamatti/notes-app-golang/ 
```

Then cd into the notes-app-golang directory and run 

```
docker compose up
```

If you don't have docker installed you first need to [download](https://www.docker.com/) it.

## Usage <a name="usage"></a>

When the docker containers are running you can start playing with the app.

If you go to [localhost/5050](http://localhost/5050) you're going to see the pgadming app. 
There you can login using: 
- username: admin@domain.com 
- password: admin 

Yo interact with the api listening on [localhost/8080](http://localhost/8080) I recommend using a tool to create requests like [postman](https://www.postman.com/).

Once there you first need to login. I've created a super-user to start with. 
Make a post request to http://localhost/8080/api/login with a body of 

```json
{
    "username":"admin",
    "password":"admin"
}
```

and you'll get a token response. This is a bearer token to acces the rest of the endpoints.
Theses are

- api/users/me (GET): to get the current user authenticated with the given token.
- api/register (POST): to create new users with **username** and **password** (and optionally a is_superuser boolean field.)
- api/notes (GET): to get all the notes from the current user loged in. Make sure you pass the authentication token as a bearer token in the header of the request.
- api/notes (POST): to create a new note with a **name** and a **content** fields.
- api/notes/{id} (GET): to get a particular note by its ID passed in the request path.
- api/notes/{id} (PUT): to update an existing note with the passed ID.
- api/notes/{id} (DELETE): to delete an especific note with the passed ID.

If you are loged in as the a superuser you and also acces the users information:
- api/users (GET): to get all the users.
- api/users (POST): to create a new user with a **username** and a **password** fields.
- api/users/{id} (GET): to get a particular user by its ID passed in the request path.
- api/users/{id} (PUT): to update an existing user with the passed ID.
- api/users/{id} (DELETE): to delete an especific user with the passed ID.


## Authors <a name = "authors"></a>

- [@vitostamatti](https://github.com/vitostamatti) - Idea & Initial work


