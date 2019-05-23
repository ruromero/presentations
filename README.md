# Presentations

Resources and sample projects used in different presentations

## Ansible Operator Example

```{bash}
operator-sdk new restaurant-ansible-operator --api-version=restaurant.ruben.redhat.com/v1alpha1 --kind=Restaurant --type=ansible
```

## Helm Operator Example

```{bash}
operator-sdk new restaurant-helm-operator --api-version=restaurant.ruben.redhat.com/v1alpha1 --kind=Restaurant --type=helm
```

## Go Operator Example

```{bash}
operator-sdk new restaurant-go-operator
cd restaurant-go-operator
operator-sdk add api --api-version=restaurant.ruben.redhat.com/v1alpha1 --kind=Restaurant
```