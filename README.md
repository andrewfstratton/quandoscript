# quandoscript
Script engine for Quando

- Script format mostly follows https://en.wikipedia.org/wiki/Wirth_syntax_notation
  - `{a}` is any, i.e. 0+
  - `+a+` is many, i.e. 1+
  - `[a]` is optional, i.e. 0 or 1
  - `(a|b)` are groups; without () at top level
  - `(")` is double quote character
  - `"+"` is (punctuation) character

Here is the (growing) script definition :
- line = {action}
- action = word {. word} _[params]_ +whitespace+
- params = "(" param {, param} ")"
- param = word "=" value
- word = letter {letter|digit}
- value = number | range | boolean | string
- string = (") {character} (")
- boolean = true | false
- letter = a..z | A..Z
- digit = 0..9
- character = UNICODE | prefix_character
- prefix_character = "\\" (") | "\\" "\\"


Notes:

- _* is 0..many, + is 1..many_
```mermaid
graph LR
    stop([" "])
    start([" "])---block--> stop
```

```mermaid
graph LR
    stop([" "])
    l-->r
    block::---l["{"]-->lnl["↵"]-->action-->nl-->r["}"]--> stop
    nl["↵"]-->action
```

```mermaid
graph LR
    stop([" "])
    word --> config
    action::--- word-->config --> run --> stop
    word-->.word-->word
```

```mermaid
graph LR
    stop([" "])
    word:: --- letter-->stop
    letter-->digit
    digit-->letter
    letter-->r[" "]
    r-->letter
```

```mermaid
graph LR
    stop([" "])
    config::---params-->stop
```

```mermaid
graph LR
    stop([" "])
    run::---params-->stop
```

```mermaid
graph LR
    stop([" "])
    l --> r
    params:: --- l["("]--> param-->r[")"]-->stop
    param-->c[","]-->param
```

```mermaid
graph LR
    stop([" "])
    param:: --- word --> eq["="]-->value-->stop
```

```mermaid
graph LR
    stop([" "])
    value:: --- l[" "]
    l--- boolean -->stop
    l--- string -->stop
    l--- number -->stop
    l--- variable  -->stop
    l--- block -->stop
```

```mermaid
graph LR
    stop([" "])
    boolean::--- true --> stop
    boolean::--- false --> stop
```

```mermaid
graph LR
    stop([" "])
    string::--- l[#quot;] -->character --> r[#quot;]--> stop
```

```mermaid
graph LR
    stop([" "])
    letter::--- a..z--> stop
    letter::--- A..Z--> stop
```

```mermaid
graph LR
    stop([" "])
    digit::--- 0..9--> stop
```

```mermaid
graph LR
    stop([" "])
    character::--- UNICODE--> stop
    character::--- prefix_character --> stop
```

```mermaid
graph LR
    stop([" "])
    prefix_character::--- l["\"]-->slash["\"]--> stop
    l-->q[#quot;]-->stop
    l-->n-->stop
```

```mermaid
graph LR
    stop([" "])
    map:: --> stop
```
