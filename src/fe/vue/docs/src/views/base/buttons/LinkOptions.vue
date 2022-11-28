<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="link-option">
      <router-link to="#link-option"></router-link>
      Link Option
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.btn-link</code> Bootstrap Button class with any other button
      classes to set a link buttons without side paddings:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <a
          href="#"
          class="btn btn-link btn-color-muted btn-active-color-primary me-5 mb-2"
          >Muted</a
        >
        <a
          href="#"
          class="btn btn-link btn-color-gray-500 btn-active-color-primary me-5 mb-2"
          >Gray 500</a
        >
        <a
          href="#"
          class="btn btn-link btn-color-info btn-active-color-primary me-5 mb-2"
          >Primary</a
        >
        <a
          href="#"
          class="btn btn-link btn-color-success btn-active-color-primary me-5 mb-2"
          >Success</a
        >
        <a
          href="#"
          class="btn btn-link btn-color-danger btn-active-color-primary me-5 mb-2"
          >Danger</a
        >
        <a
          href="#"
          class="btn btn-link btn-color-warning btn-active-color-primary me-5 mb-2"
          >Warning</a
        >
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <CodeHighlighter lang="html"
        >{{`<a
          href="#"
          class="btn btn-link btn-color-muted btn-active-color-primary me-5 mb-2"
          >Muted</a
        >
        <a
          href="#"
          class="btn btn-link btn-color-gray-500 btn-active-color-primary me-5 mb-2"
          >Gray 500</a
        >
        <a
          href="#"
          class="btn btn-link btn-color-info btn-active-color-primary me-5 mb-2"
          >Primary</a
        >
        <a
          href="#"
          class="btn btn-link btn-color-success btn-active-color-primary me-5 mb-2"
          >Success</a
        >
        <a
          href="#"
          class="btn btn-link btn-color-danger btn-active-color-primary me-5 mb-2"
          >Danger</a
        >
        <a
          href="#"
          class="btn btn-link btn-color-warning btn-active-color-primary me-5 mb-2"
          >Warning</a
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
