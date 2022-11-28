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
      Use <code>.rating</code> and <code>.rating-label</code> classes for
      displaying star based ratings. This examples uses a star icon from Duotune
      Svg Icons.
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <div class="rating">
          <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </div>
          <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </div>
          <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </div>
          <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </div>
          <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </div>
        </div>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html" :field-height="400"
      >{{`<div class="rating">
        <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </div>
        <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </div>
        <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </div>
        <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </div>
        <div class="rating-label checked">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
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
