<template>
 <div class="">


  <div id="form_create" class="container">
         <h1>Editar Receta</h1>
        <input  class="form-control" type="text" v-model="recipe.Title" name="title" placeholder="Titulo">
        <textarea v-model="recipe.Description" class="form-control" rows="4" cols="50" name="description" placeholder="descripción">
        </textarea>
        <h2>Ingredientes</h2>

        <div class="list_element" v-for="ingredient, i in this.Ingredients">
            <div v-if="ingredient.Deleted == 0">
                <div class='row'>
                    <div class="col-md-10">

                        <input class="form-control" v-model="ingredient.Description" v-on:keyup.13="editable()" v-if="is_editable"/>
                        <span v-else @click="editable()">{{ingredient.Description}}</span>


                    </div>
                    <div class="col-md-2">
                    <!--<button id="button1id" name="button1id" class="btn btn-info" @click="editable()">Editar</button>-->
                    <button id="button2id" name="button2id" class="btn btn-danger" @click="deleteIngredient(i)">Borrar</button>
                    </div>
                </div>


            </div>
        </div>


        <div class="input-group mb-3">
                <input  class="form-control" type="text" v-model="ingredient_description" name="title" placeholder="ej dos manzanas" v-on:keyup.13="setIngredient(ingredient_description)">

          <div class="input-group-append">
            <input  @click="setIngredient(ingredient_description)" id="anadir_ingredient" class="btn btn-outline-secondary" type="button" value="añadir ingrediente">

          </div>
        </div>


        <input  @click="updateRecipe(recipe)" class="btn btn-primary" type="button" value="Actualizar">


  </div>
 </div>
</template>

<script>
import {config} from '../config.js'
 export default{
      name:'update_recipe',
      data (){
           return{
            url: config.url,
            title:'Update Recipe',
            //load the data of the recipe in the form
            Ingredients: [],
            ingredient_description:"",
            recipe : {},
            is_editable: false,
            size:0

           }
      },
      mounted(){
            this.getRecipe()
      },

      methods: {

          getRecipe() {


                  fetch(this.url+"/recipes/"+this.$route.params.id)
                  .then(response => response.json())
                  .then((data) => {

                    this.recipe = data[0];
                    this.Ingredients = data[0].Ingredients


                  })
           },

          updateRecipe(recipe) {
                //console.log(recipe)
              recipe["Ingredients"] = this.Ingredients
              fetch(this.url+"/recipes/update/"+recipe.Id, {
                body: JSON.stringify(recipe),
                method: "PUT",
                headers: {
                  "Content-Type": "application/json",

                },
              })
              .then(() => {
                //this.error = null;

              })
              //console.log("cambia")
              this.$router.push({name:'View',params:{id:recipe.Id}})


            },
            setIngredient(description_ingredient){
                console.log("descrip"+description_ingredient)
                this.size = Object.keys(this.Ingredients).length
                //using $set for also update the render when change the array
                this.$set(this.Ingredients, this.size, {"Id":""+this.size, "Description":description_ingredient, "Deleted":0})
                console.log(this.Ingredients)
            },

            deleteIngredient(id){
                //this.Ingredients.splice(id,1)
                this.Ingredients[id].Deleted = 1
                console.log(this.Ingredients[id].Deleted)
            },

            updateIngredient(id,description){
                this.Ingredients[id].Description=description
                console.log(this.Ingredients)
            },

            editable(){
                if(this.is_editable == true){
                    this.is_editable = false
                }
                else{
                    this.is_editable = true
                }
            }

      }
 }
</script>

<style scoped>
#form_create{
    width:500px;
    margin-top: 80px;
}

.form-control{
    margin-bottom: 20px;
}

#anadir_ingredient{
    height: 38px;
}
.list_element{
    padding-bottom: 4px;
}
</style>
