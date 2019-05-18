package com.redhat.ruben.examples.restaurant.model;

public class Restaurant {

    private String name;
    private String location;
    private String foodType;
    private String contact;

    public String getName() {
        return name;
    }

    public Restaurant setName(String name) {
        this.name = name;
        return this;
    }

    public String getLocation() {
        return location;
    }

    public Restaurant setLocation(String location) {
        this.location = location;
        return this;
    }

    public String getFoodType() {
        return foodType;
    }

    public Restaurant setFoodType(String foodType) {
        this.foodType = foodType;
        return this;
    }

    public String getContact() {
        return contact;
    }

    public Restaurant setContact(String contact) {
        this.contact = contact;
        return this;
    }

}
