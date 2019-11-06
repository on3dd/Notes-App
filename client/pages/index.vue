<template>
  <v-app id="inspire">
    <Navbar></Navbar>

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
    <Footer></Footer>
  </v-app>
</template>

<script>
    import Navbar from "~/components/Navbar";
    import Footer from "~/components/Footer";

    export default {
        components: {
            Navbar,
            Footer,
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
