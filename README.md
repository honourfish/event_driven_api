# event_driven_api
Event Driven webapi

## Overview

Attempt to make a basic event driven api that uses a RESTful http interface, ensuring it still
conforms to REST.

## Design

1. the user make a request of the server.
2. the server passes the request to a handler.
3. the handler generates a correlation id.
4. the handler pushes the request data (message) onto a request message queue with the correlation id as metadata.
5. the handler then responds with a 202 accepted and the correlation id.
6. the request message queue then pushes the message to any relevant subscribed worker service.
7. the worker service does work on the message, and pushes a response message with the correlation id as metadata, onto a response message queue.
8. the user, at some point, makes another request to the server with the correlation id.
9. the server passes the request to a handler.
10. the handler asks the response message queue, if it has any messages with the correct correlation id.
11. if it does, the handler takes the message, processes it, and responds accordingly.
12. if it doesn't, the handler response with a suitable error code (404, 409...etc).
