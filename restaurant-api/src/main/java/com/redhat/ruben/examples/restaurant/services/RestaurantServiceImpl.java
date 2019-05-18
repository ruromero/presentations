package com.redhat.ruben.examples.restaurant.services;

import javax.annotation.PostConstruct;
import javax.enterprise.context.ApplicationScoped;

import com.redhat.ruben.examples.restaurant.model.Restaurant;
import org.eclipse.microprofile.config.inject.ConfigProperty;

@ApplicationScoped
public class RestaurantServiceImpl implements RestaurantService {

    @ConfigProperty(name = "restaurant_name", defaultValue = "Quarkus Deli")
    String name;

    @ConfigProperty(name = "restaurant_location", defaultValue = "Granada")
    String location;

    @ConfigProperty(name = "restaurant_food_type", defaultValue = "Fat free Java apps")
    String foodType;

    @ConfigProperty(name = "restaurant_contact", defaultValue = "958 66 24 42")
    String contact;

    private Restaurant restaurant;

    @PostConstruct
    public void init() {
        this.restaurant = new Restaurant()
                .setName(name)
                .setLocation(location)
                .setFoodType(foodType)
                .setContact(contact);
    }

    @Override
    public Restaurant get() {
        return restaurant;
    }
}
