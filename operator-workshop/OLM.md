# Operator Lifecycle Manager

## Install OLM Manually

```{bash}
go get https://github.com/operator-framework/operator-lifecycle-manager.git
cd $GOPATH/src/github.com/operator-framework/operator-lifecycle-manager
```

From [Installing OLM](https://github.com/operator-framework/operator-lifecycle-manager/blob/master/Documentation/install/install.md)

```{bash}
kubectl create -f deploy/upstream/quickstart/crds.yaml
kubectl create -f deploy/upstream/quickstart/olm.yaml
```

### Start OLM console

```{bash}
$ make run-console-local
build flag -mod=vendor only valid when using modules
Running script to run the OLM console locally:
. ./scripts/run_console_local.sh
Trying to pull repository quay.io/openshift/origin-console ...
sha256:0e5a314d067862b9974d6bed6fb6e071c551820c34f02d0a4f4977fda2d6d1db: Pulling from quay.io/openshift/origin-console
Digest: sha256:0e5a314d067862b9974d6bed6fb6e071c551820c34f02d0a4f4977fda2d6d1db
Status: Image is up to date for quay.io/openshift/origin-console:latest
Using https://192.168.39.170:8443
The OLM is accessible via web console at:
https://localhost:9000/
```

Note: I had to remove the `-v` flag in the `command` command and it doesn't listen to `https` but to `http`

## Generate the Cluster Service Version

```{bash}
operator-sdk olm-catalog gen-csv --csv-version 0.1.0 --update-crds
```

This command will generate the csv with the version provided and will generate the crds under `deploy/olm-catalog`

Check [generating-a-csv](https://github.com/operator-framework/operator-sdk/blob/master/doc/user/olm-catalog/generating-a-csv.md) and [building-your-csv](https://github.com/operator-framework/operator-lifecycle-manager/blob/master/Documentation/design/building-your-csv.md) for more details

The resulting CSV will take the information from the following files (can be customized):

* Roles: role.yaml
* Deployments: operator.yaml
* Custom Resource (CR's): crds/<group>_<version>_<kind>_cr.yaml
* Custom Resource Definitions (CRD's): crds/<group>_<version>_<kind>_crd.yaml

## Customize your Cluster Service Version

Required fields that MUST be filled in:

* spec.customresourcedefinitions.owned[*].description
* spec.customresourcedefinitions.owned[*].displayName
* spec.minKubeVersion (can be 0.0.0 or a specific version like v1.14.0)

Recommended fields to fill in:

* spec.customresourcedefinitions.owned[*].resources: Resources to show in the console
* spec.customresourcedefinitions.owned[*].statusDescriptors and specDescriptors: Descriptors are intended to indicate various properties of a field in order to make decisions about their content. For example, this can drive connecting two operators together (e.g. connecting the connection string from a mysql instance to a consuming application) and be used to drive rich interactions in a UI.
* spec.description: Markdown where you can describe what the operator does and how to use it
* spec.icon: Find a fancy one
* spec.keywords: To help users find your operator
* spec.links: Static links to the documentation or your repository
* spec.installModes: What mode of installation namespacing OLM should use. MultiNamespace is not yet supported by SDK Operators
* spec.maintainers: Human or organiztional entities maintaining the Operator
* spec.provider: Usually an organization
* spec.maturity: The Operator's maturity

Find all in the [docs](https://github.com/operator-framework/operator-sdk/blob/master/doc/user/olm-catalog/generating-a-csv.md#csv-fields)

Note: Check how your CSV looks like in the OLM using the [OperatorHub's Preview](https://operatorhub.io/preview)

## Deploy your Cluster Resource Version manually

### Create the OperatorGroup

From [Operator Groups](https://github.com/operator-framework/operator-lifecycle-manager/blob/master/Documentation/design/operatorgroups.md)

An OperatorGroup is an OLM resource that provides rudimentary multitenant configuration to OLM installed operators.

The following command creates an Operator group whose only target is the current namespace

```{bash}
cat <<EOF | kubectl create -f -
kind: OperatorGroup
apiVersion: operators.coreos.com/v1
metadata:
  name: myoperatorgroup
spec:
  targetNamespaces:
  - `kubectl get sa default -o jsonpath='{.metadata.namespace}'`
EOF
```

### Satisfy all the pre-requisites

* minKubeVersion
* CRD is present
* Role
* Role Binding
* Service Account

```{bash}
kubectl create -f deploy/crds/restaurant_v1alpha1_restaurant_crd.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml
kubectl create -f deploy/service_account.yaml
```

### Create the Cluster Resource Version

```{bash}
kubectl create -f deploy/olm-catalog/restaurant-operator/0.1.0/restaurant-operator.v0.1.0.clusterserviceversion.yaml
```

## Create a new Custom Resource

* Manually

```{bash}
kubectl create -f deploy/crds/restaurant_v1alpha1_restaurant_cr.yaml
```

* In the console

Installed Operators -> Restaurant Operator -> Create New

The example listed is the one present in alm-examples, generated from `deploy/crds/restaurant_v1alpha1_restaurant_cr.yaml`

## Build your own Catalog of Operators

* Clone the [operator-registry](https://github.com/operator-framework/operator-registry) project.
* Create the package file

```{yaml}
packageName: restaurant
channels:
- name: alpha
  currentCSV: restaurant-operator.v0.1.0
defaultChannel: alpha
```

* Copy your operator bundle (package + versioned CRDs and CSVs) to the `manifests` folder
* Build the registry

```{bash}
buildah build-using-dockerfile -t quay.io/ruben/workshop-catalog:latest -f upstream-example.Dockerfile .
```

* Push it to your quay.io repository (make the repository public)

```{bash}
podman push quay.io/ruben/workshop-catalog:latest
```

* Create the `CatalogSource`

```{yaml}
apiVersion: operators.coreos.com/v1alpha1
kind: CatalogSource
metadata:
  name: workshop-catalog
  namespace: default
spec:
  sourceType: grpc
  image: quay.io/ruben/workshop-catalog:v1
```

* Verify the `packagemanifest` for `restaurant` exists

```{bash}
kubectl get packagemanifests
```

* Create a subscription

```{yaml}
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: restaurant-subscription
  namespace: default
spec:
  channel: alpha
  name: restaurant
  installPlanApproval: Automatic
  source: workshop-catalog
  sourceNamespace: default
```

The subscription will create the CSV, Operator group, role, role binding and service account in your namespace

## Upgrade to a new version

### Generate the new Cluster Resource Version

```{bash}
operator-sdk olm-catalog gen-csv --csv-version 0.2.0 --from-version 0.1.0 --update-crds
```

Pay attention at the new folder `deploy/olm-catalog/0.2.0` and in the CSV there is the following attribute `replaces: restaurant-operator.v0.1.0`

I have updated the `restaurant.package.yaml` file as follows:

```{yaml}
packageName: restaurant
channels:
- name: alpha
  currentCSV: restaurant-operator.v0.2.0
- name: stable
  currentCSV: restaurant-operator.v0.1.0
defaultChannel: alpha
```

That means v0.2.0 is the alpha channel with the most recent version and v0.1.0 is the stable version. By default the alpha channel will be used.

I will create a new catalog registry with the updated bundle and update the `CatalogSource` spec.image to point to `quay.io/ruben/workshop-catalog:v2`

If the Approval is set to `Automatic` the operator will be automatically upgraded to the new version. If not, you will have to review the install plan and approve it.
