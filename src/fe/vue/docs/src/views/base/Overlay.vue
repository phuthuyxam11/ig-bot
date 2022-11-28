<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <BasicExample></BasicExample>

      <ContentBlocking></ContentBlocking>

      <CardBlocking></CardBlocking>

      <ModalBlocking></ModalBlocking>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/overlay/Overview.vue";
import BasicExample from "@/views/base/overlay/BasicExample.vue";
import ContentBlocking from "@/views/base/overlay/ContentBlocking.vue";
import CardBlocking from "@/views/base/overlay/CardBlocking.vue";
import ModalBlocking from "@/views/base/overlay/ModalBlocking.vue";

export default defineComponent({
  name: "base-overlay",
  components: {
    Overview,
    BasicExample,
    ContentBlocking,
    CardBlocking,
    ModalBlocking,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Overlay", ["Base"]);
    });
  },
});
</script>
