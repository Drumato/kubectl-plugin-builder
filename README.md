# kubectl-plugin-builder

this is a scaffold tool for developing a kubectl plugin.  

## Roadmap

- [ ] generate source-license-header at `kubectl-plugin-builder generate`
- [ ] commands
  - [ ] `add CMD_NAME` ... add command that add a new command to the yaml
    - `--parent` flag
- [ ] set cmd aliases at the generation
- [ ] set cmd args at the generation
- [ ] api client support by default(`cli-runtime/pkg/client`)
- [ ] config flags support by default(`cli-runtime/pkg/genericclioptions`)
- [ ] auto testcode generation from yaml
- [ ] kubectl-plugin-builder's test
