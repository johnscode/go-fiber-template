
# go-fiber-template

A simple golang app template using fiber and the 'handlebars' template engine. 
This app template makes it pretty easy to get a server and a set of pages up and running.

This app template is already setup with logging and configuration using the server environment.
Environment is used for configuration rather than files to discourage the use of secrets files 
(_which is a whole other conversation_)

Zerolog is used for logging its efficiency rather than the builtin log module.

The handlebars engine is used rather than golang's html template as it is easier to support multiple layouts.
It's not too hard to change the template engine, if you prefer one of the other supported ones. 
See [Fiber template docs](https://docs.gofiber.io/guide/templates) for more info.

## To Do

- dockerize it
