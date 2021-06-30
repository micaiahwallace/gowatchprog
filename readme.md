# GoWatchProg

## Project Status: ⚠️ early development

This module was created to help manage background services on machines by providing the following core feature sets.

- Install to a proper location based on installation context (user, system, allusers)
- Register service autostart based on service run context (user, system, allusers)
- Watchdog autorestart (retry count, retry delay, retry delay increase)
- Install updates remotely

## Install

```shell
go get github.com/micaiahwallace/gowatchprog
```

## Documentation

View the documentation on [pkg.go.dev](https://pkg.go.dev/github.com/micaiahwallace/gowatchprog)

## Example 

You can see gowatchprog in action on one of my other projects [GoScreenMonit](https://github.com/micaiahwallace/goscreenmonit)

## Roadmap

- [x] Windows support (Full context set not implemented for install / startup)
- [ ] Remote update feature
- [ ] Mac OS support
- [ ] Linux support
- [ ] Features accessible from command line utility (for non-go services)
