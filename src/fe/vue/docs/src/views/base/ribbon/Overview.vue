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
      Ribbon is a custom micro element that can be used on a card element for
      additional flair.
    </div>
    <!--end::Block-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "ribbon-overview",
  components: {},
  setup() {
    return {
      appName,
    };
  },
});
</script>
