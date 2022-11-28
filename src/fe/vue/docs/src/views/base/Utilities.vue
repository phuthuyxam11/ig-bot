<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <ThemeColor></ThemeColor>

      <Spacing></Spacing>

      <FontSize></FontSize>

      <LineHeight></LineHeight>

      <LetterSpacing></LetterSpacing>

      <WidthAndHeight></WidthAndHeight>

      <ZIndex></ZIndex>

      <CSSVarsInJavascript></CSSVarsInJavascript>
    </div>
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/utilities/Overview.vue";
import ThemeColor from "@/views/base/utilities/ThemeColor.vue";
import Spacing from "@/views/base/utilities/Spacing.vue";
import FontSize from "@/views/base/utilities/FontSize.vue";
import LineHeight from "@/views/base/utilities/LineHeight.vue";
import LetterSpacing from "@/views/base/utilities/LetterSpacing.vue";
import WidthAndHeight from "@/views/base/utilities/WidthAndHeight.vue";
import ZIndex from "@/views/base/utilities/ZIndex.vue";
import CSSVarsInJavascript from "@/views/base/utilities/CSSVarsInJavascript.vue";

export default defineComponent({
  name: "base-utilities",
  components: {
    Overview,
    ThemeColor,
    Spacing,
    FontSize,
    LineHeight,
    LetterSpacing,
    WidthAndHeight,
    ZIndex,
    CSSVarsInJavascript,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Utilities", ["Base"]);
    });
  },
});
</script>
