<template>
  <v-app id="inspire">
    <v-content style="min-height: 100vh;">
      <v-container
        class="my-12"
      >
        <v-row
          justify="center"
          no-gutters
          class="mx-md-6 px-2"
        >
          <v-col
            :sm="10"
            :md="8"
            :lg="6"
            class="px-2"
          >
            <div class="mb-3">
              <span
                v-bind:class="{ 'd-none': isEditing }"
                class="d-block display-2 display-sm-3"
              >{{ user.name }}</span>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
  export default {
    components: {

    },
    props: {
      source: String,
    },

    async asyncData({$axios, params}) {
      const user = await $axios.$get(`http://localhost:8080/api/v1/users/${params.id}`)
      return { user: user }
    },

    data: () => ({
      drawer: null,
      isEditing: false
    }),
  }
</script>

<style scoped>

</style>
