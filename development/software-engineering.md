### Programming Terminology 
---
##### Compiled vs Interpreted Language
```
Compiled Language
  a) Compiler translates program to machine code before execution.
  b) Time is required to compile code, but execution is fast during runtime.

Interpreted Language
  a) Interpreter reads and executes program without prior compilation.
  b) Execution is slow during runtime.
```

##### Weak vs Strong Typing
```
Weak Typing   : Interpreter allows for operations with incompatible types ("1" + 1 = 2).
Strong Typing : Interpreter does not allow operations with incompatible types ("1" + 1).
```

##### Statically vs Dynamically Typed Language
```
Statically Typed Language
  a) Type errors are checked during compile time. and the type of a variable cannot change.
  b) Type of a variable cannot change during runtime.

Dynamically Typed Language
  a) Type errors are checked during runtime.
  b) Type of a variable may change during runtime.
```

##### Declarative vs Imperative Programming
```
Declarative Programming : Practice of telling the program what you want it to do without explaining implementation details.
Imperative Programming  : Practice of telling the program how you want it to achieve something through implementation details.
```

### Object Oriented Programming
---
##### Abstraction 
```
Concept of hiding all but important features and giving ability to conceptualize things at a higher level to reduce complexity.
``` 

##### Encapsulation 
```
Bundling data fields as private to prevent open access and changes.
```

##### Inheritance
```
Concept of passing down instance variables and methods of a parent class to a child class.
As an example, a parent class could be "Dog" and a child class could be "Retriever".
This allows code to be reusable and define a common protocol for groups of classes.
```

##### Polymorphism
```
Triangle shape = new Triangle() 
Shape shape = new Triangle()           triangle is a shape 
Shape shape = new Circle()             circle is a shape  
Triangle shape = new Shape()           triangle is a shape, but a shape is not a triangle (invalid)

public void run(Shape shape) {}        objects with type shape, triangle, and circle can all be passed to the parameter

Advantages:
  1) Flexibility : ability to not have to rewrite code when a new subclass is defined 
```

##### Composition 
```
defines a has-a relationship

public class Point3D {
  public Point3D(int x, int y, int z) {
    p = new Point2D(x,y);
  }
}
```

##### Dependency Injection 
```
Suppose you have a class that depends on another object 

public Car() {
  Wheel wheel = new CanadianWheel();
  Windows window = new CanadianWindow();
}

Wheel and windows are 'dependencies' of car because they are needed to construct the car 

What if we want to create multiple objects of car, but with different wheels, companies, and materials?

Instead of instantiating and setting dependencies within the class, one can create a dependency from the outside 

This exterior dependency (e.g. wheel) can then be injected into the class car  

public Car(Wheel wh, Windows wi) {
  this.wheel = wh;
  this.windows = wi;
}

Advantages :
  1) Makes testing easier as dependency classes do not have to be instantiated 
  2) Allows for independent development of classes 
  3) Can easily switch out different dependencies to the class 
  4) Can use the dependencies in other classes 
```

##### Object Relational Mapping
```
Ability to write relational database queries through OOP languages by making them compatible with each other. 

ex) SELECT * FROM users WHERE type = "test"; --> var user = orm("users").where({type: "test"});
```

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

##### OpenAPI
```
OpenAPI is a specification for building and documenting REST APIs typically in the form of YAML or JSON. 
It is a standard that allows developers to describe the functionality of their APIs in both human and machine readable format.
```

##### Status Codes
```
1xx : Request received and understood 
2xx : Request by client was received, understood, and accepted 
3xx : Client must take additional actions 
4xx : Client screwed up (for wrong GET,DELETE requests)
5xx : Server screwed up
```
