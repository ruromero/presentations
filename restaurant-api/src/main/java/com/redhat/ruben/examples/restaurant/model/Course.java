package com.redhat.ruben.examples.restaurant.model;

import java.util.Set;

public class Course {

    private String name;
    private String description;
    private Set<String> allergens;

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

    public Set<String> getAllergens() {
        return allergens;
    }

    public void setAllergens(Set<String> allergens) {
        this.allergens = allergens;
    }
}
