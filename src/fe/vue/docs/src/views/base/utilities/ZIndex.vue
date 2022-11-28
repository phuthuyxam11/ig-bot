<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="z-index">
      <router-link to="#z-index"></router-link>
      Z-Index
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.z-index-{value}</code> class to set an element's z-index:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      Where <code>value</code> is one of:
      <ul class="pt-3">
        <template v-for="i in 6" :key="i">
          <li class="py-2">
            <code>{{ zIndexPoints[i - 1] }}</code> - for classes that set
            z-index
            <code>{{ zIndexValues[i - 1] }}</code>
          </li>
        </template>
      </ul>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <CodeHighlighter lang="html"
        >{{`
        <div class="z-index-n1"></div>
        <div class="z-index-n2"></div>
        <div class="z-index-0"></div>
        <div class="z-index-1"></div>
        <div class="z-index-2"></div>
        <div class="z-index-3"></div>
        `}}</CodeHighlighter
      >
    </div>
    <!--end::Code-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";

export default defineComponent({
  name: "z-index",
  components: {
    CodeHighlighter,
  },
  setup() {
    const zIndexValues = ["-1", "-2", "0", "1", "2", "3"];
    const zIndexPoints = ["n1", "n2", "0", "1", "2", "3"];

    return {
      zIndexValues,
      zIndexPoints,
    };
  },
});
</script>
