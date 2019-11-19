<template>
  <v-form
    id="form"
    v-model="valid"
    lazy-validation
    onsubmit="return false;"
  >
    <span class="display-3 text-center d-block mb-3">Загрузка работы</span>
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

    <v-select
      v-model="category"
      :rules="categoryRules"
      :items="categories"
      item-text="name"
      label="Категория"
      required
      name="category"
      @change="getSubjects"
    ></v-select>

    <v-select
      v-model="subject"
      :rules="subjectRules"
      :items="subjects"
      item-text="name"
      label="Предмет"
      required
      name="subject"
      @change="getTeachers"
    ></v-select>

    <v-select
      v-model="teacher"
      :rules="teacherRules"
      :items="teachers"
      item-text="name"
      label="Преподаватель"
      required
      name="teacher"
    ></v-select>

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
      @click="this.submit"
    >
      Отправить
    </v-btn>
  </v-form>
</template>

<script>
    import axios from 'axios'

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
            categories: ["Выберите категорию"],

            subject: '',
            subjectRules: [
                v => !!v || 'Поле должно быть заполнено',
            ],
            subjects: ["Выберите предмет"],

            teacher: '',
            teacherRules: [
                v => !!v || 'Поле должно быть заполнено',
            ],
            teachers: ["Выберите преподавателя"],

            checkbox: false,
        }),

        methods: {
            getCategories: function() {
                axios.get('http://localhost:8080/api/v1/categories')
                    .then(response => {
                        this.categories = response.data

                        this.subjects = []
                        this.subject = ''

                        this.teachers = []
                        this.teacher = ''
                    })
                    .catch(err => {
                        console.error(err)
                    })
            },
            getSubjects: function() {
                let categoryIdx = this.categories.indexOf(this.categories.find(el => el.name == this.category))
                if (categoryIdx == -1) return false
                // console.log("subject id = ", this.categories[categoryIdx].subject)
                axios.get("http://localhost:8080/api/v1/subjects", {
                    params: {
                        id: this.categories[categoryIdx].subject
                    }
                })
                    .then(response => {
                        this.subjects = response.data
                        this.subject = ''

                        this.teachers = []
                        this.teacher = ''
                    })
                    .catch(err => {
                        console.error(err)
                    })
            },
            getTeachers: function() {
                let subjectIdx = this.subjects.indexOf(this.subjects.find(el => el.name == this.subject))
                if (subjectIdx == -1) return false
                // console.log("teacher id = ", this.subjects[subjectIdx].id)
                axios.get("http://localhost:8080/api/v1/teachers", {
                    params: {
                        id: this.subjects[subjectIdx].id
                    }
                })
                    .then(response => {
                        this.teachers = response.data
                        this.teacher = ''
                    })
                    .catch(err => {
                        console.error(err)
                    })
            },
            submit: function () {
                let categoryIdx = this.categories.indexOf(this.categories.find(el => el.name == this.category))
                if (categoryIdx == -1) return console.error("Error: Categories are empty!");

                let subjectIdx = this.subjects.indexOf(this.subjects.find(el => el.name == this.subject))
                if (subjectIdx == -1) return console.error("Error: Subjects are empty!");;

                let data = new FormData()

                data.append("author", 1)
                data.append("category_id", this.categories[categoryIdx].id)
                data.append("teacher_id", this.subjects[subjectIdx].id)

                // data.append("posted_at", Date.now())
                data.append("title", this.title)
                data.append("description", this.description)
                data.append("file", this.file)

                axios.post("http://localhost:8080/api/v1/notes", data)
                    .then(response => {
                        if (response.status == 200) {
                            this.$nuxt.$router.replace({ path: `/notes/${response.data.id}`})
                        }
                    })
                    .catch(err => {
                        console.error(err)
                    })
            }
        },

        created() {
            this.getCategories()
        }

    }

</script>

<style></style>
