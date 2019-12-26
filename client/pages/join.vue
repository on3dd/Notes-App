<template>
  <v-app id="inspire">
    <v-content style="min-height: 100vh">
      <v-container class="my-12">
        <v-row justify="center" no-gutters class="mx-md-6">
          <v-col :sm="10" :md="6" align="center">
            <JoinForm></JoinForm>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import Navbar from "../components/Navbar"
import Footer from "../components/Footer";
import JoinForm from "../components/JoinForm";

export default {
  components: {
    JoinForm
  },
  name: "join.vue"
};
</script>

<style scoped>
</style>
