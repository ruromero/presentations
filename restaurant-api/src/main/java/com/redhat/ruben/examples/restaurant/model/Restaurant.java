package com.redhat.ruben.examples.restaurant.model;

public class Restaurant {

    private String name;
    private String location;

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
}
