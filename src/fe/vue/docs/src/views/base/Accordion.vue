<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <BasicExample></BasicExample>

      <AnimatedIcon></AnimatedIcon>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/accordion/Overview.vue";
import BasicExample from "@/views/base/accordion/BasicExample.vue";
import AnimatedIcon from "@/views/base/accordion/AnimatedIcon.vue";

export default defineComponent({
  name: "base-accordion",
  components: {
    Overview,
    BasicExample,
    AnimatedIcon,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Accordion", ["Base"]);
    });
  },
});
</script>
