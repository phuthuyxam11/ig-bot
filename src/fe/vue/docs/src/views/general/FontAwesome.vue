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
import Overview from "@/views/general/font-awesome/Overview.vue";
import Usage from "@/views/general/font-awesome/Usage.vue";
import IconColors from "@/views/general/font-awesome/IconColors.vue";
import IconSizes from "@/views/general/font-awesome/IconSizes.vue";

export default defineComponent({
  name: "font-awesome",
  components: {
    Overview,
    Usage,
    IconColors,
    IconSizes,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Font Awesome", ["General", "Icons"]);
    });
  },
});
</script>
