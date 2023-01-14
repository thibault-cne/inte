import { mount } from "@vue/test-utils";
import IDigit from "../digit.vue";

jest.useFakeTimers();

test("Test single digit", () => {
  const wrapper = mount(IDigit, {
    props: {
      digit: "1",
    },
  });

  const elem = wrapper.get('[data-test="digit"]');

  expect(elem.html({ raw: true })).toBe(
    '<span style="--value: 0;" data-test="digit" data-v-16d2bcb1=""></span>'
  );

  jest.runTimersToTime(2500);

  expect(elem.html({ raw: true })).toBe(
    '<span style="--value: 1;" data-test="digit" data-v-16d2bcb1=""></span>'
  );
});
