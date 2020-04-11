# Benthos custom plugin

Create a *benthos* custom plugin


### Install

```
go get
```

```
go build
```

### Run

```
./sarcasm -c config.yaml
```

### Test

```
http post localhost:4195/post "id"="foo1" "content"="this is totally sarcastic /s"
```
