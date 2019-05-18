package com.redhat.ruben.examples.restaurant.services;

import java.io.IOException;
import java.io.InputStream;

import javax.annotation.PostConstruct;
import javax.enterprise.context.ApplicationScoped;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import com.redhat.ruben.examples.restaurant.model.Menu;
import org.eclipse.microprofile.config.inject.ConfigProperty;

@ApplicationScoped
public class MenuServiceImpl implements MenuService {

    private static final String DEFAULT_RESTAURANT = "Andalusian";

    @ConfigProperty(name = "restaurant_type", defaultValue = DEFAULT_RESTAURANT)
    String type;

    private Menu menu;

    @PostConstruct
    public void loadMenu() {
        InputStream is = this.getClass().getClassLoader().getResourceAsStream(String.format("menu-%s.yaml", type.toLowerCase()));
        ObjectMapper mapper = new ObjectMapper(new YAMLFactory());
        try {
            this.menu = mapper.readValue(is, Menu.class);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public Menu get() {
        return menu;
    }
}
