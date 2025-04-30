# quandoscript
Script engine for Quando

Here is the (growing) script definition, note :

- xxx:: is the definition

```mermaid
graph LR
    stop([ ])
    start-->stop
    start([ ])-->sequence--> stop
```

```mermaid
graph LR
    stop([ ])
    sequence::---lines--> nl["↵"]--> stop
```

```mermaid
graph LR
    stop([ ])
    lines::---line--> nl["↵"]--> stop
    nl-->line
```

```mermaid
graph LR
    stop([ ])
    line::---id--> nl["↵"]--> stop
    id-->spacer-->action-->nl
```

```mermaid
graph LR
    stop([ ])
    id::--- digit-->stop
    digit-->s( )-->digit
```

```mermaid
graph LR
    stop([ ])
    digit::--- 0..9--> stop
```

```mermaid
graph LR
    stop([ ])
    spacer::---l[ ]-->s["#quot; #quot;"]-->r[ ]-->stop
    l-->tab-->r
    r-->l
```

```mermaid
graph LR
    stop([ ])
    word --> config
    action::--- word-->config --> run --> stop
    word-->dot["."]-->word
```

```mermaid
graph LR
    stop([ ])
    word:: --- letter-->stop
    letter-->l[ ]-->rl["letter"]-->r[ ]-->stop
    r-->l
    l-->u["_"]-->r
    l-->digit-->r
```

```mermaid
graph LR
    stop([ ])
    letter::---l[ ]--> a..z--> stop
    l--> A..Z--> stop
```

```mermaid
graph LR
    stop([ ])
    config::---params-->stop
```

```mermaid
graph LR
    stop([ ])
    run::---params-->stop
```

```mermaid
graph LR
    stop([ ])
    params:: --- l["("] --> r[")"]-->stop
    l--> param--> r
    param-->c[","]-->param
```

```mermaid
graph LR
    stop([ ])
    param:: --- word --> eq["="]-->value-->stop
```

```mermaid
graph LR
    stop([ ])
    value:: --- l[ ]
    l--> boolean -->stop
    l--> string -->stop
    l--> number -->stop
    l--> variable -->stop
    l--> id -->stop
```

```mermaid
graph LR
    stop([ ])
    boolean::---l[ ]--> true --> stop
    l-->false --> stop
```

```mermaid
graph LR
    stop([ ])
    string::--- l[#quot;] -->character --> r[#quot;]--> stop
```

```mermaid
graph LR
    stop([ ])
    character::---l[ ]--> UNICODE--> stop
    l-->prefix_character --> stop
```

```mermaid
graph LR
    stop([ ])
    prefix_character::--- l["\"]-->slash["\"]--> stop
    l-->q[#quot;]-->stop
    l-->n-->stop
    l-->t-->stop
```

```mermaid
graph LR
    stop([ ])
    variable:: ---word --> stop
```
