[WIP] jc
====

peco-like interactive JSON processing.

- jq-like syntax
- peco-like interactivity.

You can process JSON like...

![](https://camo.githubusercontent.com/6ed15cca08fd6972d12e67ee1f1fe84caa14744b/687474703a2f2f7065636f2e6769746875622e696f2f696d616765732f7065636f2d64656d6f2d70732e676966)

Syntax
---

```
> .data[0]
{
  "id": "X999_Y999",
  "from": {..},
  "message": "Looking forward to 2010!",
  "actions": [..],
  "type": "status",
  "created_time": "2010-08-02T21:27:44+0000",
  "updated_time": "2010-08-02T21:27:44+0000"
}

[1] .id
[2] .from ->
[3] .message
[4] .actions
[5] .type ->
[6] .created_time
[7] .updated_time
[9] .data[1] ->
```

```
> .data[0].from.name
"Tom Brady"

[0] .data[0].from
[9] .id
```

for following JSON

```json
{
   "data": [
      {
         "id": "X999_Y999",
         "from": {
            "name": "Tom Brady", "id": "X12"
         },
         "message": "Looking forward to 2010!",
         "actions": [
            {
               "name": "Comment",
               "link": "http://www.facebook.com/X999/posts/Y999"
            },
            {
               "name": "Like",
               "link": "http://www.facebook.com/X999/posts/Y999"
            }
         ],
         "type": "status",
         "created_time": "2010-08-02T21:27:44+0000",
         "updated_time": "2010-08-02T21:27:44+0000"
      },
      {
         "id": "X998_Y998",
         "from": {
            "name": "Peyton Manning", "id": "X18"
         },
         "message": "Where's my contract?",
         "actions": [
            {
               "name": "Comment",
               "link": "http://www.facebook.com/X998/posts/Y998"
            },
            {
               "name": "Like",
               "link": "http://www.facebook.com/X998/posts/Y998"
            }
         ],
         "type": "status",
         "created_time": "2010-08-02T21:27:44+0000",
         "updated_time": "2010-08-02T21:27:44+0000"
      }
   ]
}
```

Use case?
---

- API client development
- Everything dealing with JSON!
