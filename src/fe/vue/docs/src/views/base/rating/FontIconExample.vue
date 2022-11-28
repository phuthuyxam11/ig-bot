<template>
  <!--begin::Section-->
  <div class="pt-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="font-icons-example">
      <router-link to="#font-icons-example"></router-link>
      Font Icons Example
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      As a star icon any font icons can be used as well. This examples uses a
      star icon from Bootstrap Icons:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rating">
        <div class="rating-label me-2 checked">
          <i class="bi bi-star fs-1"></i>
        </div>
        <div class="rating-label me-2 checked">
          <i class="bi bi-star fs-1"></i>
        </div>
        <div class="rating-label me-2 checked">
          <i class="bi bi-star fs-1"></i>
        </div>
        <div class="rating-label me-2 checked">
          <i class="bi bi-star fs-1"></i>
        </div>
        <div class="rating-label me-2 checked">
          <i class="bi bi-star fs-1"></i>
        </div>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
      >{{`<div class="rating">
        <div class="rating-label me-2 checked">
          <i class="bi bi-star fs-1"></i>
        </div>
        <div class="rating-label me-2 checked">
          <i class="bi bi-star fs-1"></i>
        </div>
        <div class="rating-label me-2 checked">
          <i class="bi bi-star fs-1"></i>
        </div>
        <div class="rating-label me-2 checked">
          <i class="bi bi-star fs-1"></i>
        </div>
        <div class="rating-label me-2 checked">
          <i class="bi bi-star fs-1"></i>
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
