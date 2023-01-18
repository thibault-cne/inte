import { mount } from "@vue/test-utils";
import { afterEach, beforeEach, describe, it, vi } from "vitest";
import IDigit from "../digit.vue";

describe("Test single digit", () => {
  beforeEach(() => {
    vi.useFakeTimers();
  });

  afterEach(() => {
    vi.resetAllMocks();
  });

  it("Test single digit on init", () => {
    const wrapper = mount(IDigit, {
      props: {
        digit: "1",
      },
    });

    const elem = wrapper.get('[data-test="digit"]');

    expect(elem.html({ raw: true })).toBe(
      '<span style="--value: 0;" data-test="digit" data-v-16d2bcb1=""></span>'
    );
  });
});
