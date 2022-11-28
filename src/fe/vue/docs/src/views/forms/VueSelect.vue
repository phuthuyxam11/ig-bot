<template>
  <!--begin::Card-->
  <div class="card card-custom">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>
      <Example1></Example1>
      <Example2></Example2>
      <Example3></Example3>
      <Example4></Example4>
      <Example5></Example5>
      <Example6></Example6>
      <Example7></Example7>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/forms/vue-select/Overview.vue";
import Example1 from "@/views/forms/vue-select/Example1.vue";
import Example2 from "@/views/forms/vue-select/Example2.vue";
import Example3 from "@/views/forms/vue-select/Example3.vue";
import Example4 from "@/views/forms/vue-select/Example4.vue";
import Example5 from "@/views/forms/vue-select/Example5.vue";
import Example6 from "@/views/forms/vue-select/Example6.vue";
import Example7 from "@/views/forms/vue-select/Example7.vue";

export default defineComponent({
  name: "forms-select",
  components: {
    Overview,
    Example1,
    Example2,
    Example3,
    Example4,
    Example5,
    Example6,
    Example7,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Vue Multiselect", ["Forms"]);
    });
  },
});
</script>
