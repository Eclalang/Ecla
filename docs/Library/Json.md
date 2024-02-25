# Json

Json library implements encoding and decoding of JSON.

## Index

- [Marshal(content map[string]string | []map[string]string) (string, error)](#marshal)
- [Unmarshal(content string) (map[string]string | []map[string]string, error)](#unmarshal)

## Marshal
```
function marshal(content map[string]string | []map[string]string) (string, error)
```
Converts a map to JSON string

### Example :
```ecla
import "console";
import "json";

function testMarshal() {
    var mapStr map[string]string = {"name": "Ecla", "members": "7", "language": "Golang"};
    var str string = json.marshal(mapStr);
    console.println(str);

    var arrMap []map[string]string = [{"name": "Ecla", "members": "7", "language": "Golang"}, {"name": "Ecla", "members": "7", "language": "Golang"}];
    str = json.marshal(arrMap);
    console.println(str);
}
```

##  Unmarshal
```
function unmarshal(content string) (map[string]string | []map[string]string, error)
```
Converts JSON string to any type

### Example :
```ecla
import "console";
import "json";

function testUnmarshal() {
    var str string = "{\"name\": \"Ecla\", \"members\": \"7\", \"language\": \"Golang\"}";
    var obj map[string]string = json.unmarshal(str);
    console.println(obj);

    var strArr string = "[{\"name\": \"Ecla\", \"members\": \"7\", \"language\": \"Golang\"}, {\"name\": \"Test\", \"members\": \"12\", \"language\": \"Python\"}]";
    var objArr []map[string]string = json.unmarshal(strArr);
    console.println(objArr);
}
```