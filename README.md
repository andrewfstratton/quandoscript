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
    sequence::---line--> nl["↵"]--> stop
    nl-->line
```

```mermaid
graph LR
    stop([ ])
    line::---id--> nl["↵"]-->stop
    id-->spacer-->action-->nl
```

```mermaid
graph LR
    stop([ ])
    id::---digit-->stop
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
    action::--- word--> params --> stop
```

```mermaid
graph LR
    stop([ ])
    word:: --- letter-->stop
    letter-->l( )-->rl["letter"]-->r( )-->stop
    r-->l
    l-->dot["."]-->r
    l-->u["_"]-->r
    l-->digit-->r
```

```mermaid
graph LR
    stop([ ])
    letter::---l( )--> a..z--> stop
    l--> A..Z--> stop
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
    param:: --- word --> value-->stop
    word --> variable --> stop
```

```mermaid
graph LR
    stop([ ])
    variable:: ---l["="] --> word --> stop
```

```mermaid
graph LR
    stop([ ])
    value:: --- l[ ]
    l--> b[!]--> boolean -->stop
    l--> q[#quot;] -->characters --> qr[#quot;] -->stop
    l--> h["\#"]-->FLOATING_POINT -->stop
    l--> c[":"]--> id -->stop
```

```mermaid
graph LR
    stop([ ])
    boolean::---l( )--> true --> stop
    l-->false --> stop
```

```mermaid
graph LR
    stop([ ])
    characters::---l[ ]--> UNICODE-->r[ ]--> stop
    l-->prefix_character -->r
    r-->l
```

```mermaid
graph LR
    stop([ ])
    prefix_character::--- l["\"]-->slash["\"]--> stop
    l-->q[#quot;]-->stop
    l-->n-->stop
    l-->t-->stop
```
