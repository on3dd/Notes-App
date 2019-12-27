<template>
  <v-col>
    <v-form id="form" v-model="valid" lazy-validation onsubmit="return false;">
      <span class="display-2 display-sm-3 text-center d-block mb-3">Авторизация</span>
      <v-text-field
        v-model="username"
        :rules="usernameRules"
        label="Логин"
        required
        name="username"
      ></v-text-field>

      <v-text-field
        v-model="password"
        :rules="passwordRules"
        label="Пароль"
        required
        name="password"
        type="password"
      ></v-text-field>

      <v-checkbox
        v-model="rememberMe"
        label="Оставаться в системе"
      ></v-checkbox>

      <div class="mb-3 sign-up">
        <nuxt-link nuxt to="/join">Еще не зарегестрированы?</nuxt-link>
      </div>

      <div>
        <v-btn
          x-large color="success"
          class="mx-4"
          :disabled="!valid"
          type="submit"
          @click="submit"
        >
          Войти
        </v-btn>
      </div>
    </v-form>
  </v-col>
</template>

<script>

  import axios from "axios";

  export default {
    name: "LoginForm",
    data: () => ({
      valid: false,
      username: '',
      usernameRules: [
        v => !!v || "Поле должно быть заполнено"
      ],

      password: '',
      passwordRules: [
        v => !!v || "Поле должно быть заполнено"
      ],

      rememberMe: false,
    }),

    methods: {
      submit: function () {
        let data = new FormData()

        data.append("name", this.username)
        data.append("password", this.password)

        axios.post("http://localhost:8080/api/v1/login", data)
          .then(response => {
            if (response.status == 200) {
              this.$nuxt.$router.replace({ path: `/notes`})
            }
          })
          .catch(err => {
            console.error(err)
          })
      }
    }
  };
</script>

<style>
  .v-input__slot {
    margin-bottom: 0 !important;
  }
  .sign-up {
    text-align: left;
  }
  .sign-up a {
    text-decoration: none;
  }
  .sign-up a:hover {
    text-decoration: underline;
  }

</style>
