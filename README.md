# go_config
simplify reading config file of json format

Suppose you have a config file like this:
```
{
    "string": "8081",
    "int": 10,
    "stringList": ["a", "abc"],
    "intList": [1, 2],
}
```

# Before


You will:
1. define a struct 
2. use json.unmarshal()
3. use it 

# Now

You load config file though 
```
conf := config.LoadConfig(FileName)
```
Then use it
```
result := conf.Get("KeyName")
```
# Install

```
go get -u github.com/ssor/go_config
```