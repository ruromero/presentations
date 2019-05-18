package com.redhat.ruben.examples.restaurant.model;

import java.util.List;

public class Menu {

    private String name;
    private String description;
    private List<Course> starters;
    private List<Course> main;
    private List<Course> desserts;
    private Float price;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public List<Course> getStarters() {
        return starters;
    }

    public void setStarters(List<Course> starters) {
        this.starters = starters;
    }

    public List<Course> getMain() {
        return main;
    }

    public void setMain(List<Course> main) {
        this.main = main;
    }

    public List<Course> getDesserts() {
        return desserts;
    }

    public void setDesserts(List<Course> desserts) {
        this.desserts = desserts;
    }

    public Float getPrice() {
        return price;
    }

    public void setPrice(Float price) {
        this.price = price;
    }
}
