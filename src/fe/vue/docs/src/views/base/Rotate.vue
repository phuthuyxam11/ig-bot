<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <Basic></Basic>

      <Toggle></Toggle>

      <Collapse></Collapse>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/rotate/Overview.vue";
import Basic from "@/views/base/rotate/Basic.vue";
import Toggle from "@/views/base/rotate/Toggle.vue";
import Collapse from "@/views/base/rotate/Collapse.vue";

export default defineComponent({
  name: "base-indicator",
  components: {
    Overview,
    Basic,
    Toggle,
    Collapse,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Rotate", ["Base"]);
    });
  },
});
</script>
