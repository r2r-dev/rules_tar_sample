# `rules_tar` example repository
This project serves as an example of how to build a Golang project using `Bazel`, `Gazelle` and optionally `Nix`.

Project consists of two main parts. First a cli tool for creating a tar file from a list of inputs, and printing out file sizes. Secondly two Bazel rules utilizing said cli tool.

## Prerequisites:
This project requires UNIX-based environment and Bazel `v4.2.1`.
Depending on your host OS setup, you can fetch all of necessary dependencies using one of the following methods
#### Nix/NixOS: 
```
nix-shell
```
####  Docker/Bazelisk: 
```
docker run -v $(pwd):$(pwd) -w $(pwd) -it nickdecooman/debian-bazelisk:62443c9 bash
```

## cli tool
You can build and run a standalone version of cli tool using:
```
bazel run //cmd:tarchiver --         --help
bazel run //cmd:tarchiver -- archive --help
bazel run //cmd:tarchiver -- size    --help
```

## Bazel rules
Additionally, this project implements two simple rules utilizing the above cli tool.

### pkg_tar
a build rule accepting source files as input and producing a tar archive as an output, with optional prefixed directory structure. This rule is based on `archive` subcommand of `tararchiver` utility.

*Example*: 
```
bazel build //examples:archive
tar tvf bazel-bin/examples/archive.tar
```

#### file_size
 an executable rule, accepting a single source file and printing out its size in bytes. This rule is based on `size` subcommand of `tararchiver utility`.

*Example*:
```
bazel run //examples:size
```
