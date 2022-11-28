<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <CustomExample></CustomExample>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/carousel/Overview.vue";
import CustomExample from "@/views/base/carousel/CustomExample.vue";

export default defineComponent({
  name: "base-carousel",
  components: {
    Overview,
    CustomExample,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Carousel", ["Base"]);
    });
  },
});
</script>
