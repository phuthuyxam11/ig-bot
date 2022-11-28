<template>
  <!--begin::Section-->
  <div class="pt-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="basic-example">
      <router-link to="#basic-example"></router-link>
      Basic Example
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.alert</code> with <code>.alert-{color}</code> classes to set
      the alert's style defined with <code>$theme-colors</code> mapped in
      <code>src/assets/sass/components/_variables.scss</code>:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10 pb-0 d-flex flex-column">
        <div class="alert alert-primary d-flex align-items-center p-5 mb-10">
          <span class="svg-icon svg-icon-2hx svg-icon-primary me-4">
            <inline-svg src="media/icons/duotune/general/gen048.svg" />
          </span>
          <div class="d-flex flex-column">
            <h4 class="mb-1 text-primary">This is an alert</h4>
            <span
              >The alert component can be used to highlight certain parts of
              your page for higher content visibility.</span
            >
          </div>
        </div>
        <div class="alert alert-success d-flex align-items-center p-5 mb-10">
          <span class="svg-icon svg-icon-2hx svg-icon-success me-4">
            <inline-svg src="media/icons/duotune/general/gen048.svg" />
          </span>

          <div class="d-flex flex-column">
            <h4 class="mb-1 text-success">This is an alert</h4>
            <span
              >The alert component can be used to highlight certain parts of
              your page for higher content visibility.</span
            >
          </div>
        </div>
        <div class="alert alert-info d-flex align-items-center p-5 mb-10">
          <span class="svg-icon svg-icon-2hx svg-icon-info me-4">
            <inline-svg src="media/icons/duotune/general/gen048.svg" />
          </span>

          <div class="d-flex flex-column">
            <h4 class="mb-1 text-info">This is an alert</h4>
            <span
              >The alert component can be used to highlight certain parts of
              your page for higher content visibility.</span
            >
          </div>
        </div>
        <div class="alert alert-warning d-flex align-items-center p-5 mb-10">
          <span class="svg-icon svg-icon-2hx svg-icon-warning me-4">
            <inline-svg src="media/icons/duotune/general/gen048.svg" />
          </span>

          <div class="d-flex flex-column">
            <h4 class="mb-1 text-warning">This is an alert</h4>
            <span
              >The alert component can be used to highlight certain parts of
              your page for higher content visibility.</span
            >
          </div>
        </div>
        <div class="alert alert-danger d-flex align-items-center p-5 mb-10">
          <span class="svg-icon svg-icon-2hx svg-icon-danger me-4">
            <inline-svg src="media/icons/duotune/general/gen048.svg" />
          </span>

          <div class="d-flex flex-column">
            <h4 class="mb-1 text-danger">This is an alert</h4>
            <span
              >The alert component can be used to highlight certain parts of
              your page for higher content visibility.</span
            >
          </div>
        </div>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
        >{{`<div class="alert alert-primary d-flex align-items-center p-5 mb-10">
          <span class="svg-icon svg-icon-2hx svg-icon-primary me-4">
            <inline-svg src="media/icons/duotune/general/gen048.svg" />
          </span>
        <div class="d-flex flex-column">
          <h4 class="mb-1 text-primary">This is an alert</h4>
          <span
          >The alert component can be used to highlight certain parts of
              your page for higher content visibility.</span
          >
        </div>
      </div>`}}</CodeHighlighter
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
  name: "basic-example",
  components: {
    CodeHighlighter,
  },
});
</script>
