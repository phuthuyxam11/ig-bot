<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <CustomFormControl></CustomFormControl>

      <FlushFormControls></FlushFormControls>

      <FormLabels></FormLabels>

      <CustomFormSelect></CustomFormSelect>

      <CustomCheckboxAndRadio></CustomCheckboxAndRadio>

      <CustomSwitch></CustomSwitch>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/forms/Overview.vue";
import CustomFormControl from "@/views/base/forms/CustomFormControl.vue";
import FlushFormControls from "@/views/base/forms/FlushFormControls.vue";
import FormLabels from "@/views/base/forms/FormLabels.vue";
import CustomFormSelect from "@/views/base/forms/CustomFormSelect.vue";
import CustomCheckboxAndRadio from "@/views/base/forms/CustomCheckboxAndRadio.vue";
import CustomSwitch from "@/views/base/forms/CustomSwitch.vue";

export default defineComponent({
  name: "base-forms",
  components: {
    Overview,
    CustomFormControl,
    FlushFormControls,
    FormLabels,
    CustomFormSelect,
    CustomCheckboxAndRadio,
    CustomSwitch,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Forms", ["Base"]);
    });
  },
});
</script>
