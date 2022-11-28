<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <Basic></Basic>

      <LightStyle></LightStyle>

      <SquareBadge></SquareBadge>

      <CircleBadge></CircleBadge>

      <BootstrapButtons></BootstrapButtons>

      <InsideButtons></InsideButtons>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/badges/Overview.vue";
import Basic from "@/views/base/badges/Basic.vue";
import LightStyle from "@/views/base/badges/LightStyle.vue";
import SquareBadge from "@/views/base/badges/SquareBadge.vue";
import CircleBadge from "@/views/base/badges/CircleBadge.vue";
import BootstrapButtons from "@/views/base/badges/BootstrapButtons.vue";
import InsideButtons from "@/views/base/badges/InsideButtons.vue";

export default defineComponent({
  name: "base-badges",
  components: {
    Overview,
    Basic,
    LightStyle,
    SquareBadge,
    CircleBadge,
    BootstrapButtons,
    InsideButtons,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Badges", ["Base"]);
    });
  },
});
</script>
