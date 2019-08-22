import jester, json

router router:
  get "/users/?":
    let js = %* []
    resp js
  get "/users/@user_id/?":
    resp @"user_id"

proc main() =
  let settings = newSettings(port=Port(8080))
  var jester = initJester(router, settings=settings)
  jester.serve()

when isMainModule:
  main()
