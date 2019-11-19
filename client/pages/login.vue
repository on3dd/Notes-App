<template>
  <v-app id="inspire">
    <Navbar></Navbar>
    <v-content >
      <v-container class="my-12">
        <v-row justify="center" no-gutters class="mx-md-6">
          <v-col :sm="10" :md="6" align="center">
            <LoginForm></LoginForm>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
    <Footer></Footer>
  </v-app>
</template>

<script>
import Navbar from "../components/Navbar"
import Footer from "../components/Footer";
import LoginForm from "../components/LoginForm";

export default {
  components: {
    Navbar,
    Footer,
    LoginForm
  },
  name: "login.vue"
};
</script>

<style scoped>
</style>