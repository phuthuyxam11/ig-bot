<template>
  <!--begin::Section-->
  <div class="pb-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="bootstrap-buttons">
      <router-link to="#bootstrap-buttons"></router-link>
      Bootstrap Buttons
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      <b>{{ appName }}</b> customizes the
      <a
        href="https://getbootstrap.com/docs/5.0/components/buttons/"
        class="fw-bold"
        >Bootstrap Buttons</a
      >&nbsp; through the SASS variables in
      <code>src/assets/sass/core/components/_variables.scss</code> and adds
      additional options in <code>src/assets/sass/components/buttons/</code>.
    </div>
    <!--end::Block-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "buttons-overview",
  components: {},
  setup() {
    return {
      appName,
    };
  },
});
</script>
