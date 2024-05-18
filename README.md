# go-exp

## How Go Generic Helps Communication in JSON Format 

### We can replace `interface{}` with generic type
Sometimes our application could have to communicate with other applications through JSON format. Assuming both of sender and receiver use Go as their programming language, the receiver would probably reuse the data structure provided by sender to unmarshel the payload. 

However, if the sender defined the data field as the `interface{}` type, the receiver can only use 
 type `map[string]interface{}` [1] to assert the `interface{}` field. It could largely break the developer experience for the receiver in many perspectives. So, replacing the `interface{}` with generic type helps the receiver without asserting the interface but efficiently re-use the structure. 

See the [example](generic_response_wrapping.go). 

[1]: https://pkg.go.dev/encoding/json#Unmarshal