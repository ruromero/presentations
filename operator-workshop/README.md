# Operators workshop

## Requirements

* Git
* Go version 1.12.x
* operator-sdk version v0.8.x
* docker version 17.03+ or buildah version 1.8.2+
* kubectl version v1.11.3+.
* Access to a Kubernetes v1.11.3+ cluster

Note: This guide uses [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/#installation) version v1.0.1+ as the local Kubernetes cluster and [quay.io](quay.io) for the public registry.

### Enable the NGINX Ingress controller

```{bash}
$ minikube addons enable ingress
✅  ingress was successfully enabled
```

## Implement your own operator

### Create your first go operator

Note: You can chose between using dep or modules. Modules are used by default. In next versions it will also be added support by default for external vendoring (i.e. `--vendor=false`).

```{bash}
operator-sdk new restaurant-operator
cd restaurant-operator
```

After that review the generated code and resources. Pay special attention to the following:

* deploy folder containing yaml files to easily run your operator on kubernetes
* cmd/manager The manager register the scheme for all custom resources under `pkg/apis` and run all controlers under `pkg/controller`
* pkg/controller The controllers implementation
* pkg/apis The custom types definitions

Check the [project scaffolding layout](https://github.com/operator-framework/operator-sdk/blob/master/doc/project_layout.md) for more details

### Let's add a new CRD

```{bash}
operator-sdk add api --api-version=restaurant.workshop.redhat.com/v1aplpha1 --kind=Restaurant
```

After that we can see the following changed:

* deploy/crds is created with a basic CRD (including OpenApiV3Schema validation) and a sample CR
* pkg/apis now has a folder containing a file `<my_kind>_types.go` with the CRD struct

### Create the controller for our CRD

```{bash}
operator-sdk add controller --api-version=restaurant.workshop.redhat.com/v1alpha1 --kind=Restaurant
```

Let's review:

* `pkg/controller/<my_kind>/<my_kind>_controller.go` has a default implementation that creates a `busybox` pod

### Define the CRD

Go to the `pkg/apis/<my_kind>/<my_kind>_types.go` file and define your custom resource.

When finished, run the following to update the generated code:

```{bash}
operator-sdk generate k8s
```

Note: Run this command each time you change the types

### Generate the OpenAPI validation

```{bash}
operator-sdk generate openapi
```

Verify that your custom resource definition has a generated.

Note: Run this command only when you want to replace the CRD with the new generated values. **This will overwrite any customization**

### Implement the controller

1. Watch for changes in owned objects: Just add here the ones you feel necessary.
1. Implement the reconcile loop: You will receive notification changes for all the watched objects (CustomResources and OwnedObjects). Create the owned objects that are missing or replace them if the expected objects don't match with the existing ones.
1. Update the `role.yaml` in case different permissions are required (e.g. create `ingress` or `route`)

### Run the Operator locally

```{bash}
kubectl create -f deploy/crds/restaurant_v1alpha1_restaurant_crd.yaml
go mod vendor
operator-sdk up local --namespace=default
```

Now you can create an example CR. Notice that the generated CR is not valid.

### Build the Operator

```{bash}
$ operator-sdk build --image-builder buildah quay.io/ruben/restaurant-operator:latest
...
INFO[0008] Operator build complete.
```

Push the image to your quay.io repository (make sure the repository is public for this workshop)

```{bash}
podman push quay.io/ruben/restaurant-operator:latest
```

### Deploy the Operator in your Kubernetes cluster

Before deploying the Operator make sure you replaced all the placeholders:

```{bash}
sed -i 's|REPLACE_IMAGE|quay.io/ruben/restaurant-operator:latest|g' deploy/operator.yaml
```

Create the CRD (if you haven't done it yet)

```{bash}
kubectl create -f deploy/crds/restaurant_v1alpha1_restaurant_crd.yaml
operator-sdk up local --namespace=default
```

Create all the necessary resources

```{bash}
$ kubectl create -f deploy
deployment.apps/restaurant-operator created
role.rbac.authorization.k8s.io/restaurant-operator created
rolebinding.rbac.authorization.k8s.io/restaurant-operator created
serviceaccount/restaurant-operator created
```

## Advanced topics

* [Testing](Testing.md)
* [OLM integration](OLM.md)

## Useful commands

* To show which Kind belongs to which apiGroup.

```{bash}
$ kubectl api-resources
NAME                              SHORTNAMES   APIGROUP                         NAMESPACED   KIND
...
configmaps                        cm                                            true         ConfigMap
pods                              po                                            true         Pod
podtemplates                                                                    true         PodTemplate
deployments                       deploy       apps                             true         Deployment
ingresses                         ing          extensions                       true         Ingress
roles                                          rbac.authorization.k8s.io        true         Role
restaurants                                    restaurant.workshop.redhat.com   true         Restaurant
```

* To show all the Api Versions

```{bash}
$ kubectl api-versions
...
apps/v1
...
rbac.authorization.k8s.io/v1
rbac.authorization.k8s.io/v1beta1
restaurant.workshop.redhat.com/v1alpha1
...
v1
```
