# CLI version in the wild

This document describe what approach was implemented by some popular CLI to display its version.

Tested CLIs:
- [argo](https://github.com/argoproj/argo-workflows/blob/v3.3.9/version.go)
- [docker](https://github.com/docker/cli/blob/v20.10.17/cli/command/system/version.go)
- [eksctl](https://github.com/weaveworks/eksctl/blob/v0.108.0/cmd/eksctl/version.go)
- [fass-cli](https://github.com/openfaas/faas-cli/blob/0.14.5/commands/version.go)
- [gh](https://github.com/cli/cli/blob/v2.14.4/pkg/cmd/version/version.go)
- [golangci-lint](https://github.com/golangci/golangci-lint/blob/v1.48.0/pkg/commands/version.go)
- [helm](https://github.com/helm/helm/blob/v3.9.3/cmd/helm/version.go)
- [k3d](https://github.com/k3d-io/k3d/blob/v5.4.4/version/version.go)
- [kubectl](https://github.com/kubernetes/kubectl/blob/v0.24.4/pkg/cmd/version/version.go)
- [minikube](https://github.com/kubernetes/minikube/blob/v1.26.1/cmd/minikube/cmd/version.go)
- [porter](https://github.com/getporter/porter/blob/v1.0.0-beta.3/pkg/porter/version.go)
- [rancher](https://github.com/rancher/cli/blob/v2.6/main.go#L75)
- [terraform](https://github.com/hashicorp/terraform/blob/v1.2.8/internal/command/version.go)
- [vault](https://github.com/hashicorp/vault/blob/v1.11.2/command/version.go)

## How to get the version
<!-- 14 == 14 -->

### via command (only)

5 of 14
- [argo](https://github.com/argoproj/argo-workflows/blob/v3.3.9/version.go)
- [eksctl](https://github.com/weaveworks/eksctl/blob/v0.108.0/cmd/eksctl/version.go)
- [fass-cli](https://github.com/openfaas/faas-cli/blob/0.14.5/commands/version.go)
- [kubectl](https://github.com/kubernetes/kubectl/blob/v0.24.4/pkg/cmd/version/version.go)
- [minikube](https://github.com/kubernetes/minikube/blob/v1.26.1/cmd/minikube/cmd/version.go)

### via flag (only)

1 of 14

| CLI Name | Flags |
|----------|-------|
| [rancher](https://github.com/rancher/cli/blob/v2.6/main.go#L75) | `--version`, `-v` |


### via flag and command (both)

8 of 14

| CLI Name | Flags |
|----------|-------|
| [docker](https://github.com/docker/cli/blob/v20.10.17/cli/command/system/version.go) | `--version`, `-v` |
| [gh](https://github.com/cli/cli/blob/v2.14.4/pkg/cmd/version/version.go) | `--version` |
| [golangci-lint](https://github.com/golangci/golangci-lint/blob/v1.48.0/pkg/commands/version.go) | `--version` |
| [helm](https://github.com/helm/helm/blob/v3.9.3/cmd/helm/version.go) | `--version`, `-v` |
| [k3d](https://github.com/k3d-io/k3d/blob/v5.4.4/version/version.go) | `--version` |
| [porter](https://github.com/getporter/porter/blob/v1.0.0-beta.3/pkg/porter/version.go) | `--version`, `-v` |
| [terraform](https://github.com/hashicorp/terraform/blob/v1.2.8/internal/command/version.go) | `--version`, `-v` |
| [vault](https://github.com/hashicorp/vault/blob/v1.11.2/command/version.go) | `--version`, `-v` |

### Overall flag names popularity

| Name | Number of usage |
|-----|---------|
| `--version` | 9 |
| `-v` | 6 |

## Upgrade notice

## What data is collected

## Output

### goTemplate

- `docker version --format/-f`
  ```
  ```

- `helm version --template`
  ```
  ```

### json

- `eksctl version -ojson`
  ```json
  {"Version":"0.108.0","PreReleaseID":"","Metadata":{"BuildDate":"2022-08-12T08:12:00Z","GitCommit":"728d4ad"},"EKSServerSupportedVersions":["1.19","1.20","1.21","1.22"]}
  ```

- `golangci-lint version --format json`
  ```json
  {"version":"1.48.0","commit":"2d8fea8","date":"2022-08-04T09:19:19Z"}
  ```

- `kubectl version -ojson`
  ```json
  {
    "clientVersion": {
      "major": "1",
      "minor": "24",
      "gitVersion": "v1.24.0",
      "gitCommit": "4ce5a8954017644c5420bae81d72b09b735c21f0",
      "gitTreeState": "clean",
      "buildDate": "2022-05-03T13:46:05Z",
      "goVersion": "go1.18.1",
      "compiler": "gc",
      "platform": "darwin/amd64"
    },
    "kustomizeVersion": "v4.5.4",
    "serverVersion": {
      "major": "1",
      "minor": "23",
      "gitVersion": "v1.23.6+k3s1",
      "gitCommit": "418c3fa858b69b12b9cefbcff0526f666a6236b9",
      "gitTreeState": "clean",
      "buildDate": "2022-04-28T22:16:18Z",
      "goVersion": "go1.17.5",
      "compiler": "gc",
      "platform": "linux/amd64"
    }
  }
  ```

- `minikube version --ojson`
  ```json
  {"commit":"62e108c3dfdec8029a890ad6d8ef96b6461426dc","minikubeVersion":"v1.26.1"}
  ```

- `porter version -ojson`
  ```json
  {
    "name": "porter",
    "version": "v1.0.0-beta.3",
    "commit": "f7aa99dc"
  }
  ```

- `terraform version --json`
  ```json
  {
    "terraform_version": "1.2.8",
    "platform": "darwin_amd64",
    "provider_selections": {},
    "terraform_outdated": false
  }
  ```

### plain

- `argo version`
  ```text
  argo: v3.3.9+5db53aa.dirty
    BuildDate: 2022-08-10T02:08:30Z
    GitCommit: 5db53aa0ca54e51ca69053e1d3272e37064559d7
    GitTreeState: dirty
    GitTag: v3.3.9
    GoVersion: go1.19
    Compiler: gc
    Platform: darwin/amd64
  ```

- `docker version`
  ```text
  Client:
    Version:           20.10.9
    API version:       1.41
    Go version:        go1.16.8
    Git commit:        c2ea9bc
    Built:             Thu Nov 18 21:15:43 2021
    OS/Arch:           darwin/amd64
    Context:           default
    Experimental:      true
  
  Server: Docker Desktop 4.8.2 (79419)
    Engine:
      Version:          20.10.14
      API version:      1.41 (minimum version 1.12)
      Go version:       go1.16.15
      Git commit:       87a90dc
      Built:            Thu Mar 24 01:46:14 2022
      OS/Arch:          linux/amd64
      Experimental:     false
    containerd:
      Version:          1.5.11
      GitCommit:        3df54a852345ae127d1fa3092b95168e4a88e2f8
    runc:
      Version:          1.0.3
      GitCommit:        v1.0.3-0-gf46b6ba
    docker-init:
      Version:          0.19.0
      GitCommit:        de40ad0
  ```

- `gh version`
  ```text
  gh version 2.14.4 (2022-08-10)
  https://github.com/cli/cli/releases/tag/v2.14.4
  ```

- `golangci-lint version`
  ```text
  golangci-lint has version 1.48.0 built from 2d8fea8 on 2022-08-04T09:19:19Z
  ```

- `helm version`
  ```text
  version.BuildInfo{Version:"v3.9.3", GitCommit:"414ff28d4029ae8c8b05d62aa06c7fe3dee2bc58", GitTreeState:"clean", GoVersion:"go1.19"}
  ```

- `k3d version`
  ```text
  k3d version v5.4.4
  k3s version v1.23.8-k3s1 (default)
  ```

- `minikube version`
  ```text
  minikube version: v1.26.1
  commit: 62e108c3dfdec8029a890ad6d8ef96b6461426dc
  ```

- `porter version`
  ```text
  porter v1.0.0-beta.3 (f7aa99dc)
  ```

- `rancher --version`
  ```text
  rancher version v2.6.7
  ```

- `terraform version`
  ```text
  Terraform v1.2.8
  on darwin_amd64
  ```

- `vault version`
  ```text
  Vault v1.11.2 (3a8aa12eba357ed2de3192b15c99c717afdeb2b5), built 2022-07-29T09:48:47Z
  ```

### short

- `argo version --short`
  ```text
  argo: v3.3.9+5db53aa.dirty
  ```

- `docker --version/-v`
  ```text
  Docker version 20.10.9, build c2ea9bc
  ```

- `eksctl version`
  ```text
  0.108.0
  ```

- `fass-cli version --short-version`
  ```text
  0.14.5
  ```

- `golangci-lint version --format short`
  ```text
  1.48.0
  ```

- `helm version --short`
  ```text
  ```

- `kubectl version`
  ```text
  Client Version: v1.24.0
  Kustomize Version: v4.5.4
  ```

- `minikube version --short`
  ```text
  ```

### system

- `minikube version --components -oyaml`
  ```
  buildctl: buildctl github.com/moby/buildkit v0.10.3 c8d25d9a103b70dc300a4fd55e7e576472284e31
  commit: 62e108c3dfdec8029a890ad6d8ef96b6461426dc
  containerd: containerd containerd.io 1.6.6 10c12954828e7c7c9b6e0ea9b0c02b01407d3ae1
  cri-dockerd: ""
  crictl: crictl version v1.21.0
  crio: crio version 1.24.1
  crun: crun version UNKNOWN
  ctr: ctr containerd.io 1.6.6
  docker: Docker version 20.10.17, build 100c701
  dockerd: Docker version 20.10.17, build a89b842
  minikubeVersion: v1.26.1
  podman: podman version 3.4.2
  runc: runc version 1.1.2
  ```

- `porter version --system`
  ```
  ```

### yaml

- `kubectl version -oyaml`
  ```yaml
  clientVersion:
    buildDate: "2022-05-03T13:46:05Z"
    compiler: gc
    gitCommit: 4ce5a8954017644c5420bae81d72b09b735c21f0
    gitTreeState: clean
    gitVersion: v1.24.0
    goVersion: go1.18.1
    major: "1"
    minor: "24"
    platform: darwin/amd64
  kustomizeVersion: v4.5.4
  serverVersion:
    buildDate: "2022-04-28T22:16:18Z"
    compiler: gc
    gitCommit: 418c3fa858b69b12b9cefbcff0526f666a6236b9
    gitTreeState: clean
    gitVersion: v1.23.6+k3s1
    goVersion: go1.17.5
    major: "1"
    minor: "23"
    platform: linux/amd64
  ```

- `minikube version --oyaml`
  ```yaml
  commit: 62e108c3dfdec8029a890ad6d8ef96b6461426dc
  minikubeVersion: v1.26.1
  ```

----

### JSON

### YAML

### Go template

### Short

1 of 13

- `kubectl version --short` (deprecated, will be default in the future)
  ```
  Client Version: v1.24.0
  Kustomize Version: v4.5.4
  ```

## Frameworks built-in support

### Cobra

```go
cmd := &cobra.Command{
  Use:   "example",
  Short: "An example CLI built with github.com/spf13/cobra",

  // Version defines the version for this command. If this value is non-empty and the command does not
  // define a "version" flag, a "version" boolean flag will be added to the command and, if specified,
  // will print content of the "Version" variable. A shorthand "v" flag will also be added if the
  // command does not define one.
  Version: "1.0.3",
}
```

###

Interesting approach to print the version of deps: https://github.com/hashicorp/terraform/blob/v1.3.0-dev/version/dependencies.go
