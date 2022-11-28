<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <!--begin::Section-->
      <div class="py-5">
        <!--begin::Heading-->
        <h1 class="anchor fw-bolder mb-5" id="setup_api">
          <router-link to="#setup_api"></router-link>
          Setup API
        </h1>
        <!--end::Heading-->

        <!--begin::Block-->
        <div class="py-5">
          By default all http requests made by axios are sent to
          <code
            >https://preview.keenthemes.com/metronic8-staging/laravel/api</code
          >
          url, you can update the default url in file <code>.env</code> variable
          <code>VUE_APP_API_URL</code>. This variable is used to set the baseURL
          for axios instance in <code>src/core/services/ApiService.ts</code>.
        </div>
        <!--end::Block-->

        <div class="pb-5">
          <CodeHighlighter lang="javascript">{{
            `ApiService.vueInstance.axios.defaults.baseURL = process.env.VUE_APP_API_URL;`
          }}</CodeHighlighter>
        </div>

        <!--begin::Block-->
        <div class="py-5">
          Our default api url used in the project is the same as a api interface
          provided in our laravel theme. You can find all available api
          endpoints in file
          <code>laravel/routes/api.php</code>.
        </div>
        <!--end::Block-->

        <!--begin::Block-->
        <div class="py-5">
          <div class="pb-3">
            To get started with laravel api locally follow steps below:
          </div>
          <ol>
            <li class="py-3">
              Run laravel theme using
              <a
                href="https://preview.keenthemes.com/metronic8/laravel/documentation/getting-started/build?demo=demo1"
                >quick start doc</a
              >.
            </li>
            <li class="py-3">Keep the server running.</li>
            <li class="py-3">
              In another command line run vue app with
              <router-link to="/build">vue quick start doc</router-link>.
            </li>
            <li class="py-3">
              By default laravel app is served on url
              <a href="http://127.0.0.1:8000">http://127.0.0.1:8000</a> so to
              use the local api you need to update
              <code>VUE_APP_API_URL</code> variable in file <code>.env</code> to
              <code>http://127.0.0.1:8000/api</code>.
            </li>
          </ol>
        </div>
        <!--end::Block-->
      </div>
      <!--end::Section-->
    </div>
    <!--end::Card Body-->
  </div>
  <!--begin::Card-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";

export default defineComponent({
  name: "api-setup",
  components: {
    CodeHighlighter,
  },
});
</script>
