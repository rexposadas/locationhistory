# Running the app

You can run this as as simply as:

```bash	
        go build -o ./locationservice && ./locationservice
```

Optionally, you can set environment variables. 

```bash

	export HISTORY_SERVER_LISTEN_ADDR=8999
	export LOCATION_HISTORY_TTL_SECONDS=60

	go build -o ./locationservice && ./locationservice
```

# Design decisions

### Timestamps
There was no need to keep the timestamp of when we received the order, since the slice already keeps our history in chronological order. 

### Delete

The DELETE method should return a message if the user tries to deletes
a non-existant order. But since that's not a requirement and I am time boxing this to 2-3 hours, I didn't include it. 


### Pruning expired locations.

Pruning expired locations is done when getting the list. This is because this is the only time it matters.  Alternatively, we can have a routing which cleans that list, but that implemenation may be problematic once the list gets very large.
