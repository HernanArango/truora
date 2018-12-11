<template>
  <div id="app">

  <div class="navbar navbar-expand-lg fixed-top navbar-dark bg-primary">
      <div class="container">
        <router-link class="navbar-brand" to="/">Recetas</router-link>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarResponsive">
          <ul class="navbar-nav">

            <li class="nav-item">
              <router-link class="nav-link" to="/create">Crear</router-link>
            </li>
            <!--<form class="form-inline my-2 my-md-0">-->
                    <div class="col">
                        <div class="row"><input class="form-control" @input="onChange" placeholder="Buscar" v-model="text_search" type="text"></div>
                        <div id="autocomplete" class="row">
                              <ul>
                              <li @click="setSearch(recipe.Title)" class="autocomplete-results" v-for="recipe in recipes">
                                {{ recipe.Title }}
                              </li>
                            </ul>
                        </div>
                    </div>
             <!--</form>-->
             <!--<button id="button2id" name="button2id" class="btn btn-danger">Borrar</button>-->
             <router-link class="btn btn-danger" :to="{ name: 'Search', params: {text_search: this.text_search } }">Buscar</router-link>
          </ul>
        </div>
      </div>
    </div>
    <router-view/>
  </div>
</template>
<style>
#autocomplete{
    background-color: #F3F3F3;
    position: fixed;
    width: 233px;
}
#autocomplete li{
    list-style: none;
    margin-left: 10px;
}
#autocomplete li:hover{
    background-color: #dde1e6;
    cursor:pointer;

}
#autocomplete ul{
    margin-left: -39px;
    width: 116%;
}
#app{
    margin-top: 90px;
}


</style>

<script>
import {config} from './config.js'
export default {
  name: 'App',
  data () {
    return {
      url: config.url,
      text_search:"",
      recipes : ""
    }
  },

  methods: {

      searchRecipe(text_search) {

          fetch(this.url+"/recipes/search/"+text_search)
          .then(response => response.json())
          .then((data) => {
              this.recipes = data
            //this.$router.push({name:'Recipes',params:{recipe:data}})
            //console.log(data)
          })
      },


      onChange(){

          //validate some invalid character
          var text = this.text_search
          var text_clean = text.replace(/[.*+?^${}()|[\]\\]/g, '');
          this.searchRecipe(text_clean);


      },

      setSearch(title){
          this.text_search = title
          //clean the results
          this.recipes = null

      },

      handleClickOutside(evt) {          
          this.recipes = null

      },




 },

 mounted() {
   document.addEventListener('click', this.handleClickOutside);
 },
 destroyed() {
   document.removeEventListener('click', this.handleClickOutside);
 }
}
</script>
