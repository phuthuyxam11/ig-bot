<template>
  <!--begin::Section-->
  <div class="pt-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="circle-example">
      <router-link to="#circle-example"></router-link>
      Circle Example
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">Pagination with circle buttons:</div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <ul class="pagination pagination-circle">
          <li class="page-item previous disabled">
            <a href="#" class="page-link"><i class="previous"></i></a>
          </li>
          <li class="page-item"><a href="#" class="page-link">1</a></li>
          <li class="page-item active"><a href="#" class="page-link">2</a></li>
          <li class="page-item"><a href="#" class="page-link">3</a></li>
          <li class="page-item"><a href="#" class="page-link">4</a></li>
          <li class="page-item"><a href="#" class="page-link">5</a></li>
          <li class="page-item"><a href="#" class="page-link">6</a></li>
          <li class="page-item next">
            <a href="#" class="page-link"><i class="next"></i></a>
          </li>
        </ul>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
        >{{`<ul class="pagination pagination-circle">
        <li class="page-item previous disabled"><a href="#" class="page-link"><i class="previous"></i></a></li>
        <li class="page-item "><a href="#" class="page-link">1</a></li>
        <li class="page-item active"><a href="#" class="page-link">2</a></li>
        <li class="page-item "><a href="#" class="page-link">3</a></li>
        <li class="page-item "><a href="#" class="page-link">4</a></li>
        <li class="page-item "><a href="#" class="page-link">5</a></li>
        <li class="page-item "><a href="#" class="page-link">6</a></li>
        <li class="page-item next"><a href="#"  class="page-link"><i class="next"></i></a></li>
      </ul>`}}</CodeHighlighter
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
