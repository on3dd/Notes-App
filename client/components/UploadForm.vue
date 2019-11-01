<template>
  <v-form
    ref="form"
    v-model="valid"
    lazy-validation
  >
    <v-text-field
      v-model="title"
      :rules="titleRules"
      label="Название"
      required
    ></v-text-field>

    <v-textarea
      v-model="description"
      :rules="descriptionRules"
      label="Описание"
      hint="Введите описание работы"
    ></v-textarea>

    <v-text-field
      v-model="category"
      :rules="categoryRules"
      label="Категория"
      required
    ></v-text-field>

    <v-text-field
      v-model="subject"
      :rules="subjectRules"
      label="Предмет"
      required
    ></v-text-field>

    <v-text-field
      v-model="teacher"
      :rules="teacherRules"
      label="Преподаватель"
      required
    ></v-text-field>

    <v-checkbox
      v-model="checkbox"
      :rules="[v => !!v || 'Вы должны принять условия пользования!']"
      label="Я согласен с условиями пользования сервисом"
      required
    ></v-checkbox>

    <v-btn
      color="success"
      class="mr-4"
    >
      Отправить
    </v-btn>
  </v-form>
</template>

<script>
    export default {
        name: "UploadForm",
        data: () => ({
            valid: true,
            title: '',
            titleRules: [
                v => !!v || 'Поле быть заполнено',
                v => (v && v.length >= 10) || 'Название должно содержать как минимум 10 символов',
            ],

            description: '',
            descriptionRules: [],

            category: '',
            categoryRules: [
                v => !!v || 'Поле быть заполнено',
            ],

            subject: '',
            subjectRules: [
                v => !!v || 'Поле быть заполнено',
            ],

            teacher: '',
            teacherRules: [
                v => !!v || 'Поле быть заполнено',
            ],

            checkbox: false,
        }),

        methods: {
            validate () {
               if (this.$refs.form.validate()) {
                    this.snackbar = true
                }
            },
            reset () {
                this.$refs.form.reset()
            },
            resetValidation () {
                this.$refs.form.resetValidation()
            },
        },

    }
</script>

<style scoped>

</style>
