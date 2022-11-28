<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <BasicExample></BasicExample>

      <LineSeparator></LineSeparator>

      <DotSeparator></DotSeparator>

      <WithoutSeparator></WithoutSeparator>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/breadcrumb/Overview.vue";
import BasicExample from "@/views/base/breadcrumb/BasicExample.vue";
import LineSeparator from "@/views/base/breadcrumb/LineSeparator.vue";
import DotSeparator from "@/views/base/breadcrumb/DotSeparator.vue";
import WithoutSeparator from "@/views/base/breadcrumb/WithoutSeparator.vue";

export default defineComponent({
  name: "base-breadcrumb",
  components: {
    Overview,
    BasicExample,
    LineSeparator,
    DotSeparator,
    WithoutSeparator,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Breadcrumb", ["Base"]);
    });
  },
});
</script>
