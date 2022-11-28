<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <Basic></Basic>

      <Toggle></Toggle>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/indicator/Overview.vue";
import Basic from "@/views/base/indicator/Basic.vue";
import Toggle from "@/views/base/indicator/Toggle.vue";

export default defineComponent({
  name: "base-indicator",
  components: {
    Overview,
    Basic,
    Toggle,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Indicator", ["Base"]);
    });
  },
});
</script>
