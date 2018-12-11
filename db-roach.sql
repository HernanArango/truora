CREATE DATABASE recipes;


CREATE TABLE IF NOT EXISTS recipes.recipe (
  idrecipe SERIAL,
  title string(200) NULL,
  description String NULL,
  PRIMARY KEY (idrecipe));



CREATE TABLE IF NOT EXISTS recipes.ingredients (
  idingredients SERIAL,
  description string(500) NULL,
  recipe_idrecipe INT64 NOT NULL,
  PRIMARY KEY (idingredients),
  INDEX fk_ingredients_recipe_idx (recipe_idrecipe ASC),
  CONSTRAINT fk_ingredients_recipe
    FOREIGN KEY (recipe_idrecipe)
    REFERENCES recipes.recipe (idrecipe)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);