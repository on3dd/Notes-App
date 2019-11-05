<template>
  <v-app id="inspire">
    <Navbar></Navbar>
    <v-content style="min-height: 100vh">

    </v-content>
    <Footer></Footer>
  </v-app>
</template>

<script>
    import Navbar from "~/components/Navbar";
    import Footer from "~/components/Footer";

    export default {
        name: "index.vue",
        components: {
            Navbar,
            Footer,
        }
    }
</script>

<style scoped>

</style>
