package restaurant

import (
	"context"

	"github.com/go-logr/logr"
	v1 "github.com/ruromero/presentations/restaurant-operator/pkg/apis/restaurant/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_restaurant")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Restaurant Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileRestaurant{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("restaurant-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Restaurant
	err = c.Watch(&source.Kind{Type: &v1.Restaurant{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Restaurant
	ownedObjects := []runtime.Object{
		&appsv1.Deployment{},
		&corev1.Service{},
	}
	for _, ownedObject := range ownedObjects {
		err = c.Watch(&source.Kind{Type: ownedObject}, &handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &v1.Restaurant{},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileRestaurant{}

// ReconcileRestaurant reconciles a Restaurant object
type ReconcileRestaurant struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Restaurant object and makes changes based on the state read
// and what is in the Restaurant.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileRestaurant) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Restaurant")

	// Fetch the Restaurant instance
	instance := &v1.Restaurant{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Reconcile the Deployment object
	result, err := r.reconcileDeployment(instance, reqLogger)
	if err != nil {
		return result, err
	}

	// Reconcile the Service object
	result, err = r.reconcileService(instance, reqLogger)
	if err != nil {
		return result, err
	}
	return reconcile.Result{}, nil
}

func (r *ReconcileRestaurant) reconcileDeployment(cr *v1.Restaurant, reqLogger logr.Logger) (reconcile.Result, error) {
	// Define a new Deployment object
	deployment := newDeploymentForCR(cr)

	// Set Restaurant instance as the owner and controller
	if err := controllerutil.SetControllerReference(cr, deployment, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Deployment already exists
	found := &appsv1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
		err = r.client.Create(context.TODO(), deployment)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Deployment created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Deployment already exists - don't requeue
	reqLogger.Info("Skip reconcile: Deployment already exists", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
	return reconcile.Result{}, nil
}

func (r *ReconcileRestaurant) reconcileService(cr *v1.Restaurant, reqLogger logr.Logger) (reconcile.Result, error) {
	// Define a new Service object
	service := newServiceForCR(cr)

	// Set Restaurant instance as the owner and controller
	if err := controllerutil.SetControllerReference(cr, service, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Service already exists
	found := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Service", "Service.Namespace", service.Namespace, "Service.Name", service.Name)
		err = r.client.Create(context.TODO(), service)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Service created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Service already exists - don't requeue
	reqLogger.Info("Skip reconcile: Service already exists", "Service.Namespace", found.Namespace, "Service.Name", found.Name)
	return reconcile.Result{}, nil
}

// newDeploymentForCR returns a deployment with the same name/namespace as the cr
func newDeploymentForCR(cr *v1.Restaurant) *appsv1.Deployment {
	replicas := cr.Spec.Deployment.Replicas
	if replicas == 0 {
		replicas = 1
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
			Labels: map[string]string{
				"app": cr.Name,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: getLabels(cr),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: getLabels(cr),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "restaurant",
							Image:   "quay.io/ruben/restaurant-api:latest",
							Command: []string{"./application", "-Dquarkus.http.host=0.0.0.0"},
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: 8080,
								},
							},
							Env: []corev1.EnvVar{
								{
									Name:  "RESTAURANT_NAME",
									Value: cr.Spec.Name,
								}, {
									Name:  "RESTAURANT_LOCATION",
									Value: cr.Spec.Location,
								}, {
									Name:  "RESTAURANT_FOOD_TYPE",
									Value: cr.Spec.FoodType,
								}, {
									Name:  "RESTAURANT_CONTACT",
									Value: cr.Spec.Contact,
								},
							},
							Resources: cr.Spec.Deployment.Resources,
						},
					},
				},
			},
		},
	}
}

func newServiceForCR(cr *v1.Restaurant) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
			Labels:    getLabels(cr),
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"deployment": cr.Name,
			},
			Ports: []corev1.ServicePort{{
				Name:       "http",
				Port:       8080,
				TargetPort: intstr.FromInt(8080),
			}},
		},
	}
}

func getLabels(cr *v1.Restaurant) map[string]string {
	return map[string]string{
		"app":        cr.Name,
		"deployment": cr.Name,
	}
}
