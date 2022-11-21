# grapefruit

very simple and naive implementation of the CQRS pattern, it was created mainly to create a service using mongoDB and elastisearch.  
the "recorder" is responsible for "command" (but like a simple CRUD) and saves data to mongoDB, and then sends data via rabbitMQ to "viewer" responsible for the "query" part and storing data in elasticsearch.   
the application operates on the **Object** model.