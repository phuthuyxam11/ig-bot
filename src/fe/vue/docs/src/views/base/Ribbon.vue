<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>
      <BasicExample></BasicExample>
      <!--      <RoundedCorners></RoundedCorners>-->
      <!--      <VerticalRibbons></VerticalRibbons>-->
      <!--      <ClipRibbons></ClipRibbons>-->
      <!--      <TriangleRibbons></TriangleRibbons>-->
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/ribbon/Overview.vue";
import BasicExample from "@/views/base/ribbon/BasicExample.vue";
// import RoundedCorners from "@/views/base/ribbon/RoundedCorners.vue";
// import VerticalRibbons from "@/views/base/ribbon/VerticalRibbons.vue";
// import ClipRibbons from "@/views/base/ribbon/ClipRibbons.vue";
// import TriangleRibbons from "@/views/base/ribbon/TriangleRibbons.vue";

export default defineComponent({
  name: "base-badges",
  components: {
    Overview,
    BasicExample,
    // RoundedCorners,
    // VerticalRibbons,
    // ClipRibbons,
    // TriangleRibbons,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Alerts", ["Base"]);
    });
  },
});
</script>
