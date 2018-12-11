<template>

    <div id="general_content" class="container">
          <h1 v-if="this.recipes.length == 0">No hay resultados disponibles</h1>
          <div v-for="recipe, i in this.recipes" class="list_recipe">
              <h1>{{recipe.Title}}</h1>
              <p>{{recipe.Description}}</p>
              <button id="button3id" name="button3id" class="btn btn-primary" @click="viewRecipe(recipe.Id)">Ver</button>
              <button id="button1id" name="button1id" class="btn btn-success" @click="updateRecipe(recipe.Id)">Editar</button>
              <button class="btn btn-danger" v-confirm="{ok: dialog => deleteRecipe(recipe.Id,i),  message: '¿Esta usted seguro de eliminar esta receta?'}">Borrar</button>


          </div>
    </div>
</template>

<script>
import {config} from '../config.js'
export default {
  name: 'Recipes',

  data () {
    return {
      url: config.url,
      page: 0,
      msg: 'Welcome to Your Vue.js App',
      recipes:[],



    }
  },
  watch: {
    '$route' (to, from) {
      // react to route changes...
      this.main()
    }
},

  mounted(){

      this.main()
      this.scroll()

  },

  methods: {

      main(){

          if(this.$route.params.text_search){
              this.searchRecipes(this.$route.params.text_search)

          }
          else{

              this.getRecipes();

          }
      },
      getRecipes() {
          fetch(this.url+"?page="+this.page)
          .then(response => response.json())
          .then((data) => {
              if (this.page == 0) {
                    this.recipes = data;
              }else{
                  for (var i = 0; i < data.length; i++) {
                      this.recipes.push(data[i]);
                  }
              }



          })
      },
      searchRecipes(text_search) {
          fetch(this.url+"/recipes/search/"+text_search)
          .then(response => response.json())
          .then((data) => {
            this.recipes = data;

          })
      },

      deleteRecipe(id_recipe,id_list){


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
             console.log("borrando")

          })


          //delete element from list of recipes
          this.recipes.splice( id_list, 1 );
      },

      //go to view (details of a recipe)
      viewRecipe(idRecipe){
        this.$router.push({name:'View',params:{id:idRecipe}})
      },
      //go to update
      updateRecipe(id){
        this.$router.push({name:'Update',params:{id:id}})
    },

    //calculate the end of the document and load new data
    scroll() {
      window.onscroll = () => {
        let bottomOfWindow = document.documentElement.scrollTop + window.innerHeight === document.documentElement.offsetHeight;

        if (bottomOfWindow) {
          //this.page++
          var size = this.recipes.length -1
          //send the last id for the pagination
          this.page = this.recipes[size].Id
          this.getRecipes()

        }
      };
  },


  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#general_content{
  width:630px;
}
.list_recipe{
  margin-top: 50px;
}

</style>
