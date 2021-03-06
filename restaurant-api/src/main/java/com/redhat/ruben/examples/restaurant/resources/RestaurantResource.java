package com.redhat.ruben.examples.restaurant.resources;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.CompletionStage;

import javax.inject.Inject;
import javax.ws.rs.Consumes;
import javax.ws.rs.GET;
import javax.ws.rs.NotFoundException;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

import com.redhat.ruben.examples.restaurant.model.Menu;
import com.redhat.ruben.examples.restaurant.model.Restaurant;
import com.redhat.ruben.examples.restaurant.services.MenuService;
import com.redhat.ruben.examples.restaurant.services.RestaurantService;
import org.omg.CosNaming.NamingContextPackage.NotFound;

@Path("/api")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public class RestaurantResource {

    @Inject
    MenuService menuService;

    @Inject
    RestaurantService restaurantService;

    @GET
    @Path("/menu")
    public CompletionStage<Menu> getMenu() {
        return CompletableFuture.supplyAsync(() -> {
            Menu menu = menuService.get();
            if (menu == null) {
                throw new NotFoundException("The menu has not been loaded");
            }
            return menu;
        });
    }

    @GET
    public CompletionStage<Restaurant> getRestaurant() {
        return CompletableFuture.supplyAsync(() -> {
            return restaurantService.get();
        });
    }
}