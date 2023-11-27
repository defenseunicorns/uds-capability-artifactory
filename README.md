# uds-capability-artifactory

Bigbang [artifactory](https://repo1.dso.mil/big-bang/apps/third-party/jfrog-platform) deployed via flux by zarf

## Deployment Prerequisites

### Artifactory Capability

The Artifactory Capability expects the pieces listed below to exist in the cluster before being deployed.

#### General

- Create `artifactory` namespace
- Label `artifactory` namespace with `istio-injection: enabled`

#### Database

- A Postgres database is running on port `5432` and accessible to the cluster
- This database can be logged into via the username `artifactory`
- This database instance has a psql database created named `artifactorydb`
- The `artifactory` user has read/write access to the above mentioned database
- Create `artifactory-postgres` service in `artifactory` namespace that points to the psql database
- Create `artifactory-postgres` secret in `artifactory` namespace with the key `password` that contains the password to the `artifactory` user for the psql database

## Deploying

### Deploy Everything

#### Via Makefile and local package

```bash
# This will run build/all cluster/reset and deploy/all. Follow the breadcrumbs in the Makefile to see what and how its doing it.
make all
```

### From GHCR OCI Via Zarf

```bash
zarf package deploy ghcr.io/defenseunicorns/uds-capability/artifactory:x.x.x-amd64
```

## Building

### Use zarf to login to the needed registries i.e. registry1.dso.mil

```bash
# Download Zarf
make build/zarf

# Login to the registry
set +o history

# registry1.dso.mil (To access registry1 images needed during build time)
export REGISTRY1_USERNAME="YOUR-USERNAME-HERE"
export REGISTRY1_TOKEN="YOUR-TOKEN-HERE"
echo $REGISTRY1_TOKEN | build/zarf tools registry login registry1.dso.mil --username $REGISTRY1_USERNAME --password-stdin

set -o history
```

### Creating the Package

```bash
make build/uds-capability-artifactory
```
