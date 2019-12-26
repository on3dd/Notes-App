<template>
  <v-app id="inspire">
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
            align="center"
          >
            <UploadForm></UploadForm>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
    import Navbar from "../components/Navbar";
    import UploadForm from "../components/UploadForm";
    import Footer from "../components/Footer";

    export default {
        components: {
            UploadForm,
        },
        name: "upload.vue"
    }
</script>

<style scoped>

</style>
