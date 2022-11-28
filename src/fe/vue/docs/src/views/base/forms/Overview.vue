<template>
  <!--begin::Section-->
  <div class="pb-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="overview">
      <router-link to="#overview"></router-link>
      Overview
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      <b>{{ appName }}</b> customizes the
      <a
        href="https://getbootstrap.com/docs/5.0/forms/overview/"
        class="fw-bold"
        >Bootstrap Forms</a
      >&nbsp; through the SASS variables in
      <code>src/assets/sass/core/components/_variables.scss</code> and adds
      additional options for the form elements in
      <code>src/assets/sass/components/forms/</code>.
    </div>
    <!--end::Block-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "forms-overview",
  components: {},
  setup() {
    return {
      appName,
    };
  },
});
</script>
