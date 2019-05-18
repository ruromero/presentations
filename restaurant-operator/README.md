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

## Add a new Controller

```{bash}
operator-sdk add controller --api-version=restaurant.ruben.redhat.com/v1alpha1 --kind=Restaurant
```