<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="line-height">
      <router-link to="#line-height"></router-link>
      Line Height
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      <b>{{ appName }}</b> extends
      <a href="#" class="fw-bold">Bootstrap Line Height Classes</a>&nbsp; to
      provide additional line height options by overriding the line height SASS
      variable.
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      Use <code>.lh-{value}</code> class to set an element's line height where
      <code>value</code> is one of:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="pt-3 pb-5">
      <template v-for="(key, value, i) in lineHeightValues" :key="i">
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
        <template v-for="(key, value, i) in lineHeightValues" :key="i">
          <div :class="`lh-${value}`" class="py-5">
            Example text <br />
            for <code>.lh-{{ value }}</code>
          </div>
        </template>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <CodeHighlighter lang="html"
        >{{`
        <div class="lh-1">Example text for .lh-1</div>
        <div class="lh-sm">Example text for .lh-sm</div>
        <div class="lh-base">Example text for .lh-base</div>
        <div class="lh-lg">Example text for .lh-lg</div>
        <div class="lh-xl">Example text for .lh-xl</div>
        <div class="lh-xxl">Example text for .lh-xxl</div>
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
import { appName } from "@/core/helpers/documentation";

export default defineComponent({
  name: "line-height",
  components: {
    CodeHighlighter,
  },
  setup() {
    const lineHeightValues = {
      "1": 1,
      sm: 1.25,
      base: 1.5,
      lg: 1.75,
      xl: 2,
      xxl: 2.25,
    };

    return {
      appName,
      lineHeightValues,
    };
  },
});
</script>
