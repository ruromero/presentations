package com.redhat.ruben.examples.restaurant.services;

import java.io.File;
import java.io.IOException;

import javax.annotation.PostConstruct;
import javax.enterprise.context.ApplicationScoped;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import com.redhat.ruben.examples.restaurant.model.Menu;
import org.eclipse.microprofile.config.inject.ConfigProperty;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@ApplicationScoped
public class MenuServiceImpl implements MenuService {

    private static final Logger logger = LoggerFactory.getLogger(MenuServiceImpl.class);

    private static final String MENU_FILE = "menu.yaml";

    @ConfigProperty(name = "data_path", defaultValue = ".")
    String dataPath;

    private Menu menu;

    @PostConstruct
    public void loadMenu() {
        ObjectMapper mapper = new ObjectMapper(new YAMLFactory());
        String fileName = String.format("%s/%s", dataPath, MENU_FILE);
        try {
            this.menu = mapper.readValue(new File(fileName), Menu.class);
            logger.debug("Loaded menu {}", menu);
        } catch (IOException e) {
            logger.error("Unable to read menu from file {}", fileName, e);
        }
    }

    public Menu get() {
        return menu;
    }
}
