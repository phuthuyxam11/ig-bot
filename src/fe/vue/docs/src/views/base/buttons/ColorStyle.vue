<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="color-style">
      <router-link to="#color-style"></router-link>
      Color Style
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.btn-color-{color}</code> class to set a button's text and icon
      colors only:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <a href="#" class="btn btn-bg-dark btn-color-white me-2 mb-2">White</a>
        <a href="#" class="btn btn-bg-light btn-color-primary me-2 mb-2"
          >Primary</a
        >
        <a href="#" class="btn btn-bg-dark btn-color-secondary me-2 mb-2"
          >Secondary</a
        >
        <a href="#" class="btn btn-bg-dark btn-color-light me-2 mb-2">Light</a>
        <a href="#" class="btn btn-bg-light btn-color-success me-2 mb-2"
          >Success</a
        >
        <a href="#" class="btn btn-bg-light btn-color-info me-2 mb-2">Info</a>
        <a href="#" class="btn btn-bg-light btn-color-warning me-2 mb-2"
          >Warning</a
        >
        <a href="#" class="btn btn-bg-light btn-color-danger me-2 mb-2"
          >Danger</a
        >
        <a href="#" class="btn btn-bg-light btn-color-dark me-2 mb-2">Dark</a>
        <a href="#" class="btn btn-bg-light btn-color-muted me-2 mb-2">Muted</a>
        <a href="#" class="btn btn-bg-dark btn-color-gray-100 me-2 mb-2"
          >Gray-100</a
        >
        <a href="#" class="btn btn-bg-dark btn-color-gray-200 me-2 mb-2"
          >Gray-200</a
        >
        <a href="#" class="btn btn-bg-dark btn-color-gray-300 me-2 mb-2"
          >Gray-300</a
        >
        <a href="#" class="btn btn-bg-light btn-color-gray-400 me-2 mb-2"
          >Gray-400</a
        >
        <a href="#" class="btn btn-bg-light btn-color-gray-500 me-2 mb-2"
          >Gray-500</a
        >
        <a href="#" class="btn btn-bg-light btn-color-gray-600 me-2 mb-2"
          >Gray-600</a
        >
        <a href="#" class="btn btn-bg-light btn-color-gray-700 me-2 mb-2"
          >Gray-700</a
        >
        <a href="#" class="btn btn-bg-light btn-color-gray-800 me-2 mb-2"
          >Gray-800</a
        >
        <a href="#" class="btn btn-bg-light btn-color-gray-900 me-2 mb-2"
          >Gray-900</a
        >
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
        >{{`<a href="#" class="btn btn-bg-dark btn-color-white">White</a>
<a href="#" class="btn btn-bg-light btn-color-primary">Primary</a>
<a href="#" class="btn btn-bg-dark btn-color-secondary">Secondary</a>
<a href="#" class="btn btn-bg-dark btn-color-light">Light</a>
<a href="#" class="btn btn-bg-light btn-color-success">Success</a>
<a href="#" class="btn btn-bg-light btn-color-info">Info</a>
<a href="#" class="btn btn-bg-light btn-color-warning">Warning</a>
<a href="#" class="btn btn-bg-light btn-color-danger">Danger</a>
<a href="#" class="btn btn-bg-light btn-color-dark">Dark</a>
<a href="#" class="btn btn-bg-light btn-color-muted">Muted</a>
<a href="#" class="btn btn-bg-dark btn-color-gray-100">Gray-100</a>
<a href="#" class="btn btn-bg-dark btn-color-gray-200">Gray-200</a>
<a href="#" class="btn btn-bg-dark btn-color-gray-300">Gray-300</a>
<a href="#" class="btn btn-bg-light btn-color-gray-400">Gray-400</a>
<a href="#" class="btn btn-bg-light btn-color-gray-500">Gray-500</a>
<a href="#" class="btn btn-bg-light btn-color-gray-600">Gray-600</a>
<a href="#" class="btn btn-bg-light btn-color-gray-700">Gray-700</a>
<a href="#" class="btn btn-bg-light btn-color-gray-800">Gray-800</a>
<a href="#" class="btn btn-bg-light btn-color-gray-900">Gray-900</a>`}}</CodeHighlighter
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
  name: "buttons-overview",
  components: {
    CodeHighlighter,
  },
});
</script>
