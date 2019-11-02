<template>
  <v-form
    method="post"
    action="localhost:8080/api/v1/addNote"
    id="form"
    v-model="valid"
    lazy-validation

  >
    <v-text-field
      v-model="title"
      :rules="titleRules"
      label="Название"
      required
      name="title"
    ></v-text-field>

    <v-textarea
      v-model="description"
      :rules="descriptionRules"
      label="Описание"
      hint="Введите описание работы"
      name="descirption"
    ></v-textarea>

    <v-file-input
      v-model="file"
      :rules="fileRules"
      label="Файлы"
      accept="application/pdf"
      required
      name="file"
    ></v-file-input>

    <v-text-field
      v-model="category"
      :rules="categoryRules"
      label="Категория"
      required
      name="category"
    ></v-text-field>

    <v-text-field
      v-model="subject"
      :rules="subjectRules"
      label="Предмет"
      required
      name="subject"
    ></v-text-field>

    <v-text-field
      v-model="teacher"
      :rules="teacherRules"
      label="Преподаватель"
      required
      name="teacher"
    ></v-text-field>

    <v-checkbox
      v-model="checkbox"
      :rules="[v => !!v || 'Вы должны принять условия пользования!']"
      label="Я согласен с условиями пользования сервисом"
      required
    ></v-checkbox>

    <v-btn
      color="success"
      class="mx-4"
      :disabled="!valid"
      type="submit"
    >
      Отправить
    </v-btn>
  </v-form>
</template>

<script>
    export default {
        name: "UploadForm",
        data: () => ({
            valid: false,
            title: '',
            titleRules: [
                v => !!v || 'Поле должно быть заполнено',
                v => (v && v.length >= 10) || 'Название должно содержать как минимум 10 символов',
            ],

            file: null,
            fileRules: [
                v => !!v || "Поле должно быть заполнено"
            ],

            description: '',
            descriptionRules: [],

            category: '',
            categoryRules: [
                v => !!v || 'Поле должно быть заполнено',
            ],

            subject: '',
            subjectRules: [
                v => !!v || 'Поле должно быть заполнено',
            ],

            teacher: '',
            teacherRules: [
                v => !!v || 'Поле должно быть заполнено',
            ],

            checkbox: false,
        }),

        methods: {
            // submit() {
            //     form.submit()
            // }
        },

    }
</script>

<style scoped>

</style>
