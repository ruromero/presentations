package com.redhat.ruben.examples.restaurant.model;

import java.util.List;

public class Menu {

    private List<Course> starters;
    private List<Course> main;
    private List<Course> desserts;

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

}
