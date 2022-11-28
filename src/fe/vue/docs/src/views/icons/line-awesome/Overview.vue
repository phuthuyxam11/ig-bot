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
      Swap Font Awesome for modern
      <a href="https://icons8.com/line-awesome" class="fw-bold" target="_blank"
        >Line Awesome Icons</a
      >&nbsp;in one line of code and use with <b>{{ appName }}</b> right away.
    </div>
    <!--end::Block-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "line-awesome-overview",
  components: {},
  setup() {
    return {
      appName,
    };
  },
});
</script>
