# bazel_go_sample

bazel go sample

## Install bazel

for macOS
```
brew install bazelisk
```

## Usage

### run binary
```
bazelisk run //:gazelle-update-repos # create deps.bzl
bazelisk run //:gazelle # import go libraries for bazel
bazelisk run //:hello # execute binary
```

### run docker image

```
bazelisk run //:multi_stage_build
bazelisk run //:go_image
```
