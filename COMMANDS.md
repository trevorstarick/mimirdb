# Boolean Transactions
## AND
Keeps existing keys if they are in the result of the following query
`[a, b, c] AND [b, c] = [b, c]`

## END
Ends the transaction and returns all of the currently valid keys

## OR
Adds the result of the following query to the transaction
`[a] OR [b, c] = [a, b, c,]`

## NAND
Removes the result of the following query form the transaction
`[a, b, c] NAND [b, c] = [a]`

## START [ALL] [NONE]
Starts the transaction with either all of the data from the instance (filter out) or none (focus on)


# Connection
## AUTH password
Authenticates with Server

## LIST
Lists instances

## PING
Returns `PONG`

## QUIT
Quits Connection

## SELECT instance
Switches instance


# Geo
## BOUNDS nelatitude nelongitude swlatitude swlongitude [LIMIT limit]
Returns all keys that fall inside the bounding box

## POINT latitude longitude [RADIUS radius] [LIMIT limit]
Returns all keys that fall inside the radius


# Time
## TSCAN start [end]
Returns all keys from start to end. If no end is provided, return until now


# Keys
## BIND parent child [child...]
Binds the children with parent element

## DEL key [key...]
Deletes the key(s) and unbinds the children from them

## NUKE key
Recursivly deletes key and all children/grand-children, unbinding them from any parent

## EXISTS key [key...]
Returns an array of booleans indicative of if the key(s) exists

## GET key [key...]
Returns and array of values based off of the keys requested

## RANDOM number
Returns `n` number of random keys

## SET key value
Sets the key's value. If key doesn't exist, create a new one and set it's value

## UNBIND parent child [child...]
Unbinds the children from the parent

## WATCH key [key...]
Watches a/an key(s) and prints a changelog


# Server Admin
## CLIENTS
Lists all currently connected CLIENTS

## FLUSHALL
Deletes everything

## FLUSHDB
Deletes everything in current instance

## INFO [section]
Returns information on server, or instance. If no section is provided, return all info

## SAVE
Writes content to disk

## SHUTDOWN
Stops the server

## TIME
Returns current time on server

# Utils
## UUID number
Returns `n` random UUID(s)