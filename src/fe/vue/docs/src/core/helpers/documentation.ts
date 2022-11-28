import { computed } from "vue";
import ClipboardJS from "clipboard";
/* eslint-disable  @typescript-eslint/no-explicit-any */

/**
 * Return name of the app
 * @returns {string}
 */
export const appName = computed(() => {
  return process.env.VUE_APP_NAME;
});

/**
 * Return version of the app
 * @returns {string}
 */
export const appVersion = computed(() => {
  return process.env.VUE_APP_VERSION;
});

/**
 * Return full name of the app
 * @returns {string}
 */
export const appFullName = computed(() => {
  return process.env.VUE_APP_FULL_NAME;
});

/**
 * Return logo of the app
 * @returns {string}
 */
export const appLogo = computed(() => {
  return process.env.VUE_APP_LOGO;
});

/**
 * Return bootstrap doc link
 * @returns {string}
 */
export const bootstrapDocLink = computed(() => {
  return process.env.VUE_APP_BOOTSTRAP_DOCS_LINK;
});

/**
 * Return sass components path
 * @returns {string}
 */
export const sassComponentsPath = computed(() => {
  return process.env.VUE_APP_SASS_COMPONENTS_PATH;
});

/**
 * Return link of the app preview
 * @returns {string}
 */
export const previewLink = computed(() => {
  return process.env.VUE_APP_PREVIEW_LINK;
});

/**
 * Returns link to purchase product
 * @returns {string}
 */
export const purchaseLink = computed(() => {
  return process.env.VUE_APP_PURCHESE_LINK;
});

/**
 * Returns link to youtube channel
 * @returns {string}
 */
export const youtubeLink = computed(() => {
  return process.env.VUE_APP_YOUTUBE_LINK;
});

/**
 * Returns link to github
 * @returns {string}
 */
export const githubLink = computed(() => {
  return process.env.VUE_APP_GITHUB_LINK;
});

/**
 * Returns link to twitter
 * @returns {string}
 */
export const twitterLink = computed(() => {
  return process.env.VUE_APP_TWITTER_LINK;
});

/**
 * Returns link to instagram
 * @returns {string}
 */
export const instagramLink = computed(() => {
  return process.env.VUE_APP_INSTAGRAM_LINK;
});

/**
 * Returns link to facebook
 * @returns {string}
 */
export const facebookLink = computed(() => {
  return process.env.VUE_APP_FACEBOOK;
});

/**
 * Returns link to dribble
 * @returns {string}
 */
export const dribbleLink = computed(() => {
  return process.env.VUE_APP_DRIBBBLE;
});

//code copy button initialization
export const useCopyClipboard = () => {
  const _init = (element) => {
    let elements = element;

    if (typeof elements === "undefined") {
      elements = document.querySelectorAll(".highlight");
    }

    if (elements && elements.length > 0) {
      for (let i = 0; i < elements.length; ++i) {
        const highlight = elements[i];
        const copy = highlight.querySelector(".highlight-copy");

        if (copy) {
          const clipboard = new ClipboardJS(copy, {
            target: (trigger): any => {
              const highlight = trigger.closest(".highlight");

              if (highlight) {
                let el: Element | null =
                  highlight.querySelector(".tab-pane.active");

                if (el == null) {
                  el = highlight.querySelector(".highlight-code");
                }

                return el as Element;
              }

              return highlight;
            },
          });

          clipboard.on("success", (e) => {
            const caption = e.trigger.innerHTML;

            e.trigger.innerHTML = "copied";
            e.clearSelection();

            setTimeout(function () {
              e.trigger.innerHTML = caption;
            }, 2000);
          });
        }
      }
    }
  };

  return {
    init: (element?) => {
      _init(element);
    },
  };
};
