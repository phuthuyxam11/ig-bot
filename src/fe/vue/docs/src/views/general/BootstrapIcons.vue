<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>
      <Usage></Usage>
      <IconColors></IconColors>
      <IconSizes></IconSizes>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/general/bootstrap-icons/Overview.vue";
import Usage from "@/views/general/bootstrap-icons/Usage.vue";
import IconColors from "@/views/general/bootstrap-icons/IconColors.vue";
import IconSizes from "@/views/general/bootstrap-icons/IconSizes.vue";

export default defineComponent({
  name: "bootstrap-icons",
  components: {
    Overview,
    Usage,
    IconColors,
    IconSizes,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Bootstrap Icons", ["General", "Icons"]);
    });
  },
});
</script>
