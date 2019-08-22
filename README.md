# httpserver examples

A comparison few simple httpservers written in a variety of languages, using their standard libraries and popular frameworks.

## Go

Out of the box the support for basic web things are fairly straight forward, as you would expect from a language that is explicitly Web First by design. The notable exception is the lack of a way to cleanly deal with trailing slashes on paths. I was only able to accomplish on the `Echo` example by using middleware, though I did not dig too deeply into Go's regex (as you will understand later).

## Nim

The basic Nim httpserver is a bit more verbose than Go. One thing that was sorely lacking was a way to parse path parameters, which is silly because it is such a common task for a web developer. In the end I resorted to regex, which is simple enough but leaves the router feeling a bit more cryptic than I would like.

Where Nim outshines Go is in its metaprogramming capabilities, which you see in action on the `Jester` example. It provides a very simple and intuitive DSL. Trailing slashes was simple but not automagic.

## Conclusion

Nim, all the way.

As a language, Nim feels less 'polished' than Go, but that is actually part of its charm. It asks you to step into it, understand it more deeply. It invites you inside of its world. I geuninely did not mind the exercise of parsing path params with regex, because the language made me interested to learn it. Metaprogramming is what makes Nim feel like a nimble language. The only code smell I've encountered with Nim is its polluted namespace.

Go is a bland language. After a few hours of programming it you are left with the sense that you have always known how to program Go, and I don't think that is a good thing. The language feels unchallenging to the point where you have to wonder why you are even bothering. I applaud the effort to make web-related tasks first class citizens of the standard library. I want to like Go because it does make you feel productive, but in the end it left me underwhelmed and uninspired.
