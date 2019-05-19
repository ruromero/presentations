package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RestaurantSpec defines the desired state of Restaurant
// +k8s:openapi-gen=true
type RestaurantSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Info       Info           `json:"info"`
	Menu       Menu           `json:"menu"`
	Deployment DeploymentSpec `json:"deployment"`
}

// DeploymentSpec defines how the application will be deployed in the cluster
type DeploymentSpec struct {
	HostnameSuffix string                      `json:"hostnameSuffix"`
	Replicas       int32                       `json:"replicas,omitempty"`
	Resources      corev1.ResourceRequirements `json:"resources,omitempty"`
}

// Info defines the Restaurant
type Info struct {
	Name        string `json:"name"`
	FoodType    string `json:"foodType" yaml:"foodType"`
	PhoneNumber string `json:"phoneNumber,omitempty" yaml:"phoneNumber,omitempty"`
	Address     string `json:"address,omitempty"`
}

// Menu defines what is served in the restaurant
type Menu struct {
	Starters []Course `json:"starters"`
	Main     []Course `json:"main"`
	Desserts []Course `json:"desserts"`
}

// Course defines each individual meal
type Course struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Price       float32  `json:"price"`
	Allergens   []string `json:"allergens,omitempty"`
}

// RestaurantStatus defines the observed state of Restaurant
// +k8s:openapi-gen=true
type RestaurantStatus struct {

	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Restaurant is the Schema for the restaurants API
// +k8s:openapi-gen=true
type Restaurant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestaurantSpec   `json:"spec,omitempty"`
	Status RestaurantStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RestaurantList contains a list of Restaurant
type RestaurantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Restaurant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Restaurant{}, &RestaurantList{})
}
