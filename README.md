# kubectl-plugin-builder

this is a scaffold tool for developing a kubectl plugin.  
You can see the official documentation [here](./docs/introduction.md).
This project is an experimental stage yet,  
but I'm welcome your contributions/reports!  

## How to use

```bash
$ git clone https://github.com/Drumato/kubectl-plugin-builder
$ cd kubectl-plugin-builder
$ make
$ mv bin/kubectl-plugin-builder <PATH> # /usr/bin, etc
```

## Features

- `kubectl-plugin-builder new` command for initialization of a plugin's project
- `kubectl-plugin-builder generate` command for code-generation declaratively

## Roadmap

- [ ] commands
  - [ ] `add CMD_NAME` ... add command that add a new command to the yaml
    - `--parent` flag
- [ ] api client support by default(`cli-runtime/pkg/client`)
- [ ] config flags support by default(`cli-runtime/pkg/genericclioptions`)
- [ ] auto testcode generation from yaml
- [ ] kubectl-plugin-builder's test

## [LICENSE](./LICENSE)

```license
MIT License

Copyright (c) 2021 Drumato
```
