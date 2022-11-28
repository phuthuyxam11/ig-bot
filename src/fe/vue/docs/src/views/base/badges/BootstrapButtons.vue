<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="bootstrap-buttons">
      <router-link to="#bootstrap-buttons"></router-link>
      Bootstrap Buttons
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use a badge with <code>.badge-circle</code> class in a Bootstrap Button
      element:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10 pb-0 d-flex flex-wrap">
        <button class="btn btn-primary position-relative me-5">
          Button
          <span
            class="position-absolute top-0 start-0 translate-middle badge badge-circle badge-danger"
            >5</span
          >
        </button>

        <button class="btn btn-success position-relative me-5">
          Button
          <span
            class="position-absolute top-100 start-0 translate-middle badge badge-circle badge-warning"
            >5</span
          >
        </button>

        <button class="btn btn-danger position-relative me-5">
          Button
          <span
            class="position-absolute top-0 start-100 translate-middle badge badge-circle badge-primary"
            >5</span
          >
        </button>

        <button class="btn btn-warning position-relative">
          Button
          <span
            class="position-absolute top-100 start-100 translate-middle badge badge-circle badge-success"
            >5</span
          >
        </button>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
      >{{`<button class="btn btn-primary position-relative me-5">
        Button <span class="position-absolute top-0 start-0 translate-middle  badge badge-circle badge-danger">5</span>
      </button>

        <button class="btn btn-success position-relative me-5">
          Button <span class="position-absolute top-100 start-0 translate-middle  badge badge-circle badge-warning">5</span>
        </button>

        <button class="btn btn-danger position-relative me-5">
          Button <span class="position-absolute top-0 start-100 translate-middle  badge badge-circle badge-primary">5</span>
        </button>

        <button class="btn btn-warning position-relative">
          Button <span class="position-absolute top-100 start-100 translate-middle  badge badge-circle badge-success">5</span>
        </button>`}}</CodeHighlighter
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
  name: "bootstrap-buttons",
  components: {
    CodeHighlighter,
  },
});
</script>
