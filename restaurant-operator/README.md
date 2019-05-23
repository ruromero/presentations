# Instructions

## Create the Operator skeleton

```{bash}
operator-sdk new restaurant-operator
```

## Add a new Custom Resource Definition

```{bash}
operator-sdk add api --api-version=restaurant.ruben.redhat.com/v1alpha1 --kind=Restaurant
```

## Define the Spec and Status

In `restaurant_types.go`

```{go}
type RestaurantSpec struct {
    ...
}
```

## Generate CRD

Runs the kube-openapi OpenAPIv3 code generator for all Custom Resource Definition (CRD) API tagged fields under pkg/apis/....

Note: This command must be run every time a tagged API struct or struct field for a custom resource type is updated.

```{bash}
operator-sdk generate openapi
```

## Generate K8S

After modifying the *_types.go file always run the following command to update the generated code for that resource type

```{bash}
operator-sdk generate k8s
```

## Add a new Controller

```{bash}
operator-sdk add controller --api-version=restaurant.ruben.redhat.com/v1alpha1 --kind=Restaurant
```

## Implement Operator

* Update pod definition
* Create service
* Create route/ingress
* Create configmap
* Create deployment
* Watch for changes in Deployment/Service/Configmap

* Update status with restaurant //TODO

## Build operator

```{bash}
operator-sdk build quay.io/ruben/restaurant-operator:0.1.0
```

Replace the IMAGE placeholder

```{bash}
sed -i 's|REPLACE_IMAGE|quay.io/ruben/restaurant-operator:0.1.0|g' deploy/operator.yaml
```

Push the image to the registry

```{bash}
docker push quay.io/ruben/restaurant-operator:0.1.0
```

## Deployment

### Kubernetes

Prerequisites:

* Kubernetes cluster (e.g. Minikube) with an ingress controller
* OLM installed
* OLM console running

```{bash}
minikube start
minikube addons enable ingress
```

or

```{bash}
make run-local
```

### Openshift

Prerequisites:

* Openshift 3.11 (with OLM) or 4.x

## OLM - CSV Generation

```{bash}
operator-sdk olm-catalog gen-csv --csv-version 0.1.0 --update-crds
```

## Demo resources

Watch for the phoneNumber to be updated

```{bash}
url=`kubectl get ingress --no-headers | awk '{print$2}'`
watch "curl -s $url/api/ | jq -r .phoneNumber "
```