<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <BaseButtons></BaseButtons>

      <HoverEffect></HoverEffect>

      <LightStyle></LightStyle>

      <BackgroundStyle></BackgroundStyle>

      <OutlineDashedStyle></OutlineDashedStyle>

      <ColorStyle></ColorStyle>

      <IconAndTextStyles></IconAndTextStyles>

      <ActiveStyle></ActiveStyle>

      <ActiveLightStyle></ActiveLightStyle>

      <ActiveColorStyle></ActiveColorStyle>

      <ActiveIconAndTextColorStyles></ActiveIconAndTextColorStyles>

      <Icons></Icons>

      <SocialButtons></SocialButtons>

      <LinkOptions></LinkOptions>

      <UtilityClasses></UtilityClasses>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/buttons/Overview.vue";
import BaseButtons from "@/views/base/buttons/BaseButtons.vue";
import HoverEffect from "@/views/base/buttons/HoverEffect.vue";
import LightStyle from "@/views/base/buttons/LightStyle.vue";
import BackgroundStyle from "@/views/base/buttons/BackgroundStyle.vue";
import OutlineDashedStyle from "@/views/base/buttons/OutlineDashedStyle.vue";
import ColorStyle from "@/views/base/buttons/ColorStyle.vue";
import IconAndTextStyles from "@/views/base/buttons/IconAndTextStyles.vue";
import ActiveStyle from "@/views/base/buttons/ActiveStyle.vue";
import ActiveLightStyle from "@/views/base/buttons/ActiveLightStyle.vue";
import ActiveColorStyle from "@/views/base/buttons/ActiveColorStyle.vue";
import ActiveIconAndTextColorStyles from "@/views/base/buttons/ActiveIconAndTextColorStyles.vue";
import Icons from "@/views/base/buttons/Icons.vue";
import SocialButtons from "@/views/base/buttons/SocialButtons.vue";
import LinkOptions from "@/views/base/buttons/LinkOptions.vue";
import UtilityClasses from "@/views/base/buttons/UtilityClasses.vue";

export default defineComponent({
  name: "base-buttons",
  components: {
    Overview,
    BaseButtons,
    HoverEffect,
    LightStyle,
    BackgroundStyle,
    OutlineDashedStyle,
    ColorStyle,
    IconAndTextStyles,
    ActiveStyle,
    ActiveLightStyle,
    ActiveColorStyle,
    ActiveIconAndTextColorStyles,
    Icons,
    SocialButtons,
    LinkOptions,
    UtilityClasses,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Buttons", ["Base"]);
    });
  },
});
</script>
