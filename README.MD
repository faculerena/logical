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

