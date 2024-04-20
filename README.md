# go-exp

## How Go Generic Helps Communication in JSON Format 

### We can replace `interface{}` with generic type
In some use cases, we could have to communicate with other applications through JSON format. If we assume both of sender and receiver use Go as their programming language. The receiver would probably want to use the same data structure to unmarshel the payload from the sender. However, if the sender used the `interface{}` as the type of data field, the sender can only get the `map[string]interface{}` [1] for the data in json object. It could largely break the developer experience in receiver side. So, replacing the `interface{}` with generic type helps the receiver without asserting the interface and even using the map structure. 


See the [example](generic_response_wrapping.go). 

[1]: https://pkg.go.dev/encoding/json#Unmarshal