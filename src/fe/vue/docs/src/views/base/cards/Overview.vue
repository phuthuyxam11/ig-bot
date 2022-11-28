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
      <b>{{ appName }}</b> extends
      <a
        href="https://getbootstrap.com/docs/5.0/components/card/"
        class="fw-bold"
        >Bootstrap Card</a
      >
      through the SASS variables in
      <code>src/assets/sass/components/_variables.scss</code> and adds
      additional options in <code>src/assets/sass/components/_card.scss/</code>.
    </div>
    <!--end::Block-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "cards-overview",
  components: {},
  setup() {
    return {
      appName,
    };
  },
});
</script>
