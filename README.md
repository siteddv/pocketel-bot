# PockeTel Bot

**PockeTel** is a Telegram bot that allows you to save links in the application <a href="https://app .getpocket.com/">Pocket</a>. We can say that he is a small client for this service.

To work with the Pocket API, a self-written SDK is used - <a href="https://github.com/siteddv/golang-pocket-sdk ">golang-pocket-sdk</a>.

<a href="https://github.com/bolt/bolt">Bolt DB</a> is used as storage.

To implement user authorization, an HTTP server is launched together with the bot on port 80, to which a redirect from Pocket occurs when the user is successfully authorized.

When the server accepts the request, it generates an Access Token via the Pocket API for the user and stores it in storage.

### Stack:
- Go 1.15
- BoltDB
- Docker (for deployment)

### To launch the application:

```
make run
```

### Have a nice experience with repository! 
#### If you have any questions and suggestions, please contact me via email siteddv@gmail.com
