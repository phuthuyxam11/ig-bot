<template>
  <!--begin::Section-->
  <div class="pt-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="collapse">
      <router-link to="#collapse"></router-link>
      Collapse
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use Rotate with <code>.collapsible</code> and
      <code>.collapsed</code> classes to toggle a block using Bootstrap Collapse
      and display:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <!--begin::Toggle link-->
        <a
          data-bs-toggle="collapse"
          href="#kt_toggle_block"
          class="btn btn-link btn-color-muted btn-active-color-primary rotate collapsible collapsed"
        >
          Toggle block
          <span class="svg-icon svg-icon-3 rotate-n180 ms-1"><svg></svg></span>
        </a>
        <!--end::Toggle link-->

        <!--begin::Toggle block-->
        <div id="kt_toggle_block" class="collapse">
          <div class="rounded bg-light p-10">.....</div>
        </div>
        <!--end::Toggle block-->
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
        >{{`<!--begin::Toggle link-->
<a data-bs-toggle="collapse" href="#kt_toggle_block" class="btn btn-link btn-color-muted btn-active-color-primary rotate collapsible collapsed">
        Toggle block
        <span class="svg-icon svg-icon-3 rotate-n180 ms-1"><svg></svg></span>
      </a>
        <!--end::Toggle link-->

        <!--begin::Toggle block-->
        <div id="kt_toggle_block" class="collapse">
          <div class="rounded bg-light p-10">
            .....
          </div>
        </div>
        <!--end::Toggle block-->`}}</CodeHighlighter
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
  name: "dot-separator",
  components: {
    CodeHighlighter,
  },
});
</script>
