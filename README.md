# go-basics

Collection of basic functions e.g. for nice looking println's or logging, basic functions to make programming more easy.

See `demo_app.go` for how to use this library package.

```bash
$ go run demo_app.go

Our pretty printer
------------------
👉 DBUG: Debug Mode is on
💬 A Message
🏁 YAY - Success Message
❕ INFO: Info Message
👉 DBUG: Debug Message
🤭 WARN: Warn Message
💣 EROR: Error Message(an error)


phuslu pretty printer
---------------------
2021-12-11 19:58:03 DBG > Debug Mode is on
2021-12-11 19:58:03 INF > A Message
2021-12-11 19:58:03 INF > Success Message
2021-12-11 19:58:03 INF > Info Message
2021-12-11 19:58:03 DBG > Debug Message
2021-12-11 19:58:03 WRN > Warn Message
2021-12-11 19:58:03 ERR > error="an error" Error Message


phuslu json printer
-------------------
{"time":"2021-12-11 19:58:03","level":"debug","message":"Debug Mode is on"}
{"time":"2021-12-11 19:58:03","level":"info","message":"A Message"}
{"time":"2021-12-11 19:58:03","level":"info","message":"Success Message"}
{"time":"2021-12-11 19:58:03","level":"info","message":"Info Message"}
{"time":"2021-12-11 19:58:03","level":"debug","message":"Debug Message"}
{"time":"2021-12-11 19:58:03","level":"warn","message":"Warn Message"}
{"time":"2021-12-11 19:58:03","level":"error","error":"an error","message":"Error Message"}


End Demo
--------
👉 DBUG: Debug Mode is on
💬 ☠ Exit☠  - Have a nice Day
exit status 1
```

Don't hit me - it's my first golang project... ^_^v
