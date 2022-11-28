<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <BasicExample></BasicExample>

      <FontIconExample></FontIconExample>

      <StarStates></StarStates>

      <EditModeExample></EditModeExample>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/rating/Overview.vue";
import BasicExample from "@/views/base/rating/BasicExample.vue";
import FontIconExample from "@/views/base/rating/FontIconExample.vue";
import StarStates from "@/views/base/rating/StarStates.vue";
import EditModeExample from "@/views/base/rating/EditModeExample.vue";

export default defineComponent({
  name: "base-rating",
  components: {
    Overview,
    BasicExample,
    FontIconExample,
    StarStates,
    EditModeExample,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Rating", ["Base"]);
    });
  },
});
</script>
