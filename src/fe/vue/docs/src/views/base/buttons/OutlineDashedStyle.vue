<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="outline-dashed-style">
      <router-link to="#outline-dashed-style"></router-link>
      Outline Dashed Style
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.btn-outline</code>and <code>.btn-outline-dashed</code>Bootstrap
      Button classes with the standard outline theme color
      <code>.btn-outline-{color}</code>and
      <code>.btn-active-light-{color}</code>classes to have a button with
      outline dashed border style:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <a
          href="#"
          class="btn btn-outline btn-outline-dashed btn-outline-default me-2 mb-2"
          >Default</a
        >
        <a
          href="#"
          class="btn btn-outline btn-outline-dashed btn-outline-primary btn-active-light-primary me-2 mb-2"
          >Primary</a
        >
        <a
          href="#"
          class="btn btn-outline btn-outline-dashed btn-outline-success btn-active-light-success me-2 mb-2"
          >Success</a
        >
        <a
          href="#"
          class="btn btn-outline btn-outline-dashed btn-outline-info btn-active-light-info me-2 mb-2"
          >Info</a
        >
        <a
          href="#"
          class="btn btn-outline btn-outline-dashed btn-outline-warning btn-active-light-warning me-2 mb-2"
          >Warning</a
        >
        <a
          href="#"
          class="btn btn-outline btn-outline-dashed btn-outline-danger btn-active-light-danger me-2 mb-2"
          >Danger</a
        >
        <a
          href="#"
          class="btn btn-outline btn-outline-dashed btn-outline-dark btn-active-light-dark me-2 mb-2"
          >Dark</a
        >
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
        >{{`<a href="#" class="btn btn-outline btn-outline-dashed btn-outline-default me-2 mb-2">Default</a>
<a href="#" class="btn btn-outline btn-outline-dashed btn-outline-primary btn-active-light-primary me-2 mb-2">Primary</a>
<a href="#" class="btn btn-outline btn-outline-dashed btn-outline-success btn-active-light-success me-2 mb-2">Success</a>
<a href="#" class="btn btn-outline btn-outline-dashed btn-outline-info btn-active-light-info me-2 mb-2">Info</a>
<a href="#" class="btn btn-outline btn-outline-dashed btn-outline-warning btn-active-light-warning me-2 mb-2">Warning</a>
<a href="#" class="btn btn-outline btn-outline-dashed btn-outline-danger btn-active-light-danger me-2 mb-2">Danger</a>
<a href="#" class="btn btn-outline btn-outline-dashed btn-outline-dark btn-active-light-dark me-2 mb-2">Dark</a>`}}</CodeHighlighter>
    </div>
    <!--end::Code-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";

export default defineComponent({
  name: "buttons-overview",
  components: {
    CodeHighlighter,
  },
});
</script>
