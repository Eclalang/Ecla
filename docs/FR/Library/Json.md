# Json

La librairie Json implémente l'encodage et le décodage en JSON.

## Index

- [Marshal(content map[string]string | []map[string]string) (string, error)](#marshal)
- [Unmarshal(content string) (map[string]string | []map[string]string, error)](#unmarshal)

## Marshal
```
function marshal(content map[string]string | []map[string]string) (string, error)
```
Convertit une map en JSON string

### Exemple :
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
Convertit un string JSON en n'importe quel type

### Exemple :
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