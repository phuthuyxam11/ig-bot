<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>
      <ImageUsage></ImageUsage>
      <InlineUsage></InlineUsage>
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
import Overview from "@/views/icons/duotune/Overview.vue";
import ImageUsage from "@/views/icons/duotune/ImageUsage.vue";
import InlineUsage from "@/views/icons/duotune/InlineUsage.vue";
import IconColors from "@/views/icons/duotune/IconColors.vue";
import IconSizes from "@/views/icons/bootstrap-icons/IconSizes.vue";

export default defineComponent({
  name: "duotune-page",
  components: {
    Overview,
    ImageUsage,
    InlineUsage,
    IconColors,
    IconSizes,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Duotune", ["Icons"]);
    });
  },
});
</script>
