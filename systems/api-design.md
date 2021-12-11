### API Design
---
##### What is an API
```
1) defines interactions between software such as types of calls/requests, how they are made, data formats, conventions 
2) way of communicating between applications 
3) allows for servers/systems to communicate with one another such that automation is possible 
4) allows information to be shared as a service 
```

##### API Design
```
CRUD operations : Create, Read, Update, Delete

Examples of Entity Definitions 
1) Payment
  - id : uuid 
  - customer_id : uuid 
  - restaurant_id : uuid 
  - amount : int
  - status : enum ["success", "pending", "failed"]

2) Restaurant 
  - id : uuid 
  - name : string 
  - address : string 
  - account : Account 
 
Example of Payment.json and Payment object
{id : "abac1123-bfsdg", customer_id : "bdfsx-123cvxc", restaurant-id : "bac1123-bfsdg", amount : 2000, status : pending}

Example of Restaurant.json and Restaurant object 
{id : "bac1123-bfsddg", name : "Papa Johns", address : "4005 Chestnut Street", account : {Bank : ___, Account No. ___}}


Example of Endpoint Definitions 
1) Payment
  - Payment createPayment(payment: Payment)
    path : POST /v1/payments
  - Payment getPayment(id: uuid)
    path : GET /v1/payments/id 
  - Payment updatePayment(id: uuid, updatedPayment: Payment) 
    path : UPDATE /v1/payments/id
  - Payment[] listPayments(offset: int, limit: int) --> Pagination 
    path : GET /v1/payments

2) Restaurant
  - Restaurant createRestaurant(restaurant: Restaurant)
  - Restaurant getRestaurant(id: uuid)
  - Restaurant deleteRestaurant(id: uuid)
  
Pagination : limit the response of a potentially larger response, usually when retrieving huge lists  
```

##### REST Principles
```
REST principles 
  1) verbs : GET (read), POST (create), PUT/PATCH (update entire/partial), DELETE (delete)
  
REST parameters
broken into endpoint(site) + query parameter(condition)  
ex) api.com/cars?type=SUV&year=2019

REST practices
1) Use "nouns" and not "verbs"
2) Use "plural" for list of items 
3) Use "camel case"
```

##### REST Mappings
```
One-To-Many Mapping
api.com/tickets/145/messages/4 --> find 4th message for 145th ticket 
a single ticket has N unique messages associated with that ticket

Many-To-Many Mapping
api.com/groups/200/users/56 --> find user of id 56 in 200th group
a user might also be in different groups    
```

##### Status Codes
```
1xx : Request received and understood 
2xx : Request by client was received, understood, and accepted 
  1) 201 Resource Created (for POST methods)
  2) 202 Accepted 
  3) 204 No content (for DELETE methods)
3xx : Client must take additional actions 
4xx : Client screwed up (for wrong GET,DELETE requests)
5xx : Server screwed up
  1) 500 Internal Server Error
  2) 504 Gateway Timeout
```