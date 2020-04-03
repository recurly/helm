# Helm
This is a fork of the https://github.com/helm/helm project. We're forking to patch a bug in Tiller.

* [How to build for Recurly](#how-to-build-for-recurly)
* [Modifications to the Repo](#modifications-to-the-repo)
 
# How to build for Recurly <a name="how-to-build-for-recurly"/></a>

**Build locally**:
```bash
$ make bootstrap build
```

**Build the docker image**:
```bash
$ docker build .
```

# Modifications to the Repo <a name="modifications-to-the-repo"/></a>
Modifications to the original repo from most recent to oldest change:

* Set `staging` as the default branch
* Patched `pkg/kube/client.go` to include nil checks
* Created a new root `Dockerfile` that will build docker image
* Updated Makefile `build` target to build the docker image
* Checking in dependencies to avoid having to deal with failures
* Rewrote the readme to provide clear instructions for Recurly
* Removed some of the base docs to make is simpler to understand
* Created the new `staging` branch which is based off `release-2.16`
