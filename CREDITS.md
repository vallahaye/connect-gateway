# Credits

## Connect-Gateway v0

- **Valentin Lahaye** ([@vallahaye](https://github.com/vallahaye))<br/>
  *Original author of the Connect-Gateway*

- **Johannes Brüderl** ([@birdayz](https://github.com/birdayz))<br/>
  *Fixed some major issues when integrating the Connect-Gateway into [Redpanda Console](https://github.com/redpanda-data/console) ([PR#26](https://github.com/vallahaye/connect-gateway/pull/26), [PR#31](https://github.com/vallahaye/connect-gateway/pull/31))*

- **Emile Friot** ([@emileFRT](https://github.com/emileFRT))<br/>
  *Helped refining the idea and tested several implementations*

- **The Buf development team** ([@bufbuild](https://github.com/bufbuild))<br/>
  *For their awesome software solutions and some codes taken directly from their code base:*
  - Support for interceptors through a [`HandlerOption`](https://pkg.go.dev/go.vallahaye.net/connect-gateway#HandlerOption) and the internal interceptors chaining logic
  - Wrap comments in generated files
