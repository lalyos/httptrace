## Why

Httptrace library aims to make debugging http/https traffic tracing easy.

- Have you ever wished that your favourite golang tool, or library gave you more insight?
- Do you want to create a tool ot library and give users the ability to see all http traffic?

Read on ...

## Installation

Like a good gopher village citizen:
```
go get github.com/lalyos/httptrace
```

## Usage

Add one single extara `import` to your library, or to the tool you are hacking on:
```
import _ "github.com/lalyos/httptrace"
```

Now if you rebuild the tool. You are able to switch on the traffic tracer:

```
DEBUG=1 mytool
```
The tool will output to **stderr** all http traffic.


Or for permanently swith:
```
export DEBUG=1
mytool
```

## Configuration

If you don’t like the environment variable name `DEBUG` you can customize it:

```
import "github.com/lalyos/httptrace"

func init() {
	httptrace.DEBUG_ENV_NAME = "MYTOOL_TRACE"
}
```

Now you can switch the trace with the modified environment variable:
```
export MYTOOL_TRACE=1
mytool
```

