<template>
  <v-app id="inspire">
    <Navbar></Navbar>
    <v-content style="min-height: 100vh">
      <v-container
        class="my-12"
      >
        <v-row
          justify="center"
          no-gutters
          class="mx-md-6"
        >
          <v-col
            :sm="10"
            :md="6"
          >
            <span class="d-none d-md-block display-3 mb-3 text-truncate">{{ note.title }}</span>
            <span class="d-block d-md-none display-2 mb-3 text-truncate">{{ note.title }}</span>
            <span class="headline font-weight-regular d-block mb-2">{{note.descirption}}</span>
            <span class="d-none d-md-block title mb-2">
                <a class="author font-weight-regular mr-2" href="">{{ author.name }},</a>
                <span class="font-weight-light">{{ note.posted_at }}</span>
              </span>
            <span class="d-block d-md-none title mb-2">
                <span class="d-block font-weight-regular mb-2">Автор:
                  <a href="" class="author">{{ author.name }}</a>
                </span>
                <span class="d-block font-weight-light">{{ note.posted_at }}</span>
              </span>
            <span class="title mb-2 d-block font-weight-regular">Предмет: <a class="subject"
                                                                             href="">{{subject.name}}</a></span>
            <span class="title mb-2 d-block font-weight-regular">Преподаватель: <a class="teacher" href="">{{teacher.name}}</a></span>
            <div class="my-6">
              <v-btn x-large color="primary"><a style="color:white;text-decoration:none;" :href="note.link">Отрыть работу</a></v-btn>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
    <Footer></Footer>
  </v-app>
</template>

<script>
    import Navbar from "~/components/Navbar";
    import Footer from "~/components/Footer";
    import axios from 'axios'

    export default {
        data: () => ({
            note: '',
            author: '',
            category: '',
            subject: '',
            teacher: '',

        }),
        components: {
            Navbar,
            Footer,
        },
        mounted() {
            console.log(this.$route.params.id)

            // Time formatting options
            const options = {
                year: 'numeric',
                month: 'long',
                day: 'numeric',
                timezone: 'UTC',
                hour: 'numeric',
                minute: 'numeric',
                second: 'numeric'
            };

            axios.get('http://localhost:8080/api/v1/getNote', {params: {id: this.$route.params.id}})
                .then(response => {
                    console.log(response.data)
                    this.note = response.data

                    let timestamp = Date.parse(this.note.posted_at)
                    this.note.posted_at = new Date(timestamp).toLocaleString("ru", options)

                    axios.get('http://localhost:8080/api/v1/getUser', {params: {id: this.note.author_id}})
                        .then(response => {
                            console.log(response.data)
                            this.author = response.data
                        })
                        .catch(err => {
                            console.log(err)
                        })

                    axios.get('http://localhost:8080/api/v1/getCategory', {params: {id: this.note.category_id}})
                        .then(response => {
                            console.log(response.data)
                            this.category = response.data

                            axios.get('http://localhost:8080/api/v1/getSubject', {params: {id: this.category.subject}})
                                .then(response => {
                                    console.log(response.data)
                                    this.subject = response.data
                                })
                                .catch(err => {
                                    console.log(err)
                                })
                        })
                        .catch(err => {
                            console.log(err)
                        })

                    axios.get('http://localhost:8080/api/v1/getTeacher', {params: {id: this.note.teacher_id}})
                        .then(response => {
                            console.log(response.data)
                            this.teacher = response.data
                        })
                        .catch(err => {
                            console.log(err)
                        })
                })
                .catch(err => {
                    console.log(err)
                })
        },
        validate({params}) {
            return /^\d+$/.test(params.id)
        }
    }
</script>

<style>
  .author, .subject, .teacher {
    text-decoration: none;
  }

  .author:hover, .subject:hover, .teacher:hover {
    text-decoration: underline;
  }
</style>
