<template>
  <v-app id="inspire">
    <v-content style="min-height: 100vh">
      <v-container class="my-12">
        <v-row justify="center" no-gutters class="mx-md-6">
          <v-col :sm="10" :md="6" align="center">
            <LoginForm></LoginForm>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import LoginForm from "../components/LoginForm";

export default {
  components: {
    LoginForm
  },
  name: "login.vue"
};
</script>

<style scoped>
</style>
