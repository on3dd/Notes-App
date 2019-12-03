<template>
  <v-form id="form" v-model="valid" lazy-validation onsubmit="return false;">
    <span class="display-3 text-center d-block mb-3">Регистрация</span>
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

    <v-text-field
      v-model="repeatPassword"
      :rules="repeatPasswordRules"
      label="Повторите пароль"
      required
      name="repeatPassword"
      type="repeatPassword"
    ></v-text-field>

    <v-btn 
      x-large color="success" 
      class="mx-4" 
      :disabled="!valid" 
      type="submit" 
      @click="() => {}"
      >
        Зарегистрироваться
      </v-btn>
  </v-form>
</template>

<script>
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

            repeatPassword: '',
            repeatPasswordRules: [
                v => !!v || "Поле должно быть заполнено"
            ]
        }),

        methods: {

        }
    };
</script>

<style>
  .v-input__slot {
    margin-bottom: 0 !important;
  }
</style>
