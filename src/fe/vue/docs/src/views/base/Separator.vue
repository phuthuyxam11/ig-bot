<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <BasicExample></BasicExample>

      <Sizes></Sizes>

      <Colors></Colors>

      <Styles></Styles>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/separator/Overview.vue";
import BasicExample from "@/views/base/separator/BasicExample.vue";
import Sizes from "@/views/base/separator/Sizes.vue";
import Colors from "@/views/base/separator/Colors.vue";
import Styles from "@/views/base/separator/Styles.vue";

export default defineComponent({
  name: "base-separator",
  components: {
    Overview,
    BasicExample,
    Sizes,
    Colors,
    Styles,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Separator", ["Base"]);
    });
  },
});
</script>
