# Athghnó
### Athghnó /**ah**-ɣn̪ˠoː/ — Activity or work that needs to be redone, usually because it wasn't done right the first time.

Athgnó is an ActivityPub server written in Go, currently a work in progress.
The purpose is to create a spec-conformant server with room for extension,
making it capable of running any potential ActivityPub 'app'.

## Why?
I want to get better at Go, and I want to understand how ActivityPub works better.
I also think that existing servers are too tightly coupled to specific clients and wanted to see if I could make something more general purpose from scratch.
## What's with the name?
It's Irish
## How do I use this?
Currently, you don't. It's not ready for use. However, you can run the current WIP application by doing the following:
### Clone the repo:
`git clone https://github.com/mkgiles/athghno && cd athghno`
### Build the application
`go build .`
### Copy it to your server (optional)
You may need to cross-compile if your server uses a different architecture to your development machine
### Create a .env file with the required parameters
`ATHGHNO_HOSTNAME=<Your server's hostname>`\
`ATHGHNO_PORT=<Port Number>`\
If you need access to port 80, use `setcap` to grant the application access to privileged network ports, do not run as root, it is not safe.
### Run the server
`./athghno`
## What is done:
- ActivityStreams2 API
- ActivityPub extension to AS2 API
- Creating, storing, and retrieving AS2 objects in a persistent key/value store
- Signed requests to other servers for actor information
- A dummy Webfinger mock API that lies
## What is not done:
- Everything else
### Okay but seriously
- The rest of the Webfinger API
- ActivityPub S2S Protocol (currently being worked on)
- SSL authentication
- Activity side-effects
- ActivityPub C2S API
- Extensibility API
- Security hardening and any optimisation
## What's next?
- Providing endpoints for other ActivityPub servers to query and submit objects
- Webfinger API
- Potentially updating json dependency to json-ld to leverage LD specific features
- Fetching and caching data from other servers.
- Handling of Collections
- Activity delivery
- Activity side-effects
- All of C2S API
- Exposing the internal API in such a way that extensions can be written for specific ActivityPub extensions
## Contributing
Forks and PRs are welcome. If you have suggestions for extensions to support (such as mastodon's toot vocabulary, or pixelfed's stories) feel free to create an issue explaining why.
Keep in mind this is an exploratory project, not intended for production use.