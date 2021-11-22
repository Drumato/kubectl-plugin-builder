# kubectl-plugin-builder

this is a scaffold tool for developing a kubectl plugin.  

## Roadmap

- [ ] commands
  - [x] new command
  - [x] generate command
    - [ ] support nested command architecture
    - [ ] refactoring
  - [ ] add command that add a new command to the yaml
    - [ ] `--parent` flag
- [ ] api client support by default(`cli-runtime/pkg/client`)
- [ ] auto testcode gereration from yaml
