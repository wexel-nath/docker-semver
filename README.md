# docker-semver [![Codefresh build status](https://g.codefresh.io/api/badges/pipeline/wexel/Utils%2Fsemver?branch=master&key=eyJhbGciOiJIUzI1NiJ9.NWVjODc4OTNhNjJlNzAwZDU0ZDUwYmIx.wCWD3CFfgsDRWQY-K2COubziUQaUEJuZ3sjKHQ1roEQ&type=cf-1)](https://g.codefresh.io/pipelines/semver/builds?repoOwner=wexel-nath&repoName=docker-semver&serviceName=wexel-nath%2Fdocker-semver&filter=trigger:build~Build;branch:master;pipeline:5edd7c3f8b63eaf20d2e9437~semver)
Automate docker image [Semantic Versioning](https://semver.org/) with commit messages.

Given a docker IMAGE_TAGS_URL, docker-semver will find the current patch semantic version,
then increment the version based on the provided GIT_COMMIT_MSG 

### Example
A docker repository has the following tags:
* 2.2.1, 2.2, 2, latest
* 1.7.10, 1.7, 1

A commit is pushed with the message: "MINOR: new functionality X"

The following tags will be returned by docker-semver
* MAJOR: 2
* MINOR: 2.3
* PATCH: 2.3.0

### Commit message formats
docker-semver reads the first word in the commit message. It is case-insensitive.
##### Examples
* MAJOR: incompatible API changes
* minor: add functionality in a backwards compatible manner
* patch some fixes <-- colon is optional
* some commit message <-- increments a patch version by default 

### Notes
* If no semantic versions are found, docker-semver will return the tags:
  * 1
  * 1.0
  * 1.0.1

* By default, docker-semver will fetch up to 100 tags from the IMAGE_TAGS_URL.
Set IMAGE_TAGS_LIMIT to override this.

### Usage
```sh
docker run \
    -e IMAGE_TAGS_URL='https://hub.docker.com/v2/repositories/wexel/docker-semver/tags' \
    -e GIT_COMMIT_MSG='some commit message' \
    wexel/docker-semver \
    docker-semver

// Output
1
1.0
1.0.2
```

##### Handling the result
```sh
TAGS=$(docker run \
    -e IMAGE_TAGS_URL='https://hub.docker.com/v2/repositories/wexel/docker-semver/tags' \
    -e GIT_COMMIT_MSG='some commit message' \
    wexel/docker-semver \
    docker-semver)
[ $(echo "$TAGS" | wc -w) != "3" ] && exit 1

for tag in $TAGS; do
    [ -z "$MAJOR_TAG" ] && export MAJOR_TAG="$tag" && continue
    [ -z "$MINOR_TAG" ] && export MINOR_TAG="$tag" && continue
    [ -z "$PATCH_TAG" ] && export PATCH_TAG="$tag" && break
done

echo $MAJOR_TAG // 1
echo $MINOR_TAG // 1.0
echo $PATCH_TAG // 1.0.2
```
