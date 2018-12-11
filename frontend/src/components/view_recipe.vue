<template>
    <div id="general_content" class="container" style="margin-top:50px;">

    <h1 >{{recipe.Title}}</h1>

        <p>{{recipe.Description}}</p>
        <br><h3>Ingredientes</h3><br>
        <li v-for="ingredient, i in recipe.Ingredients">
            {{ingredient.Description}}

        </li>
        <router-link class="btn btn-success" :to="{ name: 'Update', params: {id: recipe.Id } }">Editar</router-link>
        <button class="btn btn-danger" v-confirm="{ok: dialog => deleteRecipe(recipe.Id),  message: '¿Esta usted seguro de eliminar esta receta?'}">Borrar</button>


    </div>
</template>

<script>
import {config} from '../config.js'
export default {
  name: 'view_recipe',

  data () {
    return {
      url: config.url,
      recipe:[],
      //temporal fix
      count:0



    }
  },
  watch: {
      recipe:function(newRecipe,oldRecipe){

            if(this.count<5){
                this.getRecipe();
                this.count++
            }

      }
  },
  mounted() {
      this.getRecipe();
  },

  methods: {

      getRecipe() {
              //this.recipe = this.$route.params.recipe;

              fetch(this.url+"/recipes/"+this.$route.params.id)
              .then(response => response.json())
              .then((data) => {
                //console.log("data1 "+data[0].Description)
                this.recipe = data[0];
                


              })

       },

       deleteRecipe(id_recipe){

           console.log("borrando"+this.url)
           //delete element Request
           fetch(this.url+"/recipes/delete/"+id_recipe, {
             body: JSON.stringify(id_recipe),
             method: "DELETE",
             headers: {
               "Content-Type": "application/json",
             },
           })
           .then(() => {
             //this.error = null;
             //console.log("Error deleting")
             this.$router.push({name:'Recipes'})

           })



       },


  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  #general_content{
    margin-top:50px;
    width:630px;
  }
</style>
