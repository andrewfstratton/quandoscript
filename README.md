# quandoscript
Script engine for Quando

Here is the (growing) script definition, note :

- xxx:: is the definition

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
    word-->dot["."]-->word
```

```mermaid
graph LR
    stop([" "])
    word:: --- letter-->stop
    letter-->l[" "]-->rl["letter"]-->r[" "]-->stop
    r-->l
    l-->u["_"]-->r
    l-->digit-->r
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
