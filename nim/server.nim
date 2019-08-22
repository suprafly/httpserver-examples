import asynchttpserver, asyncdispatch, uri, json, re
import strutils

var server = newAsyncHttpServer()

proc matchRoute(req: Request, req_method: string, pattern: Regex) : bool  =
  var path = req.url.path
  path.removeSuffix("/")
  var matches: array[1, string]
  if ($req.reqMethod == req_method) and (find(path, pattern, matches) > -1):
    true
  else:
    false

proc getPathParam(req: Request, pattern: Regex) : string  =
  var path = req.url.path
  path.removeSuffix("/")
  var matches: array[1, string]
  if find(path, pattern, matches) > -1:
    matches[0]
  else:
    ""

proc jsonHeaders() : HttpHeaders  =
  newHttpHeaders([("Content-Type","application/json")])

proc cb(req: Request) {.async.} =
  if matchRoute(req, "GET", re"/users$"):
    let js = %* []
    await req.respond(Http200, $js, jsonHeaders())
  elif matchRoute(req, "GET", re"/users/([0-9]+)$"):
    let user_id = getPathParam(req, re"/users/([0-9]+)$")
    await req.respond(Http200, $user_id, jsonHeaders())
  else:
    await req.respond(Http400, "error")

echo "Server started on http://localhost:8080"

waitFor server.serve(Port(8080), cb)
