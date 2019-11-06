<template>
<v-app id="inspire">
  <v-navigation-drawer
      v-model="drawer"
      app
    >
      <v-list dense>
        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-home</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Главная</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item link>
          <v-list-item-action>
                      <v-icon x-large dark>folder_open</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Работы</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item link>
          <v-list-item-action>
            <v-icon>mdi-contact-mail</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>О нас</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      app
      color="indigo"
      dark
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title>Американские конспекты</v-toolbar-title>
    </v-app-bar>
</v-app>
</template>

<script>
export default {
  name: "Navbar",
  props: {
      source: String,
    },
    data: () => ({
      drawer: null,
    }),
};
</script>

<style>
  .v-application--wrap {
    min-height: 0;
  }
</style>
