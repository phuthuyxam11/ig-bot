<template>
  <!--begin::Section-->
  <div class="pt-10">
    <!--begin::Heading-->
    <h1 class="anchor fw-bolder mb-5" id="edit-mode-example">
      <router-link to="#edit-mode-example"></router-link>
      Edit Mode Examples
    </h1>
    <!--end::Heading-->

    <!--begin::Block-->
    <div class="py-5">
      Enable editable rating with Label and Radio input elements. Hover and
      click on the below rating to set a rating value that can be submitted
      along with attached radio inputs:
    </div>
    <!--end::Block-->

    <!--begin::Block-->
    <div class="py-5">
      <div class="rounded border p-10">
        <div class="rating">
          <!--begin::Reset rating-->
          <label
            class="btn btn-light fw-bolder btn-sm rating-label me-3"
            for="kt_rating_input_0"
            >Reset</label
          >
          <input
            class="rating-input"
            name="rating"
            value="0"
            checked="checked"
            type="radio"
            id="kt_rating_input_0"
          />
          <!--end::Reset rating-->
          <!--begin::Star 1-->
          <label class="rating-label" for="kt_rating_input_1">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </label>
          <input
            class="rating-input"
            name="rating"
            value="1"
            type="radio"
            id="kt_rating_input_1"
          />
          <!--end::Star 1-->
          <!--begin::Star 2-->
          <label class="rating-label" for="kt_rating_input_2">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </label>
          <input
            class="rating-input"
            name="rating"
            value="2"
            type="radio"
            id="kt_rating_input_2"
          />
          <!--end::Star 2-->
          <!--begin::Star 3-->
          <label class="rating-label" for="kt_rating_input_3">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </label>
          <input
            class="rating-input"
            name="rating"
            value="3"
            type="radio"
            id="kt_rating_input_3"
          />
          <!--end::Star 3-->
          <!--begin::Star 4-->
          <label class="rating-label" for="kt_rating_input_4">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </label>
          <input
            class="rating-input"
            name="rating"
            value="4"
            type="radio"
            id="kt_rating_input_4"
          />
          <!--end::Star 4-->
          <!--begin::Star 5-->
          <label class="rating-label" for="kt_rating_input_5">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
          </label>
          <input
            class="rating-input"
            name="rating"
            value="5"
            type="radio"
            id="kt_rating_input_5"
          />
          <!--end::Star 5-->
        </div>
      </div>
    </div>
    <!--end::Block-->

    <!--begin::Code-->
    <div class="py-5">
      <!--prettier-ignore-->
      <CodeHighlighter lang="html" :field-height="400"
        >{{`<div class="rating">
        <!--begin::Reset rating-->
        <label
            class="btn btn-light fw-bolder btn-sm rating-label me-3"
            for="kt_rating_input_0"
        >Reset</label
        >
        <input
            class="rating-input"
            name="rating"
            value="0"
            checked="checked"
            type="radio"
            id="kt_rating_input_0"
        />
        <!--end::Reset rating-->
        <!--begin::Star 1-->
        <label class="rating-label" for="kt_rating_input_1">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </label>
        <input
            class="rating-input"
            name="rating"
            value="1"
            type="radio"
            id="kt_rating_input_1"
        />
        <!--end::Star 1-->
        <!--begin::Star 2-->
        <label class="rating-label" for="kt_rating_input_2">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </label>
        <input
            class="rating-input"
            name="rating"
            value="2"
            type="radio"
            id="kt_rating_input_2"
        />
        <!--end::Star 2-->
        <!--begin::Star 3-->
        <label class="rating-label" for="kt_rating_input_3">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </label>
        <input
            class="rating-input"
            name="rating"
            value="3"
            type="radio"
            id="kt_rating_input_3"
        />
        <!--end::Star 3-->
        <!--begin::Star 4-->
        <label class="rating-label" for="kt_rating_input_4">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </label>
        <input
            class="rating-input"
            name="rating"
            value="4"
            type="radio"
            id="kt_rating_input_4"
        />
        <!--end::Star 4-->
        <!--begin::Star 5-->
        <label class="rating-label" for="kt_rating_input_5">
            <span class="svg-icon svg-icon-1">
              <inline-svg src="media/icons/duotune/general/gen029.svg" />
            </span>
        </label>
        <input
            class="rating-input"
            name="rating"
            value="5"
            type="radio"
            id="kt_rating_input_5"
        />
        <!--end::Star 5-->
      </div>`}}</CodeHighlighter
      >

      <!--begin::Block-->
      <div class="py-5">
        Rating with default value can be setup through a radio input's
        <code>checked="checked"</code> state. This examples shows a rating with
        3 stars default value:
      </div>
      <!--end::Block-->
    </div>
    <!--end::Section-->
  </div>
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
