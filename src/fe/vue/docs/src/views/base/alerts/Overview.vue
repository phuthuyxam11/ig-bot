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
      {{ appFullName }} customizes the
      <a href="https://getbootstrap.com/docs/5.1/components/alerts/"
        >Bootstrap Alerts</a
      >
      through creative use of Bootstrap utilities and with handpicked icons.
    </div>
    <!--end::Block-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { appFullName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "alerts-overview",
  components: {},
  setup() {
    return {
      appFullName,
    };
  },
});
</script>
