goconnect
=========

Go has a powerful set of built in libraries to serve HTTP requests,
which are an excellent building block for building a full web
application.  Goconnect take ideas from Sencha Connect and tries
to map them to Go constructs to create easy to use composable web
applications. The goal is to try and take some of the friction out
of web development in Go and make adding functionality easier.

It's takes the general idea from connect of how to layer requests
through different specific chains of handlers that invoke the next
handler in the chain or stop the request. This is an old idea.

Connect has the general interface

	interface { request, response, next }

I've mapped this to the Go Handler interface and added a next object:

	interface { response, request, next }

To keep it as simple to use as possible next is a closure around
the next handler in the chain. Right now this is generated for every
request.

I've only ported a few of the middleware plugins as a proof of concept.

 * logger - log requests / responses
 * basicauth - require basic auth
 * static - serve static files using Go net/http built in file server
 * limit - limit content-length of the request

The next important thing to try and implement is add Session
capability through Gorilla / Session.

