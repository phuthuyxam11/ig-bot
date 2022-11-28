<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="custom-form-control">
      <router-link to="#custom-form-control"></router-link>
      Custom Form Control
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.form-control-solid</code> class with
      <code>.form-control</code> to set solid background style for a form
      control:
    </div>
    <!--end::Block-->

    <div class="py-5">
      <div class="rounded border p-10">
        <div class="mb-10">
          <label class="form-label">Default input</label>
          <input
            type="text"
            class="form-control"
            placeholder="name@example.com"
          />
        </div>
        <div class="mb-10">
          <label class="form-label">Solid background style</label>
          <input
            type="text"
            class="form-control form-control-solid"
            placeholder="name@example.com"
          />
        </div>
        <div class="mb-10">
          <label class="form-label">White background style</label>
          <input
            type="text"
            class="form-control form-control-white"
            placeholder="name@example.com"
          />
        </div>
        <!--begin::Code-->
        <div class="py-0">
          <!--prettier-ignore-->
          <CodeHighlighter lang="html"
          >{{`<input type="text" class="form-control" placeholder="name@example.com"/>
<input type="text" class="form-control form-control-solid" placeholder="name@example.com"/>
<input type="text" class="form-control form-control-white" placeholder="name@example.com"/>`}}</CodeHighlighter
          >
        </div>
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
