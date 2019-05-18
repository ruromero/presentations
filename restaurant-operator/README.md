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

```{bash}
operator-sdk generate openapi
```

## Generate K8S

```{bash}
operator-sdk generate k8s
```

## Add a new Controller

```{bash}
operator-sdk add controller --api-version=restaurant.ruben.redhat.com/v1alpha1 --kind=Restaurant
```

## Implement Operator

## Build operator

## OLM - CSV Generation

```{bash}
operator-sdk olm-catalog gen-csv --csv-version 0.1.0 --update-crds
```