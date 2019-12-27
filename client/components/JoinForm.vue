<template>
  <v-col>
    <v-form id="form" v-model="valid" lazy-validation onsubmit="return false;">
      <span class="display-2 display-sm-3 text-center d-block mb-3">Регистрация</span>
      <v-text-field
        v-model="username"
        :rules="usernameRules"
        label="Логин"
        required
        name="username"
      />

      <v-text-field
        v-model="email"
        :rules="emailRules"
        label="Email"
        required
        name="email"
        type="email"
      />

      <v-textarea
        v-model="about"
        :rules="aboutRules"
        label="О себе"
        name="about"
        rows="2"
      />

      <v-text-field
        v-model="password"
        :rules="passwordRules"
        label="Пароль"
        required
        name="password"
        type="password"
      />

      <v-text-field
        v-model="repeatPassword"
        :rules="repeatPasswordRules"
        label="Повторите пароль"
        required
        name="repeatPassword"
        type="password"
      />

      <div class="mb-3 sign-in">
        <nuxt-link nuxt to="/login">Уже зарегестрированы?</nuxt-link>
      </div>

      <v-btn
        x-large color="success"
        class="mx-4"
        :disabled="!valid"
        type="submit"
        @click="submit"
      >
        Зарегистрироваться
      </v-btn>
    </v-form>
  </v-col>
</template>

<script>
  import axios from "axios";

  export default {
    name: "JoinForm",
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

      email: '',
      emailRules: [
        v => !!v || "Поле должно быть заполнено"
      ],

      about: '',
      aboutRules: [],

      repeatPassword: '',
      repeatPasswordRules: [
        v => !!v || "Поле должно быть заполнено"
      ]
    }),

    methods: {
      submit: function () {
        if (this.password != this.repeatPassword) {
          return console.log("Password does not match!")
        }

        let data = new FormData()

        data.append("name", this.username)
        data.append("email", this.email)
        data.append("about", this.about)
        data.append("password", this.password)

        axios.post("http://localhost:8080/api/v1/join", data)
          .then(response => {
            if (response.status == 200) {
              this.$nuxt.$router.replace({ path: `/users/${response.data.id}`})
            }
          })
          .catch(err => {
            console.error(err)
          })
      }

    },
  }
  ;
</script>

<style>
  .v-input__slot {
    margin-bottom: 0 !important;
  }

  .sign-in {
    text-align: left;
  }
  .sign-in a {
    text-decoration: none;
  }
  .sign-in a:hover {
    text-decoration: underline;
  }
</style>
