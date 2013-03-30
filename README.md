ms.go
=====

Port of [ms.js](https://github.com/guille/ms.js) to Go.


func Parse(str string) (float64, error)
------------------------------------------------

```go
ms.Parse("1d")      // 86400000
ms.Parse("10h")     // 36000000
ms.Parse("2h")      // 7200000
ms.Parse("1m")      // 60000
ms.Parse("5s")      // 5000
ms.Parse("100")     // 100
```

func Short(ms float64) string
-----------------------------

```go
ms.Short(60000)     // "1m"
ms.Short(2 * 60000) // "2m"
```

func Long(ms float64) string
----------------------------

```go
ms.Long(60000)      // "1 minute"
ms.Long(2 * 60000)  // "2 minutes"
```

