# Gourmet: a gourmet X11 swallower
Gourmet hides your current window before launching an external program and unhides it after quitting.
Gourmet was inspired by [devour](https://github.com/salman-abedin/devour).

## Dependencies

- Go (build)
- Xlib (client-side header files)

## Usage

```
$ gourmet <cmd> <args>
```
or
```
$ gourmet -s <cmd> <args>
```

## Installation

#### Git

```
$ git clone https://github.com/thalting/gourmet.git && cd gourmet && go install
```
