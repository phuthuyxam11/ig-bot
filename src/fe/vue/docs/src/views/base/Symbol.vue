<template>
  <!--begin::Card-->
  <div class="card">
    <!--begin::Card Body-->
    <div class="card-body fs-6 py-15 px-10 py-lg-15 px-lg-15 text-gray-700">
      <Overview></Overview>

      <Avatar></Avatar>

      <AvatarBackground></AvatarBackground>

      <AvatarRatio></AvatarRatio>

      <Label></Label>

      <Badge></Badge>

      <Styles></Styles>

      <Sizes></Sizes>

      <Group></Group>
    </div>
    <!--end::Card Body-->
  </div>
  <!--end::Card-->
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { setCurrentPageBreadcrumbs } from "@/core/helpers/breadcrumb";
import Overview from "@/views/base/symbol/Overview.vue";
import Avatar from "@/views/base/symbol/Avatar.vue";
import AvatarBackground from "@/views/base/symbol/AvatarBackground.vue";
import AvatarRatio from "@/views/base/symbol/AvatarRatio.vue";
import Label from "@/views/base/symbol/Label.vue";
import Badge from "@/views/base/symbol/Badge.vue";
import Styles from "@/views/base/symbol/Styles.vue";
import Sizes from "@/views/base/symbol/Sizes.vue";
import Group from "@/views/base/symbol/Group.vue";

export default defineComponent({
  name: "base-symbol",
  components: {
    Overview,
    Avatar,
    AvatarBackground,
    AvatarRatio,
    Label,
    Badge,
    Styles,
    Sizes,
    Group,
  },
  setup() {
    onMounted(() => {
      setCurrentPageBreadcrumbs("Symbol", ["Base"]);
    });
  },
});
</script>
