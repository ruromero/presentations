package com.redhat.ruben.examples.restaurant;

import javax.json.Json;
import javax.json.JsonObject;

import io.quarkus.test.junit.QuarkusTest;
import io.restassured.http.ContentType;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;
import static org.hamcrest.CoreMatchers.equalTo;

@QuarkusTest
public class RestaurantResourceTest {

    @Test
    public void testRestaurantEndpoint() {
        given()
                .when().get("/api")
                .then()
                .assertThat()
                .statusCode(200)
                .contentType(ContentType.JSON)
                .body("name", equalTo("Casa Ruben"))
                .body("location", equalTo("C/ Altramuz 22, Motril"))
                .body("foodType", equalTo("Andalusian"));
    }


    @Test
    public void testMenuEndpoint() {
        given()
                .when().get("/api/menu")
                .then()
                .assertThat()
                .statusCode(200)
                .contentType(ContentType.JSON)
                .body("starters[0].name", equalTo("Jamón ibérico"))
                .body("starters[0].price", equalTo(15f));
    }
}