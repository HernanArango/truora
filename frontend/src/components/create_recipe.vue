<template>
 <div class="">


  <div id="form_create" class="container">
         <h1>Crear Receta</h1>
         <p v-if="error_title.length">
            <b style="color:red;">{{error_title}}</b>
          </p>
        <input  class="form-control" type="text" v-model="title" name="title" placeholder="Titulo">
        <b style="color:red;">{{error_description}}</b>
        <textarea v-model="description" class="form-control" rows="4" cols="50" name="description" placeholder="descripción">
        </textarea>
        <h2>Ingredientes</h2>
        <ul>
            <li v-for="ingredient, i in this.Ingredients">

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
            </li>
        </ul>
        <b style="color:red;">{{error_ingredients}}</b>
        <div class="input-group mb-3">

                <input  class="form-control" type="text" v-model="ingredient_description" name="title" placeholder="ej dos manzanas" v-on:keyup.13="setIngredient(ingredient_description)">

          <div class="input-group-append">
            <input  @click="setIngredient(ingredient_description)"  id="anadir_ingredient" class="btn btn-outline-secondary" type="button" value="añadir ingrediente">

          </div>
        </div>

        <input  @click="validate()" class="btn btn-primary" type="button" value="Guardar">


  </div>
 </div>
</template>

<script>
import {config} from '../config.js'
 export default{
      name:'create_recipe',
      data (){
           return{
            url: config.url,
            title:null,
            description:null,
            recipe:{},
            Ingredients:[],
            ingredient_description:"",
            error_title: "",
            error_description: "",
            error_ingredients : "",
            size:0,
            is_editable: false,

           }
      },

      methods: {

          validate(){
              this.error_title = '';
              this.error_description = '';
              this.error_ingredients = '';
              if (!this.title) {
                  this.error_title = 'Por favor llene este campo';
              }else if (!this.description) {
                  this.error_description = 'Por favor llene este campo';
              }else if (this.Ingredients.length == 0) {
                  this.error_ingredients = 'Debe ingresar por lo menos un ingrediente';
              }
              else{
                  this.setRecipe()
              }
          },

          setRecipe() {
              this.recipe.Title = this.title
              this.recipe.Description = this.description

              this.recipe["Ingredients"] = this.Ingredients
              fetch(this.url+"/recipes", {
                body: JSON.stringify(this.recipe),
                method: "POST",
                headers: {
                  "Content-Type": "application/json",
                },
              })
              .then(response => response.json())
              .then((data) => {

                 console.log(data);
                 this.$router.push({name:'View',params:{id:data.Id}})


             })




            },

            setIngredient(description_ingredient){
                console.log("descrip"+description_ingredient)
                this.size = Object.keys(this.Ingredients).length
                //this.Ingredients[this.size] ={"Id":this.size, "Description":description_ingredient}
                //using $set for also update the render when change the array
                this.$set(this.Ingredients, this.size, {"Id":""+this.size, "Description":description_ingredient, "Deleted":0})
                this.description_ingredient = ""
                console.log(this.Ingredients)
            },

            deleteIngredient(id){
                //delete this.Ingredients[id];
                this.Ingredients.splice(id,1)
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
    margin-bottom: 20px!important;
}

#anadir_ingredient{
    height: 38px;
}
</style>
