<template>
  <!--begin::Section-->
  <div class="py-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="css-variables-in-js">
      <router-link to="#css-variables-in-js"></router-link>
      Access CSS variables in Javascript
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      <ul class="py-0">
        <li class="py-2">
          <code>--bs-light-primary</code>- Light version of primary color hex
          value.
        </li>
        <li class="py-2">
          <code>--bs-active-primary</code>- Active version of primary color hex
          value.
        </li>
        <li class="py-2">
          <code>--bs-light-success</code>- Light version of success color hex
          value.
        </li>
        <li class="py-2">
          <code>--bs-active-success</code>- Active version of success color hex
          value.
        </li>
        <li class="py-2">
          <code>--bs-light-info</code>- Light version of info color hex value.
        </li>
        <li class="py-2">
          <code>--bs-active-info</code>- Active version of info color hex value.
        </li>
        <li class="py-2">
          <code>--bs-light-warning</code>- Light version of warning color hex
          value.
        </li>
        <li class="py-2">
          <code>--bs-active-warning</code>- Active version of warning color hex
          value.
        </li>
        <li class="py-2">
          <code>--bs-light-danger</code>- Light version of danger color hex
          value.
        </li>
        <li class="py-2">
          <code>--bs-active-danger</code>- Active version of danger color hex
          value.
        </li>
        <li class="py-2">
          <code>--bs-light-dark</code>- Light version of dark color hex value.
        </li>
        <li class="py-2">
          <code>--bs-active-dark</code>- Active version of dark color hex value.
        </li>
        <li class="py-2">
          <code>--bs-gray-100</code>- Gray 100 color hex value.
        </li>
        <li class="py-2">
          <code>--bs-gray-200</code>- Gray 200 color hex value.
        </li>
        <li class="py-2">
          <code>--bs-gray-300</code>- Gray 300 color hex value.
        </li>
        <li class="py-2">
          <code>--bs-gray-400</code>- Gray 400 color hex value.
        </li>
        <li class="py-2">
          <code>--bs-gray-500</code>- Gray 500 color hex value.
        </li>
        <li class="py-2">
          <code>--bs-gray-600</code>- Gray 600 color hex value.
        </li>
        <li class="py-2">
          <code>--bs-gray-700</code>- Gray 700 color hex value.
        </li>
        <li class="py-2">
          <code>--bs-gray-800</code>- Gray 800 color hex value.
        </li>
        <li class="py-2">
          <code>--bs-gray-900</code>- Gray 900 color hex value.
        </li>
        <li class="py-2">
          <code>--bs-xs</code>- Extra small breakpoint width value - 0.
        </li>
        <li class="py-2">
          <code>--bs-sm</code>- Extra small breakpoint width value - 576px.
        </li>
        <li class="py-2">
          <code>--bs-md</code>- Medium small breakpoint width value - 768px.
        </li>
        <li class="py-2">
          <code>--bs-lg</code>- Large small breakpoint width value - 992px.
        </li>
        <li class="py-2">
          <code>--bs-xl</code>- Extra large small breakpoint width value -
          1200px.
        </li>
        <li class="py-2">
          <code>--bs-xxl</code>- Double extra large breakpoint width value -
          1400px.
        </li>
        <li class="py-2">
          <code>--bs-font-sans-serif</code>- Main font family. E.g:
          <code>"Roboto", sans-serif</code>.
        </li>
      </ul>
    </div>
    <!--end::Block-->

    <div class="py-5">
      To access these variables, simply call it with this function
      <code>getCSSVariableValue("css variable name")</code>. <br />For example
      to access the primary color, CSS responsive breakpoint value or the theme
      font family you can use the below code:
    </div>

    <!--begin::Code-->
    <div class="py-5">
      <CodeHighlighter lang="javascript">{{
        `import { getCSSVariableValue } from "@/assets/ts/_utils";

const primaryColor = getCSSVariableValue("--bs-primary");
const breakpointLg = getCSSVariableValue("--bs-lg");
const fontFamily = getCSSVariableValue("--bs-font-sans-serif");`
      }}</CodeHighlighter>
    </div>
    <!--end::Code-->
  </div>
  <!--end::Section-->
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CodeHighlighter from "@/components/highlighters/CodeHighlighter.vue";

export default defineComponent({
  name: "line-height",
  components: {
    CodeHighlighter,
  },
});
</script>
