# AML

AML (Arith Metic Language) is an interpreted math equation language. You can either use the interactive shell by running
the interpreter with no arguments, or with a file ending in .aml by passing the name of the file as an argument.

Build:

```console
cd src
go build
```

Run the interpreter with no arguments to enter the interactive shell:

```console
% ./aml
Enter "q" to quit

> _
```

Run with an AML file:

```console
./aml file.aml
```

Run tests after building:

```console
go run tests.go
```

---

## Syntax

AML can interpret human readable equations, no fancy syntax needed: 

```aml
2 + 2
```

### However, there are some rules...

Each equation should be on its own line.

Basic arithmetic operators should only be in between integers (or decimals).

| Operator   | Action                       |
|------------|------------------------------|
| +          | Add                          |
| -          | Subtract                     |
| *, x, X, × | Multiply                     |
| /          | Divide                       |
| %          | Modulus (division remainder) |

Expressions (a statement of integers and operators) may be inside of closures and calculated seperately, but operands
must follow and precede closures.

```aml
5 * (8 + 4) / 8
```

Comments are allowed, but must either be on their own line or at the very end of a line. Comments are denoted using `#`.

```aml
# This is a comment

2 + 2 # equals 5

# This is a
# multiline
# comment
```
