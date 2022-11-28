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
      Use the <code>.ribbon</code> class within the
      <code>.card-header</code> element of a <code>.card</code> component. Then
      add the ribbon content within the <code>.card-header</code> This example
      uses a simple ribbon on a card component.
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <!--begin::Card-->
      <div class="card card-bordered">
        <div class="card-header ribbon">
          <div class="card-title">Ribbon Example</div>
          <div class="ribbon-label bg-primary">Ribbon</div>
        </div>
        <div class="card-body">
          Lorem Ipsum is simply dummy text of the printing and typesetting
          industry. Lorem Ipsum has been the industry's standard dummy text ever
          since the 1500s, when an unknown printer took a galley of type and
          scrambled it to make a type specimen book. It has survived not only
          five centuries, but also the leap into electronic typesetting,
          remaining essentially unchanged. It was popularised in the 1960s with
          the release of Letraset sheets containing Lorem Ipsum passages, and
          more recently with desktop publishing software like Aldus PageMaker
          including versions of Lorem Ipsum.
        </div>
      </div>
      <!--end::Card-->
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
      >{{`<div class="card card-bordered">
        <div class="card-header ribbon">
        <div class="card-title">Ribbon Example</div>
        <div class="ribbon-label bg-primary">Ribbon</div>
      </div>
        <div class="card-body">...</div>
    </div>`}}</CodeHighlighter>
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
