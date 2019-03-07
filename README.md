# Service Architecture is based on ArdanLabs Service Starter Kit
https://github.com/ardanlabs/service

## ID is a UUID getting used for Primary Index
URL: http://0.0.0.0:3000/v1/user/{id}

## Secondary Index is implemented as a Unique Index on Line 142 in handlers/users.go
URL: http://0.0.0.0:3000/v1/user?email={email}

## Most of the validations have to be taken care of
Added todo notes for the same

## Things todo apart
- [ ] Adding Tracing
- [ ] Adding Authentication to POST request
- [ ] More Elaborated Error Handling
- [ ] Handle Debug server more properly
- [ ] gracefully shutdown the server on error
- [ ] Test cases
- [ ] Proper seperation of business logic
- [ ] profiling setup

## Dev Setup
docker build -t loomx:0.0.1 .

docker run -p 3000:3000 loomx:0.0.1

## POST requests
URL: http://0.0.0.0:3000/v1/user

Request Body: {
	"username": "htgyl",
	"email": "hitesh@udacity.com"
}

Content-Type: application/json

## Folder Structure
* Internal folder mostly contains platform related stuff
* handlers contains all the routings
* main contains only configuration setup
