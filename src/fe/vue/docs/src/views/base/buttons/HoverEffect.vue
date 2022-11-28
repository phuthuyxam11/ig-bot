<template>
  <!--begin::Section-->
  <div class="pb-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="hover-effect">
      <router-link to="#hover-effect"></router-link>
      Hover Effect
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.btn-hover-{effect}</code> class to set various hover effects
      for a button:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <a href="#" class="btn btn-primary btn-hover-rise me-3">Rise</a>
        <a href="#" class="btn btn-danger btn-hover-scale me-3">Scale</a>
        <a href="#" class="btn btn-success btn-hover-rotate-end me-3"
          >Rotate to end</a
        >
        <a href="#" class="btn btn-warning btn-hover-rotate-start"
          >Rotate to start</a
        >
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <CodeHighlighter lang="html"
        >{{`<a href="#" class="btn btn-primary btn-hover-rise me-5">Rise</a>
        <a href="#" class="btn btn-danger btn-hover-scale me-5">Scale</a>
        <a href="#" class="btn btn-success btn-hover-rotate-end"
          >Rotate to end</a
        >
        <a href="#" class="btn btn-warning btn-hover-rotate-start"
          >Rotate to start</a
        >`}}</CodeHighlighter
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
  name: "hover-effect",
  components: {
    CodeHighlighter,
  },
});
</script>
