# BACKEND IMPLEMENTATION

### Endpoint to Subscribe to webhook events. Can subscribe for more events via the same endpoint 
### Endpoint to 'create' webhooks (can send from 1 to 10,000 webhooks at once)
### Endpoint to get last 10 webhooks proceed for this user
### A register endpoint which is basically just to input a username (no password, no serious authentication going on)
### When a user is logged in, he cannot login elsewhere (should I implement this?)
### the timestamp, username and ID would be used to create idempotency keys so the same webhook isn't stored twice
### Sending multiple means the endpoint is hit multiple times so the endpoint has to return as fast as posible. Ideally less than 3 second, but at worst less than a minute
