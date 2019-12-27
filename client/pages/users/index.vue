<template>
  <v-app id="inspire">
    <v-content style="min-height: 100vh;">
      <v-container
        class="fill-height"
        fluid
      >
        <v-row
          align="center"
          justify="center"
        >
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
  export default {
    components: {

    },
    props: {
      source: String,
    },

    data: () => ({
      drawer: null,
    }),
  }
</script>

<style scoped>

</style>
