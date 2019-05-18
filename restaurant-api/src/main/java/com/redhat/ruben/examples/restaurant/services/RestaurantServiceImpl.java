package com.redhat.ruben.examples.restaurant.services;

import javax.annotation.PostConstruct;
import javax.enterprise.context.ApplicationScoped;

import com.redhat.ruben.examples.restaurant.model.Restaurant;
import org.eclipse.microprofile.config.inject.ConfigProperty;

@ApplicationScoped
public class RestaurantServiceImpl implements RestaurantService {

    @ConfigProperty(name = "restaurant_name")
    String name;

    @ConfigProperty(name = "restaurant_location")
    String location;

    private Restaurant restaurant;

    @PostConstruct
    public void init() {
        this.restaurant = new Restaurant()
                .setName(name)
                .setLocation(location);
    }

    @Override
    public Restaurant get() {
        return restaurant;
    }
}
