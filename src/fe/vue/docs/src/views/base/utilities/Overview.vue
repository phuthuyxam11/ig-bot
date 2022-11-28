<template>
  <!--begin::Section-->
  <div class="pb-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="overview">
      <router-link to="/utilities#overview"></router-link>
      Overview
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      <b>{{ appName }}</b> fully utilizes the
      <a href="https://getbootstrap.com/docs/5.0/utilities/api/" class="fw-bold"
        >Bootstrap Utility API</a
      >&nbsp; in <code>src/assets/sass/core/components/_utilities.scss</code> to
      generate additional utility classes to cover its own design system. The
      customized utility SASS variables are defined in
      <code>src/assets/sass/core/components/_variables.scss</code> and in
      <code>src/assets/sass/core/components/_variables.custom.scss</code>
      you can override any variable for your needs.
    </div>
    <!--end::Block-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "utilities-overview",
  components: {},
  setup() {
    return {
      appName,
    };
  },
});
</script>
