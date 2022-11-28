<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="letter-spacing">
      <router-link to="#letter-spacing"></router-link>
      Letter Spacing
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.ls-{value}</code> class to set an element's letter spacing
      where <code>value</code> is one of:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="pt-3 pb-5">
      <template v-for="(key, value, i) in lineSpacingValues" :key="i">
        <li class="py-2">
          <code>{{ value }}</code> - for classes that set line height to
          <code>{{ key }}</code>
        </li>
      </template>
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <template v-for="(key, value, i) in lineSpacingValues" :key="i">
          <div :class="`ls-${value}`" class="py-5">
            Example text for <code>.ls-{{ value }}</code>
          </div>
        </template>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <CodeHighlighter lang="html"
        >{{`
        <div class="ls-1">Example text for .ls-1</div>
        <div class="ls-2">Example text for .ls-2</div>
        <div class="ls-3">Example text for .ls-3</div>
        <div class="ls-4">Example text for .ls-4</div>
        <div class="ls-5">Example text for .ls-5</div>
        `}}</CodeHighlighter
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
  name: "letter-spacing",
  components: {
    CodeHighlighter,
  },
  setup() {
    const lineSpacingValues = {
      "1": 0.1,
      "2": 0.115,
      "3": 0.125,
      "4": 0.25,
      "5": 0.5,
    };

    return {
      lineSpacingValues,
    };
  },
});
</script>
