# HTTP Handler (Presentation layer)
The HTTP Handlers handle HTTP requests sent to a specific path.

If the request method is POST/PUT/PATCH and the request body is not empty, in the handler function the payload is bound/extracted§§§ to the DTO model. 
The `bindAndValidate` function takes the payload from the request and tries to unmarshal it into the target DTO model. After a successful binding, the data is validated based on the configuration defined as tags/annotations§§§ on the DTO model. If one of these steps fails, the handler returns an error.  
When the target DTO model receives the payload data, the handler function passes it to a usecase function. The use case implements the necessary business logic and returns the appropriate DTO response model and an error value.

In order to have a clear separation between **presentation**, **domain** and **data source layers**, the http handlers should call **only** usecases. No direct calls to 3rd party services or db queries are allowed here.