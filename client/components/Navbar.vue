<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      app
    >
      <v-list dense>
        <v-list-item v-for="(item, index) in items" :key="index" link nuxt :to="item.path">
          <v-list-item-action>
            <v-icon>{{item.icon}}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title style="font-size: 15px">{{item.text}}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-divider></v-divider>

        <v-list-item link nuxt to="/upload">
          <v-list-item-action>
            <v-icon>mdi-upload</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title style="font-size: 15px">Загрузка</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"/>
      <v-toolbar-title>Конспекты</v-toolbar-title>
      <v-spacer></v-spacer>

      <v-btn link nuxt to="/login"
             outlined tile
             color="primary"
      >
        <v-icon class="mr-1">mdi-account-circle</v-icon>
        <span>Войти</span>
      </v-btn>
    </v-app-bar>
  </v-app>
</template>

<script>
  import 'material-design-icons-iconfont/dist/material-design-icons.css'

  export default {
    name: "Navbar",
    props: {
      source: String,
    },
    data: () => ({
      drawer: true,
      items: [
        { path: "/", icon: "mdi-home", text: "Главная" },
        { path: "/notes", icon: "mdi-folder-open", text: "Работы" },
        { path: "/users", icon: "mdi-account-group", text: "Пользователи" }
      ]
    }),
  };

</script>

<style>
  .v-application--wrap {
    min-height: 0;
  }

  .v-list-item--active .v-icon {
    color: #1976d2 !important;
  }

</style>
