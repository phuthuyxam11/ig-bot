<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <CustomCard></CustomCard>

      <CardScroll></CardScroll>

      <FlushBorders></FlushBorders>

      <ResetSidePaddings></ResetSidePaddings>

      <CardBorder></CardBorder>

      <CardScretch></CardScretch>

      <UtilityClasses></UtilityClasses>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/cards/Overview.vue";
import CustomCard from "@/views/base/cards/CustomCard.vue";
import CardScroll from "@/views/base/cards/CardScroll.vue";
import FlushBorders from "@/views/base/cards/FlushBorders.vue";
import ResetSidePaddings from "@/views/base/cards/ResetSidePaddings.vue";
import CardBorder from "@/views/base/cards/CardBorder.vue";
import CardScretch from "@/views/base/cards/CardStretch.vue";
import UtilityClasses from "@/views/base/cards/UtilityClasses.vue";

export default defineComponent({
  name: "base-cards",
  components: {
    Overview,
    CustomCard,
    CardScroll,
    FlushBorders,
    ResetSidePaddings,
    CardBorder,
    CardScretch,
    UtilityClasses,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Cards", ["Base"]);
    });
  },
});
</script>
