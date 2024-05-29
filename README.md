# rqbit-telegram
A Telegram bot to remotely control rqbit instance.

# Installation

Make sure you've latest version of golang installed<br/>
1.Clone the repo <br/>
2. cd into the repo <br/>
3. run `go mod download` or `go mod tidy` or both <br/>
4. copy the sample.config.toml to config.toml and fill with deatils <br/>
5. `go build -o rqbtg ./main.go` <br/>
6. run the binay `./rqbtg`

# Usage
## Available Commands

```
/start - start the bot with basic information
/add "metalink" - pass magnet link or http .torrent link and bot will add it to client.
If you want to to add .torrent file just send any .torrent file to the bot (no command requires) and bot will add it to the client.
```

## Contributing

Pull requests are welcome. For major changes, please open [an issue](https://github.com/joybiswas007/rqbit-telegram/issues/new) first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

Show your support by starring [⭐️](https://github.com/joybiswas007/rqbit-telegram/stargazers) this project!
