<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <BasicExample></BasicExample>

      <FullScreen></FullScreen>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/modal/Overview.vue";
import BasicExample from "@/views/base/modal/BasicExample.vue";
import FullScreen from "@/views/base/modal/FullScreen.vue";

export default defineComponent({
  name: "base-modal",
  components: {
    Overview,
    BasicExample,
    FullScreen,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Modal", ["Base"]);
    });
  },
});
</script>
