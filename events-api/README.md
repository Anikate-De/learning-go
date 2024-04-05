# Events REST API

This is a REST API for managing events. Events are stored in a SQLite database.
An event has the following attributes:

- Name
- Date
- Location
- Description

It has the following endpoints:

- `GET /events`: Get a list of all events
- `GET /events/<id>`: Get a single event by ID
- `POST /events`: Create a new bookable event, requires authentication
- `PUT /events/<id>`: Update an event, requires authentication, can only be done by the author
- `DELETE /events/<id>`: Delete an event, requires authentication, can only be done by the author
- `POST /signup`: Register a new user
- `POST /login`: Login to the API, returns a JWT token
- `POST /events/<id>/register`: Book a place at an event, requires authentication
- `DELETE /events/<id>/register`: Cancel a booking, requires authentication
