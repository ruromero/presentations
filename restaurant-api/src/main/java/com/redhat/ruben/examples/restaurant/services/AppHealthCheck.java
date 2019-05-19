package com.redhat.ruben.examples.restaurant.services;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;

import org.eclipse.microprofile.health.Health;
import org.eclipse.microprofile.health.HealthCheck;
import org.eclipse.microprofile.health.HealthCheckResponse;
import org.eclipse.microprofile.health.HealthCheckResponseBuilder;

@Health
@ApplicationScoped
public class AppHealthCheck implements HealthCheck {

    @Inject
    RestaurantService restaurantService;
    @Inject
    MenuService menuService;

    @Override
    public HealthCheckResponse call() {
        HealthCheckResponseBuilder response = HealthCheckResponse.named("Restaurant health check");
        if (menuService.get() == null || restaurantService.get() == null) {
            return response.down().build();
        }
        return response.up().build();
    }
}
