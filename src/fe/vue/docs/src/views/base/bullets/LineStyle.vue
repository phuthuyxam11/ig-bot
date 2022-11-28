<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="line-style">
      <router-link to="#line-style"></router-link>
      Line Style
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.bullet-line</code> to apply a line style bullet:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <div class="d-flex flex-column">
          <li class="d-flex align-items-center py-2">
            <span class="bullet bullet-line me-5"></span> Item 1
          </li>
          <li class="d-flex align-items-center py-2">
            <span class="bullet bullet-line bg-danger me-5"></span> Item 2
          </li>
          <li class="d-flex align-items-center py-2">
            <span class="bullet bullet-line bg-success me-5"></span> Item 3
          </li>
          <li class="d-flex align-items-center py-2">
            <span class="bullet bullet-line bg-info me-5"></span> Item 4
          </li>
          <li class="d-flex align-items-center py-2">
            <span class="bullet bullet-line bg-primary me-5"></span> Item 5
          </li>
        </div>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <CodeHighlighter lang="html"
        >{{`
        <div class="d-flex flex-column">
          <li class="d-flex align-items-center py-2">
            <span class="bullet bullet-line me-5"></span> Item 1
          </li>
          <li class="d-flex align-items-center py-2">
            <span class="bullet bullet-line bg-danger me-5"></span> Item 2
          </li>
          <li class="d-flex align-items-center py-2">
            <span class="bullet bullet-line bg-success me-5"></span> Item 3
          </li>
          <li class="d-flex align-items-center py-2">
            <span class="bullet bullet-line bg-info me-5"></span> Item 4
          </li>
          <li class="d-flex align-items-center py-2">
            <span class="bullet bullet-line bg-primary me-5"></span> Item 5
          </li>
        </div>
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
  name: "line-style",
  components: {
    CodeHighlighter,
  },
});
</script>
