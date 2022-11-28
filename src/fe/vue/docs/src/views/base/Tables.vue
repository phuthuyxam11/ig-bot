<template>
  <!--begin::Card-->
  <div class="card">
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

      <Example8></Example8>

      <Example9></Example9>

      <Example10></Example10>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/tables/Overview.vue";
import Example1 from "@/views/base/tables/Example1.vue";
import Example2 from "@/views/base/tables/Example2.vue";
import Example3 from "@/views/base/tables/Example3.vue";
import Example4 from "@/views/base/tables/Example4.vue";
import Example5 from "@/views/base/tables/Example5.vue";
import Example6 from "@/views/base/tables/Example6.vue";
import Example7 from "@/views/base/tables/Example7.vue";
import Example8 from "@/views/base/tables/Example8.vue";
import Example9 from "@/views/base/tables/Example9.vue";
import Example10 from "@/views/base/tables/Example10.vue";

export default defineComponent({
  name: "base-tables",
  components: {
    Overview,
    Example1,
    Example2,
    Example3,
    Example4,
    Example5,
    Example6,
    Example7,
    Example8,
    Example9,
    Example10,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Tables", ["Base"]);
    });
  },
});
</script>
