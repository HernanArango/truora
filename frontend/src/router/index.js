import Vue from 'vue'
import Router from 'vue-router'
import Recipes from '@/components/Recipes'
import Create_recipe from '@/components/create_recipe'
import View_recipe from '@/components/view_recipe'
import Update_recipe from '@/components/update_recipe'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Recipes',
      component: Recipes

  },
  {
    path: '/search/:text_search',
    name: 'Search',
    component: Recipes

  },

  {
      path: '/create',
      name: 'Create',
      component: Create_recipe

    },
    {
        path: '/view/:id',
        name: 'View',
        component: View_recipe

    },
    {
        path: '/update/:id',
        name: 'Update',
        component: Update_recipe

    },
  ]
})
