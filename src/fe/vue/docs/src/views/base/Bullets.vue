<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <Basic></Basic>

      <Colors></Colors>

      <DotStyle></DotStyle>

      <VerticalStyle></VerticalStyle>

      <LineStyle></LineStyle>

      <Sizes></Sizes>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/bullets/Overview.vue";
import Basic from "@/views/base/bullets/Basic.vue";
import Colors from "@/views/base/bullets/Colors.vue";
import DotStyle from "@/views/base/bullets/DotStyle.vue";
import VerticalStyle from "@/views/base/bullets/VarticalStyle.vue";
import LineStyle from "@/views/base/bullets/LineStyle.vue";
import Sizes from "@/views/base/bullets/Sizes.vue";

export default defineComponent({
  name: "base-bullets",
  components: {
    Overview,
    Basic,
    Colors,
    DotStyle,
    VerticalStyle,
    LineStyle,
    Sizes,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Bullets", ["Base"]);
    });
  },
});
</script>
