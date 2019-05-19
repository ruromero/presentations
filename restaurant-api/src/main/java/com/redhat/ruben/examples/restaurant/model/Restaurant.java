package com.redhat.ruben.examples.restaurant.model;

public class Restaurant {

    private String name;
    private String address;
    private String foodType;
    private String phoneNumber;

    public String getName() {
        return name;
    }

    public Restaurant setName(String name) {
        this.name = name;
        return this;
    }

    public String getAddress() {
        return address;
    }

    public Restaurant setAddress(String address) {
        this.address = address;
        return this;
    }

    public String getFoodType() {
        return foodType;
    }

    public Restaurant setFoodType(String foodType) {
        this.foodType = foodType;
        return this;
    }

    public String getPhoneNumber() {
        return phoneNumber;
    }

    public Restaurant setPhoneNumber(String phoneNumber) {
        this.phoneNumber = phoneNumber;
        return this;
    }

}
