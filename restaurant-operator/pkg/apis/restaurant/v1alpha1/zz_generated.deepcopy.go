// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Course) DeepCopyInto(out *Course) {
	*out = *in
	if in.Allergens != nil {
		in, out := &in.Allergens, &out.Allergens
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Course.
func (in *Course) DeepCopy() *Course {
	if in == nil {
		return nil
	}
	out := new(Course)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Menu) DeepCopyInto(out *Menu) {
	*out = *in
	if in.Starters != nil {
		in, out := &in.Starters, &out.Starters
		*out = make([]Course, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Main != nil {
		in, out := &in.Main, &out.Main
		*out = make([]Course, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Desserts != nil {
		in, out := &in.Desserts, &out.Desserts
		*out = make([]Course, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Menu.
func (in *Menu) DeepCopy() *Menu {
	if in == nil {
		return nil
	}
	out := new(Menu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Restaurant) DeepCopyInto(out *Restaurant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Restaurant.
func (in *Restaurant) DeepCopy() *Restaurant {
	if in == nil {
		return nil
	}
	out := new(Restaurant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Restaurant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestaurantList) DeepCopyInto(out *RestaurantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Restaurant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestaurantList.
func (in *RestaurantList) DeepCopy() *RestaurantList {
	if in == nil {
		return nil
	}
	out := new(RestaurantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RestaurantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestaurantSpec) DeepCopyInto(out *RestaurantSpec) {
	*out = *in
	in.Menu.DeepCopyInto(&out.Menu)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestaurantSpec.
func (in *RestaurantSpec) DeepCopy() *RestaurantSpec {
	if in == nil {
		return nil
	}
	out := new(RestaurantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestaurantStatus) DeepCopyInto(out *RestaurantStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestaurantStatus.
func (in *RestaurantStatus) DeepCopy() *RestaurantStatus {
	if in == nil {
		return nil
	}
	out := new(RestaurantStatus)
	in.DeepCopyInto(out)
	return out
}
