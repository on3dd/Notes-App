<template>
  <v-form id="form" v-model="valid" lazy-validation onsubmit="return false;">
    <span class="display-3 text-center d-block mb-3">Авторизация</span>
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

    <v-btn color="success" class="mx-4" :disabled="!valid" type="submit" @click="() => {}">Войти</v-btn>
  </v-form>
</template>

<script>

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
            ]
        }),

        methods: {}
    };
</script>

<style scoped>
</style>
