# PROJECT AIM

### Have a frontend to test the speed of processing webhook. Just a way to see what the systemm is doing
### User should be able to dtermine the number of webhooks they can send (Up to 10,000 per request cycle. Don't want to handle larger scale cause of funds)
### User can retry a webhook using he webhook ID
### User can see their webhok public key
### User can subcribe for a webhook. Only uses the user name so there's no auth or anything of the sort.
### Webhooks older than a month are deleted automatically
### There would be three events that a user can subscribe for. Each event has a static payload with the exception of a few fields that are user centric
### User can determine if they would use the rust backend or the go backend right from startup. Can only use one at a time (Both backends use the same db)
### User will see the last 10 webhooks and their status on their webhook display manual
### Rust and Go implementations should be architecturally the same
