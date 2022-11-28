<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="custom-form-select">
      <router-link to="#custom-form-select"></router-link>
      Custom Form Select
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.form-select-solid</code> class with
      <code>.form-select</code> to set solid background style for a form select:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <div class="mb-10">
          <label class="form-label">Default select</label>
          <select class="form-select" aria-label="Select example">
            <option>Open this select menu</option>
            <option value="1">One</option>
            <option value="2">Two</option>
            <option value="3">Three</option>
          </select>
        </div>
        <div class="mb-10">
          <label class="form-label">Solid background style</label>
          <select
            class="form-select form-select-solid"
            aria-label="Select example"
          >
            <option>Open this select menu</option>
            <option value="1">One</option>
            <option value="2">Two</option>
            <option value="3">Three</option>
          </select>
        </div>
        <div class="mb-10">
          <label class="form-label">White background style</label>
          <select
            class="form-select form-select-white"
            aria-label="Select example"
          >
            <option>Open this select menu</option>
            <option value="1">One</option>
            <option value="2">Two</option>
            <option value="3">Three</option>
          </select>
        </div>
        <!--begin::Code-->
        <!--prettier-ignore-->
        <CodeHighlighter lang="html" field-height="400"
        >{{`<select class="form-select" aria-label="Select example">
  <option>Open this select menu</option>
  <option value="1">One</option>
  <option value="2">Two</option>
  <option value="3">Three</option>
</select>

  <select class="form-select form-select-solid" aria-label="Select example">
    <option>Open this select menu</option>
    <option value="1">One</option>
    <option value="2">Two</option>
    <option value="3">Three</option>
  </select>

  <select class="form-select form-select-white" aria-label="Select example">
    <option>Open this select menu</option>
    <option value="1">One</option>
    <option value="2">Two</option>
    <option value="3">Three</option>
  </select>`}}</CodeHighlighter
        >
        <!--end::Code-->
      </div>
    </div>
    <!--end::Block-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";

export default defineComponent({
  name: "custom-form-select",
  components: {
    CodeHighlighter,
  },
});
</script>
