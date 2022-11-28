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
        class="fw-bolder me-1"
        href="https://getbootstrap.com/docs/5.0/content/tables/"
        >Bootstrap Tables</a
      >
      to provide wider variants for multi-purpose usage as static table and
      integrated with
      <a href="https://datatables.net/" class="fw-bolder me-1" target="_blank"
        >Datatables</a
      >
      as well.
    </div>
    <!--end::Block-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "tables-overview",
  components: {},
  setup() {
    return {
      appName,
    };
  },
});
</script>
