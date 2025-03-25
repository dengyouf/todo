```
  main  go  run . --help                                                ✔  306  00:01:38 
Usage of /Users/dengyouf/Library/Caches/go-build/2e/2e4fbe5164346ece753be495b95f7c8d2c17dacb143edbe9623d6301e4e87a97-d/todo:
  -add string
        Add a new todo entry
  -del int
        Delete a todo entry (default -1)
  -edit string
        Edit a todo entry
  -list
        List all todo entries
  -toggle int
        Toggle a todo entry (default -1)

```

```
  main  go  run . -list                                                 ✔  307  00:01:46 
┌───┬────────────────┬───────────┬───────────────────────────┬───────────────────────────┐
│ # │     Title      │ Completed │        Created At         │       Completed At        │
├───┼────────────────┼───────────┼───────────────────────────┼───────────────────────────┤
│ 0 │ Go to ShangHai │ ❌        │ 2025-03-25T23:53:23+08:00 │                           │
│ 1 │ Go to Beijing  │ ✅        │ 2025-03-25T23:53:43+08:00 │ 2025-03-25T23:53:54+08:00 │
└───┴────────────────┴───────────┴───────────────────────────┴───────────────────────────┘
  main   go run . -add "Go to ChangSha"                                  ✔  308  00:02:16 
adding todos...
  main  go run . -toggle 2                                            ✔  309  00:02:32 
  main  go  run . -list                                               ✔  310  00:02:42 
┌───┬────────────────┬───────────┬───────────────────────────┬───────────────────────────┐
│ # │     Title      │ Completed │        Created At         │       Completed At        │
├───┼────────────────┼───────────┼───────────────────────────┼───────────────────────────┤
│ 0 │ Go to ShangHai │ ❌        │ 2025-03-25T23:53:23+08:00 │                           │
│ 1 │ Go to Beijing  │ ✅        │ 2025-03-25T23:53:43+08:00 │ 2025-03-25T23:53:54+08:00 │
│ 2 │ Go to ChangSha │ ✅        │ 2025-03-26T00:02:32+08:00 │ 2025-03-26T00:02:42+08:00 │
└───┴────────────────┴───────────┴───────────────────────────┴───────────────────────────┘           
```
