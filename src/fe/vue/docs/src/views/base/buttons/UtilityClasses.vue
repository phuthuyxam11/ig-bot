<template>
  <!--begin::Section-->
  <div class="pb-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="utility-classes">
      <router-link to="#utility-classes"></router-link>
      Utility Classes
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.btn-flush</code> class to remove paddings, borders, background
      and rounded corners:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <a href="#" class="btn btn-flush">Flushed button</a>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
        >{{`<a href="#" class="btn btn-flush">Flushed button</a>`}}</CodeHighlighter
      >
    </div>
    <!--end::Code-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.btn-flex</code> class to vertically center button inner
      elements with display: inline-flex align-items: center:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <a href="#" class="btn btn-flex btn-primary px-6">
        <span class="svg-icon svg-icon-2x">
          <inline-svg src="media/icons/duotune/general/gen006.svg" />
        </span>
        <span class="d-flex flex-column align-items-start ms-2">
          <span class="fs-3 fw-bolder">Caption</span>
          <span class="fs-7">Some description</span>
        </span>
      </a>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
        >{{`<a href="#" class="btn btn-flex btn-primary px-6">
        <span class="svg-icon svg-icon-2x">
          <inline-svg src="media/icons/duotune/general/gen006.svg" />
        </span>
        <span class="d-flex flex-column align-items-start ms-2">
          <span class="fs-3 fw-bolder">Caption</span>
          <span class="fs-7">Some description</span>
        </span>
      </a>`}}</CodeHighlighter
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
  name: "utility-classes",
  components: {
    CodeHighlighter,
  },
});
</script>
