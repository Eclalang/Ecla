# Json

Json library implements encoding and decoding of JSON.

## Index

- [Marshal(content any) string](#marshal)
- [Unmarshal(content string) any](#unmarshal)

## Marshal
```
function marshal(content any) string
```
Converts any type to JSON string

### Example :
```ecla
import "console";
import "json";

function testMarshal() {
    var mapStr map[string]string = {"name": "Ecla", "members": "7", "language": "Golang"};
    var str string = json.marshal(mapStr);
    console.println(str);
}
```

##  Unmarshal
```
function unmarshal(content string) any
```
Converts JSON string to any type

### Example :
```ecla
import "console";
import "json";

function testUnmarshal() {
    var str string = "{'name': 'Ecla', 'members': '7', 'language': 'Golang'}";
    var obj map[string]string = json.unmarshal(str);
    console.println(obj);
}
```