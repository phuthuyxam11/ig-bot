<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="form-labels">
      <router-link to="#form-labels"></router-link>
      Form Labels
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.required</code>class with <code>.form-label</code>or any custom
      element to indicate input's mandatory status:
    </div>
    <!--end::Block-->

    <div class="py-5">
      <div class="rounded border p-10">
        <div class="mb-10">
          <label for="exampleFormControlInput1" class="required form-label"
            >Example</label
          >
          <input
            type="email"
            class="form-control form-control-solid"
            placeholder="Example input"
          />
        </div>
        <!--begin::Code-->
        <!--prettier-ignore-->
        <CodeHighlighter lang="html"
        >{{`<div class="mb-10">
          <label for="exampleFormControlInput1" class="required form-label">Example</label>
          <input type="email" id="exampleFormControlInput1" class="form-control form-control-solid" placeholder="Example input"/>
        </div>`}}</CodeHighlighter
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
