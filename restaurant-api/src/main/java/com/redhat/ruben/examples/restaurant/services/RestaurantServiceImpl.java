package com.redhat.ruben.examples.restaurant.services;

import java.io.File;
import java.io.IOException;

import javax.annotation.PostConstruct;
import javax.enterprise.context.ApplicationScoped;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import com.redhat.ruben.examples.restaurant.model.Restaurant;
import org.eclipse.microprofile.config.inject.ConfigProperty;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@ApplicationScoped
public class RestaurantServiceImpl implements RestaurantService {

    private static final Logger logger = LoggerFactory.getLogger(RestaurantServiceImpl.class);

    private static final String INFO_FILE = "info.yaml";

    @ConfigProperty(name = "data_path", defaultValue = ".")
    String dataPath;

    private Restaurant restaurant;

    @PostConstruct
    public void loadMenu() {
        ObjectMapper mapper = new ObjectMapper(new YAMLFactory());
        String fileName = String.format("%s/%s", dataPath, INFO_FILE);
        try {
            this.restaurant = mapper.readValue(new File(fileName), Restaurant.class);
            logger.debug("Loaded restaurant info {}", restaurant);
        } catch (IOException e) {
            logger.error("Unable to read restaurant info from file {}", fileName, e);
        }
    }

    @Override
    public Restaurant get() {
        return restaurant;
    }
}
