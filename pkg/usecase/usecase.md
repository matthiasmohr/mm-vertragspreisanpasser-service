# Use cases (Domain Layer)
The _Domain Layer_ implements all the business logic. 

The use cases should always take a `ctx` of type `context.Context` as a parameter. 
The `ctx` is needed for the following reasons:
- The `ctx` contains the request related data like `requestID` or depending on the use case resource identifiers or the name of the client. Furthermore, all log messages should be logged with the context to include the request data stored in the context in the log message. This will helps with debugging and analyzing issues.
- If performing 3rd party calls, always use the context in the request. This will allow canceling the request if the called party cancels the request on its side (UI).

The use case usually receives a request DTO as an incoming parameter and transfers it into the needed domain model for internal handling.  
The return parameters usually consists of a type DTO response model and an error value. 
