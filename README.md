# traefik-dynamic-mux
Combines different Traefik dynamic configurations into a single JSON configuration, which can be used via the `http` provider. Powered by [Fiber](https://github.com/gofiber/fiber).

## Usage
Draft your `Muxfile` as below, then configure `http://<host>/mux` as the HTTP JSON dynamic configuration. Ensure no other JSON configuration is set. Default port is 9393.

A Docker image for ARM64 and AMD64 is also available at `czlucius/traefik-dynamic-mux`.

## Muxfile
traefik-dynamic-mux uses a special Dockerfile inspired configuration file to make it easy to modify and alter the configuration.
The `Muxfile` defines the configuration. It has a very simple syntax, `COMMAND arg0, arg1, ...`

An example:
```
JSONURL http://localhost:3000/webhooks/traefik/main.json
YAMLFILE /etc/traefik/config.yml
RELABEL_ENTRYPOINT web http
RELABEL_ENTRYPOINT websecure https
APPLYCERTRESOLVER letsencrypt https [override]
DEFAULTCERT path/to/cert.crt path/to/cert.key
```
The order of the commands matters, the file will be read from top to bottom; bottom rules will apply after the top rules have been applied.

The default location is at `/etc/traefik/Muxfile`, else you can specify a location by supplying the `MUXFILE` environment variable.

More commands will be added as project progresses; the list of commands are below.

## Commands
Each line of a `Muxfile` is for a command, which has this structure:
```
COMMAND arg0, arg1, ....
```

- `JSONURL <url>`
Pulls JSON configuration from the specified URL.
e.g. `JSONURL http://localhost:3000/webhooks/traefik/main.json`

- `YAMLFILE <path>`
Pulls YAML configuration from the specified path.
e.g. `YAMLFILE /etc/traefik/config1.yml`

- `APPLYCERTRESOLVER <certresolver name> <entrypoint> [override]`
Applies certresolver for all routes with specified entrypoint.
The 3rd argument is optional and accepts a literal string `override`, which will override existing certresolvers if they are present, else it will remain unchanged.
e.g. `APPLYCERTRESOLVER letsencrypt websecure override`

- `RELABEL_ENTRYPOINT <from> <to>`
Renames entrypoint from `<from>` to `<to>`
e.g. `RELABEL_ENTRYPOINT https websecure`

- `DEFAULTCERT <certfile> <keyfile>`
Applies a default certificate.


## Comments
To comment a line, add `#` at the beginning (no spaces at the start), like so:
`# this is a comment`
