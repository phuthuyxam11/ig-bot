<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="flush-form-controls">
      <router-link to="#flush-form-controls"></router-link>
      Flush Form Control
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.form-control-flush</code>class with
      <code>.form-control</code>to remove the default background-color, borders,
      and rounded corners of a form input.
    </div>
    <!--end::Block-->

    <div class="py-5">
      <div class="rounded border p-10">
        <div class="mb-10">
          <label for="exampleFormControlInput1" class="form-label"
            >Example</label
          >
          <input
            type="email"
            id="exampleFormControlInput1"
            class="form-control form-control-flush"
            placeholder="Example input"
          />
        </div>
        <!--begin::Code-->
        <!--prettier-ignore-->
        <CodeHighlighter lang="html"
        >{{`<input type="email" class="form-control form-control-flush" placeholder="Example input"/>`}}</CodeHighlighter
        >
        <!--end::Code-->
      </div>
    </div>
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";

export default defineComponent({
  name: "custom-form-control",
  components: {
    CodeHighlighter,
  },
});
</script>
