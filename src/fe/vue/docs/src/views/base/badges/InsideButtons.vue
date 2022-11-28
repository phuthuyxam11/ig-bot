<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="inside-button">
      <router-link to="#inside-button"></router-link>
      Inside Button
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use a badge with <code>.badge-circle</code> classes inside button:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10 pb-0 d-flex flex-wrap">
        <button class="btn btn-primary me-5">
          Notifications
          <span class="badge badge-circle badge-warning ms-2">5</span>
        </button>

        <button class="btn btn-success me-5">
          Messages <span class="badge badge-circle badge-danger ms-2">4</span>
        </button>

        <button class="btn btn-danger me-5">
          Orders <span class="badge badge-circle badge-primary ms-2">4</span>
        </button>

        <button class="btn btn-danger me-5">
          Profile <span class="badge badge-circle badge-white ms-2">3</span>
        </button>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html"
        >{{`<button class="btn btn-primary me-5">
        Notifications <span class="badge badge-circle badge-warning ms-2">5</span>
      </button>

        <button class="btn btn-success me-5">
          Messages <span class="badge badge-circle badge-danger ms-2">4</span>
        </button>

        <button class="btn btn-danger me-5">
          Orders <span class="badge badge-circle badge-primary ms-2">4</span>
        </button>

        <button class="btn btn-danger me-5">
          Profile <span class="badge badge-circle badge-white ms-2">3</span>
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
