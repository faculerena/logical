# Trivaluated logic expression parser

*Note: This DOES NOT apply the "NOT" operator. If someone has an idea on how to implement it and wants to do it, please do!*


Simple logical expression parser and evaluator. It accepts this format:

```text
"(...) operation (...)"
```

Examples
```text
exp1 = "(a&b)|c"
exp2 = "(a>b)|(c&d)"
```

Given an expression, it'll ask for the values of each variable. Enter T, F or U as needed.
It'll then evaluate the expression and print the result.


Here's more info and another examples of the parser (the latest 3 doesn't evaluate the expression). I took inspiration and got helped by the authors!
[parser by iglosiggio (javascript)](https://gist.github.com/iglosiggio/b9f0452123160fe147488bf348047605)

[Haskell parser by nsanmartin (Haskell)](https://gist.github.com/nsanmartin/e7edd30d3689a14a0edfa4caab0098e2)

[Python parser by fsan (Python)](https://tildegit.org/fsan/amor-plastico/src/branch/master/main.py)
