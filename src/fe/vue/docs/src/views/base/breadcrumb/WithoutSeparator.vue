<template>
  <!--begin::Section-->
  <div class="pt-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="without-separator">
      <router-link to="#without-separator"></router-link>
      Without Separator
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Add <code>.breadcrumb-separatorless</code> to remove the separator.
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <ol class="breadcrumb breadcrumb-separatorless text-muted fs-6 fw-bold">
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
        >{{`<ol class="breadcrumb breadcrumb-separatorless text-muted fs-6 fw-bold">
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
