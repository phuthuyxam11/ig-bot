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
      Use an ordered or unordered list with linked list items to create a
      minimally styled breadcrumb. Use our utilities to add additional styles as
      desired.
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <ol class="breadcrumb text-muted fs-6 fw-bold">
          <li class="breadcrumb-item pe-3">
            <a href="#" class="pe-3">Home</a>
          </li>
          <li class="breadcrumb-item pe-3">
            <a href="#" class="pe-3">Library</a>
          </li>
          <li class="breadcrumb-item px-3 text-muted">Active</li>
        </ol>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
        >{{`<ol class="breadcrumb text-muted fs-6 fw-bold">
        <li class="breadcrumb-item pe-3"><a href="#" class="pe-3">Home</a></li>
        <li class="breadcrumb-item pe-3"><a href="#" class="pe-3">Library</a></li>
        <li class="breadcrumb-item px-3 text-muted">Active</li>
      </ol>`}}</CodeHighlighter
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
